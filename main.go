package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"time"
)

var directories []Directory
var myHeader Header
var app = tview.NewApplication()
var headerText = tview.NewTextView()

func StartScan(path string) {
	startTime := time.Now()
	fmt.Printf("Looking for Node Modules")
	err := CheckNodeModulesInDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	endTime := time.Now()
	myHeader.time = endTime.Sub(startTime)
}

func main() {
	var scanPath string
	if len(os.Args) > 1 {
		scanPath = os.Args[1]
	} else {
		fmt.Fprintf(os.Stderr, "No parameters were provided. Please provide the path to scan")
		os.Exit(1)
	}

	StartScan(scanPath)
	flex, dirList := ConfigureApp(app)

	// Populate header details
	headerText.SetTextColor(tcell.ColorGreen)
	SetHeaderMessage()
	PopulateList(dirList)
	doneChan := make(chan error)
	PopulateWithSizes(dirList, doneChan)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

/*func RefreshList(dirList *tview.List) {
	for index, dir := range directories {
		myHeader.totalSize += dir.size
		dirString := fmt.Sprintf("%s              [Calculating...⌛️]", dir.path)
		if dir.calculating == false {
			dirString = fmt.Sprintf("%s              %d MB", dir.path, dir.size)
		}
		dirList.SetItemText(index, dirString, "")
	}
}*/
