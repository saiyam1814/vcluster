package cli

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/loft-sh/log"
	"github.com/loft-sh/vcluster/pkg/cli/flags"
	"github.com/loft-sh/vcluster/pkg/constants"
)

func ResumeDocker(ctx context.Context, globalFlags *flags.GlobalFlags, vClusterName string, log log.Logger) error {
	containerName := getControlPlaneContainerName(vClusterName)

	// check if container exists
	exists, running, err := checkDockerContainerState(ctx, containerName)
	if err != nil {
		return fmt.Errorf("failed to check container state: %w", err)
	}

	if !exists {
		return fmt.Errorf("vCluster container %s not found", containerName)
	}

	if running {
		log.Infof("vCluster %s is already running", vClusterName)
		return nil
	}

	// start the container
	log.Infof("Resuming vCluster %s...", vClusterName)
	err = startDockerContainerByName(ctx, containerName)
	if err != nil {
		return fmt.Errorf("failed to resume vCluster: %w", err)
	}
	if err := restoreMountPropagation(ctx, containerName); err != nil {
		return fmt.Errorf("failed to restore mount propagation on control plane: %w", err)
	}

	// start the nodes
	nodes, err := findDockerContainer(ctx, constants.DockerNodePrefix+vClusterName+".")
	if err != nil {
		return fmt.Errorf("failed to find vCluster nodes: %w", err)
	}
	for _, node := range nodes {
		log.Infof("Starting node %s from vCluster %s...", node.Name, vClusterName)
		workerContainer := getWorkerContainerName(vClusterName, node.Name)
		err = startDockerContainerByName(ctx, workerContainer)
		if err != nil {
			return fmt.Errorf("failed to start vCluster node: %w", err)
		}
		if err := restoreMountPropagation(ctx, workerContainer); err != nil {
			return fmt.Errorf("failed to restore mount propagation on node %s: %w", node.Name, err)
		}
	}

	// start the load balancers
	loadBalancers, err := findDockerContainer(ctx, constants.DockerLoadBalancerPrefix+vClusterName+".")
	if err != nil {
		return fmt.Errorf("failed to find vCluster load balancers: %w", err)
	}
	for _, loadBalancer := range loadBalancers {
		log.Infof("Starting load balancer %s from vCluster %s...", loadBalancer.Name, vClusterName)
		err = startDockerContainerByName(ctx, constants.DockerLoadBalancerPrefix+vClusterName+"."+loadBalancer.Name)
		if err != nil {
			return fmt.Errorf("failed to start vCluster load balancer: %w", err)
		}
	}

	log.Donef("Successfully resumed vCluster %s", vClusterName)
	return nil
}

func startDockerContainerByName(ctx context.Context, containerName string) error {
	args := []string{"start", containerName}
	output, err := exec.CommandContext(ctx, "docker", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker start failed: %w, output: %s", err, string(output))
	}
	return nil
}

// restoreMountPropagation re-applies rshared mount propagation on a
// privileged container. The flag is set during vcluster create but lost
// when the container is stopped. Without it kubelet cannot propagate
// mounts (e.g. bpf, cgroupv2, netns) to pods.
func restoreMountPropagation(ctx context.Context, containerName string) error {
	output, err := exec.CommandContext(ctx, "docker", "exec", containerName, "mount", "--make-rshared", "/").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to restore mount propagation: %w, output: %s", err, string(output))
	}
	return nil
}
