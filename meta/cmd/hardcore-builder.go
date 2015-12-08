package main

import (
	"os"
	"flag"
	"io/ioutil"
	"github.com/byorty/hardcore/meta"
	"encoding/xml"
	"path/filepath"
	"github.com/byorty/hardcore/utils"
	"github.com/byorty/hardcore/meta/plugin"
	"github.com/byorty/log4go"
)

var (
	plugins = []plugin.Plugin{
		plugin.NewIncludePlugin(),
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
				for _, pl := range plugins {
					pl.Do(env)
				}
				log.Debug(config)
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
