package assertion

import "net/url"

// URL returns a URL parsed from the string and panics if it fails to parse.
func URL(str string) *url.URL {
	parsed, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return parsed
}
