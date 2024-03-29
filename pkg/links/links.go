package links

import (
	"os/user"
	"path/filepath"
	"sort"

	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
)

func GetAllLinks() ([]helpers.Link, error) {
	currentUser, err := user.Current()
	if err != nil {
		return []helpers.Link{}, err
	}

	homeDir := currentUser.HomeDir
	folderName := ".qlwiz"
	folderPath := filepath.Join(homeDir, folderName)
	filename := "config.yaml"
	filePath := filepath.Join(folderPath, filename)
	config, err := helpers.ReadConfig(filePath)
	if err != nil {
		return []helpers.Link{}, err
	}
	sort.Slice(config.Links, func(i, j int) bool {
		return config.Links[i].Name < config.Links[j].Name
	})
	return config.Links, nil

}
