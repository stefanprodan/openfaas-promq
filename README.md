# openfaas-promq

OpenFaaS function that executes Prometheus queries

### Usage

![OpenFaaS](https://github.com/stefanprodan/openfaas-promq/blob/master/screens/invoke-ui.jpg)

Deploy:

```bash
$ faas-cli deploy -f ./promq.yml --gateway=http://<GATEWAY-IP> 
```

Invoke:

```bash
$ echo -n '{"server": "http://prometheus.openfaas:9090", "query": "sum(increase(gateway_function_invocation_total[1h]))  by (function_name)", "start": "5 hours ago", "end": "now", "step": "1h","format": "table"}' | faas-cli invoke promq --gateway=<GATEWAY-IP>

       time| function_name:certinfo| function_name:nodeinfo| function_name:promq
 1512293537|               0.000000|            1001.390821|                    
 1512297137|               1.001391|            1001.390821|                    
 1512300737|               0.000000|               0.000000|                    
 1512304337|               0.000000|               0.000000|                    
 1512307937|               0.000000|               0.000000|                    
 1512311537|               3.004172|               3.004172|           13.148504
```

Parameters:

* server: default `http://prometheus.openfaas:9090` (can also be set with `PROMETHEUS_URL` environment variable)
* query: promQL format
* start: default `1 hour ago`
* end: default `now`
* step: default `15s`
* format: `table` or `json`

```json
{
  "server": "http://prometheus.openfaas:9090",
  "query": "sum(increase(gateway_function_invocation_total[1h]))  by (function_name)",
  "start": "5 hours ago",
  "end": "1 hour ago",
  "step": "1h",
  "format": "table"
}
```

