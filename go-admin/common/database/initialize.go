package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"go-admin/common/global"
	log "go-admin/core/logger"
	"go-admin/core/sdk"
	toolsConfig "go-admin/core/sdk/config"
	"go-admin/core/sdk/pkg"
	mycasbin "go-admin/core/sdk/pkg/casbin"
	toolsDB "go-admin/core/tools/database"
	. "go-admin/core/tools/gorm/logger"
)

// Setup 配置数据库
func Setup() {
	for k := range toolsConfig.DatabasesConfig {
		setupSimpleDatabase(k, toolsConfig.DatabasesConfig[k])
	}
}

func setupSimpleDatabase(host string, c *toolsConfig.Database) {
	if global.Driver == "" {
		global.Driver = c.Driver
	}
	log.Infof("%s => %s", host, pkg.Green(c.Source))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		log.Fatal(pkg.Red(c.Driver+" connect error: "), err)
	} else {
		log.Info(pkg.Green(c.Driver + " connect success!"))
	}

	e := mycasbin.Setup(db, "")
	sdk.Runtime.SetDb(host, db)
	sdk.Runtime.SetCasbin(host, e)
}
