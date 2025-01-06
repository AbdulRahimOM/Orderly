package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "resp_code"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func CustomLogger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()
	if err != nil {
		return err
	}

	duration := time.Since(start)
	path := c.Path()

	fmt.Printf("%s | %d | %v | %s | %s | %s | %s\n", time.Now().Format("15:04:05"), c.Response().StatusCode(), duration, c.IP(), c.Method(), path, "-")

	// Record duration metrics
	httpRequestDuration.WithLabelValues(path).Observe(duration.Seconds())

	// Only process if the response content type is JSON
	if string(c.Response().Header.ContentType()) != fiber.MIMEApplicationJSON {
		//record the count of requests with no idea of the response code
		httpRequestsTotal.WithLabelValues(path,"-").Inc()
	} else {

		var response struct {
			RespCode string `json:"resp_code"`
		}

		if err := json.Unmarshal(c.Response().Body(), &response); err != nil {
			fmt.Println("Error parsing JSON response:", err)
			return err
		}

		// Record the count of requests with the response code
		httpRequestsTotal.WithLabelValues(path, response.RespCode).Inc()
	}

	return err
}
