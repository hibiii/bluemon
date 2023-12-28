package main

// Represents a single version of a release for one version of Minecraft.
type Version struct {
	Name       string `json:"name"`
	Downloads  int    `json:"-"`
	ModrinthId string `json:"modrinthId"`
}
