package cmd

import (
	"fmt"

	"x-ui/config"
	"x-ui/internal/database"
	"x-ui/internal/web/service"

	"github.com/spf13/cobra"
)

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Manage SSL certificates",
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing
	},
}

func init() {
	rootCmd.AddCommand(certCmd)

	certCmd.Flags().StringVar(&webCertFile, "cert", "", "Path to public key file")
	certCmd.Flags().StringVar(&webKeyFile, "key", "", "Path to private key file")
	certCmd.Flags().BoolVar(&reset, "reset", false, "Remove certificate")

	certCmd.Run = func(cmd *cobra.Command, args []string) {
		if reset {
			UpdateCert("", "")
		} else {
			UpdateCert(webCertFile, webKeyFile)
		}
	}
}

func UpdateCert(publicKey string, privateKey string) {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	if (privateKey != "" && publicKey != "") || (privateKey == "" && publicKey == "") {
		settingService := service.SettingService{}
		err = settingService.SetCertFile(publicKey)
		if err != nil {
			fmt.Println("set certificate public key failed:", err)
		} else {
			fmt.Println("set certificate public key success")
		}

		err = settingService.SetKeyFile(privateKey)
		if err != nil {
			fmt.Println("set certificate private key failed:", err)
		} else {
			fmt.Println("set certificate private key success")
		}
	} else {
		fmt.Println("both public and private key should be entered.")
	}
}
