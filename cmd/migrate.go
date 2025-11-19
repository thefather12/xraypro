package cmd

import (
	"fmt"
	"log"

	"x-ui/config"
	"x-ui/internal/database"
	"x-ui/internal/web/service"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database from other/old x-ui",
	Run: func(cmd *cobra.Command, args []string) {
		MigrateDb()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func MigrateDb() {
	inboundService := service.InboundService{}

	err := database.InitDB(config.GetDBPath())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start migrating database...")
	inboundService.MigrateDB()
	fmt.Println("Migration done!")
}
