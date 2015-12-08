package main

import (
	"os"
	"fmt"
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
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			var config meta.Configuration
			err = xml.Unmarshal(data, &config)
			if err == nil {
				config.MetaPath = filepath.Dir(filename)
				config.AbsPath, _ = filepath.Abs(filepath.Join(config.MetaPath, ".."))
				for _, pl := range plugins {
					pl.Do(&config)
				}
				log.Info(config)
				fmt.Println(config)
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
