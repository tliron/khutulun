package commands

import (
	"os"

	"github.com/spf13/cobra"
	clientpkg "github.com/tliron/khutulun/client"
	"github.com/tliron/kutil/util"
)

func init() {
	hostCommand.AddCommand(hostListCommand)
}

var hostListCommand = &cobra.Command{
	Use:   "list",
	Short: "List hosts in a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := clientpkg.NewClientFromConfiguration(configurationPath, clusterName)
		util.FailOnError(err)
		util.OnExitError(client.Close)

		hosts, err := client.ListHosts()
		util.FailOnError(err)
		if len(hosts) > 0 {
			err = Transcriber().Print(hosts, os.Stdout, format)
			util.FailOnError(err)
		}
	},
}
