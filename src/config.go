package src

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	ViewRoute string
	EditRoute string
	SaveRoute string

	viewRouteConfig = "routes.ViewRoute"
	editRouteConfig = "routes.EditRoute"
	saveRouteConfig = "routes.SaveRoute"
)

func LoadConfig() {
	setDefaults()

	viper.SetConfigName("conf")
	viper.AddConfigPath("./src/")
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
}

func setConfigVars() {
	ViewRoute = fmt.Sprint(viper.Get(viewRouteConfig))
	EditRoute = fmt.Sprint(viper.Get(editRouteConfig))
	SaveRoute = fmt.Sprint(viper.Get(saveRouteConfig))
}
