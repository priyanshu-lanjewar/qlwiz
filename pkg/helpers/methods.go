package helpers

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func WriteConfig(path string, config Config) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
