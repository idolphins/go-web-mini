package controller

import (
	"fmt"
	"osstp-go-hive/app/admin/dao"
	"osstp-go-hive/app/admin/model"
	"osstp-go-hive/app/admin/vo"
	pkg_response "osstp-go-hive/pkg/response"

	"github.com/gin-gonic/gin"
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

	user := model.User{
		Username:     req.Username,
		Password:     req.Password,
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Nickname:     &req.Nickname,
		Introduction: &req.Introduction,
		Status:       1,
		Creator:      "系统",
	}
	userDao := dao.NewUserDao()

	userDao.SearchUser(c, user.Username, user.Mobile)

	err := userDao.CreateUser(&user)
	if err != nil {
		fmt.Println("创建用户失败: " + err.Error())
		pkg_response.Fail(c, nil, "创建用户失败: "+err.Error())
		return
	}
	pkg_response.Success(c, nil, "创建用户成功")

}
