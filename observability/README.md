# Observability

In my last years in my DevOps journey, one of my weak points is the observability.

## Logging: 
I'm ok. Everything that is printed to stdout can be collected with tools like Fluent bit to ingest to few monitoring tools like Cloudwatch or Grafana.

## Metrics:
I'm ok. From Prometheus we can measure CPU, Memory or storage and send this metrics to a monitoring tool such as Grafana.

## Tracing: 
I know about ElasticSearch a little. But instrumenting applications is my weakness. Could be due to lack of experience.

So, today I will start with it. Let's see the targets.

### What's the goal?

Learn how to instrument Golang applications with:

- ElasticSearch
- OpenTelemetry

### Why ElasticSearch?

It's commonly used across companies and also in my job is used. So I could be able to help to improve current instrumentation.

### Why OpenTelemetry?

In my last job was very mentioned as an standard library for applications instrumentation



### Tools to reach out further:

- [Signoz](https://github.com/SigNoz/signoz)


https://grafana.com/docs/grafana-cloud/monitor-applications/application-observability/setup/instrument/go/#set-up-a-telemetry-data-collector
