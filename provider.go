package core

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

type BasePlan interface {
	Apply()
}

type BaseVM interface {
	GetName() string
	Description(string) BaseVM
	Files(func()) BaseVM
	Units(func()) BaseVM
	Passwd(func()) BaseVM
	// Networks(func()) BaseVM
	// Disks(func()) BaseVM

	// Requires(...interface{}) BaseVM
	// Wants(...interface{}) BaseVM
	// After(...interface{}) BaseVM

	// TODO change to Notify!
	// OnCreated(f func(vm BaseVM)) BaseVM
	// OnUpdated(f func(vm BaseVM)) BaseVM
	// OnDestroyed(f func(vm BaseVM)) BaseVM
}

type BaseFile interface {
	// TODO Root(n string) BaseFile
	GetPath() string
	Mode(m uint32) BaseFile // FileMode?
	Uid(uid int) BaseFile
	Gid(gid int) BaseFile
	Contents(v interface{}) BaseFile

	// TODO move into a different interface?
	Filter(fn ValueFilter) BaseFile
}

type BaseURLFile interface {
	BaseFile
	Compression(string) BaseURLFile
	Verification(string, string) BaseURLFile
}

type BaseUnit interface {
	GetName() string
	Enable(b bool) BaseUnit
	Mask(b bool) BaseUnit
	Contents(c interface{}) BaseUnit

	// TODO move into a different interface?
	Filter(fn ValueFilter) BaseUnit

	DropIns(fn func()) BaseUnit
}

type BaseUnitDropIn interface {
	GetName() string
	Contents(c interface{}) BaseUnitDropIn

	// TODO move into a different interface?
	Filter(fn ValueFilter) BaseUnitDropIn
}

type BaseUser interface {
	GetName() string
	PasswordHash(string) BaseUser
	SSHAuthorizedKeys(keys ...string) BaseUser
	Uid(uint) BaseUser
	Homedir(string) BaseUser
	NoCreateHome(bool) BaseUser
	PrimaryGroup(string) BaseUser
	Groups(...string) BaseUser
	NoUserGroup(bool) BaseUser
	System(bool) BaseUser
	NoLogInit(bool) BaseUser
	Shell(string) BaseUser
}
type Value interface {
	String() string
}

type StringValue string

func (s StringValue) String() string {
	return string(s)
}

type URLValue struct {
	URL                      url.URL
	Compression              string
	VerificationHashFunction string
	VerificationHashSum      string
}

func (u URLValue) String() string {
	return "" // Ignore...
}

type ValueFilter func(Value) Value

func File(path string) Value {
	if !filepath.IsAbs(path) {
		path = filepath.Join(WorkingDir, path)
	}
	if f, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			panic(Errorf("File `%s` does not exist", path))
		} else {
			panic(Errorf("File `%s` unexpected error: %v", path, err))
		}
	} else if f.IsDir() {
		panic(Errorf("File `%s` is a directory", path))
	}
	return localFile{path: path}
}

type localFile struct {
	path  string
	cache *string
}

func (f localFile) String() string {
	if f.cache == nil {
		v, err := ioutil.ReadFile(f.path)
		if err != nil {
			panic(Errorf("File `%s` unexpected error: %v", f.path, err))
		}
		str := string(v)
		f.cache = &str
	}
	return string(*f.cache)
}

func Template(path string) Value {
	return StringValue("")
}

// func Download(url, hash string) Value {
// 	return StringValue("")
// }

// TODO...
func Decrypt(path Value) Value {
	return path
}
