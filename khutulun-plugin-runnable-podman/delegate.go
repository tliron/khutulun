package main

import cloutpkg "github.com/tliron/puccini/clout"

const servicePrefix = "khutulun"

//
// Delegate
//

type Delegate struct {
	host string
}

// delegate.Delegate interface
func (self *Delegate) ProcessService(namespace string, serviceName string, phase string, clout *cloutpkg.Clout, coercedClout *cloutpkg.Clout) (*cloutpkg.Clout, error) {
	switch phase {
	case "schedule":
		return self.Schedule(namespace, serviceName, clout, coercedClout)

	case "reconcile":
		return self.Reconcile(namespace, serviceName, clout, coercedClout)

	default:
		panic(phase)
	}
}