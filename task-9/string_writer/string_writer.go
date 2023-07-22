package string_writer

import (
	"io"
)

func PrintStrings(w io.Writer, asgs ...any) {
	for _, a := range asgs {
		if s, ok := a.(string); ok {
			w.Write([]byte(s))
		}
	}
}
