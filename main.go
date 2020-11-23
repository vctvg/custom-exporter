package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
)

type CustomManager struct {
	Label string
	CustomDescription *prometheus.Desc
}

func (c *CustomManager)RealScratch()(RealMetric map[string]float64){
	RealMetric = map[string]float64{
		"127.0.0.1": rand.Float64()*100,
	}
	return
}

func (c *CustomManager)Describe(ch chan<- *prometheus.Desc){
	ch <- c.CustomDescription
}

func (c *CustomManager)Collect(ch chan<- prometheus.Metric){
	RealMetric := c.RealScratch()
	for host, num := range RealMetric{
		ch <- prometheus.MustNewConstMetric(
			c.CustomDescription,
			prometheus.GaugeValue,
			num,
			host,
		)
	}
}

func NewCustomManager() *CustomManager{
	return &CustomManager{
		CustomDescription: prometheus.NewDesc(
			"test_value",
			"help information",
			[]string{"host"},
			nil,
		),
	}
}

func main(){
	worker := NewCustomManager()
	reg := prometheus.NewPedanticRegistry()

	reg.MustRegister(worker)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8080", nil)
}