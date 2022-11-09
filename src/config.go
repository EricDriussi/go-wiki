package src

// TODO.refactor

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	ViewRoute string
	EditRoute string
	SaveRoute string
)

func LoadConfig() {
	viper.SetDefault("routes.EditRoute", "/wiki/edit/")
	viper.SetDefault("routes.SaveRoute", "/wiki/save/")
	viper.SetDefault("routes.ViewRoute", "/wiki/view/")

	viper.SetConfigName("conf")
	viper.AddConfigPath("./src/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	ViewRoute = fmt.Sprint(viper.Get("routes.ViewRoute"))
	EditRoute = fmt.Sprint(viper.Get("routes.EditRoute"))
	SaveRoute = fmt.Sprint(viper.Get("routes.SaveRoute"))

	viper.OnConfigChange(func(e fsnotify.Event) {
		if e.Op.String() == "WRITE" {
			fmt.Println("Config file updated, reloading data...")
			ViewRoute = fmt.Sprint(viper.Get("routes.ViewRoute"))
			EditRoute = fmt.Sprint(viper.Get("routes.EditRoute"))
			SaveRoute = fmt.Sprint(viper.Get("routes.SaveRoute"))
		}
	})
	viper.WatchConfig()
}
