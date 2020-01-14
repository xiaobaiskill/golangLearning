prometheus
===

### Metrics
* Counter
```
累加器，  只能加不能减。 例如： 任务执行数， 错误数，创建次数 
```

* Gauge
```
可变数。  例如： cpu利用率，内存使用情况，正在执行的个数，当前访问数，一秒执行任务量。。。
```

* Histogram
```
柱状图， 用于观察结果的采样
```

* Summary
```
表示一段时间内数据采样结果~~~~
```

### Vec
* 不加Vec
```
NewCounter(opts CounterOpts) Counter
NewGauge(opts GaugeOpts) Gauge
NewSummary(opts SummaryOpts) Summary
NewHistogram(opts HistogramOpts) Histogram 
```
* 加Vec
```
NewCounterVec(opts CounterOpts, labelNames []string) *CounterVec
NewGaugeVec(opts GaugeOpts, labelNames []string) *GaugeVec
NewSummaryVec(opts SummaryOpts, labelNames []string) *SummaryVec
NewHistogramVec(opts HistogramOpts, labelNames []string) *HistogramVec
```
* 区别
```
加vec 的可以有固定种类 多个分支
不加vec 只有一个分支
```

### 自定义监听
```
register := promethesu.NewRegister()
register.MustRegister(...)

func metricsHandle(w http.ResponseWriter, r *http.Request) {
	promhttp.HandlerFor(register, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}
```