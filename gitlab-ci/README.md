This directory contains:
* Namespace (`presentation-gitlab-k8s`).
* ServiceAccount for GitLab CI deployment.
* RBAC Role and RoleBinding for this Namespace.
* TLS Certificate Secret for this Namespace (for K8s TLS Ingress).
* ServiceMonitor (in `monitoring/` directory) for PrometheusOperator monitoring.

The ServiceAccount and RBAC manifests are required for clusters with RBAC enabled.
You will then need the ServiceAccount token for the GitLab Kubernetes integration/GitLab Kubernetes Cluster feature.

For more information, see the repo root [README.md](../README.md).
