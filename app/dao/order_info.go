// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-eshop/app/dao/internal"
)

// orderInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type orderInfoDao struct {
	*internal.OrderInfoDao
}

var (
	// OrderInfo is globally public accessible object for table order_info operations.
	OrderInfo orderInfoDao
)

func init() {
	OrderInfo = orderInfoDao{
		internal.NewOrderInfoDao(),
	}
}

// Fill with you ideas below.