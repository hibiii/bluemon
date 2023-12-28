package main

import (
	"fmt"
)

func main() {
	config, err := ConfigReadFromFile("config.json")
	if err != nil {
		fmt.Printf("Error: could not read file: %s\n", err)
		return
	}

	mc := NewModrinthClient()
	for i := 0; i < len(config.Versions); i++ {
		ver := &config.Versions[i]
		err = mc.GetDownloadsForVersion(ver)
		if err != nil {
			fmt.Printf("Error: could not get download count for version %s: %s\n", ver.Name, err)
		}
	}
	for i := 0; i < len(config.Versions); i++ {
		ver := config.Versions[i]
		if ver.Downloads < 0 {
			fmt.Printf("%s: unsuccessful\n", ver.Name)
		} else {
			fmt.Printf("%s: %d downloads\n", ver.Name, ver.Downloads)
		}
	}
}
