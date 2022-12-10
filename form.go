package orderedform

import (
	"net/url"
	"strings"
)

// Entry represents an entry in a Form.
type Entry [2]string

// Key returns the entry's key.
func (e Entry) Key() string {
	return e[0]
}

// Value returns the entry's value.
func (e Entry) Value() string {
	return e[1]
}

// Form is a form that maintains its correct order as new key value pairs are added.
type Form []Entry

// Add adds the specified key to the form with the specified value.
func (form *Form) Add(k, v string) {
	*form = append(*form, Entry{k, v})
}

// Encoder is a function capable of encoding an entry.
type Encoder func(builder *strings.Builder, entry Entry)

// Encode encodes Form using the given encoder function.
//
// An Encoder is responsible for encoding an Entry to text format
// and PlaintextEncoder is a reasonable choice if you want the form
// encoded "as-is". If you want to URL encode the keys and values,
// use URLEncoder.
func (form *Form) Encode(encoder Encoder) string {
	var builder strings.Builder
	for _, entry := range *form {
		if builder.Len() > 0 {
			builder.WriteString("&")
		}

		encoder(&builder, entry)
	}

	return builder.String()
}

// PlaintextEncoder is an Encoder that encodes an Entry without modification.
func PlaintextEncoder(builder *strings.Builder, entry Entry) {
	builder.WriteString(entry.Key())
	builder.WriteString("=")
	builder.WriteString(entry.Value())
}

// URLEncoder is an Encoder that encodes an Entry as URL-encoded values.
func URLEncoder(builder *strings.Builder, entry Entry) {
	builder.WriteString(url.QueryEscape(entry.Key()))
	builder.WriteString("=")
	builder.WriteString(url.QueryEscape(entry.Value()))
}
