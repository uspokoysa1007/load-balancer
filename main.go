package main

import "fmt"

func main() {
	fmt.Println("load-balancer v0.1.0 — scaffold")
	// TODO:
	// - reverse proxy на net/http/httputil
	// - стратегии: round-robin, least-connections, random
	// - health-check бэкендов
	// - конфиг-файл (YAML/JSON), retries, таймауты
}
