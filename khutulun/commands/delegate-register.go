package commands

import (
	contextpkg "context"

	"github.com/spf13/cobra"
)

func init() {
	delegateCommand.AddCommand(delegateRegisterCommand)
	delegateRegisterCommand.Flags().StringVarP(&unpack, "unpack", "u", "auto", "unpack archive (\"tar\", \"tgz\", \"zip\", \"auto\" or \"false\")")
}

var delegateRegisterCommand = &cobra.Command{
	Use:   "register [DELEGATE NAME] [[CONTENT PATH or URL]]",
	Short: "Register a delegate",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		registerPackage(contextpkg.TODO(), namespace, "delegate", getPluginArgs(args))
	},
}
