package api

import (
	"cps/models"
	"easygf/app"
	"easygf/db"
	"easygf/error_code"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

//统计代理余额
func AgentMoney(r *ghttp.Request) {

	c := app.NewApp(r)

	var sumMoney float64
	var user models.User

	userID := r.GetInt("user_id")

	unix := time.Now().Unix()

	if err := db.DB().
		Raw("select count(money) from rebate_log where user_id = ? and status = ? and confirm_time < ?", userID, 3, unix).
		Scan(&sumMoney).Error; err != nil {
		fmt.Println("计算代理商余额错误")
		c.ReturnJson(error_code.SYSTEM_ERROR, "统计代理商返利总额失败")
	}

	if e := db.DB().Where("user_id = ?", userID).First(&user).Error; e != nil {
		fmt.Println("查询用户失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "查询用户失败")
	}

	user.Money = user.Money + sumMoney

	if e := db.DB().Save(&user).Error; e != nil {
		fmt.Println("修改用户余额失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "修改代理商余额失败")
	}

	c.ReturnSuccess("agentMoney", user.Money)

}
