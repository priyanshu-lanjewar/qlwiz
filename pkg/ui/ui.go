package ui

import (
	"github.com/getlantern/systray"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/icon"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/links"
	"github.com/skratchdot/open-golang/open"
)

func launchLink(linkItem *systray.MenuItem, link helpers.Link) {
	go func() {
		for {
			<-linkItem.ClickedCh
			open.Run(link.Link)
		}
	}()
}

func OnReady() {
	systray.SetTitle("qlwiz app")
	systray.SetTemplateIcon(icon.AppIcon, icon.AppIcon)
	systray.AddMenuItem("Links", "")
	systray.AddSeparator()
	link, _ := links.GetAllLinks()
	for _, l := range link {
		item := systray.AddMenuItem(l.Name, l.Link)
		item.SetIcon(icon.LinkIcon)
		launchLink(item, l)
	}
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the app")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}
