package pull

import (
	"cps/models"
	"easygf/app"
	"easygf/db"
	"github.com/gogf/gf/net/ghttp"
)

//user_id下级充值明细
func RechargeDetails(r *ghttp.Request) {
	c := app.NewApp(r)

	userID := r.GetInt64("user_id")
	categoryID := r.GetInt("category_id")

	var rebateLogs []models.RebateLog

	qs := db.DB().Where("user_id = ? and status = ? and category = ?", userID, 3, categoryID)

	qs.Offset(c.GetOffset()).Limit(c.GetLimit()).Order("create_time desc").Find(&rebateLogs)
	qs.Count(&c.Count)

	c.PaginationList("data", rebateLogs)
	c.ReturnSuccess(rebateLogs)
}
