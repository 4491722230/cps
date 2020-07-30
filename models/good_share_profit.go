package models

import "github.com/jinzhu/gorm"

type GoodShareProfit struct {
	gorm.Model
	GoodID int64   //商品id
	Gain1  float64 //上级代理返利比率
	Gain2  float64 //上上级代理返利比率

}
