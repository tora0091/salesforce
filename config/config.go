package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config is confing data structure
type Config struct {
	Interface string
	Version   string
	Username  string
	Password  string
	Login     struct {
		URL string
		XML string
	}
	URL struct {
		Base     string
		Describe string
		Job      string
		Close    string
		Batch    string
		Result   string
	}
	XML struct {
		Insert string
		Update string
		Delete string
		Close  string
	}
}

// C is global variable
var C Config

// NewConfig return config object
func NewConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatal(err)
	}

	C.Login.URL = replaceString(C.Login.URL)
	C.URL.Base = replaceString(C.URL.Base)
	C.URL.Describe = replaceString(C.URL.Describe)
	C.URL.Job = replaceString(C.URL.Job)
	C.URL.Close = replaceString(C.URL.Close)
	C.URL.Batch = replaceString(C.URL.Batch)
	C.URL.Result = replaceString(C.URL.Result)
}

func replaceString(oldString string) string {
	var newString string
	newString = strings.Replace(oldString, "[salesforce_interface]", C.Interface, 1)
	return strings.Replace(newString, "[salesforce_version]", C.Version, 1)
}
