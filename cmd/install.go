/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var programsToInstall = map[string]string{
	"nodejs (nvm)": "curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash",
	"java (sdkman)": "curl -s \"https://get.sdkman.io\" | bash",
	"zsh": "sudo apt install zsh -y",
}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [programas...]",
	Short: "Instala programas de desenvolvimento",
	Long: `Este comando instala uma lista de programas de desenvolvimento comuns.
Exemplos:
foundation-cli install nodejs zsh sdkman`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Iniciando a instalação dos programas ⚙️: ")
		fmt.Println("---")

//		for _, program := range args {
//			installProgram(program)
//		}

		fmt.Println("---")
		color.Green("Instalação concluída com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
