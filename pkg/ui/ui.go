package ui

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/getlantern/systray"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/groups"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/helpers"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/icon"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/links"
	"github.com/priyanshu-lanjewar/qlwiz/pkg/manager"
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
	link, _ := links.GetAllLinks()
	if len(link) != 0 {
		systray.AddMenuItem("Links", "")
		systray.AddSeparator()
		for _, l := range link {
			item := systray.AddMenuItem(l.Name, l.Link)
			item.SetIcon(icon.LinkIcon)
			launchLink(item, l)
		}
		systray.AddSeparator()
	}
	grps, _ := groups.GetAllGroups()
	fmt.Println(grps)
	if len(grps) != 0 {
		systray.AddMenuItem("Groups", "")
		systray.AddSeparator()
		for _, g := range grps {
			grpItem := systray.AddMenuItem(g.Name, "")
			grpItem.SetIcon(icon.GroupIcon)
			for _, item := range g.Member {
				subItem := grpItem.AddSubMenuItem(item.Name, item.Link)
				subItem.SetIcon(icon.LinkIcon)
				launchLink(subItem, item)
			}
		}
		systray.AddSeparator()
	}
	manage := systray.AddMenuItem("Manage", "")
	refresh := manage.AddSubMenuItem("Refresh", "")
	edit := manage.AddSubMenuItem("Edit Config", "")
	exit := manage.AddSubMenuItem("Exit", "")
	go func() {
		for {
			select {
			case <-exit.ClickedCh:
				systray.Quit()
			case <-edit.ClickedCh:
				manager.EditConfig()
			case <-refresh.ClickedCh:
				systray.Quit()
				command := exec.Command("qlwiz", "launch-ui-cmd")
				command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_ = command.Start()
			}
		}
	}()
}
