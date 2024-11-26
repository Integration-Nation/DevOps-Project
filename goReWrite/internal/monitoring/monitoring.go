package monitoring

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"gorm.io/gorm"
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

	TotalUsers = promauto.With(registry).NewGauge(
		prometheus.GaugeOpts{
			Namespace: "whoknows",
			Name:      "total_users",
			Help:      "Total number of registered users.",
		},
	)

	ActiveUsers = promauto.With(registry).NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "whoknows",
			Name:      "active_users",
			Help:      "Number of active users in different time periods.",
		},
		[]string{"period"}, // period can be "daily", "weekly", "monthly"
	)
	SQLQueryDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "sql_query_duration_seconds",
		Help:    "Duration of SQL queries in seconds",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 10), // Fra 1 ms til ~1 sekund
	},
	[]string{"query_type"}, // Labels som "SELECT", "INSERT", "UPDATE", "DELETE"
)
)

// Metrik-definitions
var (
	DBActiveConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_active_connections",
		Help: "Number of active connections to the database.",
	})
	DBIdleConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_idle_connections",
		Help: "Number of idle connections to the database.",
	})
	DBInUseConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_in_use_connections",
		Help: "Number of in-use connections to the database.",
	})
	DBMaxOpenConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_max_open_connections",
		Help: "Maximum number of open connections to the database.",
	})
)


func init() {
	// Register metrics with the default registry
	prometheus.MustRegister(
		SystemDiskUsage,
		SystemMemoryUsage,
		SystemCPUUsage,
		HTTPConcurrentRequests,
		HTTPTotalErrors,
		TotalUsers,
		ActiveUsers,
		DBActiveConnections,
		SQLQueryDuration,
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

// UpdateTotalUsers updates the total number of registered users
func UpdateTotalUsers(count int) {
	TotalUsers.Set(float64(count))
}

// UpdateActiveUsers updates the number of active users for a specific time period
func UpdateActiveUsers(period string, count int) {
	ActiveUsers.WithLabelValues(period).Set(float64(count))
}
func UpdateDBMetrics(db *gorm.DB) {
	// Hent den underliggende databaseforbindelse
	sqlDB, err := db.DB()
	if err != nil {
		// Hvis vi ikke kan få sql.DB, log fejlen og returnér
		return
	}

	stats := sqlDB.Stats()

	// Opdater Prometheus-metrikker
	DBActiveConnections.Set(float64(stats.OpenConnections))
	DBIdleConnections.Set(float64(stats.Idle))
	DBInUseConnections.Set(float64(stats.InUse))
	DBMaxOpenConnections.Set(float64(sqlDB.Stats().MaxOpenConnections))
}

func ObserveSQLQuery(queryType string, start time.Time) {
	duration := time.Since(start).Seconds()
	SQLQueryDuration.WithLabelValues(queryType).Observe(duration)
}
