<div align="center">
  <a href="https://www.vcluster.com" target="_blank">

<picture>
      <!-- For Dark Mode -->
      <source media="(prefers-color-scheme: dark)" srcset="docs/static/media/vcluster_horizontal_orange_white.svg">
      <!-- For Light Mode -->
      <source media="(prefers-color-scheme: light)" srcset="docs/static/media/vcluster_horizontal_orange_black.svg">
      <!-- Fallback -->
      <img alt="vCluster Logo" src="docs/static/media/vcluster_horizontal_orange_white.svg" width="600">
</picture>	  

  </a>

<br/><br/>

### Flexible Tenancy For Kubernetes and AI Infra

[![GitHub stars](https://img.shields.io/github/stars/loft-sh/vcluster?style=for-the-badge&logo=github&color=orange)](https://github.com/loft-sh/vcluster/stargazers)
[![Slack](https://img.shields.io/badge/Slack-4.2K+-blue?style=for-the-badge&logo=slack)](https://slack.loft.sh/)
[![License](https://img.shields.io/github/license/loft-sh/vcluster?style=for-the-badge)](https://github.com/loft-sh/vcluster/blob/main/LICENSE)
[![Contributors](https://img.shields.io/github/contributors/loft-sh/vcluster?style=for-the-badge&logo=github)](https://github.com/loft-sh/vcluster/graphs/contributors)

**[Website](https://www.vcluster.com)** • **[Quickstart](https://www.vcluster.com/docs/get-started/)** • **[Documentation](https://www.vcluster.com/docs/what-are-virtual-clusters)** • **[Blog](https://loft.sh/blog)** • **[Slack](https://slack.loft.sh/)**

</div>

---

## What is vCluster?

**vCluster** creates fully functional virtual Kubernetes clusters that run inside namespaces of a host cluster. Each virtual cluster has its own API server, runs on shared or dedicated infrastructure, and gives you flexible tenancy options—from simple namespaces to fully dedicated clusters.

**40M+ virtual clusters deployed** by companies like Adobe, CoreWeave, Atlan, and NVIDIA.

<div align="center">

![vCluster gif](./docs/static/media/vcluster-github-gif-1280.gif)

</div>

---

## 🚀 Quick Start

```bash
# Install vCluster CLI
brew install loft-sh/tap/vcluster

# Create a virtual cluster
vcluster create my-vcluster --namespace team-x

# Connect to the virtual cluster
vcluster connect my-vcluster --namespace team-x

# Use kubectl as usual - you're now in your virtual cluster!
kubectl get namespaces
```

**Prerequisites:** A running Kubernetes cluster and `kubectl` configured.

👉 **[Full Quickstart Guide](https://www.vcluster.com/docs/get-started)**

---

## 🎯 Use Cases

### GPU & AI Infrastructure

Build your internal GPU cloud with vCluster. Developers get fast, secure access to GPUs, and your organization maximizes utilization of every card—without sacrificing isolation.

```yaml
# vcluster.yaml - GPU tenant isolation
sync:
  fromHost:
    nodes:
      enabled: true
      selector:
        labels:
          nvidia.com/gpu: "true"
```

### Bare Metal Kubernetes

Run Kubernetes on bare metal with zero VMs. Virtual clusters and virtual nodes give you isolation without the expensive overhead.

```yaml
# vcluster.yaml - Standalone on bare metal
controlPlane:
  standalone:
    enabled: true
    joinNode: 
      enabled: true
privateNodes:
  enabled: true
```

### Platform Engineering

Build secure, scalable, multi-tenant Kubernetes environments. Empower every team with isolated, self-service access—without managing more physical clusters.

```yaml
# vcluster.yaml - Multi-tenant platform
networking:
  podCIDR: "10.244.0.0/16"
policies:
  networkPolicy:
    enabled: true
  resourceQuota:
    enabled: true
```

---

## 🏗️ Tenancy Models

vCluster supports flexible tenancy to match your isolation requirements:

| Model | Description | Use Case |
|-------|-------------|----------|
| **Shared Nodes** | Virtual clusters share the host's nodes and plugins | Dev environments, cost optimization |
| **Dedicated Nodes** | Virtual clusters run on their own set of host-assigned nodes | Production workloads, tenant isolation |
| **Private Nodes** | Fully separate nodes with their own CNI, CSI, and control | Maximum isolation, compliance |

<details>
<summary><strong>Shared Nodes Example</strong></summary>

```yaml
# vcluster.yaml
sync:
  fromHost:
    nodes:
      enabled: false  # Use pseudo nodes
```
</details>

<details>
<summary><strong>Dedicated Nodes Example</strong></summary>

```yaml
# vcluster.yaml
sync:
  fromHost:
    nodes:
      enabled: true
      selector:
        labels:
          tenant: tenant-1
```
</details>

<details>
<summary><strong>Private Nodes Example</strong></summary>

```yaml
# vcluster.yaml
privateNodes:
  enabled: true
controlPlane:
  service:
    spec:
      type: NodePort  # or LoadBalancer
```
</details>

---

## ✨ Key Features

### 🔒 Private Nodes <sup>NEW</sup>
Assign dedicated worker nodes to virtual clusters with complete isolation. Workloads run directly on private nodes with their own CNI, CSI, and networking stack.

### ⚡ Auto Nodes <sup>NEW</sup>
Dynamic autoscaling powered by [Karpenter](https://karpenter.sh/). Scale virtual clusters across public cloud, private cloud, hybrid, and bare metal—eliminating cloud vendor lock-in.

### 🖥️ Standalone Mode <sup>NEW</sup>
Run vCluster directly on bare metal or VMs without a host cluster. Deploy the control plane on dedicated nodes for maximum isolation.

### 🔄 Resource Syncing
Sync any Kubernetes resource between virtual and host clusters. Built-in support for pods, services, secrets, configmaps, and CRDs.

### 🔌 Integrations
Native integrations with cert-manager, external-secrets, KubeVirt, Istio, and metrics-server.

### 📊 High Availability
Run multiple replicas with leader election. Use embedded etcd or external databases (PostgreSQL, MySQL, RDS) as backing stores.

---

## 🌍 Deploy Anywhere

<table>
<tr>
<td align="center"><strong>EKS</strong></td>
<td align="center"><strong>GKE</strong></td>
<td align="center"><strong>AKS</strong></td>
<td align="center"><strong>OpenShift</strong></td>
<td align="center"><strong>Rancher</strong></td>
</tr>
<tr>
<td align="center"><strong>Bare Metal</strong></td>
<td align="center"><strong>KubeVirt</strong></td>
<td align="center"><strong>Tanzu</strong></td>
<td align="center"><strong>Private Cloud</strong></td>
<td align="center"><strong>Standalone</strong></td>
</tr>
</table>

---

## 📖 Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Host Cluster                            │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              vCluster Namespace                      │   │
│  │  ┌─────────────────────────────────────────────┐    │   │
│  │  │           Virtual Cluster                    │    │   │
│  │  │  ┌─────────────┐  ┌─────────────────────┐   │    │   │
│  │  │  │ API Server  │  │  Controller Manager │   │    │   │
│  │  │  └─────────────┘  └─────────────────────┘   │    │   │
│  │  │  ┌─────────────┐  ┌─────────────────────┐   │    │   │
│  │  │  │   Syncer    │  │      CoreDNS        │   │    │   │
│  │  │  └─────────────┘  └─────────────────────┘   │    │   │
│  │  └─────────────────────────────────────────────┘    │   │
│  │                                                      │   │
│  │  ┌────────────┐ ┌────────────┐ ┌────────────┐       │   │
│  │  │  Pod (app) │ │  Pod (app) │ │  Service   │       │   │
│  │  └────────────┘ └────────────┘ └────────────┘       │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

Each virtual cluster includes:
- **Dedicated API server** — Full Kubernetes API isolation
- **Syncer** — Bi-directional resource synchronization  
- **Optional components** — Scheduler, CoreDNS, etcd

---

## 🌟 Why vCluster?

<details>
<summary><strong>Security & Isolation</strong></summary>

- **Granular permissions** — Users operate with minimized host cluster permissions while having admin-level control within their vCluster
- **Isolated control plane** — Each vCluster has its own dedicated API server
- **Customizable policies** — Apply OPA, network policies, resource quotas per tenant
- **Separate backing stores** — SQLite, embedded etcd, or external databases
</details>

<details>
<summary><strong>Full Tenant Access</strong></summary>

- **Admin capabilities** — Deploy CRDs, create namespaces, manage cluster-scoped resources
- **Isolated networking** — Pods in different virtual clusters cannot communicate by default
- **Node management** — Assign static nodes or share node pools
</details>

<details>
<summary><strong>Cost Efficiency</strong></summary>

- **Lightweight** — Spin up clusters in seconds (vs. ~45 min for EKS)
- **Resource sharing** — Minimize infrastructure by sharing host resources
- **Single pod control plane** — Minimal operational overhead
</details>

<details>
<summary><strong>Flexibility</strong></summary>

- **Multi-version support** — Run different Kubernetes versions with version skew
- **Multiple distros** — K8s, K3s, and more
- **Adaptable storage** — SQLite to enterprise etcd to RDS
</details>

---

## 🏢 Trusted By

<table>
<tr>
<td><strong>Atlan</strong><br/>100 → 1 clusters</td>
<td><strong>Aussie Broadband</strong><br/>99% faster provisioning</td>
<td><strong>CoreWeave</strong><br/>GPU cloud at scale</td>
</tr>
<tr>
<td><strong>Adobe</strong><br/>Enhanced dev environments</td>
<td><strong>Codefresh</strong><br/>Hosted ArgoCD</td>
<td><strong>Scanmetrix</strong><br/>99% faster deployments</td>
</tr>
</table>

**Also used by:** NVIDIA, ABBYY, Lintasarta, Precisely, Shipwire, Trade Connectors, and many more.

👉 **[View Case Studies](https://www.vcluster.com/case-studies)**

---

## 📚 Learn More

### Conference Talks

| Event | Title | Link |
|-------|-------|------|
| HashiConf 2025 | GPU sharing done right with Vault and vCluster | [Watch](https://www.youtube.com/watch?v=zWx17azSqyU) |
| KubeCon EU 2023 | How We Securely Scaled Multi-Tenancy with vCluster | [Watch](https://www.youtube.com/watch?v=hFiHU6W4_z0) |
| KubeCon NA 2022 | How Adobe Planned For Scale With vCluster | [Watch](https://www.youtube.com/watch?v=p8BluR5WT5w) |
| KubeCon NA 2021 | Virtual Clusters are the Future of Multi-Tenancy | [Watch](https://www.youtube.com/watch?v=QddWNqchD9I) |

### Community Resources

| Channel | Title | Link |
|---------|-------|------|
| TeKanAid | Build Your IDP with Backstage, Crossplane, and ArgoCD | [Watch](https://www.youtube.com/watch?v=nIxl2PcEs-0) |
| DevOps Toolkit | How To Create Virtual Kubernetes Clusters | [Watch](https://www.youtube.com/watch?v=JqBjpvp268Y) |
| TechWorld with Nana | Build your Self-Service Kubernetes Platform | [Watch](https://www.youtube.com/watch?v=tt7hope6zU0) |
| Kubesimplify | Let's Learn vCluster | [Watch](https://www.youtube.com/watch?v=I4mztvnRCjs) |

👉 **[YouTube Channel](https://www.youtube.com/@vcluster)** • **[Blog](https://loft.sh/blog)**

---

## 🤝 Contributing

We love contributions! Whether it's bug fixes, new features, or documentation improvements.

```bash
# Quick start with DevPod
```

[![Open in DevPod!](https://devpod.sh/assets/open-in-devpod.svg)](https://devpod.sh/open#https://github.com/loft-sh/vcluster)

👉 **[Contributing Guide](https://github.com/loft-sh/vcluster/blob/main/CONTRIBUTING.md)**

---

## 🔗 Links

| Resource | Link |
|----------|------|
| 📖 Documentation | [vcluster.com/docs](https://www.vcluster.com/docs/what-are-virtual-clusters) |
| 💬 Slack Community | [slack.loft.sh](https://slack.loft.sh/) |
| 🌐 Website | [vcluster.com](https://www.vcluster.com) |
| 🐦 X (Twitter) | [@vcluster](https://x.com/vcluster) |
| 💼 LinkedIn | [vCluster](https://www.linkedin.com/company/vcluster) |
| 💬 Chat with Expert | [Start Chat](https://start-chat.com/slack/Loft-Labs/NnQl1M) |

---

## 📜 License

vCluster is licensed under the **[Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0)**.

---

<div align="center">

**© 2025 [Loft Labs](https://loft.sh). All rights reserved.**

Made with ❤️ by the vCluster community.

⭐ **Star us on GitHub** — it helps!

</div>
