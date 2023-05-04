package configs

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func configEnvCreate() {
	file, err := generates.FilesCreate("./apps/configs/environment.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, configEnvContent())
		fmt.Println("Config Env File Created")
	}
}

func configEnvContent() string {
	var text = `package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWTKEY"); found {
		app.JWTKEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSERNAME"); found {
		app.DBUSERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPASS = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		app.DBPORT = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = false
	}

	err2 := viper.Unmarshal(&app)
	if err2 != nil {
		log.Println("error parse config : ", err2.Error())
		return nil
	}

	if isRead {
		viper.AddConfigPath("./")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.DBUSERNAME = viper.Get("DBUSERNAME").(string)
		app.DBPASS = viper.Get("DBPASS").(string)
		app.DBHOST = viper.Get("DBHOST").(string)
		app.DBPORT, _ = viper.Get("DBPORT").(string)
		app.DBNAME = viper.Get("DBNAME").(string)
		app.JWTKEY = viper.Get("JWTKEY").(string)
	}

	return &app
}
	
	`

	return text
}
