package misc

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/BurntSushi/xdg"

	"github.com/u-root/wingo/logger"
)

var ConfigPaths = xdg.Paths{
	Override:  "",
	XDGSuffix: "wingo",
}

var DataPaths = xdg.Paths{
	Override:  "",
	XDGSuffix: "wingo",
}

var ScriptPaths = xdg.Paths{
	Override:  "",
	XDGSuffix: "wingo",
}

func ConfigFile(name string) string {
	fpath, err := ConfigPaths.ConfigFile(name)
	if err != nil {
		_, ok := FileMap[name]
		if !ok {
			logger.Error.Fatalf("ConfigFile(%q): %v and not in map", name, err)
		}
		return name
	}
	return fpath
}

func DataFile(name string) ([]byte, error) {
	fpath, err := DataPaths.DataFile(name)
	if err != nil {
		b, ok := FileMap[name]
		if !ok {
			return nil, fmt.Errorf("%v: OS got %v and there is no builtin file", name, err)
		}
		return b, nil
	}
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	FileMap[name] = b
	return b, nil
}

func MustDataFile(name string) []byte {
	b, err := DataFile(name)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	return b
}

func ScriptPath(name string) string {
	fpath, err := ScriptPaths.ConfigFile(path.Join("scripts", name, name))
	if err != nil {
		fpath, err = ScriptPaths.ConfigFile(path.Join("scripts", name))
		if err != nil {
			logger.Warning.Println(err)
			return ""
		}
	}
	return fpath
}

func ScriptConfigPath(name string) string {
	fname := fmt.Sprintf("%s.cfg", name)
	fp, err := ScriptPaths.ConfigFile(path.Join("scripts", name, fname))
	if err != nil {
		logger.Warning.Println(err)
		return ""
	}
	return fp
}
