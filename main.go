/**
 * @Author: dy
 * @Description:
 * @File: main
 * @Date: 2022/12/22 15:07
 */
package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cluster = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cluster",
		Help: "集群总数",
	})
	cluster_state_started = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cluster_state_started",
		Help: "集群状态正常",
	})

	cluster_state_failed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cluster_state_failed",
		Help: "集群状态异常",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cluster, cluster_state_started, cluster_state_failed)
}

func main() {
	go func() {
		for true {
			cluster.Set(float64(randInt()))
			cluster_state_started.Set(float64(randInt()))
			cluster_state_failed.Set(float64(randInt()))
			time.Sleep(5 * time.Second)
		}

	}()
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9100", nil))
}

func randInt() int64 {
	return int64(rand.Intn(5000))
}
