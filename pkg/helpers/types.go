package helpers

type Link struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}

type Group struct {
	Name   string   `yaml:"name"`
	Member []string `yaml:"member"`
}

type Config struct {
	Links  []Link   `yaml:"links"`
	Groups []Group  `yaml:"groups"`
}