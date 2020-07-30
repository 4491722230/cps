package push

import (
	"cps/models"
	"easygf/app"
	"easygf/db"
	"easygf/error_code"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

//推送到cms商品的返利比率
func PushGoodShareProfit(r *ghttp.Request) {
	c := app.NewApp(r)

	var goodShareProfit models.GoodShareProfit

	goodID := r.GetInt64("good_id")
	gain1 := r.GetFloat64("gain1")
	gain2 := r.GetFloat64("gain2")

	goodShareProfit.GoodID = goodID
	goodShareProfit.Gain1 = gain1
	goodShareProfit.Gain2 = gain2

	db.DB().Begin()

	if err := db.DB().Save(&goodShareProfit).Error; err != nil {
		fmt.Print(err)
		c.ReturnJson(error_code.SYSTEM_ERROR, goodShareProfit)
	}

	c.ReturnSuccess()

}
