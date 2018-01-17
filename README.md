# presentation-gitlab-k8s
These are the example files for my presentation about GitLab + Kubernetes for Continuous Integration and Delivery.
The files are also mainly used in my GitLab post series about the GitLab CI runners.

**INFO** This isn't the best way to deploy to K8s, this is a more of an example how simple it can be.

The presentation can be found here: [Kubernetes - WYNTK - GitLab CI + Kubernetes Presentation](https://edenmal.net/2017/07/31/Kubernetes-WYNTK-GitLab-CI-Kubernetes-Presentation/).
The blog post these files were specifically used in is here: [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.net/2017/07/04/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/).

All my blog posts around GitLab and Kubernetes can be found here:
* [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.net/2017/07/04/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/)
* [GitLab + Kubernetes: Running CI Runners in Kubernetes](https://edenmal.net/2017/08/31/GitLab-Kubernetes-Running-CI-Runners-in-Kubernetes/)
* [GitLab + Kubernetes: GitLab on top of Kubernetes](https://edenmal.net/2017/11/04/GitLab-Kubernetes-GitLab-on-top-of-Kubernetes/)

## Using these files
You have to replace the following strings in all files:
* `gitlab.edenmal.moe` with your GitLab address (example `gitlab.example.com`)
* `registry.edenmal.moe` with your Docker Registry address (example `registry.example.com`)

Then you can just import the files/repository.

## GitLab Docs References
* GitLab Kubernetes Integration Docs: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html
* GitLab Kubernetes Integration Docs Environment variables: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html#deployment-variables

As of GitLab `10.3` the Kubernetes Integration is marked as deprecated, the following docs show the new feature
called Clusters:
* GitLab 10.3 release - Kubernetes integration service: https://about.gitlab.com/2017/12/22/gitlab-10-3-released/#kubernetes-integration-service
* GitLab Clusters Feature Docs: https://docs.gitlab.com/ce/user/project/clusters/index.html

