package main

import (
	"errors"
	"net/url"
	"strings"
  "fmt"
)

type Track struct {
	URL     string
	Artists string
	Title   string
}

func isURL(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func toURL(stub string) string {
	if isURL(stub) == false {
    // Assume the stub is a YouTube ID if it is not a URL
  	return "https://www.youtube.com/watch?v=" + stub
	}
	return stub
}

func TrackFrom(s string) (Track, error) {
	fields := strings.Split(s, "\t")
  fmt.Println(fields)
	if len(fields) == 3 {
	  track := Track{
      URL: toURL(fields[0]),
      Artists: fields[1],
      Title: fields[2],
    }
    return track, nil
	}
	return Track{}, errors.New("not enough fields in line")
}
