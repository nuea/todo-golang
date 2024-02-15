package internal

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nuea/todo-golang/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server struct {
	cfg *config.AppConfig
	gin *gin.Engine
}

func (s *Server) withHealthCheck() *Server {
	s.gin.GET("/hc", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	return s
}

func (s *Server) withMetrics() *Server {
	s.gin.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
	return s
}

func (s *Server) Run() {
	s.gin.Run(":" + s.cfg.System.Port)
}

// ProvideServer
func ProvideServer(cfg *config.AppConfig) *Server {
	ge := gin.New()
	ge.Use(cors.Default())
	ge.Use(otelgin.Middleware(cfg.System.ServiceName))

	sv := &Server{
		cfg: cfg,
		gin: ge,
	}

	return sv.withHealthCheck().withMetrics()
}
