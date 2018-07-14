package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jawher/mow.cli"
)

func main() {

	type feature1 struct {
		Enable     bool
		validators []string
	}

	type feature2 struct {
		Enable bool
	}

	type tomlConfig struct {
		Title string
		F1    feature1 `toml:"feature1"`
		F2    feature2 `toml:"feature2"`
	}

	//this will create an app for the cli using mow.cli
	app := cli.App("gallactic", "Read the configuration files from the working directory")

	// This part will i.e spec will check for the number of arguments and generate the help/ usage of this automatically
	app.Spec = "-d"

	var (
		// declare the wokring directory argument as a multi-string argument
		dir = app.StringArg("DIR", "", "Gallactic Working Directory")
	)

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		configFile := *dir + "config.toml"
		dat, err := ioutil.ReadFile(configFile)
		check(err)
		fmt.Print(string(dat))
		tomlData := string(dat)
		var conf tomlConfig
		if _, err := toml.Decode(tomlData, &conf); err != nil {
			log.Fatal(err)
		}
		log.Printf("title: %s", conf.Title)
		log.Printf("Feature 1: %#v", conf.F1)
		log.Printf("Feature 2: %#v", conf.F2)

		TOMLString(conf)
	}
	// Invoke the app passing in os.Args
	app.Run(os.Args[1:])

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TOMLString(conf interface{}) string {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	err := encoder.Encode(conf)
	if err != nil {
		return fmt.Sprintf("<Could not serialise config: %v>", err)
	}
	bs := buf.Bytes()

	err = ioutil.WriteFile("output.toml", bs, 0644)

	return buf.String()
}
