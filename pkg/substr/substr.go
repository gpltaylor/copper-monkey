package substr

import (
	"fmt"
	"github.com/spf13/cobra"
)

var sidx, eidx int

func NewCmdSubstr() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "substr",
		Short:   "Obtains a substring of a given string",
		Aliases: []string{"substr"},
		//Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sstr := args[0][sidx:eidx]
			fmt.Println(sstr)
		},
	}

	cmd.Flags().IntVarP(&sidx, "startindex", "s", 0, "Start Index")
	cmd.Flags().IntVarP(&eidx, "endindex", "e", -1, "End Index")
	return cmd
}
