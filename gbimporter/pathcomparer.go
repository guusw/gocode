package gbimporter

import "strings"

// PathComparer compares 2 paths for equality based on their strings, the paths should already have already been cleaned
type PathComparer interface {
	Compare(a string, b string) bool
}

type pathComparerNormal struct{}

func (s pathComparerNormal) Compare(a string, b string) bool {
	return a == b
}

type pathComparerCaseInsensitive struct{}

func (s pathComparerCaseInsensitive) Compare(a string, b string) bool {
	return strings.EqualFold(a, b)
}

func NewPathComparer(ctx *PackedContext) PathComparer {
	switch ctx.GOOS {
	case "windows":
		return pathComparerCaseInsensitive{}
	default:
		return pathComparerNormal{}
	}
}
