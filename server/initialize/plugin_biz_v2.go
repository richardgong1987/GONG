package initialize

import (
	"github.com/gin-gonic/gin"
	_ "github.com/richardgong1987/server/plugin"
	"github.com/richardgong1987/server/utils/plugin/v2"
)

func PluginInitV2(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPluginV2(engine *gin.Engine) {
	PluginInitV2(engine, plugin.Registered()...)
}
