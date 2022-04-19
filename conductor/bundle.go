package conductor

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/danjacques/gofslock/fslock"
)

const LOCK_FILE = ".lock"

type BundleIdentifier struct {
	Namespace string `json:"namespace" yaml:"namespace"`
	Type      string `json:"type" yaml:"type"`
	Name      string `json:"name" yaml:"name"`
}

type BundleFile struct {
	Path       string
	Executable bool
}

func (self *Conductor) ListBundles(namespace string, type_ string) ([]BundleIdentifier, error) {
	if namespaces, err := self.namespaceToNamespaces(namespace); err == nil {
		var identifiers []BundleIdentifier
		for _, namespace_ := range namespaces {
			if files, err := os.ReadDir(self.getBundleTypeDir(namespace_, type_)); err == nil {
				for _, file := range files {
					if file.IsDir() {
						identifiers = append(identifiers, BundleIdentifier{
							Namespace: namespace_,
							Type:      type_,
							Name:      file.Name(),
						})
					}
				}
			} else {
				if !os.IsNotExist(err) {
					return nil, err
				}
			}
		}
		return identifiers, nil
	} else {
		return nil, err
	}
}

func (self *Conductor) ListBundleFiles(namespace string, type_ string, name string) ([]BundleFile, error) {
	if lock, err := self.lockBundle(namespace, type_, name, false); err == nil {
		defer func() {
			if err := lock.Unlock(); err != nil {
				log.Errorf("unlock: %s", err.Error())
			}
		}()

		path := self.getBundleDir(namespace, type_, name)
		length := len(path) + 1
		var files []BundleFile
		if err := filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error {
			if !entry.IsDir() {
				if stat, err := os.Stat(path); err == nil {
					files = append(files, BundleFile{
						Path:       path[length:],
						Executable: stat.Mode()&0100 != 0,
					})
				} else {
					return err
				}
			}
			return nil
		}); err == nil {
			return files, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (self *Conductor) ReadBundleFile(namespace string, type_ string, name string, path string) (fslock.Handle, io.ReadCloser, error) {
	if lock, err := self.lockBundle(namespace, type_, name, false); err == nil {
		path = filepath.Join(self.getBundleDir(namespace, type_, name), path)
		log.Debugf("reading from %q", path)
		if file, err := os.Open(path); err == nil {
			return lock, file, nil
		} else {
			if err := lock.Unlock(); err != nil {
				log.Errorf("unlock: %s", err.Error())
			}
			return nil, nil, err
		}
	} else {
		return nil, nil, err
	}
}

func (self *Conductor) DeleteBundle(namespace string, type_ string, name string) error {
	if lock, err := self.lockBundle(namespace, type_, name, false); err == nil {
		defer func() {
			if err := lock.Unlock(); err != nil {
				log.Errorf("unlock: %s", err.Error())
			}
		}()

		path := self.getBundleDir(namespace, type_, name)
		log.Infof("deleting bundle %q", path)
		// TODO: is it OK to delete the lock file while we're holding it?
		return os.RemoveAll(path)
	} else {
		return err
	}
}

func (self *Conductor) getNamespaceDir(namespace string) string {
	if namespace == "" {
		namespace = "_"
	}
	return filepath.Join(self.statePath, namespace)
}

func (self *Conductor) getBundleTypeDir(namespace string, type_ string) string {
	return filepath.Join(self.getNamespaceDir(namespace), type_)
}

func (self *Conductor) getBundleDir(namespace string, type_ string, name string) string {
	return filepath.Join(self.getBundleTypeDir(namespace, type_), name)
}

func (self *Conductor) getBundleMainFile(namespace string, type_ string, name string) string {
	switch type_ {
	case "template", "profile", "clout":
		return filepath.Join(self.getBundleDir(namespace, type_, name), type_+".yaml")
	default:
		return filepath.Join(self.getBundleDir(namespace, type_, name), name)
	}
}

func (self *Conductor) lockBundle(namespace string, type_ string, name string, create bool) (fslock.Handle, error) {
	path := filepath.Join(self.getBundleDir(namespace, type_, name), LOCK_FILE)
	blocker := newBlocker(time.Second, 5)
	if lock, err := fslock.LockSharedBlocking(path, blocker); err == nil {
		return lock, nil
	} else {
		if os.IsNotExist(err) {
			if create {
				// Touch and try again
				if err := touch(path); err == nil {
					return fslock.LockSharedBlocking(path, blocker)
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
}

func touch(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err == nil {
		if file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err == nil {
			return file.Close()
		} else {
			return err
		}
	} else {
		return err
	}
}

func newBlocker(wait time.Duration, attempts int) fslock.Blocker {
	var attempts_ int
	return func() error {
		time.Sleep(wait)
		if attempts <= 0 {
			return nil
		} else {
			attempts_++
			if attempts_ == attempts {
				return fslock.ErrLockHeld
			} else {
				return nil
			}
		}
	}
}
