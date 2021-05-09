package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type (
	errReader struct{}
	errWriter struct{}
)

func (r errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}

func (r errWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write error")
}

func (r errWriter) String() string {
	return ""
}

func Test_echoResult(t *testing.T) {
	type args struct {
		r io.Reader
		w io.Writer
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "",
			args: args{
				r: bytes.NewBuffer([]byte("¥aaa\n")),
				w: &bytes.Buffer{},
			},
			wantW: "" +
				"\u001B[31m¥\u001B[0maaa\n" +
				"============\n" +
				"[U+00A5 '¥']\n",
		},
		{
			name: "",
			args: args{
				r: bytes.NewBuffer([]byte("bbb\n")),
				w: &bytes.Buffer{},
			},
			wantW: "bbb\n",
		},
		{
			name: "",
			args: args{
				r: errReader{},
				w: &bytes.Buffer{},
			},
			wantW: "",
		},
		{
			name: "",
			args: args{
				r: bytes.NewBuffer([]byte("bbb\n")),
				w: errWriter{},
			},
			wantW: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echoResult(tt.args.r, tt.args.w)
			w, _ := tt.args.w.(fmt.Stringer)
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
		w io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				r: bytes.NewBuffer([]byte("¥aaa\n")),
				w: &bytes.Buffer{},
			},
			wantW:   "\u001B[31m¥\u001B[0maaa\n",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				r: bytes.NewBuffer([]byte("bbb\n")),
				w: &bytes.Buffer{},
			},
			wantW:   "bbb\n",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				r: errReader{},
				w: &bytes.Buffer{},
			},
			wantW:   "",
			wantErr: true,
		},
		{
			name: "",
			args: args{
				r: bytes.NewBuffer([]byte("¥aaa\n")),
				w: errWriter{},
			},
			wantW:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := echoHighlighted(tt.args.r, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("echoHighlighted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			w, _ := tt.args.w.(fmt.Stringer)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echoHighlighted() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_echoUnicodeFormat(t *testing.T) {
	type args struct {
		w io.Writer
	}
	tests := []struct {
		name    string
		args    args
		ngmap   map[string]int
		wantW   string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{w: &bytes.Buffer{}},
			ngmap:   map[string]int{},
			wantW:   "",
			wantErr: false,
		},
		{
			name:    "",
			args:    args{w: &bytes.Buffer{}},
			ngmap:   map[string]int{"¥": 0},
			wantW:   "[U+00A5 '¥']\n",
			wantErr: false,
		},
		{
			name:    "",
			args:    args{w: errWriter{}},
			ngmap:   map[string]int{"¥": 0},
			wantW:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ngmap = tt.ngmap
			err := echoUnicodeFormat(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("echoUnicodeFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			w, _ := tt.args.w.(fmt.Stringer)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echoUnicodeFormat() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
