package ghscaffoldcmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	// ...other imports...
)

// RunCLI runs the ghscaffold CLI
func RunCLI() {
	var owners string
	var file string
	var limitOutput int
	var clear string

	rootCmd := &cobra.Command{
		Use:   "ghscaffold",
		Short: "Select best GitHub repo for your project",
		Run: func(cmd *cobra.Command, args []string) {
			// ...load config, handle clear cache...
			ownerList := strings.Split(strings.ReplaceAll(owners, " ", ""), ",")
			// ...parse ghscaffold.md, fetch repos, match, print results...
		},
	}

	rootCmd.Flags().StringVar(&owners, "owners", "", "Comma separated GitHub owners")
	rootCmd.Flags().StringVar(&file, "file", "ghscaffold.md", "Project details file")
	rootCmd.Flags().IntVar(&limitOutput, "limit_output", 5, "Limit output repos")
	rootCmd.Flags().StringVar(&clear, "clear", "", "Clear cache (e.g. 'repos')")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
