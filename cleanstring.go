package main

import (
  "strings"
)

// isAlpha determines if the provided rune is part of the alphabet runes.
// isAlpha checks if the rune is in either the lowercase or uppercase rune region.
func isAlpha(r rune) bool {
  isUppercase := 'A' <= r && r <= 'Z'
  isLowercase := 'a' <= r && r <= 'z' 
  return isUppercase || isLowercase
}

// isNumber determines if the given rune is part of the numerical runes.
func isNumber(r rune) bool {
  return '0' <= r && r <= '9'
}

// toLower returns the lowercased permutation of the provided rune.
func toLower(r rune) rune {
  lower := strings.ToLower(string(r))[0]
  return rune(lower)
}

// Clean returns a string which is "cleanly formatted".
// Alphanumeric characters are generally maintained.
// Punctuation is substituted for underscores.
func Clean(str string) string {
  var sb strings.Builder  

  var R rune

  for _, r := range str {

    if isAlpha(r) {
      // Transform alphabet character to lowercase
      R = toLower(r)
    } else if isNumber(r) {
      // Keep original rune if it is a number
      R = r
    // All branches from here are non-alphanum cases
    } else if R != '_' {
      // Include underscore if underscore was not the previous rune
      R = '_'
    } else {
      // Do not include multiple unbroken underscores, leaving one underscore
      continue
    }

    sb.WriteRune(R)
  }

  newStr := sb.String()
  trimmed := strings.Trim(newStr, "_")

  return trimmed
}

// CleanSlice returns the result of cleanly-formatting then joining a slice of strings.
// Each string in the slice is first cleanly-formatted.
// All strings are then joined using "+".
func CleanSlice(strs []string) string {
  newStrs := make([]string, len(strs))

  for i, str := range strs {
    newStrs[i] = Clean(str)
  }

  return strings.Join(newStrs, "+")
}

// CleanMap returns the result of cleanly-formatting then joining a map of strings.
// Each key and value in the slice is first cleanly-formatted.
// All key-value pairs are then joined using "-".
// All pairs are then joined using "+".
func CleanMap(strs map[string]string) string {
  newStrs := make([]string, len(strs))

  i := 0

  for key, value := range strs {
    cleanKey := Clean(key)
    cleanValue := Clean(value)
    newStrs[i] = cleanKey + "-" + cleanValue
    i++
  }

  return strings.Join(newStrs, "+")
}
