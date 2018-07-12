package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetInt(key string) int {
	env := viper.GetString("ENV")
	if len(env) == 0 {
		env = "development"
	}
	result := viper.GetInt(fmt.Sprintf("%s.%s", env, key))
	fmt.Println(env, key, result)
	return result
}
func GetString(key string) string {
	env := viper.GetString("ENV")
	if len(env) == 0 {
		env = "development"
	}
	result := viper.GetString(fmt.Sprintf("%s.%s", env, key))
	fmt.Println(env, key, result)
	return result
}
func GetStringSlice(key string) []string {
	env := viper.GetString("ENV")
	if len(env) == 0 {
		env = "development"
	}
	result := viper.GetStringSlice(fmt.Sprintf("%s.%s", env, key))
	fmt.Println(env, key, result)
	return result
}
