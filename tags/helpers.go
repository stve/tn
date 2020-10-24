package tags

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bogem/id3v2"
)

// helper functions

func setPicture(coverFilename string, filename string) error {
	frontCover, err := os.Open(coverFilename)
	if err != nil {
		return err
	}
	defer frontCover.Close()

	frontCoverBytes, err := ioutil.ReadAll(frontCover)
	if err != nil {
		return err
	}

	pic := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		Picture:     frontCoverBytes,
		PictureType: id3v2.PTFrontCover,
		Description: "Cover",
	}

	// Open file and find tag in it
	tag, err := id3v2.Open(filename, id3v2.Options{Parse: false})
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

func clearTags(filename string) error {
	tag, err := id3v2.Open(filename, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("error opening mp3 file", err)
	}
	defer tag.Close()

	tag.DeleteAllFrames()
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag:", err)
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
