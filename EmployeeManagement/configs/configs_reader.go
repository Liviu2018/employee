package configsreader

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Configs models the configuration values needed by the app
type Configs struct {
	Port       int    `json:"port" yaml:"port"`
	DataSource string `json:"dataSource" yaml:"dataSource"`
}

func (c *Configs) String() string {
	return fmt.Sprintf("Configs[Port %d, DataSource %s]", c.Port, c.DataSource)
}

// ReadConfigs reads configs form a yaml file, returns a Config struct with the values it has read
func ReadConfigs() (*Configs, error) {
	path, _ := os.Getwd()

	filename := path + "/../configs/config.yaml"

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %v yaml file: %v \n", filename, err)
		return nil, err
	}

	result := Configs{}

	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		fmt.Printf("Error unmarshalling yaml file content: %v \n", err)
		return nil, err
	}

	return &result, nil
}
