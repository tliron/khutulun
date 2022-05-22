package agent

import (
	"io"
	"os"
	"strings"

	"github.com/danjacques/gofslock/fslock"
	"github.com/tliron/kutil/format"
	"github.com/tliron/kutil/logging"
	problemspkg "github.com/tliron/kutil/problems"
	cloutpkg "github.com/tliron/puccini/clout"
	"github.com/tliron/puccini/clout/js"
)

func (self *Agent) OpenClout(namespace string, serviceName string) (fslock.Handle, *cloutpkg.Clout, error) {
	if lock, err := self.lockPackage(namespace, "clout", serviceName, false); err == nil {
		cloutPath := self.getPackageMainFile(namespace, "clout", serviceName)
		log.Debugf("reading clout: %q", cloutPath)
		if clout, err := cloutpkg.Load(cloutPath, "yaml"); err == nil {
			return lock, clout, nil
		} else {
			logging.CallAndLogError(lock.Unlock, "unlock", log)
			return nil, nil, err
		}
	} else {
		return nil, nil, err
	}
}

func (self *Agent) SaveClout(serviceNamespace string, serviceName string, clout *cloutpkg.Clout) error {
	cloutPath := self.getPackageMainFile(serviceNamespace, "clout", serviceName)
	log.Infof("writing to %q", cloutPath)
	if file, err := os.OpenFile(cloutPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err == nil {
		defer logging.CallAndLogError(file.Close, "file close", log)
		return format.WriteYAML(clout, file, "  ", false)
	} else {
		return err
	}
}

func (self *Agent) CoerceClout(clout *cloutpkg.Clout) error {
	problems := problemspkg.NewProblems(nil)
	js.Coerce(clout, problems, self.urlContext, true, "yaml", true, false, false)
	return problems.ToError(true)
}

func (self *Agent) OpenFile(path string, coerceClout bool) (io.ReadCloser, error) {
	if coerceClout {
		if file, err := os.Open(path); err == nil {
			defer logging.CallAndLogError(file.Close, "file close", log)
			if clout, err := cloutpkg.Read(file, "yaml"); err == nil {
				if err := self.CoerceClout(clout); err == nil {
					if clout_, err := format.EncodeYAML(clout, "  ", false); err == nil {
						return io.NopCloser(strings.NewReader(clout_)), nil
					} else {
						return nil, err
					}
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return os.Open(path)
	}
}
