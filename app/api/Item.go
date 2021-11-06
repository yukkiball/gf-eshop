/**
 * @Author yuki
 * @Date 2021/11/6
 */

package api

import (
	"gf-eshop/library/constants"
	"gf-eshop/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Item = new(ItemApi)

type ItemApi struct{}

func (*ItemApi) CreateItem(r *ghttp.Request) {

}

func (*ItemApi) GetItem(r *ghttp.Request) {
	response.JsonExit(r, constants.Success, "test")
}

func (*ItemApi) ListItem(r *ghttp.Request) {

}

func (*ItemApi) PublishPromo(r *ghttp.Request) {

}
