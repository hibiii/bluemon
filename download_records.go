package main

// Represents the data of a record entry, a version ref -> download count map.
//
// - key string: a version ref as defined in VersionRefTable.
//
// - value float64: the download number for that specific version, total.
type Entry map[string]float64

// Represents the entry holder.
//
// - key string: an UNIX timestamp of when the check was performed.
//
// - value Entry: the total download numbers recorded at the time of checking.
type EntryTable map[string]Entry

// Represents a version reference table in order to save some bytes of storage.
// All Entry accesses must be done with a version ref.
//
// - key string: a version name, as used in `Version.Name`
//
// - value string: the version ref that's used internally in individual Entrys
type VersionRefTable map[string]string

// The download records file structure
type DownloadRecord struct {
	References VersionRefTable `json:"refs"`
	Entries    EntryTable      `json:"entries"`
}
