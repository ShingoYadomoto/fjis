package main

import (
	"bytes"
	"io"
	"testing"
)

func Test_echoResult(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "",
			args: args{r: bytes.NewBuffer([]byte("¥aaa\n"))},
			wantW: "" +
				"\u001B[31m¥\u001B[0maaa\n" +
				"============\n" +
				"[U+00A5 '¥']\n",
		},
		{
			name: "",
			args: args{r: bytes.NewBuffer([]byte("bbb\n"))},
			wantW: "bbb\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			echoResult(tt.args.r, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echoResult() = %v, want %v", gotW, tt.wantW)
			}
      ngmap = map[string]int{} 
		})
	}
}

func Test_echoHighlighted(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{r: bytes.NewBuffer([]byte("¥aaa\n"))},
			wantW:   "\u001B[31m¥\u001B[0maaa\n",
			wantErr: false,
		},
		{
			name:    "",
			args:    args{r: bytes.NewBuffer([]byte("bbb\n"))},
			wantW:   "bbb\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := echoHighlighted(tt.args.r, w)
			if (err != nil) != tt.wantErr {
				t.Errorf("echoHighlighted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echoHighlighted() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_echoUnicodeFormat(t *testing.T) {
	tests := []struct {
		name    string
		ngmap   map[string]int
		wantW   string
		wantErr bool
	}{
		{
			name:  "",
			ngmap: map[string]int{},
			wantW: "",
		},
		{
			name:  "",
			ngmap: map[string]int{"¥": 0},
			wantW: "[U+00A5 '¥']\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			ngmap = tt.ngmap
			err := echoUnicodeFormat(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("echoUnicodeFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echoUnicodeFormat() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
