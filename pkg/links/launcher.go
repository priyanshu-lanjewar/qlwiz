package links

import (
	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
	"github.com/skratchdot/open-golang/open"
)

func Launch(link helpers.Link) error {
	var err error
	url := link.Link
	err = open.Run(url)

	if err != nil {
		return err
	}

	return nil
}
