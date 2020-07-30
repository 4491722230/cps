package router

import (
	"github.com/gogf/gf/net/ghttp"
)

func Router(s *ghttp.Server) {
	pull := s.Group("/api/pull")
	pull.GET("")
	push := s.Group("/api/push")
}
