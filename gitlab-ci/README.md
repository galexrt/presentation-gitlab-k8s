This directory contains:

* `namespace.yaml` - Namespace (`presentation-gitlab-k8s`).
* `secrets.yaml` - TLS Certificate Secret for this Namespace (for K8s TLS Ingress).
* `managed-cluster` -
  * RBAC Role and ClusterRoleBinding for the cluster, so that GitLab can manage the Kubernetes cluster for you. Has advantages and disadvantages, e.g., might not be a good choice when you share the cluster with multiple tenants.
* `monitoring/*` - Contains `prometheus-operator` manifests of Prometheus and ServiceMonitor to easily get some monitoring for your application started.
* `per-namespace/` - .

The RBAC + ServiceAccount manifests are required for clusters with RBAC enabled.
You will then need the ServiceAccount token for the GitLab Kubernetes integration/GitLab Kubernetes Cluster feature.

For more information, see the repo root [README.md](../README.md) and the GitLab documentation [https://git.test.code.nrw/help/user/project/clusters/index.md](https://git.test.code.nrw/help/user/project/clusters/index.md).
