package links

import (
	"os/exec"
	"runtime"

	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
)

func Launch(link helpers.Link) error {
	var err error
	url := link.Link
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		return err
	}

	return nil
}
