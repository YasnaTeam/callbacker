package client

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	Server   string       `json:"server"`
	Username string       `json:"username"`
	Routes   []RouteTable `json:"routes"`
}

type RouteTable struct {
	Route    string
	Callback string
}

func (c *Configuration) Close() {
	c.Save()
}

func (c *Configuration) Save() error {
	return saveConfiguration(c)
}

func (c *Configuration) AddRouteCallback(route, callback string) {
	c.Routes = append(c.Routes, RouteTable{route, callback})
}

func NewConfigurationFile() error {
	routes := []RouteTable{}
	configuration := &Configuration{
		"",
		"",
		routes,
	}

	return saveConfiguration(configuration)
}

func GetConfiguration() (*Configuration, error) {
	if _, err := os.Stat("routes.json"); os.IsNotExist(err) {
		NewConfigurationFile()
	}

	file, _ := os.Open("routes.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	err := decoder.Decode(configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}

func saveConfiguration(configuration *Configuration) error {
	fp, err := os.Create("routes.json")
	if err != nil {
		return err
	}
	defer fp.Close()

	information, err := json.Marshal(configuration)
	if err != nil {
		return err
	}

	_, err = fp.Write(information)
	if err != nil {
		return err
	}

	err = fp.Sync()
	if err != nil {
		return err
	}

	return nil
}