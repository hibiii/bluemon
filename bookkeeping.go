package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const download_record_file = "dls.json"

func BookkeepingLoadRecords() (*DownloadRecord, error) {
	file, err := os.OpenFile(download_record_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records := DownloadRecord{}

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if len(contents) == 0 {
		return &records, nil
	}

	err = json.Unmarshal(contents, &records)
	if err != nil {
		return nil, err
	}
	return &records, nil
}

func BookkeepingSaveRecords(r *DownloadRecord) error {
	file, err := os.OpenFile(download_record_file+".tmp", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	contents, err := json.Marshal(r)
	if err != nil {
		return err
	}

	bytesToGo := len(contents)
	for bytesToGo > 0 {
		written, err := file.Write(contents)
		if err != nil && err != io.ErrShortWrite {
			return err
		}
		bytesToGo -= written
	}

	return bookkeepingSwapFiles()
}

func bookkeepingSwapFiles() error {
	// Warning: function relies on the fact that
	// "If newpath already exists and is not a directory, Rename replaces it."
	err := os.Rename(download_record_file, download_record_file+".bkp")
	if err != nil {
		return fmt.Errorf("bookkeepingSwapFiles: 1st step failed: %s", err)
	}
	err = os.Rename(download_record_file+".tmp", download_record_file)
	if err != nil {
		return fmt.Errorf("bookkeepingSwapFiles: 2nd step failed: %s", err)
	}
	return nil
}
