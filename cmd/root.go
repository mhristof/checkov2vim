package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mhristof/checkov2vim/checkov"
	"github.com/mhristof/checkov2vim/log"
	"github.com/spf13/cobra"
)

var version = "devel"

var rootCmd = &cobra.Command{
	Use:     "checkov2vim",
	Short:   "Convert checkov logs to vim error logs",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		Verbose(cmd)

		fmt.Println(strings.Join(checkov.ToVim(bufio.NewScanner(os.Stdin)), "\n"))
	},
}

// Verbose Increase verbosity
func Verbose(cmd *cobra.Command) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		log.Panic(err)
	}

	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}
func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase verbosity")
	rootCmd.PersistentFlags().BoolP("dryrun", "n", false, "Dry run")
}

// Execute The main function for the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
