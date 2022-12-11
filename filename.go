package main

import (
	"strings"
)

func OutputFilename(artistsDirty string, trackNameDirty string) string {
	artists := strings.Split(artistsDirty, "&")
	artistString := CleanSlice(artists)

	trackString := Clean(trackNameDirty)

	artistTrack := map[string]string{artistString: trackString}

	return CleanMap(artistTrack)
}

func ChangeExtension(fileName string, extension string) string {
	temp := strings.Split(fileName, ".")
	str := strings.Join(temp[:len(temp)-1], ".")
	return str + "." + extension
}
