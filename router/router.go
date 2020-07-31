package router

import (
	"cps/api/pull"
	"cps/api/push"
	"github.com/gogf/gf/net/ghttp"
)

func Router(s *ghttp.Server) {
	pullGroup := s.Group("/api/pull")
	pullGroup.GET("/agentMoney", pull.AgentMoney)
	pullGroup.GET("/agentWithdrawalApplication", pull.AgentWithdrawalApplication)
	pullGroup.GET("/rechargeDetails", pull.RechargeDetails)

	pushGroup := s.Group("/api/push")
	pushGroup.POST("/pushGoodShareProfit", push.PushGoodShareProfit)
	pushGroup.POST("/pushOrder", push.PushOrder)
	pushGroup.POST("pushUser", push.PushUser)
}
