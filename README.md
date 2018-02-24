# presentation-gitlab-k8s
These are the example files for my presentation about GitLab + Kubernetes for Continuous Integration and Delivery.
The files are also mainly used in my GitLab post series about the GitLab CI runners.

**INFO** This isn't the best way to deploy to K8s, this is a more of an example how simple it can be.

The presentation can be found here: [Kubernetes - WYNTK - GitLab CI + Kubernetes Presentation](https://edenmal.moe/2017/07/31/Kubernetes-WYNTK-GitLab-CI-Kubernetes-Presentation/).
The blog post these files were specifically used in is here: [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.moe/2017/07/04/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/).

An uptodate list of all my blog posts around GitLab and Kubernetes can be found on [this page](https://edenmal.moe/tags/GitLab/).
This list is just an excerpt:
* [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.moe/2017/07/04/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/)
* [Kubernetes - WYNTK - GitLab CI + Kubernetes Presentation](https://edenmal.moe/2017/07/31/Kubernetes-WYNTK-GitLab-CI-Kubernetes-Presentation/)
* [GitLab + Kubernetes: Running CI Runners in Kubernetes](https://edenmal.moe/2017/08/31/GitLab-Kubernetes-Running-CI-Runners-in-Kubernetes/)
* [GitLab + Kubernetes: GitLab on top of Kubernetes](https://edenmal.moe/2017/11/04/GitLab-Kubernetes-GitLab-on-top-of-Kubernetes/)
* [GitLab: Use Keycloak as SAML 2.0 OmniAuth Provider](https://edenmal.moe/2018/01/16/GitLab-Keycloak-SAML-2-0-OmniAuth-Provider/)

## Using these files
You have to replace the following addresses in all files:
* `gitlab.edenmal.net` with your GitLab address (example `gitlab.example.com`).
* `registry.edenmal.net`/`registry.zerbytes.net` with your Docker Registry address (example `registry.example.com`).
* `edenmal.net` (in the Ingress manifest) with your Domain name.

You also need to create a "Docker Login" Secret named `regsecret` in the `presentation-gitlab-k8s` Namespace. The Namespace manifest is in the [`manifests/`](/manifests/) directory.
The guide for that can be found here: [Kubernetes.io - Pull an Image from a Private Registry](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/).

Then you can just import the repository into your GitLab instance and are ready to go.

## GitLab Docs References
* GitLab Kubernetes Integration Docs: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html
* GitLab Kubernetes Integration Docs Environment variables: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html#deployment-variables

As of GitLab `10.3` the Kubernetes Integration is marked as deprecated and with `10.4` it is now disabled, the following docs show the new feature called Clusters:
* GitLab 10.3 release - Kubernetes integration service: https://about.gitlab.com/2017/12/22/gitlab-10-3-released/#kubernetes-integration-service
* GitLab Clusters Feature Docs: https://docs.gitlab.com/ce/user/project/clusters/index.html

## Thanks!
Thanks to [@shadycuz - GitHub](https://github.com/shadycuz) for his comments with improvements for the code in this repository!
