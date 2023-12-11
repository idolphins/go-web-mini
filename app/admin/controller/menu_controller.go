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

type IMenuController interface {
	GetMenus(c *gin.Context)             // 获取菜单列表
	GetMenuTree(c *gin.Context)          // 获取菜单树
	CreateMenu(c *gin.Context)           // 创建菜单
	UpdateMenuById(c *gin.Context)       // 更新菜单
	BatchDeleteMenuByIds(c *gin.Context) // 批量删除菜单

	GetUserMenusByUserId(c *gin.Context)    // 获取用户的可访问菜单列表
	GetUserMenuTreeByUserId(c *gin.Context) // 获取用户的可访问菜单树
}

type MenuController struct {
	MenuDao dao.IMenuDao
}

func NewMenuController() IMenuController {
	menuDao := dao.NewMenuDao()
	menuController := MenuController{MenuDao: menuDao}
	return menuController
}

// 获取菜单列表
func (mc MenuController) GetMenus(c *gin.Context) {
	menus, err := mc.MenuDao.GetMenus()
	if err != nil {
		pkg_response.Fail(c, nil, "获取菜单列表失败: "+err.Error())
		return
	}
	pkg_response.Success(c, gin.H{"menus": menus}, "获取菜单列表成功")
}

// 获取菜单树
func (mc MenuController) GetMenuTree(c *gin.Context) {
	menuTree, err := mc.MenuDao.GetMenuTree()
	if err != nil {
		pkg_response.Fail(c, nil, "获取菜单树失败: "+err.Error())
		return
	}
	pkg_response.Success(c, gin.H{"menuTree": menuTree}, "获取菜单树成功")
}

// 创建菜单
func (mc MenuController) CreateMenu(c *gin.Context) {
	var req vo.CreateMenuRequest
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

	menu := model.Menu{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       &req.Icon,
		Path:       req.Path,
		Redirect:   &req.Redirect,
		Component:  req.Component,
		Sort:       req.Sort,
		Status:     req.Status,
		Hidden:     req.Hidden,
		NoCache:    req.NoCache,
		AlwaysShow: req.AlwaysShow,
		Breadcrumb: req.Breadcrumb,
		ActiveMenu: &req.ActiveMenu,
		ParentId:   &req.ParentId,
		Creator:    ctxUser.Username,
	}

	err = mc.MenuDao.CreateMenu(&menu)
	if err != nil {
		pkg_response.Fail(c, nil, "创建菜单失败: "+err.Error())
		return
	}
	pkg_response.Success(c, nil, "创建菜单成功")
}

// 更新菜单
func (mc MenuController) UpdateMenuById(c *gin.Context) {
	var req vo.UpdateMenuRequest
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

	// 获取路径中的menuId
	menuId, _ := strconv.Atoi(c.Param("menuId"))
	if menuId <= 0 {
		pkg_response.Fail(c, nil, "菜单ID不正确")
		return
	}

	// 获取当前用户
	ur := dao.NewUserDao()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		pkg_response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	menu := model.Menu{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       &req.Icon,
		Path:       req.Path,
		Redirect:   &req.Redirect,
		Component:  req.Component,
		Sort:       req.Sort,
		Status:     req.Status,
		Hidden:     req.Hidden,
		NoCache:    req.NoCache,
		AlwaysShow: req.AlwaysShow,
		Breadcrumb: req.Breadcrumb,
		ActiveMenu: &req.ActiveMenu,
		ParentId:   &req.ParentId,
		Creator:    ctxUser.Username,
	}

	err = mc.MenuDao.UpdateMenuById(uint(menuId), &menu)
	if err != nil {
		pkg_response.Fail(c, nil, "更新菜单失败: "+err.Error())
		return
	}

	pkg_response.Success(c, nil, "更新菜单成功")

}

// 批量删除菜单
func (mc MenuController) BatchDeleteMenuByIds(c *gin.Context) {
	var req vo.DeleteMenuRequest
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
	err := mc.MenuDao.BatchDeleteMenuByIds(req.MenuIds)
	if err != nil {
		pkg_response.Fail(c, nil, "删除菜单失败: "+err.Error())
		return
	}

	pkg_response.Success(c, nil, "删除菜单成功")
}

// 根据用户ID获取用户的可访问菜单列表
func (mc MenuController) GetUserMenusByUserId(c *gin.Context) {
	// 获取路径中的userId
	userId, _ := strconv.Atoi(c.Param("userId"))
	if userId <= 0 {
		pkg_response.Fail(c, nil, "用户ID不正确")
		return
	}

	menus, err := mc.MenuDao.GetUserMenusByUserId(uint(userId))
	if err != nil {
		pkg_response.Fail(c, nil, "获取用户的可访问菜单列表失败: "+err.Error())
		return
	}
	pkg_response.Success(c, gin.H{"menus": menus}, "获取用户的可访问菜单列表成功")
}

// 根据用户ID获取用户的可访问菜单树
func (mc MenuController) GetUserMenuTreeByUserId(c *gin.Context) {
	// 获取路径中的userId
	userId, _ := strconv.Atoi(c.Param("userId"))
	if userId <= 0 {
		pkg_response.Fail(c, nil, "用户ID不正确")
		return
	}

	menuTree, err := mc.MenuDao.GetUserMenuTreeByUserId(uint(userId))
	if err != nil {
		pkg_response.Fail(c, nil, "获取用户的可访问菜单树失败: "+err.Error())
		return
	}
	pkg_response.Success(c, gin.H{"menuTree": menuTree}, "获取用户的可访问菜单树成功")
}
