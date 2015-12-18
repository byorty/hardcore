package main

import (
	"encoding/xml"
	"flag"
	"github.com/byorty/hardcore/meta"
	"github.com/byorty/hardcore/meta/plugin"
	"github.com/byorty/hardcore/utils"
	"io/ioutil"
	"os"
	"path/filepath"
    "strings"
    "sort"
    "github.com/byorty/hardcore/meta/common"
    "github.com/byorty/hardcore/log"
)

var (
	plugins = []plugin.Plugin{
		new(plugin.Include),
		new(plugin.Init),
		new(plugin.Directory),
		new(plugin.Enum),
		new(plugin.Model),
		new(plugin.Controller),
		new(plugin.File),
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
			var config meta.Configuration
			err = xml.Unmarshal(data, &config)
			if err == nil {
                logger.Debug("success unmarshal file %s", filename)
				env := new(meta.Environment)
				env.MetaPath = filepath.Dir(filename)
				env.AbsPath, _ = filepath.Abs(filepath.Join(env.MetaPath, ".."))
				env.Configuration = &config
				env.Configuration.Files = make([]common.File, 0)
				env.Logger = logger

                parts := strings.Split(env.AbsPath, string(filepath.Separator))
                i := sort.Search(len(parts), func(x int) bool {
                    return parts[x] == "src"
                })
                if i < len(parts) && parts[i] == "src" {
                    env.ImportPart = filepath.Join(parts[i+1:]...)
                } else {
                    logger.Critical("can't find directory src")
                    os.Exit(1)
                }

				for _, pl := range plugins {
					pl.Do(env)
				}
			} else {
                logger.Critical("can't unmarshal xml file %s", filename)
				os.Exit(1)
			}
		} else {
            logger.Critical("can't read file %s", filename)
			os.Exit(1)
		}
	} else {
        logger.Critical("file %s not found", filename)
		os.Exit(1)
	}
}
