package orderedform

import (
	"net/url"
	"strings"
)

// Form is a form that maintains its correct order as new key value pairs are added.
type Form [][2]string

// Add adds the specified key to the form with the specified value.
func (form *Form) Add(k, v string) {
	*form = append(*form, [2]string{
		url.QueryEscape(k),
		url.QueryEscape(v),
	})
}

// URLEncode encodes the values into "URL encoded" form ("bar=baz&foo=quux") in the correct order
// that Add was called.
func (form *Form) URLEncode() string {
	var builder strings.Builder
	for _, v := range *form {
		if builder.Len() > 0 {
			builder.WriteString("&")
		}

		builder.WriteString(v[0])
		builder.WriteString("=")
		builder.WriteString(v[1])
	}

	return builder.String()
}
