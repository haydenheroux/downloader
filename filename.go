package main

import (
	"regexp"
	"strings"
)

func clean(str string) string {
	reg, err := regexp.Compile("[&\\'+\\(\\)\\[\\]\\.\\-]+")
	if err != nil {
		printError("regexp.Compile", err)
	}
	temp := str
	// Remove padding spaces
	temp = strings.Trim(temp, " \n")
	// Spaces -> underscores
	temp = strings.ReplaceAll(temp, " ", "_")
	// Dash -> underscores
	// Remove "bad" characters
	temp = reg.ReplaceAllString(temp, "")
	// Lowercase
	temp = strings.ToLower(temp)
	return temp
}

func CreateOutputFileName(artistsDirty string, trackNameDirty string) string {
	var trackName string

	artists := strings.Split(artistsDirty, "&")
	for index, artist := range artists {
		artists[index] = clean(artist)
	}
	artistString := strings.Join(artists, "+")

	trackName = clean(trackNameDirty)

	return artistString + "-" + trackName
}

func ChangeExtension(fileName string, extension string) string {
	temp := strings.Split(fileName, ".")
	str := strings.Join(temp[:len(temp)-1], ".")
	return str + "." + extension
}
