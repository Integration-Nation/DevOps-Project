package monitoring

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	registry = prometheus.NewRegistry()

	// System Metrics
	SystemDiskUsage = promauto.With(registry).NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "whoknows",
			Name:      "system_disk_usage_bytes",
			Help:      "Disk usage in bytes.",
		},
		[]string{"path"},
	)

	SystemMemoryUsage = promauto.With(registry).NewGauge(
		prometheus.GaugeOpts{
			Namespace: "whoknows",
			Name:      "system_memory_usage_bytes",
			Help:      "Memory usage in bytes.",
		},
	)

	SystemCPUUsage = promauto.With(registry).NewGauge(
		prometheus.GaugeOpts{
			Namespace: "whoknows",
			Name:      "system_cpu_usage_percent",
			Help:      "CPU usage percentage.",
		},
	)

	// HTTP Metrics
	HTTPConcurrentRequests = promauto.With(registry).NewGauge(
		prometheus.GaugeOpts{
			Namespace: "whoknows",
			Name:      "http_concurrent_requests",
			Help:      "Number of concurrent HTTP requests being processed.",
		},
	)

	HTTPTotalErrors = promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: "whoknows",
			Name:      "http_total_errors",
			Help:      "Total number of HTTP errors encountered.",
		},
	)
)

func init() {
	// Register metrics with the default registry
	prometheus.MustRegister(
		SystemDiskUsage,
		SystemMemoryUsage,
		SystemCPUUsage,
		HTTPConcurrentRequests,
		HTTPTotalErrors,
	)
}

func CollectSystemMetrics() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Memory Usage
		vm, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Error collecting memory metrics: %v", err)
		} else {
			SystemMemoryUsage.Set(float64(vm.Used))
			log.Printf("Memory Usage: %v bytes", vm.Used)
		}

		// Disk Usage
		usage, err := disk.Usage("/")
		if err != nil {
			log.Printf("Error collecting disk metrics: %v", err)
		} else {
			SystemDiskUsage.WithLabelValues("/").Set(float64(usage.Used))
			log.Printf("Disk Usage: %v bytes", usage.Used)
		}

		// CPU Usage
		percentages, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Printf("Error collecting CPU metrics: %v", err)
		} else if len(percentages) > 0 {
			SystemCPUUsage.Set(percentages[0])
			log.Printf("CPU Usage: %.2f%%", percentages[0])
		}
	}
}
