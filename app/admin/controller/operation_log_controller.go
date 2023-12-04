package controller

import (
	"osstp-go-hive/app/admin/dao"
	"osstp-go-hive/app/admin/vo"
	"osstp-go-hive/global"
	pkg_response "osstp-go-hive/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IOperationLogController interface {
	GetOperationLogs(c *gin.Context)             // 获取操作日志列表
	BatchDeleteOperationLogByIds(c *gin.Context) //批量删除操作日志
}

type OperationLogController struct {
	OperationLogDao dao.IOperationLogDao
}

func NewOperationLogController() IOperationLogController {
	OperationLogDao := dao.NewOperationLogDao()
	operationLogController := OperationLogController{OperationLogDao: OperationLogDao}
	return operationLogController
}

// 获取操作日志列表
func (oc OperationLogController) GetOperationLogs(c *gin.Context) {
	var req vo.OperationLogListRequest
	// 绑定参数
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
	logs, total, err := oc.OperationLogDao.GetOperationLogs(&req)
	if err != nil {
		pkg_response.Fail(c, nil, "获取操作日志列表失败: "+err.Error())
		return
	}
	pkg_response.Success(c, gin.H{"logs": logs, "total": total}, "获取操作日志列表成功")
}

// 批量删除操作日志
func (oc OperationLogController) BatchDeleteOperationLogByIds(c *gin.Context) {
	var req vo.DeleteOperationLogRequest
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
	err := oc.OperationLogDao.BatchDeleteOperationLogByIds(req.OperationLogIds)
	if err != nil {
		pkg_response.Fail(c, nil, "删除日志失败: "+err.Error())
		return
	}

	pkg_response.Success(c, nil, "删除日志成功")
}
