package agent

import (
	"strings"

	delegatepkg "github.com/tliron/khutulun/delegate"
)

func (self *Agent) DeployService(templateNamespace string, templateName string, serviceNamespace string, serviceName string) error {
	var delegate delegatepkg.Delegate
	var client *delegatepkg.DelegatePluginClient
	var err error
	if client, delegate, err = self.GetDelegate(); err == nil {
		defer client.Close()
	} else {
		return err
	}

	if _, problems, err := self.CompileTosca(templateNamespace, templateName, serviceNamespace, serviceName); err == nil {
		self.ProcessService(serviceNamespace, serviceName, delegate, "schedule")
		self.ProcessService(serviceNamespace, serviceName, delegate, "reconcile")
		return nil
	} else {
		if problems != nil {
			return problems.WithError(nil, false)
		} else {
			return err
		}
	}
}

//
// ServiceIdentifier
//

type ServiceIdentifier struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func (self *ServiceIdentifier) Equals(identifier *ServiceIdentifier) bool {
	if self == identifier {
		return true
	} else {
		return (self.Namespace == identifier.Namespace) && (self.Name == identifier.Name)
	}
}

// fmt.Stringer interface
func (self *ServiceIdentifier) String() string {
	return self.Namespace + "," + self.Name
}

//
// ServiceIdentifiers
//

type ServiceIdentifiers struct {
	List []*ServiceIdentifier
}

func NewServiceIdentifiers() *ServiceIdentifiers {
	return new(ServiceIdentifiers)
}

func (self *ServiceIdentifiers) Has(identifier *ServiceIdentifier) bool {
	for _, identifier_ := range self.List {
		if identifier_.Equals(identifier) {
			return true
		}
	}
	return false
}

func (self *ServiceIdentifiers) Add(identifiers ...*ServiceIdentifier) bool {
	var added bool
	for _, identifier := range identifiers {
		if !self.Has(identifier) {
			self.List = append(self.List, identifier)
			added = true
		}
	}
	return added
}

func (self *ServiceIdentifiers) Merge(identifiers *ServiceIdentifiers) bool {
	if identifiers != nil {
		return self.Add(identifiers.List...)
	} else {
		return false
	}
}

// fmt.Stringer interface
func (self *ServiceIdentifiers) String() string {
	var builder strings.Builder
	last := len(self.List) - 1
	for index, identifier := range self.List {
		builder.WriteString(identifier.String())
		if index != last {
			builder.WriteRune(';')
		}
	}
	return builder.String()
}
