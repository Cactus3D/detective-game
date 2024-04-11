package locations

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type YamlMap struct {
	Map struct {
		Layers []struct {
			ID    int    `yaml:"id"`
			Name  string `yaml:"name"`
			Edges []struct {
				Dst int `yaml:"dst"`
				Val int `yaml:"val"`
			} `yaml:"edges,omitempty"`
			Parent int `yaml:"parent,omitempty"`
		} `yaml:"layers"`
	} `yaml:"map"`
}

func LoadYamlMap(path string) (*YamlMap, error) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln("cannot read yamlmap file:", path)
		return nil, err
	}

	var yamlMap YamlMap

	err = yaml.Unmarshal(yamlFile, &yamlMap)

	if err != nil {
		log.Fatalln("cannot unmarshall gamemap from file:", path)
		return nil, err
	}

	return &yamlMap, nil
}
