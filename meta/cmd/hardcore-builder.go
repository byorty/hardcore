package main

import (
	"encoding/xml"
	"flag"
	"github.com/byorty/hardcore/meta"
	"github.com/byorty/hardcore/meta/plugin"
	"github.com/byorty/hardcore/utils"
	"github.com/byorty/log4go"
	"io/ioutil"
	"os"
	"path/filepath"
    "strings"
    "sort"
)

var (
	plugins = []plugin.Plugin{
		new(plugin.Include),
		new(plugin.Init),
		new(plugin.Directory),
		new(plugin.Enum),
		new(plugin.Controller),
	}
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "configuration xml file")
	flag.Parse()

	log := log4go.NewDefaultLogger(log4go.DEBUG)
	defer log.Close()

	filename = filepath.Join(utils.Pwd(), filename)
	if utils.FileExists(filename) {
		log.Debug("file %s is exists", filename)
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			log.Debug("success read file %s", filename)
			var config meta.Configuration
			err = xml.Unmarshal(data, &config)
			if err == nil {
				log.Debug("success unmarshal file %s", filename)
				env := new(meta.Environment)
				env.MetaPath = filepath.Dir(filename)
				env.AbsPath, _ = filepath.Abs(filepath.Join(env.MetaPath, ".."))
				env.Configuration = &config
				env.Logger = log

                parts := strings.Split(env.AbsPath, string(filepath.Separator))
                i := sort.Search(len(parts), func(x int) bool {
                    return parts[x] == "src"
                })
                if i < len(parts) && parts[i] == "src" {
                    env.ImportPart = filepath.Join(parts[i+1:]...)
                } else {
                    log.Critical("can't find directory src")
                    os.Exit(1)
                }

				for _, pl := range plugins {
					pl.Do(env)
				}
			} else {
				log.Critical("can't unmarshal xml file %s", filename)
				os.Exit(1)
			}
		} else {
			log.Critical("can't read file %s", filename)
			os.Exit(1)
		}
	} else {
		log.Critical("file %s not found", filename)
		os.Exit(1)
	}
}
