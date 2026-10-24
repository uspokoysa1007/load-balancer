# load-balancer

HTTP load balancer / reverse proxy на Go.

## Возможности (план)

- reverse proxy для HTTP/HTTPS;
- стратегии: round-robin, least-connections, random;
- health-check бэкендов с автоисключением упавших;
- retries и таймауты;
- конфигурация через YAML/JSON;
- метрики (запросы/сек, статусы).

## Стек

- Язык: Go
- Планируемые библиотеки: stdlib `net/http/httputil`, `gopkg.in/yaml.v3`

## Сборка и запуск

```bash
go build -o bin/load-balancer .
./bin/load-balancer -config config.yaml
```

## Лицензия

MIT
