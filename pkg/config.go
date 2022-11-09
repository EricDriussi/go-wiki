package pkg

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	ViewRoute     string
	EditRoute     string
	SaveRoute     string
	WikiPagesPath string

	viewRouteConfig     = "routes.ViewRoute"
	editRouteConfig     = "routes.EditRoute"
	saveRouteConfig     = "routes.SaveRoute"
	wikiPagesPathConfig = "resources.WikiPagesPath"
)

func LoadConfig() {
	setDefaults()

	viper.SetConfigName("conf")
	viper.AddConfigPath("./pkg/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	setConfigVars()

	viper.OnConfigChange(func(e fsnotify.Event) {
		if e.Op.String() == "WRITE" {
			fmt.Println("Config file updated, reloading data...")
			setConfigVars()
		}
	})
	viper.WatchConfig()
}

func setDefaults() {
	viper.SetDefault(viewRouteConfig, "/wiki/view/")
	viper.SetDefault(editRouteConfig, "/wiki/edit/")
	viper.SetDefault(saveRouteConfig, "/wiki/save/")
	viper.SetDefault(wikiPagesPathConfig, "assets/")
}

func setConfigVars() {
	ViewRoute = viper.GetString(viewRouteConfig)
	EditRoute = viper.GetString(editRouteConfig)
	SaveRoute = viper.GetString(saveRouteConfig)
	WikiPagesPath = viper.GetString(wikiPagesPathConfig)
}
