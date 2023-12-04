package global

import (
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	ZLog           *zap.SugaredLogger
	Validate       *validator.Validate
	Trans          ut.Translator
	CasbinEnforcer *casbin.Enforcer
)
