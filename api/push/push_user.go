package push

import (
	"cps/models"
	"easygf/app"
	"easygf/db"
	"easygf/error_code"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

//将代理商分销的用户及代理商的信息存储到cps数据库
func PushUser(r *ghttp.Request) {
	c := app.NewApp(r)

	var user models.User
	var agent models.User

	userID := r.GetUint("user_id")
	name := r.GetString("name")
	higherID := r.GetUint("higher_id")
	higherName := r.GetString("higher_name")

	user.ID = userID
	user.Name = name
	user.HigherID = higherID

	agent.ID = higherID
	agent.Name = higherName

	tx := db.DB().Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Save(&user).Error; err != nil {
		fmt.Print(err)
		c.ReturnJson(error_code.SYSTEM_ERROR, user, agent)
	}

	if err := tx.Save(&agent).Error; err != nil {
		fmt.Print(err)
		c.ReturnJson(error_code.SYSTEM_ERROR, user, agent)
	}

	tx.Commit()
	c.ReturnSuccess()

}
