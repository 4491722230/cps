package pull

import (
	"cps/models"
	"easygf/app"
	"easygf/db"
	"easygf/error_code"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// 邀请明细
func UserInvite(r *ghttp.Request) {
	c := app.NewApp(r)

	userID := r.Get("user_id")

	var user models.User
	var invicatedUser []models.User

	if err := db.DB().Where("id=?").First(&user).Error; err != nil {
		fmt.Println(err)

	}

	if user.ID == 0 || user.Agent == false {
		fmt.Print("")
		c.ReturnJson(error_code.SYSTEM_ERROR, "您不是代理商")
	}

	qs := db.DB()

	if err := db.DB().Where("higher_id = ?", userID).Find(&invicatedUser).Error; err != nil {
		fmt.Println(err)
	}

	qs.Offset(c.GetOffset()).Limit(c.GetLimit()).Find(&invicatedUser)
	qs.Count(&c.Count)

	c.PaginationList("data", invicatedUser)

}
