package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"

	"go-crdt-load-test/internal/app"
	"go-crdt-load-test/loader"
)

func main() {
	loaderConfig, err := getConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(loaderConfig); err != nil {
		log.Fatal(err)
	}
}

func getConfig(path string) (*loader.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var loaderConfig loader.Config
	err = yaml.Unmarshal(yamlFile, &loaderConfig)
	if err != nil {
		return nil, err
	}

	return &loaderConfig, nil
}
