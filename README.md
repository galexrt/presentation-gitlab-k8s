# presentation-gitlab-k8s
These are the example files for my presentation about GitLab + Kubernetes for Continuous Integration and Delivery.
These files may also be used in my other GitLab posts about GitLab CI.

![Kubernetes and GitLab](/media/kubernetes-and-gitLab.png)

**INFO** This isn't the best way to deploy to K8s, this is more of an example how simple it can be.

The presentation can be found here: [Kubernetes - WYNTK - GitLab CI + Kubernetes Presentation](https://edenmal.moe/post/2017/Kubernetes-WYNTK-GitLab-CI-Kubernetes-Presentation/).
The blog post these files were specifically used in is here: [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.moe/post/2017/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/).

An uptodate list of all my blog posts around GitLab and Kubernetes can be found on [this page](https://edenmal.moe/tags/gitlab/).
This list is just an excerpt:
* [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.moe/post/2017/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/)
* [Kubernetes - WYNTK - GitLab CI + Kubernetes Presentation](https://edenmal.moe/post/2017/Kubernetes-WYNTK-GitLab-CI-Kubernetes-Presentation/)
* [GitLab + Kubernetes: Running CI Runners in Kubernetes](https://edenmal.moe/post/2017/GitLab-Kubernetes-Running-CI-Runners-in-Kubernetes/)
* [GitLab + Kubernetes: GitLab on top of Kubernetes](https://edenmal.moe/post/2017/GitLab-Kubernetes-GitLab-on-top-of-Kubernetes/)
* [GitLab: Use Keycloak as SAML 2.0 OmniAuth Provider](https://edenmal.moe/post/2018/GitLab-Keycloak-SAML-2-0-OmniAuth-Provider/)

## Requirements
The following points are required for this repository to work correctly:
* GitLab (>= `10.4`) with the following features configured:
    * [Container Registry](https://docs.gitlab.com/ce/user/project/container_registry.html)
    * [GitLab CI](https://about.gitlab.com/features/gitlab-ci-cd/) (with working [GitLab CI Runners](https://docs.gitlab.com/ce/ci/runners/))
* [Kubernetes](https://kubernetes.io/) cluster
    * You need to be "bound" to the `admin` (`cluster-admin`) ClusterRole, see [Kubernetes.io Using RBAC Authorization - User-facing Roles](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles).
    * An Ingress controller should already been deployed, see [Kubernetes.io Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/).
* `kubectl` installed locally.
* Editor of your choice.

## Using this repository
You have to replace the following addresses in all files:
* `gitlab.zerbytes.net` with your GitLab address (e.g. `gitlab.example.com`).
* `registry.zerbytes.net` with your Docker registry address (e.g. `registry.example.com`).
* `edenmal.net` (in the Ingress manifest) with your domain name.
    * You probably also want to change the subdomain name while you are at it.
* `presentatio-gitlab-k8s` with the Namespace name of your choice.

You also need to create a "Docker Login" Secret which contains your GitLab Registry access data (e.g. Username and Access token with registry access) named `regsecret` in the Namespace `presentation-gitlab-k8s`.
A guide for that can be found here: [Kubernetes.io - Pull an Image from a Private Registry](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/).
The Namespace manifest is in the [`gitlab-ci/`](/gitlab-ci/) directory.

Then you can just import the repository into your GitLab instance and are ready to go.

## GitLab Docs References
* GitLab Kubernetes Integration Docs: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html
* GitLab Kubernetes Integration Docs Environment variables: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html#deployment-variables

As of GitLab `10.3` the Kubernetes Integration is marked as deprecated and with `10.4` it is now disabled, the following docs show the new feature called Clusters:
* GitLab 10.3 release - Kubernetes integration service: https://about.gitlab.com/2017/12/22/gitlab-10-3-released/#kubernetes-integration-service
* GitLab Clusters Feature Docs: https://docs.gitlab.com/ce/user/project/clusters/index.html

## File Structure
### Example Application
* [`main.go`](/main.go) - The Golang example application code.
* [`vendor/`](/vendor/) - Contains the Golang example application dependencies (`dep` is used).

### Kubernetes Base GitLab CI Manifests
* [`gitlab-ci/`](/gitlab-ci/)
    * [`namespace.yaml`](/gitlab-ci/namespace.yaml) - Namespace in which the GitLab CI will deploy the application.
    * [`rbac.yaml`](/gitlab-ci/rbac.yaml) - Contains GitLab CI RBAC Role, RoleBinding and ServiceAccount.
    * [`secret.yaml`](/gitlab-ci/secret.yaml) - Contains a TLS wildcard certificate for the application Ingress.

### Build Process
* [`Dockerfile`](/Dockerfile) - Contains the Docker image build instructions.
* [`.gitlab-ci.yml`](/.gitlab-ci.yml) - Contains the GitLab CI instructions.

### Deployment Manifests
* [`manifests/`](/manifests/) - Kubernetes manifests used to deploy the Docker image built in the CI pipeline.
    * [`deployment.yaml`](/manifests/deployment.yaml) - Deployment for the Docker image.
    * [`ingress.yaml`](/manifests/ingress.yaml) - Ingress for the application.
    * [`service.yaml`](/manifests/service.yaml) - Service for the application.

### Miscellaneous
* [`media/`](/media/) - Contains media for the [`README.md`](/README.md) in this repository.

## Thanks!
Thanks to [@shadycuz - GitHub](https://github.com/shadycuz) for his comments with improvements for the code in this repository!

## ToDo
- [ ] Implement app review, see https://gitlab.com/gitlab-examples/review-apps-nginx/blob/master/.gitlab-ci.yml
