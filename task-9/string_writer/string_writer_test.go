package string_writer

import (
	"bytes"
	"testing"
)

func TestPrintStrings(t *testing.T) {
	type args struct {
		asgs []any
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "Works for slice",
			args: args{
				[]any{
					true,
					"W",
					14,
					[]int{},
					60069.15,
				}},
			wantW: "W",
		},
		{
			name:  "Works for empty",
			args:  args{[]any{}},
			wantW: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			PrintStrings(w, tt.args.asgs...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintStrings() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
