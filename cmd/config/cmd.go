package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var shouldPrint bool
var configOutPath string
var shouldSave bool

var RootCmd = &cobra.Command{
	Use:   "config",
	Short: "Thùy chỉnh, hiện thị các config",
	RunE: func(cmd *cobra.Command, args []string)error{
		if configOutPath != "" {
			return writeConfig(configOutPath)
		}

		if shouldSave {
			configPath, err := saveConfig()
			if err != nil {
				return err
			}
			log.Printf("Configs was save to: %s\n", configPath)
			return nil
		}
		return printConfig()
	},
}

func printConfig () error {
	allSettings := viper.AllSettings()
	for key, value := range allSettings {
		fmt.Printf("%s: %v\n", key, value)
	}
	return nil
}

func writeConfig (configPath string) error {
	return viper.WriteConfigAs(configPath)
}

func saveConfig () (string, error){
	if err := viper.WriteConfig(); err != nil {
		return "", err
	}
	return viper.ConfigFileUsed(), nil
}

func init () {
	RootCmd.Flags().BoolVar(&shouldPrint, "print", false, "print current config")
	RootCmd.Flags().BoolVar(&shouldSave, "save", false, "save current config, only work if config file exists")
	RootCmd.Flags().StringVar(&configOutPath, "out", "", "file path to write config")
}