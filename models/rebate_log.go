package models

type RebateLog struct {
	UserID      uint    `json:"user_id"`      //获佣用户
	BuyUserID   uint    `json:"buy_user_id"`  //购买人id
	Nickname    string  `json:"nickname"`     //购买人昵称
	OrderSN     int64   `json:"order_sn"`     //订单编号
	OrderID     int64   `json:"order_id"`     //订单ID
	goodID      int     `json:"good_id"`      //商品id
	GoodName    string  `json:"good_name"`    //商品价格
	GoodsPrice  float64 `json:"goods_price"`  //订单商品总额
	CategoryID  int8    `json:"category_id"`  //类型ID
	Money       float64 `json:"money"`        //获佣金额
	Level       int8    `json:"level"`        //获佣用户级别 直接上级为1，间接上级为2
	CreateTime  int64   `json:"create_time"`  //分成记录生成时间
	Confirm     int64   `json:"confirm"`      //确定收货时间
	Status      int8    `json:"status"`       //该订单状态 1充值中 2充值失败 3充值成功
	ConfirmTime int64   `json:"confirm_time"` //确定分成或者取消时间
	Remark      string  `json:"remark"`       //备注
	PriceStatus int8    `json:"pruce_status"` //获拥状态，1为已取出，2为未取出
	//store_id
}
