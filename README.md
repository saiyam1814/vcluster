<br>

<div align="center">
  <a href="https://www.vcluster.com" target="_blank">
    <picture>
      <!-- For Dark Mode -->
      <source media="(prefers-color-scheme: dark)" srcset="docs/static/media/vcluster_horizontal_orange_white.svg">
      <!-- For Light Mode -->
      <source media="(prefers-color-scheme: light)" srcset="docs/static/media/vcluster_horizontal_orange_black.svg">
      <!-- Fallback -->
      <img alt="vCluster Logo" src="docs/static/media/vcluster_horizontal_orange_white.svg" width="500">
    </picture>
  </a>
</div>

 
[![X](https://img.shields.io/twitter/follow/loft_sh?style=social)](https://twitter.com/loft_sh)
[![Documentation](https://img.shields.io/badge/docs-vcluster-blue)](https://www.vcluster.com/docs/what-are-virtual-clusters)
[![GoDoc](https://godoc.org/github.com/loft-sh/vcluster?status.svg)](https://godoc.org/github.com/loft-sh/vcluster)
[![License](https://img.shields.io/github/license/loft-sh/vcluster)](LICENSE)
[![Join us on Slack!](docs/static/media/slack.svg)](https://slack.loft.sh/)
[![Open in DevPod!](https://devpod.sh/assets/open-in-devpod.svg)](https://devpod.sh/open#https://github.com/loft-sh/vcluster)
![Latest Release](https://img.shields.io/github/v/release/loft-sh/vcluster?style=for-the-badge&label=Latest%20Release&color=%23007ec6)
![License: Apache-2.0](https://img.shields.io/github/license/loft-sh/vcluster?style=for-the-badge&color=%23007ec6)
  <a href="https://www.vcluster.com">
    <img src="docs/static/media/vcluster_square_icon.svg" alt="Website" width="30" height="30"> Website
  </a>

Virtual clusters are fully functional Kubernetes clusters nested within a physical host cluster, offering robust isolation and flexibility ideal for multi-tenancy. With vCluster, multiple teams can work autonomously on shared infrastructure, reducing conflicts, enhancing productivity, and optimizing cost-efficiency.

<br>

### Watch our 30-second Quickstart Video:
[![Kubernetes workshop 101](https://img.youtube.com/vi/gQ-KG57ruvY/0.jpg)](https://www.youtube.com/watch?v=gQ-KG57ruvY)


## Benefits

Virtual clusters provide immense benefits for large-scale Kubernetes deployments and multi-tenancy.

<img src="docs/static/media//diagrams/vcluster-comparison.png" width="500">

### Security and Isolation

- **Granular Permissions:** Operate with minimized host cluster permissions while maintaining admin-level control within each vCluster, supporting independent CRD, RBAC, and policy management.
- **Dedicated Control Plane:** Each vCluster includes its own API server and control plane, creating an isolation boundary.
- **Custom Security Policies:** Tenants can enforce specific governance measures such as OPA policies, resource quotas, and network policies, complementing existing host security configurations.
- **Data Protection:** Choose from isolated data stores such as embedded SQLite, etcd, or external databases for enhanced data management.

### Tenant Access and Flexibility

- **Full Admin Access:** Tenants can deploy CRDs, create namespaces, and manage cluster-scoped resources without host restrictions.
- **Integrated Networking:** Configurable network policies ensure secure, isolated tenant environments with optional inter-cluster communication.
- **Node Management:** Support static node assignments and shared node pools for optimized resource allocation.

### Cost-Effective Scaling

- **Lightweight Deployment:** Launch virtual clusters in seconds compared to physical clusters with longer provisioning times (e.g., EKS ~45 minutes).
- **Resource Sharing:** Efficiently utilize underlying host resources to lower infrastructure costs.
- **Simple Operations:** A compact control plane and integrated CoreDNS simplify management, supporting scalability for large environments.

### Compatibility and Adaptability

- **Multi-Version Support:** vCluster runs different Kubernetes versions (K8s, K3s, K0s), accommodating version skews and ensuring flexibility.
- **Diverse Backing Stores:** Choose data storage solutions to meet specific performance and durability needs.
- **Wide Deployment Scope:** Deploy vCluster on EKS, GKE, AKS, OpenShift, K3s, and more, ensuring consistent operation across cloud, edge, and on-prem environments.

### Scalability Enhancements

- **API Server Load Reduction:** Each vCluster’s dedicated API server mitigates load on the host cluster’s API server, improving performance.
- **CRD Management:** Avoid CRD conflicts with independent CRD administration within each vCluster.

## Common Use Cases

### Pre-Production Environments

- **Developer Self-Service:** Empower developers with self-service virtual clusters for quick, isolated environments, reducing errors and enhancing productivity.
- **Efficient CI/CD Pipelines:** Use ephemeral clusters for isolated testing and PR previews, avoiding shared environment bottlenecks.

### Production Deployments

- **Customer-Focused Isolation:** Offer a dedicated vCluster per customer for ISVs, ensuring tenant security and smooth scalability.
- **Managed Kubernetes Services:** Build services with reduced COGS and high margins by consolidating resources with tenant-specific vClusters.

## Public Talks and Blogs

Explore our [public talks](https://www.vcluster.com/talks) and [blogs](https://loft.sh/blog) for insights from industry experts and real-world use cases, showcasing vCluster's adoption by notable companies.

## Quickstart Guide

Deploy your first vCluster with a minimal configuration:

```bash
brew install loft-sh/tap/vcluster
vcluster create my-vcluster --namespace team-x
```
For detailed steps, visit our [quickstart documentation](https://www.vcluster.com/docs/get-started).

## Contributing

Thank you for your interest in contributing! Please refer to
[CONTRIBUTING.md](https://github.com/loft-sh/vcluster/blob/main/CONTRIBUTING.md) for guidance.

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

<http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
