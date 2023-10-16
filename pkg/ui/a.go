package ui

import (
	"fmt"

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
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	// go func() {
	// 	//systray.SetTemplateIcon(icon.Data, icon.Data)
	// 	systray.SetTitle("qlwiz")
	// 	systray.SetTooltip("Quick Link Wizard")
	// 	systray.AddMenuItem("Links", "All Links")
	// 	systray.AddSeparator()
	// 	link, _ := links.GetAllLinks()
	// 	itemMap := map[*systray.MenuItem]helpers.Link{}
	// 	itemLst := []chan *systray.MenuItem{}
	// 	for _, i := range link {
	// 		it := systray.AddMenuItem(i.Name, i.Link)
	// 		itemLst = append(itemLst, make(chan *systray.MenuItem))
	// 		itemMap[it] = i
	// 	}
	// 	mChange := systray.AddMenuItem("Change Me", "Change Me")
	// 	mChecked := systray.AddMenuItemCheckbox("Unchecked", "Check Me", true)
	// 	mEnabled := systray.AddMenuItem("Enabled", "Enabled")
	// 	// Sets the icon of a menu item. Only available on Mac.
	// 	//mEnabled.SetTemplateIcon(icon.Data, icon.Data)

	// 	systray.AddMenuItem("Ignored", "Ignored")

	// 	subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
	// 	subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
	// 	subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
	// 	subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")

	// 	mUrl := systray.AddMenuItem("Open UI", "my home")
	// 	mQuit := systray.AddMenuItem("退出", "Quit the whole app")

	// 	// Sets the icon of a menu item. Only available on Mac.
	// 	//mQuit.SetIcon(icon.Data)

	// 	systray.AddSeparator()
	// 	mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")
	// 	shown := true
	// 	toggle := func() {
	// 		if shown {
	// 			subMenuBottom.Check()
	// 			subMenuBottom2.Hide()
	// 			mQuitOrig.Hide()
	// 			mEnabled.Hide()
	// 			shown = false
	// 		} else {
	// 			subMenuBottom.Uncheck()
	// 			subMenuBottom2.Show()
	// 			mQuitOrig.Show()
	// 			mEnabled.Show()
	// 			shown = true
	// 		}
	// 	}

	// 	for {

	// 		agg := make(chan *systray.MenuItem)
	// 		for _, ch := range itemLst {
	// 			go func(c chan *systray.MenuItem) {
	// 				for x := range c {
	// 					agg <- x
	// 				}
	// 			}(ch)
	// 		}

	// 		select {
	// 		case <-mChange.ClickedCh:
	// 			mChange.SetTitle("I've Changed")
	// 		case <-mChecked.ClickedCh:
	// 			if mChecked.Checked() {
	// 				mChecked.Uncheck()
	// 				mChecked.SetTitle("Unchecked")
	// 			} else {
	// 				mChecked.Check()
	// 				mChecked.SetTitle("Checked")
	// 			}
	// 		case <-mEnabled.ClickedCh:
	// 			mEnabled.SetTitle("Disabled")
	// 			mEnabled.Disable()
	// 		case <-mUrl.ClickedCh:
	// 			open.Run("https://www.getlantern.org")
	// 		case <-subMenuBottom2.ClickedCh:
	// 			panic("panic button pressed")
	// 		case <-subMenuBottom.ClickedCh:
	// 			toggle()
	// 		case <-mToggle.ClickedCh:
	// 			toggle()
	// 		case <-mQuit.ClickedCh:
	// 			systray.Quit()
	// 			fmt.Println("Quit2 now...")
	// 			return
	// 		}
	// 	}
	// }()
}
