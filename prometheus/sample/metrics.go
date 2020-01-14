package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

type metrics struct {
	register        *prometheus.Registry

	jobTotalCounter prometheus.Counter
	jobGaugeVec     *prometheus.GaugeVec
	cpugauge  prometheus.Gauge
}

func (m *metrics) InitMetrics(metricsPrefix string) {
	m.jobGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "Job",
		Help:      "正在运行中的job数",
	}, []string{"jobName"})
	m.jobTotalCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: metricsPrefix,
		Name:      "Job_total",
		Help:      "运行中的job 总数",
	})
	m.cpugauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:"cpu",
		Help:"cpu 利用率",
	})

	m.register = prometheus.NewRegistry()
	m.register.MustRegister(m.jobGaugeVec, m.jobTotalCounter,m.cpugauge)
}

func (m *metrics) AddJobTotal(num float64) {
	m.jobTotalCounter.Add(num)
}

func (m *metrics) JobInc(value string) {
	m.jobGaugeVec.With(prometheus.Labels{"jobName": value}).Inc() // inc =>+1
}
func (m *metrics) JobDnc(value string) {
	m.jobGaugeVec.With(prometheus.Labels{"jobName": value}).Dec() // Dec =>-1
}
func (m *metrics) cpu(num float64){
	m.cpugauge.Set(num)
}

var m metrics

func init() {
	m = metrics{}
	m.InitMetrics("prometheus_test")

	// 测试哦
	go func() {
		var i = 1
		var mstr = map[string]string{"jobname1": "1", "jobname2": "2", "jobname3": "3"}
		for {
			k, _ := GetRangeMap(mstr)
			if i%4 == 0 {
				m.JobDnc(k)
			} else {
				m.JobInc(k)
			}

			m.AddJobTotal(1)
			time.Sleep(time.Second)
			i++
		}
	}()

	go func(){
		for {
			m.cpu(rand.Float64() * 100)
			time.Sleep( 500 * time.Millisecond)
		}
	}()

}

