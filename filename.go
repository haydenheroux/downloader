package main

import (
	"strings"

	"github.com/haydenheroux/cleanstring"
)

func OutputFilename(artistsDirty string, trackNameDirty string) string {
	artists := strings.Split(artistsDirty, "&")
	artistString := cleanstring.CleanSlice(artists)

	trackString := cleanstring.Clean(trackNameDirty)

	artistTrack := map[string]string{artistString: trackString}

	return cleanstring.CleanMap(artistTrack)
}

func ChangeExtension(fileName string, extension string) string {
	temp := strings.Split(fileName, ".")
	str := strings.Join(temp[:len(temp)-1], ".")
	return str + "." + extension
}
