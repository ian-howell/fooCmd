package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFooCommand() *cobra.Command {
	return &cobra.Command{
		Use: "foo",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("foo")
		},
	}
}
