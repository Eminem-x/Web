package api

import (
	"github.com/spf13/cobra"
	"go-admin/common/database"
	"go-admin/common/storage"
	"go-admin/core/config/source/file"

	ext "go-admin/config"
	"go-admin/core/sdk/config"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

func setup() {
	// 注入配置拓展项
	config.ExtendConfig = &ext.ExtConfig
	// 1. 读取配置
	config.Setup(
		file.NewSource(file.WithPath(configYml)),
		database.Setup,
		storage.Setup,
	)
}

func run() error {
	return nil
}
