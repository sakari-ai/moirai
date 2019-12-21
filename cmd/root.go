package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix    = "CONFIG"
	keyStorage   = "storage"
	keyNamespace = "namespace"
)

// This represents the base command when called without any sub commands
var rootCmd = &cobra.Command{
	Use:   "moirai",
	Short: "Moirai",
	Long:  "Moirai",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)

	keyConfig := "config"

	flags := rootCmd.PersistentFlags()
	flags.String(keyConfig, "config/config.yaml", "config storage")
	flags.String(keyNamespace, "", "config namespace")

	viper.BindPFlag(keyStorage, flags.Lookup(keyConfig))
	viper.BindPFlag(keyNamespace, flags.Lookup(keyNamespace))

	viper.RegisterAlias(keyConfig, keyStorage)
}
