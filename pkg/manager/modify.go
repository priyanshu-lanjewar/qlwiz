package manager

import (
	"os/user"
	"path/filepath"

	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/links"
)

func EditConfig() error {

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	homeDir := currentUser.HomeDir
	folderName := ".qlwiz"
	folderPath := filepath.Join(homeDir, folderName)
	filename := "config.yaml"
	filePath := filepath.Join(folderPath, filename)

	return links.Launch(helpers.Link{
		Link: filePath,
	})
}
