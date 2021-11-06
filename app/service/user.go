/**
 * @Author yuki
 * @Date 2021/11/6
 */

package service

import (
	"gf-eshop/app/dao"
	"gf-eshop/app/model"
	"github.com/jinzhu/copier"
)

var User = new(userService)

type userService struct {

}

func (*userService) GetUserById(id int) (userVo *model.UserVo, err error) {
	userInfo, err := dao.UserInfo.GetUserById(id)
	userVo = &model.UserVo{}
	err = copier.Copy(userVo, userInfo)
	if err != nil {
		return nil, err
	}
	userPassword, err := dao.UserPassword.GetPasswordById(id)
	userVo.EncryptPassword = userPassword.EncrptPassword
	return
}
