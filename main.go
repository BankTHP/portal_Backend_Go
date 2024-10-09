package main

import (
    "fmt"
    "log"
    "github.com/spf13/viper"
)

func initConfig() {
    env := viper.GetString("GO_ENV")
    if env == "" {
        env = "dev"
    }

    viper.SetConfigName(fmt.Sprintf("config.%s", env))
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }
}

func main() {
    initConfig()

    fmt.Println("App Port:", viper.GetString("app.port"))
    fmt.Println("Debug Mode:", viper.GetBool("app.debug"))
}
