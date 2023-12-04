package controller

import (
	"osstp-go-hive/app/admin/dao"
	"osstp-go-hive/app/admin/model"
	"osstp-go-hive/app/admin/vo"
	"osstp-go-hive/global"
	pkg_response "osstp-go-hive/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IApiController interface {
	GetApis(c *gin.Context)             // 获取接口列表
	GetApiTree(c *gin.Context)          // 获取接口树(按接口Category字段分类)
	CreateApi(c *gin.Context)           // 创建接口
	UpdateApiById(c *gin.Context)       // 更新接口
	BatchDeleteApiByIds(c *gin.Context) // 批量删除接口
}

type ApiController struct {
	ApiDao dao.IApiDao
}

func NewApiController() IApiController {
	ApiDao := dao.NewApiDao()
	apiController := ApiController{ApiDao: ApiDao}
	return apiController
}

// @Title 获取接口列表
// @Description 获取接口列表
// @Param        id   path      int  true  "Account ID"
// @router / [GET]
func (ac ApiController) GetApis(c *gin.Context) {
	var req vo.ApiListRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		pkg_response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		pkg_response.Fail(c, nil, errStr)
		return
	}
	// 获取
	apis, total, err := ac.ApiDao.GetApis(&req)
	if err != nil {
		pkg_response.Fail(c, nil, "获取接口列表失败")
		return
	}
	pkg_response.Success(c, gin.H{
		"apis": apis, "total": total,
	}, "获取接口列表成功")
}

// @Title 获取接口树
// @Description 获取接口树(按接口Category字段分类)
// @Param        id   path      int  true  "Account ID"
// @router /tree [GET]
func (ac ApiController) GetApiTree(c *gin.Context) {
	tree, err := ac.ApiDao.GetApiTree()
	if err != nil {
		pkg_response.Fail(c, nil, "获取接口树失败")
		return
	}
	pkg_response.Success(c, gin.H{
		"apiTree": tree,
	}, "获取接口树成功")
}

// 创建接口
func (ac ApiController) CreateApi(c *gin.Context) {
	var req vo.CreateApiRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		pkg_response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		pkg_response.Fail(c, nil, errStr)
		return
	}

	// 获取当前用户
	ur := dao.NewUserDao()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		pkg_response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	api := model.Api{
		Method:   req.Method,
		Path:     req.Path,
		Category: req.Category,
		Desc:     req.Desc,
		Creator:  ctxUser.Username,
	}

	// 创建接口
	err = ac.ApiDao.CreateApi(&api)
	if err != nil {
		pkg_response.Fail(c, nil, "创建接口失败: "+err.Error())
		return
	}

	pkg_response.Success(c, nil, "创建接口成功")
	return
}

// 更新接口
func (ac ApiController) UpdateApiById(c *gin.Context) {
	var req vo.UpdateApiRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		pkg_response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		pkg_response.Fail(c, nil, errStr)
		return
	}

	// 获取路径中的apiId
	apiId, _ := strconv.Atoi(c.Param("apiId"))
	if apiId <= 0 {
		pkg_response.Fail(c, nil, "接口ID不正确")
		return
	}

	// 获取当前用户
	ur := dao.NewUserDao()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		pkg_response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	api := model.Api{
		Method:   req.Method,
		Path:     req.Path,
		Category: req.Category,
		Desc:     req.Desc,
		Creator:  ctxUser.Username,
	}

	err = ac.ApiDao.UpdateApiById(uint(apiId), &api)
	if err != nil {
		pkg_response.Fail(c, nil, "更新接口失败: "+err.Error())
		return
	}

	pkg_response.Success(c, nil, "更新接口成功")
}

// 批量删除接口
func (ac ApiController) BatchDeleteApiByIds(c *gin.Context) {
	var req vo.DeleteApiRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		pkg_response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		pkg_response.Fail(c, nil, errStr)
		return
	}

	// 删除接口
	err := ac.ApiDao.BatchDeleteApiByIds(req.ApiIds)
	if err != nil {
		pkg_response.Fail(c, nil, "删除接口失败: "+err.Error())
		return
	}

	pkg_response.Success(c, nil, "删除接口成功")
}
