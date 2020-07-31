package pull

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
	var rebateLogs []models.RebateLog

	userID := r.GetInt("user_id")

	unix := time.Now().Unix()

	if err := db.DB().
		Raw("select count(money) from rebate_log where user_id = ? and status = ? and confirm_time < and price_status = ?", userID, 3, unix, 1).
		Scan(&sumMoney).Error; err != nil {
		fmt.Println("计算代理商余额错误")
		c.ReturnJson(error_code.SYSTEM_ERROR, "统计代理商返利总额失败")
		return
	}

	if e := db.DB().Where("user_id = ?", userID).First(&user).Error; e != nil {
		fmt.Println("查询用户失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "查询用户失败")
	}

	user.Money = user.Money + sumMoney

	tx := db.DB().Begin()
	defer tx.RollbackUnlessCommitted()

	if e := tx.Save(&user).Error; e != nil {
		fmt.Println("修改用户余额失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "修改代理商余额失败")
		return
	}

	if e := db.DB().First("user_id = ?").Find(rebateLogs).Error; e != nil {
		fmt.Println("查询user_id下的用户失败" + e.Error())
		c.ReturnJson(error_code.SYSTEM_ERROR, "查询user_id下的用户失败")
		return

	}

	for _, rebateLog := range rebateLogs {
		rebateLog.PriceStatus = 2
		if err := db.DB().Save(&rebateLog).Error; err != nil {
			fmt.Println("保存rebateLog失败" + err.Error())
			c.ReturnJson(error_code.SYSTEM_ERROR, "保存rebateLog失败")
			return
		}
	}

	tx.Commit()

	c.ReturnSuccess("agentMoney", user.Money)

}
