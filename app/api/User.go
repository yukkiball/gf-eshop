/**
 * @Author yuki
 * @Date 2021/11/6
 */

package api

import (
	"gf-eshop/app/service"
	"gf-eshop/library/constants"
	"gf-eshop/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(UserApi)

type UserApi struct {
}

func (*UserApi) Login(r *ghttp.Request) {

}

func (*UserApi) Register(r *ghttp.Request) {

}

func (*UserApi) GetOtp(r *ghttp.Request) {

}

func (*UserApi) GetUser(r *ghttp.Request) {
	id := r.GetInt("id")
	userVo, err := service.User.GetUserById(id)
	//fmt.Println(userVo)
	if err != nil {
		g.Log().Errorf("获取用户失败")
	}
	response.JsonExit(r, constants.Success, "获取用户成功", userVo)
}
