package main

import (
	"fmt"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

type Directory struct {
	path        string
	size        int64
	calculating bool
}

const NodeModules = "node_modules"

func CheckNodeModulesInDir(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == NodeModules {
			directories = append(directories, Directory{path: path, size: 0})
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking the path %v: %v", path, err)
	}
	return nil
}

func DirSizeThreaded(dir Directory, dirChan chan Directory, errChan chan error) {
	go func() {
		var size int64
		//dirChan := dir
		err := filepath.Walk(dir.path, func(_ string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				size += info.Size()
			}
			return err
		})
		if err != nil {
			errChan <- err
		} else {
			dir.size = int64(bytesToMB(size))
			dirChan <- dir
		}
	}()
}

func PopulateWithSizes(dirList *tview.List, done chan error) {
	go func() {
		for index, dir := range directories {
			dirChan := make(chan Directory)
			errChan := make(chan error)
			DirSizeThreaded(dir, dirChan, errChan)
			select {
			case dirItem := <-dirChan:
				dir.calculating = false
				dir.size = dirItem.size
				myHeader.totalSize += dirItem.size
				dirString := fmt.Sprintf("%s              %d MB", dir.path, dirItem.size)
				dirList.SetItemText(index, dirString, "")
				myHeader.text = "Calculating Size..."
				SetHeaderMessage()
			case err := <-errChan:
				done <- err
				return
			}
		}
		SetHeaderMessage()
		//RefreshList(dirList)
		done <- nil
	}()

}

func RemoveDir(dirList *tview.List, index int, done chan error) {
	go func() {
		dirIndex := myHeader.selectedId
		// Delete a directory and its contents
		err := os.RemoveAll(directories[dirIndex].path)
		if err != nil {
			headerText.SetText(
				fmt.Sprintf("Time: %s     Total Size: %d\n\n(q) to quit\n(d) to delete\nInfo: %s", myHeader.time, myHeader.totalSize, err.Error()))
			done <- err
			return
		}
		myHeader.reclaimedSize += directories[dirIndex].size
		SetHeaderMessage()
		RemoveIndexPreserveOrder(directories, index)
		dirList.RemoveItem(index)

		done <- nil
	}()
}
