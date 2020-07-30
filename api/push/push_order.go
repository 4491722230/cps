package push

import (
	"cps/models"
	"easygf/app"
	"easygf/db"
	"easygf/error_code"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

//推送被代理邀请来的用户的订单
func PushOrder(r *ghttp.Request) {
	c := app.NewApp(r)

	var rebatelog models.RebateLog
	var higherRebatelog models.RebateLog
	var goodShareProfit models.GoodShareProfit
	var user models.User
	var higherUser models.User

	//userID := r.GetInt64("user_id")
	buyUserID := r.GetInt64("buy_user_id")
	nickname := r.GetString("nickname")
	orderSN := r.GetInt64("order_sn")
	orderID := r.GetInt64("order_id")
	goodID := r.GetInt64("good_id")
	goodName := r.GetString("goodName")
	categoryID := r.GetInt("category_id")
	confirm := r.GetInt64("confirm")
	status := r.GetInt("status")
	goodsPrice := r.GetFloat64("goodsPrice")
	remark := r.GetString("remark")

	if status != 3 {
		fmt.Println("订单还未付款，请付款之后再推送到cms上")
		c.ReturnJson(error_code.SYSTEM_ERROR, "还未支付")
		return
	}

	if err := db.DB().Where("good_id = ?", goodID).First(&goodShareProfit).Error; err != nil {
		fmt.Println("查询good代理分成失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "查询good代理分成失败，请先推送改商品的代理分成")
		return
	}

	if err := db.DB().Where("user_id = ?", buyUserID).First(&user).Error; err != nil {
		fmt.Println("查询用户失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "查询用户失败")
		return
	}

	if err := db.DB().Where("user_id = ?", user.HigherID).First(&higherUser).Error; err != nil {
		fmt.Println("查询用户失败")
		c.ReturnJson(error_code.SYSTEM_ERROR, "查询用户失败")
		return
	}

	rebatelog.Level = 2
	rebatelog.Money = goodsPrice * goodShareProfit.Gain1
	rebatelog.ConfirmTime = time.Now().AddDate(0, 0, 5).Unix()
	rebatelog.UserID = user.HigherID
	rebatelog.Remark = remark
	rebatelog.OrderSN = orderSN
	rebatelog.CreateTime = time.Now().Unix()
	rebatelog.GoodName = goodName
	rebatelog.GoodsPrice = goodsPrice
	rebatelog.BuyUserID = buyUserID
	rebatelog.CategoryID = int8(categoryID)
	rebatelog.Confirm = confirm
	rebatelog.OrderID = orderID
	rebatelog.Nickname = nickname

	higherRebatelog.Level = 3
	higherRebatelog.Money = goodsPrice * goodShareProfit.Gain2
	higherRebatelog.ConfirmTime = time.Now().AddDate(0, 0, 5).Unix()
	higherRebatelog.UserID = higherUser.HigherID
	higherRebatelog.Remark = remark
	higherRebatelog.OrderSN = orderSN
	higherRebatelog.CreateTime = time.Now().Unix()
	higherRebatelog.GoodName = goodName
	higherRebatelog.GoodsPrice = goodsPrice
	higherRebatelog.BuyUserID = buyUserID
	higherRebatelog.CategoryID = int8(categoryID)
	higherRebatelog.Confirm = confirm
	higherRebatelog.OrderID = orderID
	higherRebatelog.Nickname = nickname

	if err := db.DB().Save(&rebatelog).Error; err != nil {
		fmt.Print(err)
		c.ReturnJson(error_code.SYSTEM_ERROR, rebatelog)
	}

	if err := db.DB().Save(&higherRebatelog).Error; err != nil {
		fmt.Print(err)
		c.ReturnJson(error_code.SYSTEM_ERROR, rebatelog)
	}

	c.ReturnSuccess()

}
