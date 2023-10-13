package links

import (
	"os/user"
	"path/filepath"

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
	return config.Links, nil

}
