package initialize

import (
	"context"
	model "github.com/richardgong1987/server/model/system"
	"github.com/richardgong1987/server/plugin/plugin-tool/utils"
)

func Dictionary(ctx context.Context) {
	entities := []model.SysDictionary{}
	utils.RegisterDictionaries(entities...)
}
