/*
Copyright © 2024 leig HERE <leigme@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crew",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			log.Println("args is nil")
			return
		case 1:
			result := make([]string, 0)
			if strings.EqualFold(args[0], "all") {
				for _, key := range viper.AllKeys() {
					result = append(result, viper.GetString(key))
				}
			} else {
				for _, key := range viper.AllKeys() {
					if strings.HasSuffix(key, args[0]) {
						result = append(result, viper.GetString(key))
					}
				}
			}
			fmt.Printf("%s", strings.Join(result, ","))
		default:
			fmt.Printf("%s", viper.GetString(strings.Join(args, ".")))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("crew.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath())
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
