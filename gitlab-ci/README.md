This directory contains:
* Namespace (`presentation-gitlab-k8s`).
* ServiceAccount for GitLab CI deployment.
* RBAC Role and RoleBinding for this Namespace.
* TLS Certificate Secret for this Namespace (for K8s TLS Ingress).

The ServiceAccount and RBAC manifests are required for clusters with RBAC enabled.
You will use the ServiceAccount token for the GitLab Kubernetes integration/GitLab Kubernetes Cluster.

For more information, see the repo root [README.md](../README.md).
