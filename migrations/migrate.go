package main

import (
	"deall/cmd/config"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migration",
	Short: "add migration",
	Long:  `add migration to database`,
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up migration",
	Long:  "up migration to database",
	Run: func(cmd *cobra.Command, args []string) {
		migrate("up")
	},
}
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "down migration",
	Long:  "down migration to database",
	Run: func(cmd *cobra.Command, args []string) {
		migrate("down")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(upCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func migrate(command string) {
	var cmd *exec.Cmd
	var config = config.LoadConfig()
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.DBConfig.Username, config.DBConfig.Password, config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.Name)

	if command == "down" {
		cmd = exec.Command("migrate", "-path", "migrations", "-database", fmt.Sprintf("mysql://%s", dsn), "-verbose", command, "-all")
	} else {
		cmd = exec.Command("migrate", "-path", "migrations", "-database", fmt.Sprintf("mysql://%s", dsn), "-verbose", command)
	}

	err := cmd.Run()
	if err != nil {
		logrus.Error(err)
	}
}
