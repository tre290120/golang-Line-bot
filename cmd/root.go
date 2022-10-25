package cmd

import (
	"fmt"
	"os"
	"simplebot/router"
	"simplebot/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "start line bot server",
	Long:  "command to set port, config path of line bot server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		configPath, _ := cmd.Flags().GetString("configPath")
		service.InitBot(configPath)
		server := router.GetRouter()
		fmt.Printf("serve at " + port)
		server.Run(":" + port)

	},
}

func init() {
	rootCmd.Flags().StringP("configPath", "c", "", "line bot config path")
	rootCmd.Flags().StringP("port", "p", "3000", "server port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
