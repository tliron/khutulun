package agent

import (
	"fmt"

	"github.com/tliron/kutil/ard"
	cloutpkg "github.com/tliron/puccini/clout"
	cloututil "github.com/tliron/puccini/clout/util"
)

func (self *Agent) Instantiate(clout *cloutpkg.Clout, coercedClout *cloutpkg.Clout) bool {
	// TODO apply redundancy policies

	count := 1

	for _, vertex := range clout.Vertexes {
		if cloututil.IsTosca(vertex.Metadata, "NodeTemplate") {
			if cloututil.IsToscaType(vertex.Properties, "cloud.puccini.khutulun::Instantiated") {
				name, _ := ard.NewNode(vertex.Properties).Get("name").String()

				for index := 0; index < count; index++ {
					instanceName := fmt.Sprintf("%s-%d", name, index)
					cloututil.Put(
						"instances", cloututil.NewList("cloud.puccini.khutulun::Instance", ard.List{
							cloututil.NewStringMap(ard.StringMap{"name": instanceName}, "string"),
						}),
						vertex.Properties, "attributes")
				}
			}
		}
	}

	return true // changed
}
