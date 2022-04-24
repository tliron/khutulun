package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	profileCommand.AddCommand(profileRegisterCommand)
	profileRegisterCommand.Flags().StringVarP(&unpack, "unpack", "u", "auto", "unpack archive (\"auto\" or \"false\")")
}

var profileRegisterCommand = &cobra.Command{
	Use:   "register [PROFILE NAME] [[CONTENT PATH or URL]]",
	Short: "Register a profile",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		registerPackage(namespace, "profile", args)
	},
}
