package main

import (
	"os"
	"fmt"
	"flag"
	"io/ioutil"
	"github.com/byorty/hardcore/meta"
	"encoding/xml"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "configuration xml file")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filename = fmt.Sprintf("%s%s%s", pwd, string(os.PathSeparator), filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	} else {
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			var config meta.Configuration
			err = xml.Unmarshal(data, &config)
			if err == nil {
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
