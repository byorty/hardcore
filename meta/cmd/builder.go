package main

import (
	"encoding/xml"
	"flag"
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/meta/conf"
	"github.com/byorty/hardcore/meta/plugin"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	plugins = []types.Plugin{
		new(plugin.Init),
		new(plugin.Directory),
		new(plugin.Extend),
		new(plugin.Enum),
		new(plugin.Model),
		new(plugin.Exporter),
		new(plugin.Importer),
		new(plugin.Middleware),
		new(plugin.Controller),
		new(plugin.Router),
		new(plugin.DAOs),
		new(plugin.File),
		new(plugin.Formatter),
	}
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "configuration xml file")
	flag.Parse()

	logger := log.NewDefaultLogger(log.FINEST)
	defer logger.Close()

	filename = filepath.Join(utils.Pwd(), filename)
	if utils.FileExists(filename) {
		logger.Debug("file %s is exists", filename)
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			logger.Debug("success read file %s", filename)
			var config conf.Configuration
			err = xml.Unmarshal(data, &config)
			if err == nil {
				logger.Debug("success unmarshal file %s", filename)

				metaPath := filepath.Dir(filename)
				absPath, _ := filepath.Abs(filepath.Join(metaPath, ".."))
				parts := strings.Split(absPath, string(filepath.Separator))

				var i int
				hasSrcDir := false
				for j, part := range parts {
					if part == "src" {
						i = j
						hasSrcDir = true
						break
					}
				}

				if hasSrcDir {
					config.Files = make([]types.File, 0)
					env := conf.NewEnvironment(
						absPath,
						metaPath,
						filepath.Join(parts[i+1:]...),
						logger,
						&config,
					)

					for _, pl := range plugins {
						pl.Do(env)
					}
				} else {
					logger.Critical("can't find directory src")
					os.Exit(4)
				}
			} else {
				logger.Critical("can't unmarshal xml file %s", filename)
				os.Exit(3)
			}
		} else {
			logger.Critical("can't read file %s", filename)
			os.Exit(2)
		}
	} else {
		logger.Critical("file %s not found", filename)
		os.Exit(1)
	}
}
