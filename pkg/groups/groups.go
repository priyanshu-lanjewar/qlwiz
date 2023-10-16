package groups

import (
	"os/user"
	"path/filepath"
	"sort"

	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
)

func GetAllGroups() ([]helpers.Group, error) {
	currentUser, err := user.Current()
	if err != nil {
		return []helpers.Group{}, err
	}

	homeDir := currentUser.HomeDir
	folderName := ".qlwiz"
	folderPath := filepath.Join(homeDir, folderName)
	filename := "config.yaml"
	filePath := filepath.Join(folderPath, filename)
	config, err := helpers.ReadConfig(filePath)
	if err != nil {
		return []helpers.Group{}, err
	}
	sort.Slice(config.Groups, func(i, j int) bool {
		return config.Groups[i].Name < config.Groups[j].Name
	})
	return config.Groups, nil
}
