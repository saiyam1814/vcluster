<div align="center">
  <a href="https://www.vcluster.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="docs/static/media/vcluster_horizontal_orange_white.svg">
      <source media="(prefers-color-scheme: light)" srcset="docs/static/media/vcluster_horizontal_orange_black.svg">
      <img alt="vCluster" src="docs/static/media/vcluster_horizontal_orange_white.svg" width="400">
    </picture>
  </a>
  <p><strong>Flexible Tenancy For Kubernetes and AI Infra</strong></p>

[![GitHub stars](https://img.shields.io/github/stars/loft-sh/vcluster?style=for-the-badge&logo=github&color=orange)](https://github.com/loft-sh/vcluster/stargazers)
[![Slack](https://img.shields.io/badge/Slack-4.2K+-4A154B?style=for-the-badge&logo=slack&logoColor=white)](https://slack.loft.sh/)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-14K+-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/company/vcluster)
[![X](https://img.shields.io/badge/X-3.5K+-000000?style=for-the-badge&logo=x&logoColor=white)](https://x.com/loft_sh)

**[Website](https://www.vcluster.com)** • **[Quickstart](https://www.vcluster.com/docs/get-started/)** • **[Documentation](https://www.vcluster.com/docs/vcluster/introduction/what-are-virtual-clusters)** • **[Blog](https://loft.sh/blog)** • **[Slack](https://slack.loft.sh/)**

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

# Use kubectl as usual - you're now in your virtual cluster!
kubectl get namespaces
```

**Prerequisites:** A running Kubernetes cluster and `kubectl` configured.

👉 **[Full Quickstart Guide](https://www.vcluster.com/docs/get-started)**

### 🎮 Try Without Installing

No Kubernetes cluster? Try vCluster instantly in your browser:

[![Try on Killercoda](https://img.shields.io/badge/Try%20on-Killercoda-22B573?style=for-the-badge&logo=kubernetes&logoColor=white)](https://killercoda.com/vcluster)

---

## 🆕 What's New

| Version | Feature | Description |
|---------|---------|-------------|
| **v0.29** | [Standalone Mode](https://www.vcluster.com/docs/vcluster/deploy/control-plane/binary/) | Run vCluster without a host cluster—directly on bare metal or VMs |
| **v0.28** | [Auto Nodes](https://www.vcluster.com/docs/vcluster/deploy/worker-nodes/private-nodes/auto-nodes/) | Karpenter-powered dynamic autoscaling for private nodes |
| **v0.27** | [Private Nodes](https://www.vcluster.com/docs/vcluster/deploy/worker-nodes/private-nodes) | External nodes with full CNI/CSI isolation |

👉 **[Full Changelog](https://www.vcluster.com/releases)**

---

## 🎯 Use Cases

| Use Case | Description | Learn More |
|----------|-------------|------------|
| **GPU Cloud Providers** | Launch managed K8s for GPUs. Give customers isolated, production-grade Kubernetes fast. | [View →](https://www.vcluster.com/solutions/gpu-cloud-providers) |
| **Internal GPU Platform** | Maximize GPU utilization without sacrificing isolation. Self-service access for AI/ML teams. | [View →](https://www.vcluster.com/solutions/internal-gpu-platform) |
| **AI Factory** | Run AI on-prem where your data lives. Multi-tenant K8s for training, fine-tuning, inference. | [View →](https://www.vcluster.com/solutions/ai-factory) |
| **Bare Metal K8s** | Run Kubernetes on bare metal with zero VMs. Isolation without expensive overhead. | [View →](https://www.vcluster.com/solutions/bare-metal-kubernetes) |
| **Software Vendors** | Ship Kubernetes-native software. Each customer gets their own isolated virtual cluster. | [View →](https://www.vcluster.com/solutions/software-vendors) |
| **Cost Savings** | Cut Kubernetes costs by consolidating clusters. Sleep mode pauses inactive clusters. | [View →](https://www.vcluster.com/cost-savings) |

---

## 🏗️ Architectures

vCluster offers multiple deployment architectures. Each builds on the previous, offering progressively more isolation.

### Architecture Comparison

| | **Shared Nodes** | **Dedicated Nodes** | **Private Nodes** | **Standalone** |
|---|:---:|:---:|:---:|:---:|
| **Host Cluster** | Required | Required | Required | Not Required |
| **Node Isolation** | ❌ | ✅ | ✅ | ✅ |
| **CNI/CSI Isolation** | ❌ | ❌ | ✅ | ✅ |
| **Best For** | Dev/test, cost | Production | Compliance, GPU | Bare metal, edge |

👉 **[Full Architecture Guide](https://www.vcluster.com/docs/vcluster/introduction/architecture/)**

### Minimal Configuration

**Step 1:** Create a `vcluster.yaml`:

```yaml
# Shared Nodes (default) - no config needed
# Or customize for your architecture:
sync:
  fromHost:
    nodes:
      enabled: false  # Shared Nodes
```

**Step 2:** Apply it:

```bash
vcluster create my-vcluster -f vcluster.yaml
```

<details>
<summary><strong>🔹 Shared Nodes</strong> — Maximum density, minimum cost</summary>

Virtual clusters share the host cluster's nodes. Workloads run as regular pods in a namespace.

<div align="center">
<img src="./assets/vcluster-architecture-shared-nodes.png" alt="Shared Nodes Architecture" width="600">
</div>

```yaml
sync:
  fromHost:
    nodes:
      enabled: false  # Uses pseudo nodes
```
</details>

<details>
<summary><strong>🔹 Dedicated Nodes</strong> — Isolated compute on labeled node pools</summary>

Virtual clusters get their own set of labeled host nodes. Workloads are isolated but still managed by the host.

<div align="center">
<img src="./assets/vcluster-architecture-dedicated-nodes.png" alt="Dedicated Nodes Architecture" width="600">
</div>

```yaml
sync:
  fromHost:
    nodes:
      enabled: true
      selector:
        labels:
          tenant: my-tenant
```
</details>

<details>
<summary><strong>🔹 Private Nodes</strong> <sup>v0.27+</sup> — Full CNI/CSI isolation</summary>

External nodes join the virtual cluster directly with their own CNI, CSI, and networking stack. Complete workload isolation from the host cluster.

<div align="center">
<img src="./assets/vcluster-architecture-private-nodes.png" alt="Private Nodes Architecture" width="600">
</div>

👉 **[Private Nodes Docs](https://www.vcluster.com/docs/vcluster/deploy/worker-nodes/private-nodes)**

```yaml
privateNodes:
  enabled: true
controlPlane:
  service:
    spec:
      type: NodePort
```
</details>

<details>
<summary><strong>🔹 vCluster Standalone</strong> <sup>v0.29+</sup> — No host cluster required</summary>

Run vCluster without any host cluster. Deploy the control plane directly on bare metal or VMs. The highest level of isolation—vCluster becomes the cluster.

<div align="center">
<img src="./assets/vcluster-architecture-standalone.png" alt="Standalone Architecture" width="600">
</div>

👉 **[Standalone Docs](https://www.vcluster.com/docs/vcluster/deploy/control-plane/binary/)**

```yaml
controlPlane:
  standalone:
    enabled: true
    joinNode:
      enabled: true
privateNodes:
  enabled: true
```
</details>

<details>
<summary><strong>⚡ Auto Nodes</strong> <sup>v0.28+</sup> — Karpenter-powered dynamic autoscaling</summary>

Automatically provision and deprovision private nodes based on workload demand. Works across public cloud, private cloud, hybrid, and bare metal environments.

<div align="center">
<img src="./assets/vcluster-architecture-auto-nodes.png" alt="Auto Nodes Architecture" width="600">
</div>

👉 **[Auto Nodes Docs](https://www.vcluster.com/docs/vcluster/deploy/worker-nodes/private-nodes/auto-nodes/)**

```yaml
autoNodes:
  enabled: true
  nodeProvider: <provider>
privateNodes:
  enabled: true
```
</details>

---

## ✨ Key Features

| Feature | Description |
|---------|-------------|
| **🎛️ Isolated Control Plane** | Each vCluster gets its own API server, controller manager, and data store—complete Kubernetes API isolation |
| **🔗 Shared Platform Stack** | Leverage the host cluster's CNI, CSI, ingress, and other infrastructure—no duplicate platform components |
| **🔒 Security & Multi-Tenancy** | Tenants get admin access inside their vCluster while having minimal permissions on the host cluster |
| **🔄 Resource Syncing** | Bidirectional sync of any Kubernetes resource. Pods, services, secrets, configmaps, CRDs, and more |
| **💤 Sleep Mode** | Pause inactive virtual clusters to save resources. Instant wake when needed |
| **🔌 Integrations** | Native support for cert-manager, external-secrets, KubeVirt, Istio, and metrics-server |
| **📊 High Availability** | Multiple replicas with leader election. Embedded etcd or external databases (PostgreSQL, MySQL, RDS) |

---

## 🌍 Deploy Anywhere

<div align="center">

| **Cloud** | **On-Prem** | **Bare Metal** | **Virtualization** |
|:---------:|:-----------:|:--------------:|:------------------:|
| EKS | OpenShift | Standalone Mode | KubeVirt |
| GKE | Rancher | Private Nodes | Proxmox |
| AKS | Tanzu | | VMware |
| DigitalOcean | Private Cloud | | |
| Linode | | | |

</div>

---

## 🏢 Trusted By

<table>
<tr>
<td align="center"><a href="https://www.vcluster.com/case-studies/atlan"><strong>Atlan</strong></a><br/>100 → 1 clusters</td>
<td align="center"><a href="https://www.vcluster.com/case-studies/aussie-broadband"><strong>Aussie Broadband</strong></a><br/>99% faster provisioning</td>
<td align="center"><a href="https://www.vcluster.com/case-studies/coreweave"><strong>CoreWeave</strong></a><br/>GPU cloud at scale</td>
</tr>
<tr>
<td align="center"><a href="https://www.vcluster.com/case-studies/adobe"><strong>Adobe</strong></a><br/>Enhanced dev environments</td>
<td align="center"><a href="https://www.vcluster.com/case-studies/codefresh"><strong>Codefresh</strong></a><br/>Hosted ArgoCD</td>
<td align="center"><a href="https://www.vcluster.com/case-studies/scanmetrix"><strong>Scanmetrix</strong></a><br/>99% faster deployments</td>
</tr>
<tr>
<td align="center"><a href="https://www.vcluster.com/case-studies/deloitte"><strong>Deloitte</strong></a><br/>Enterprise K8s platform</td>
<td align="center"><a href="https://www.vcluster.com/case-studies/gofundme"><strong>GoFundMe</strong></a><br/>Developer productivity</td>
<td align="center"><a href="https://www.vcluster.com/case-studies/nscale"><strong>NScale</strong></a><br/>AI infrastructure</td>
</tr>
</table>

**Also used by:** NVIDIA, ABBYY, Lintasarta, Precisely, Shipwire, Trade Connectors, and many more.

👉 **[View All Case Studies](https://www.vcluster.com/case-studies)**

---

## 📚 Learn More

<details>
<summary><strong>🎤 Conference Talks</strong></summary>

| Event | Title | Speaker | Link |
|-------|-------|---------|------|
| Open Source Finance Forum 2025 | Banking on Open Source: Security in Multi-Tenant Environments | Lukas Gentele, Joshua Bucknor | [Watch](https://www.youtube.com/@vcluster) |
| HashiConf 2025 | GPU sharing done right with Vault and vCluster | Rich Burroughs | [Watch](https://www.youtube.com/watch?v=zWx17azSqyU) |
| KubeCon EU 2023 | How We Securely Scaled Multi-Tenancy with vCluster | Atlan Team | [Watch](https://www.youtube.com/watch?v=hFiHU6W4_z0) |
| KubeCon NA 2022 | How Adobe Planned For Scale With vCluster | Adobe Team | [Watch](https://www.youtube.com/watch?v=p8BluR5WT5w) |
| KubeCon NA 2021 | Virtual Clusters are the Future of Multi-Tenancy | Lukas Gentele | [Watch](https://www.youtube.com/watch?v=QddWNqchD9I) |

</details>

<details>
<summary><strong>🎬 Community Resources</strong></summary>

| Channel | Title | Link |
|---------|-------|------|
| TeKanAid | Build Your IDP with Backstage, Crossplane, and ArgoCD | [Watch](https://www.youtube.com/watch?v=nIxl2PcEs-0) |
| DevOps Toolkit | How To Create Virtual Kubernetes Clusters | [Watch](https://www.youtube.com/watch?v=JqBjpvp268Y) |
| TechWorld with Nana | Build your Self-Service Kubernetes Platform | [Watch](https://www.youtube.com/watch?v=tt7hope6zU0) |
| Kubesimplify | Let's Learn vCluster | [Watch](https://www.youtube.com/watch?v=I4mztvnRCjs) |

</details>

👉 **[YouTube Channel](https://www.youtube.com/@vcluster)** • **[Blog](https://loft.sh/blog)**

---

## 🤝 Contributing

We welcome contributions! Check out our **[Contributing Guide](https://github.com/loft-sh/vcluster/blob/main/CONTRIBUTING.md)** to get started.

---

## 🔗 Links

| Resource | Link |
|----------|------|
| 📖 Documentation | [vcluster.com/docs](https://www.vcluster.com/docs/vcluster/introduction/what-are-virtual-clusters) |
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
