package server_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benchttp/server"
)

func TestLimitBytesReader(t *testing.T) {
	testcases := []struct {
		label   string
		maxSize int64
		body    []byte
		expSize int
		expErr  string
	}{
		{
			label:   "do not limit body below max size",
			body:    []byte("a"),
			maxSize: 3,
			expSize: 1,
			expErr:  "",
		},
		{
			label:   "do not limit body equal to max size",
			body:    []byte("abc"),
			maxSize: 3,
			expSize: 3,
			expErr:  "",
		},
		{
			label:   "limit body above max size",
			body:    []byte("abcdef"),
			maxSize: 3,
			expSize: 3,
			expErr:  "http: request body too large",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.label, func(t *testing.T) {
			rr := httptest.NewRecorder()
			rq := httptest.NewRequest("", "/", ioReader(tc.body))

			var (
				gotSize int
				gotErr  string
			)

			server.LimitBytesReader(tc.maxSize)(
				http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
					b, err := io.ReadAll(r.Body)
					if err != nil {
						gotErr = err.Error()
					}
					gotSize = len(b)
				}),
			).ServeHTTP(rr, rq)

			if gotSize != tc.expSize {
				t.Errorf("body size: exp %d, got %d", tc.expSize, gotSize)
			}

			if gotErr != tc.expErr {
				t.Errorf("read error: exp %q, got %q", tc.expErr, gotErr)
			}
		})
	}
}

func ioReader(b []byte) io.Reader {
	return bytes.NewReader(b)
}
