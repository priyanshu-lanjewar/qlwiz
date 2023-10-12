package manager

import (
	"os"
	"os/user"
	"path/filepath"
)

func Reset() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	homeDir := currentUser.HomeDir
	folderName := ".qlwiz"
	folderPath := filepath.Join(homeDir, folderName)

	fileName := "config.yaml"
	filePath := filepath.Join(folderPath, fileName)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := os.RemoveAll(folderPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	return Initialize()
}
