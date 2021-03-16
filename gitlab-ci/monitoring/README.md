# monitoring
A Grafana dashboard to visualize the metrics from the demo Golang application can be found here: [https://grafana.com/dashboards/6671](https://grafana.com/dashboards/6671).

The `prometheus.yaml` contains a [prometheus-operator/prometheus-operator](https://github.com/prometheus-operator/prometheus-operator) `Prometheus` CustomResource object. It can be used together with the next file `servicemonitor.yaml`

The `servicemonitor.yaml` contains a [prometheus-operator/prometheus-operator](https://github.com/prometheus-operator/prometheus-operator) `ServiceMonitor` CustomResource. If you use it with an existing Prometheus, make sure the Prometheus is configured to pickup the `ServiceMonitor` in the namespace (`serviceMonitorNamespaceSelector:`) and the `ServiceMonitor` itself (`serviceMonitorSelector:`).

***

For more information on the Prometheus Custom Resource and prometheus-operator other CRs, checkout:

* https://github.com/prometheus-operator/prometheus-operator
* https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md
