package helpers

type Link struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}

type Group struct {
	Name   string `yaml:"name"`
	Member []Link `yaml:"member"`
}

type Config struct {
	Links  []Link  `yaml:"links"`
	Groups []Group `yaml:"groups"`
}
