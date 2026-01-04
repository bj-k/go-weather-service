package handlers

import (
	"encoding/json"
	"expvar"
	"net/http"
	"runtime/metrics"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	results := make(map[string]interface{})

	// 1. Add Custom Metrics (from expvar)
	expvar.Do(func(kv expvar.KeyValue) {
		var val interface{}
		if err := json.Unmarshal([]byte(kv.Value.String()), &val); err != nil {
			results[kv.Key] = kv.Value.String()
		} else {
			results[kv.Key] = val
		}
	})

	// 2. Add Runtime Metrics
	descs := metrics.All()
	samples := make([]metrics.Sample, len(descs))
	for i := range samples {
		samples[i].Name = descs[i].Name
	}
	metrics.Read(samples)

	for _, sample := range samples {
		switch sample.Value.Kind() {
		case metrics.KindUint64:
			results[sample.Name] = sample.Value.Uint64()
		case metrics.KindFloat64:
			results[sample.Name] = sample.Value.Float64()
		case metrics.KindFloat64Histogram:
			results[sample.Name] = "histogram"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
