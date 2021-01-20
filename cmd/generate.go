package cmd

import (
	"github.com/mhristof/checkov2vim/ale"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate the linter vimscript file required",
		Run: func(cmd *cobra.Command, args []string) {
			path, err := cmd.Flags().GetString("dest")
			if err != nil {
				panic(err)
			}

			dir, err := homedir.Expand(path)
			if err != nil {
				panic(err)
			}

			ale.GenerateCheckov(dir)
		},
	}
)

func init() {
	generateCmd.Flags().StringP("dest", "d", "~/.vim/bundle/ale/ale_linters/terraform/checkov.vim", "Destination of the vim script")

	rootCmd.AddCommand(generateCmd)
}
