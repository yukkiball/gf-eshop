// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// Item is the golang structure for table item.
type Item struct {
	Id          int     `orm:"id,primary"  json:"id"`          //
	Title       string  `orm:"title"       json:"title"`       //
	Price       float64 `orm:"price"       json:"price"`       //
	Description string  `orm:"description" json:"description"` //
	Sales       int     `orm:"sales"       json:"sales"`       //
	ImgUrl      string  `orm:"img_url"     json:"imgUrl"`      //
}

// ItemStock is the golang structure for table item_stock.
type ItemStock struct {
	Id     int `orm:"id,primary" json:"id"`     //
	Stock  int `orm:"stock"      json:"stock"`  //
	ItemId int `orm:"item_id"    json:"itemId"` //
}

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	Id         string  `orm:"id,primary"  json:"id"`         //
	UserId     int     `orm:"user_id"     json:"userId"`     //
	ItemId     int     `orm:"item_id"     json:"itemId"`     //
	ItemPrice  float64 `orm:"item_price"  json:"itemPrice"`  //
	Amount     int     `orm:"amount"      json:"amount"`     //
	OrderPrice float64 `orm:"order_price" json:"orderPrice"` //
	PromoId    int     `orm:"promo_id"    json:"promoId"`    //
}

// Promo is the golang structure for table promo.
type Promo struct {
	Id             int         `orm:"id,primary"       json:"id"`             //
	PromoName      string      `orm:"promo_name"       json:"promoName"`      //
	StartDate      *gtime.Time `orm:"start_date"       json:"startDate"`      //
	ItemId         int         `orm:"item_id"          json:"itemId"`         //
	PromoItemPrice float64     `orm:"promo_item_price" json:"promoItemPrice"` //
	EndDate        *gtime.Time `orm:"end_date"         json:"endDate"`        //
}

// SequenceInfo is the golang structure for table sequence_info.
type SequenceInfo struct {
	Name         string `orm:"name,primary"  json:"name"`         //
	CurrentValue int    `orm:"current_value" json:"currentValue"` //
	Step         int    `orm:"step"          json:"step"`         //
}

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id           int    `orm:"id,primary"      json:"id"`           //
	Name         string `orm:"name"            json:"name"`         //
	Gender       int    `orm:"gender"          json:"gender"`       // //1???????????????2????????????
	Age          int    `orm:"age"             json:"age"`          //
	Telphone     string `orm:"telphone,unique" json:"telphone"`     //
	RegisterMode string `orm:"register_mode"   json:"registerMode"` // //byphone, bywechat, by alipay
	ThirdPartyId string `orm:"third_party_id"  json:"thirdPartyId"` //
}

// UserPassword is the golang structure for table user_password.
type UserPassword struct {
	Id             int    `orm:"id,primary"      json:"id"`             //
	EncrptPassword string `orm:"encrpt_password" json:"encrptPassword"` //
	UserId         int    `orm:"user_id"         json:"userId"`         //
}
