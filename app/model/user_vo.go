/**
 * @Author yuki
 * @Date 2021/11/6
 */

package model

type UserVo struct {
	Id              int
	Name            string
	Gender          int
	Age             int
	Telphone       string
	RegisterMode    string
	ThirdPartyId    string
	EncryptPassword string
}