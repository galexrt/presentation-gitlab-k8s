This directory contains the actual manifests used to deploy the built Docker image to Kubernetes.
The manifests contain placeholders (like `__CI_ENVIRONMENT_SLUG__`) which are replaced by the GitLab CI pipeline.

For more information, see the repo root [README.md](../README.md) and the [.gitlab-ci.yml](../.gitlab-ci.yml).
