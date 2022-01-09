package cmd

import (
	"github.com/spf13/cobra"

	"github.com/myhro/feeds/copasa"
	"github.com/myhro/feeds/generator"
)

func Copasa(cmd *cobra.Command, args []string) {
	generator.Print(copasa.Command, copasa.XML)
}
