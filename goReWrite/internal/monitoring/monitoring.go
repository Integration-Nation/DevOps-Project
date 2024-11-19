package monitoring

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	ConcurrentRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_concurrent_requests",
			Help: "Number of concurrent HTTP requests being processed.",
		},
	)

	TotalErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors encountered.",
		},
	)
)

var (
	diskUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "disk_usage_bytes",
			Help: "Disk usage in bytes.",
		},
		[]string{"path"},
	)

	memoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Memory usage in bytes.",
		},
	)
	cpuUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percent",
			Help: "CPU usage percentage.",
		},
	)
)


func init() {
	prometheus.MustRegister(ConcurrentRequests, TotalErrors, diskUsage, memoryUsage, cpuUsage)
}

func CollectSystemMetrics() {
	for {
		// Collect memory usage
		vm, err := mem.VirtualMemory()
		if err == nil {
			memoryUsage.Set(float64(vm.Used))
		} else {
			log.Println("Error collecting memory metrics:", err)
		}

		// Collect disk usage
		usage, err := disk.Usage("/")
		if err == nil {
			diskUsage.WithLabelValues("/").Set(float64(usage.Used))
		} else {
			log.Println("Error collecting disk metrics:", err)
		}

		// Collect CPU usage
		percentages, err := cpu.Percent(0, false) // Aggregate CPU usage (false: overall usage, not per core)
		if err == nil && len(percentages) > 0 {
			cpuUsage.Set(percentages[0]) // Use the first value for overall CPU usage
		} else {
			log.Println("Error collecting CPU metrics:", err)
		}

		time.Sleep(10 * time.Second) // Adjust collection interval as needed
	}
}
