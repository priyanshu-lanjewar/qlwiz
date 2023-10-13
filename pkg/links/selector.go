package links

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
)

func getLinkNames(links []helpers.Link) []string {
	names := make([]string, len(links))
	for i, link := range links {
		names[i] = link.Name
	}
	return names
}

func GetLinksToLaunch(links []helpers.Link) (helpers.Link, error) {
	prompt := &survey.Select{
		Message: "Select a link:",
		Options: getLinkNames(links),
	}

	var selectedLinkIndex int
	if err := survey.AskOne(prompt, &selectedLinkIndex); err != nil {
		return helpers.Link{}, err
	}

	return links[selectedLinkIndex], nil
}
