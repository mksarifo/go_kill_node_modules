package main

import (
	"fmt"
	"time"
)

type Header struct {
	time          time.Duration
	totalSize     int64
	reclaimedSize int64
	text          string
	selectedId    int
}

const HeaderFormat = "Time: %s     Total Size: %s     Reclaimed: %s\n\n(q) to quit\n(d) to delete\nInfo: %s"

func formatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d MB", size)
	} else {
		return fmt.Sprintf("%d GB", size/1024)
	}
}

func SetHeaderMessage() {
	myHeader.text = "Use the mouse or keyboard to select a directory"

	totalSize := formatSize(myHeader.totalSize)
	reclaimedSize := formatSize(myHeader.reclaimedSize)

	headerText.SetText(
		fmt.Sprintf(HeaderFormat, myHeader.time, totalSize, reclaimedSize, myHeader.text))
}
