package handler

import "github.com/prometheus/client_golang/prometheus/promhttp"

func MetricsHandler(c *Context) {
	handler := promhttp.Handler()
	handler.ServeHTTP(c.GinContext.Writer, c.GinContext.Request)
}
