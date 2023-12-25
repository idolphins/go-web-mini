package controller

import (
	"fmt"
	"osstp-go-hive/app/admin/dao"
	"osstp-go-hive/app/admin/model"
	"osstp-go-hive/app/admin/vo"
	"osstp-go-hive/config"
	pkg_response "osstp-go-hive/pkg/response"
	pkg_util "osstp-go-hive/pkg/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IBaseController interface {
	Register(c *gin.Context)
}

type BaseController struct{}

func NewBaseController() IBaseController {
	baseController := BaseController{}
	return baseController
}

// 新用户注册
func (b BaseController) Register(c *gin.Context) {

	var req vo.CreateUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		// 数据绑定失败
		pkg_response.Fail(c, nil, err.Error())
		return
	}
	userDao := dao.NewUserDao()

	oldUser, err := userDao.SearchUser(req.Username, req.Mobile)
	if oldUser != nil {
		// 用户已经存在了
		pkg_response.Fail(c, nil, err.Error())
		return
	}

	// 前端传来的密码是rsa加密的,先解密
	// 密码通过RSA解密
	decodeData, err := pkg_util.RSADecrypt([]byte(req.Password), config.Config.System.RSAPrivateBytes)
	if err != nil {
		pkg_response.Fail(c, nil, err.Error())
		return
	}
	req.Password = string(decodeData)
	if len(req.Password) < 6 {
		pkg_response.Fail(c, nil, "密码长度至少为6位")
		return
	}

	req.Password, err = pkg_util.GenPasswd(req.Password)
	if err != nil {
		pkg_response.Fail(c, nil, "密码格式不准确")
		return
	}

	user := model.User{
		Username:     req.Username,
		Password:     req.Password,
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Nickname:     &req.Nickname,
		Introduction: &req.Introduction,
		Status:       1,
		Creator:      "注册",
		Roles: []*model.Role{
			{
				Model:   gorm.Model{ID: 3},
				Name:    "访客",
				Keyword: "guest",
				Desc:    new(string),
				Sort:    5,
				Status:  1,
				Creator: "注册",
			},
		},
	}

	// 用户不存在 开始创建
	errr := userDao.CreateUser(&user)
	if errr != nil {
		fmt.Println("创建用户失败: " + err.Error())
		pkg_response.Fail(c, nil, "创建用户失败: "+err.Error())
		return
	}
	pkg_response.Success(c, nil, "创建用户成功")

}
