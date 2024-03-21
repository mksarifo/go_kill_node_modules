package main

func RemoveIndexPreserveOrder(s []Directory, index int) []Directory {
	s = append(s[:index], s[index+1:]...)
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

func bytesToMB(size int64) float64 {
	return float64(size) / 1024.0 / 1024.0
}
