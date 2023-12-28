package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: bluemon <versionID>")
		return
	}
	version := os.Args[1]
	mc := NewModrinthClient()
	dl, err := mc.GetVersionDownloads(version)
	if err != nil {
		fmt.Printf("Error: could not get download count for version %s: %s\n", version, err)
	} else {
		fmt.Printf("Downloads: %d\n", dl)
	}
}
