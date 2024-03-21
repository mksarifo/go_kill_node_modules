package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ConfigureApp(app *tview.Application) (*tview.Flex, *tview.List) {
	var flex = tview.NewFlex()
	var dirList = tview.NewList().ShowSecondaryText(false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 100 {
			done := make(chan error)
			index := myHeader.selectedId

			dirList.SetItemText(myHeader.selectedId, directories[myHeader.selectedId].path+" [ REMOVING ⌛️]", "")
			myHeader.reclaimedSize += directories[myHeader.selectedId].size
			SetHeaderMessage()
			RemoveDir(dirList, index, done)

		}
		return event
	})
	flex.SetDirection(tview.FlexRow)
	flexContainer := tview.NewFlex()
	flex.AddItem(headerText, 0, 1, false)
	flex.AddItem(flexContainer, 0, 1, true)
	flexContainer.AddItem(dirList, 0, 1, true)

	return flex, dirList
}

func PopulateList(dirList *tview.List) {
	for index, dir := range directories {
		dir.calculating = true
		myHeader.totalSize += dir.size
		dirString := fmt.Sprintf("%s              [Calculating...⌛️]", dir.path)
		color := tcell.ColorBlack
		dirList.AddItem(dirString, "", rune(49+index), nil).SetShortcutColor(color).SetChangedFunc(HandleSelect)
	}
}

func HandleSelect(index int, s1 string, s2 string, rune rune) {
	myHeader.text = fmt.Sprintf("%s", s1)
	myHeader.selectedId = index
	SetHeaderMessage()
}
