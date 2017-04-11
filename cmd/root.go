package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(changebgr)
}

//RootCmd wht
var RootCmd = &cobra.Command{
	Use:   "SetPaper",
	Short: "setPaper is changeing background",
	Long:  `A fast background changer`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}
