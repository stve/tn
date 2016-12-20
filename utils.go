package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bogem/id3v2"
)

// helper functions

func setPicture(coverFilename string, filename string) error {
	frontCover, err := ioutil.ReadFile(coverFilename)
	if err != nil {
		return err
	}

	pic := id3v2.PictureFrame{
		Encoding:    id3v2.ENUTF8,
		MimeType:    "image/jpeg",
		Picture:     frontCover,
		PictureType: id3v2.PTFrontCover,
		Description: "Front Cover",
	}

	// Open file and find tag in it
	tag, err := id3v2.Open(filename)
	if err != nil {
		return err
	}

	tag.DeleteFrames(tag.CommonID("Attached picture"))
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag:", err)
	}

	tag.AddAttachedPicture(pic)
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag:", err)
	}

	if err = tag.Close(); err != nil {
		log.Fatal("Error while closing a tag:", err)
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return (err == nil)
}

func checkNotEmpty(str string) bool {
	return (len(str) > 0)
}
