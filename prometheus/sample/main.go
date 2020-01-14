package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func metricsHandle(w http.ResponseWriter, r *http.Request) {
	promhttp.HandlerFor(m.register, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/metrics", metricsHandle)
	http.Handle("/metrics_base", promhttp.Handler()) // 最基本的使用
	http.ListenAndServe(":8082", nil)
}
