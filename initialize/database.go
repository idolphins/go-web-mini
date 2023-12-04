package initialize

import (
	"fmt"
	"osstp-go-hive/app/admin/model"
	"osstp-go-hive/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_DB *gorm.DB
)

// 初始化mysql数据库
func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.Database,
		config.Config.Mysql.Charset,
		config.Config.Mysql.Collation,
		config.Config.Mysql.Query,
	)
	// 隐藏密码
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Config.Mysql.Username,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.Database,
		config.Config.Mysql.Charset,
		config.Config.Mysql.Collation,
		config.Config.Mysql.Query,
	)
	//Log.Info("数据库连接DSN: ", showDsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		//// 指定表前缀
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: config.Config.Mysql.TablePrefix + "_",
		//},
	})
	if err != nil {
		_Log.Panicf("初始化mysql数据库异常: %v", err)
		panic(fmt.Errorf("初始化mysql数据库异常: %v", err))
	}

	// 开启mysql日志
	if config.Config.Mysql.LogMode {
		db.Debug()
	}
	// 全局DB赋值
	_DB = db
	// 自动迁移表结构
	dbAutoMigrate()
	_Log.Infof("初始化mysql数据库完成! dsn: %s", showDsn)

	return db
}

// 自动迁移表结构
func dbAutoMigrate() {
	_DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Menu{},
		&model.Api{},
		&model.OperationLog{},
	)
}
