
通过启动 prometheus 来采集服务的各种运行指标，方便监控。

启动一个 prometheus

```bash
docker run -p 9090:9090 -v /Users/fuzhifei/Code/go/src/fwds/deploy/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```
