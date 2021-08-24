package config

var APP_NAME string

const APP_VERSION = "v1.0.12"

func LoadVersion() {
	APP_NAME = AppConfigure.Name + "-" + AppConfigure.Env
}
