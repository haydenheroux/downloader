package downloader

import (
	"errors"
	"fmt"
)

func missingDependencyError(dependency string) error {
	return fmt.Errorf("dependency %s is missing", dependency)
}

func unavailableError() error {
	return errors.New("content is unavailable")
}

func ageRestricted() error {
	return errors.New("content is age-restricted")
}
