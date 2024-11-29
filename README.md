# CloudQuery concurrency-childtable-example Source Plugin

Reproduction of the concurrency behavior for child relation tables:
- if items are added all at once, the concurrency for the child table is up to 5 (which is configurable [here](https://github.com/cloudquery/plugin-sdk/blob/698c50f7a942b45e4e7055a12bcc9e148a78661d/scheduler/scheduler.go#L24C2-L24C40) I believe)
- if items are added one by one, the child relation is processed one by one

## Screenshot

![CleanShot 2024-11-29 at 08 41 11@2x](https://github.com/user-attachments/assets/943d2eef-71c8-4ba5-b6e4-f8be0ce0c226)

## Reproduce

- Clone this repo
- Run jaeger tracing as decribed in the [cloudquery documentation](https://docs.cloudquery.io/docs/advanced-topics/monitoring/overview#opentelemetry-preview)
- Run `go build . && cloudquery sync .`
- Navigate to http://localhost:16686/search?service=cloudquery-concurrency-childtable-example and click `Find traces`

Tested with Cloudquery cli v6.3.0