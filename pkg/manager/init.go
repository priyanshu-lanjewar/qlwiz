package manager

import (
	"os"
	"os/user"
	"path/filepath"
)

func Initialize() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	homeDir := currentUser.HomeDir
	folderName := ".qlwiz"
	folderPath := filepath.Join(homeDir, folderName)

	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}

	fileName := "config.yaml"
	filePath := filepath.Join(folderPath, fileName)

	yamlContent := `links:
  - name: qlwiz
    link: https://github.com/priyanshu-lanjewar/qlwiz
groups:
  - name: qlwiz
    member:
      - qlwiz-git
`
	err = os.WriteFile(filePath, []byte(yamlContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
