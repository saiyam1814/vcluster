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
| **v0.29** | [Standalone Mode](https://www.vcluster.com/docs/vcluster/deploy/standalone) | Run vCluster without a host cluster—directly on bare metal or VMs |
| **v0.28** | [Auto Nodes](https://www.vcluster.com/docs/vcluster/deploy/worker-nodes/auto-nodes) | Karpenter-powered dynamic autoscaling for private nodes |
| **v0.27** | [Private Nodes](https://www.vcluster.com/docs/vcluster/deploy/worker-nodes/private-nodes) | External nodes with full CNI/CSI isolation |

👉 **[Full Changelog](https://github.com/loft-sh/vcluster/releases)**

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

👉 **[Full Architecture Guide](https://www.vcluster.com/docs/vcluster/introduction/architecture/)**

<details>
<summary><strong>Shared Nodes</strong> — Maximum density, minimum cost</summary>

Virtual clusters share the host cluster's nodes. Workloads run as regular pods in a namespace.

```yaml
sync:
  fromHost:
    nodes:
      enabled: false  # Uses pseudo nodes
```
</details>

<details>
<summary><strong>Dedicated Nodes</strong> — Isolated compute on labeled node pools</summary>

Virtual clusters get their own set of labeled host nodes. Workloads are isolated but still managed by the host.

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
<summary><strong>Private Nodes</strong> <sup>v0.27+</sup> — Full CNI/CSI isolation</summary>

External nodes join the virtual cluster directly with their own CNI, CSI, and networking stack. Complete workload isolation from the host cluster.

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
<summary><strong>vCluster Standalone</strong> <sup>v0.29+</sup> — No host cluster required</summary>

Run vCluster without any host cluster. Deploy the control plane directly on bare metal or VMs. The highest level of isolation—vCluster becomes the cluster.

👉 **[Standalone Docs](https://www.vcluster.com/docs/vcluster/deploy/standalone)**

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
| **⚡ Auto Nodes** | Karpenter-powered autoscaling for private nodes across any infrastructure <sup>v0.28+</sup> |

---

## 📊 Architecture Comparison

| Architecture | Host Cluster Required | Node Isolation | CNI/CSI Isolation | Best For |
|-------------|:--------------------:|:--------------:|:-----------------:|----------|
| **Shared Nodes** | ✅ | ❌ | ❌ | Dev/test, cost savings |
| **Dedicated Nodes** | ✅ | ✅ | ❌ | Production multi-tenancy |
| **Private Nodes** | ✅ | ✅ | ✅ | Compliance, GPU workloads |
| **Standalone** | ❌ | ✅ | ✅ | Bare metal, edge, air-gapped |

---

## 🌍 Deploy Anywhere

**Cloud:** EKS • GKE • AKS • DigitalOcean • Linode

**On-Prem:** OpenShift • Rancher • Tanzu • Private Cloud

**Bare Metal:** Standalone Mode • Private Nodes

**Virtualization:** KubeVirt • Proxmox • VMware

---

## 🔧 How It Works

```
┌────────────────────────────────────────────────────────────┐
│                     Host Cluster                           │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              vCluster Namespace                     │   │
│  │  ┌─────────────────────────────────────────────┐    │   │
│  │  │           Virtual Cluster                   │    │   │
│  │  │  ┌─────────────┐  ┌─────────────────────┐   │    │   │
│  │  │  │ API Server  │  │  Controller Manager │   │    │   │
│  │  │  └─────────────┘  └─────────────────────┘   │    │   │
│  │  │  ┌─────────────┐  ┌─────────────────────┐   │    │   │
│  │  │  │   Syncer    │  │      CoreDNS        │   │    │   │
│  │  │  └─────────────┘  └─────────────────────┘   │    │   │
│  │  └─────────────────────────────────────────────┘    │   │
│  │                                                     │   │
│  │  ┌────────────┐ ┌────────────┐ ┌────────────┐       │   │
│  │  │  Pod (app) │ │  Pod (app) │ │  Service   │       │   │
│  │  └────────────┘ └────────────┘ └────────────┘       │   │
│  └─────────────────────────────────────────────────────┘   │
└────────────────────────────────────────────────────────────┘
```

Each virtual cluster includes:
- **Dedicated API server** — Full Kubernetes API isolation
- **Syncer** — Bi-directional resource synchronization  
- **Optional components** — Scheduler, CoreDNS, etcd

👉 **[Detailed Architecture Docs](https://www.vcluster.com/docs/vcluster/deploy/architecture)**

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

We welcome contributions! Check out our **[Contributing Guide](https://github.com/loft-sh/vcluster/blob/main/CONTRIBUTING.md)** to get started.

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
