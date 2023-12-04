package initialize

import (
	"fmt"
	"osstp-go-hive/config"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	_CasbinEnforcer *casbin.Enforcer
)

// 初始化casbin策略管理器
func InitCasbinEnforcer() *casbin.Enforcer {
	casbinE, err := mysqlCasbin()
	if err != nil {
		_Log.Panicf("初始化Casbin失败：%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}

	_CasbinEnforcer = casbinE
	_Log.Info("初始化Casbin完成!")

	return casbinE
}

func mysqlCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(_DB)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(config.Conf.Casbin.ModelPath, a)
	if err != nil {
		return nil, err
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
