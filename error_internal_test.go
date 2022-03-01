package server

import (
	"errors"
	"testing"
)

func TestHttpErrorOf(t *testing.T) {
	for _, testcase := range []struct {
		name string
		err  error
		want *httpError
	}{
		{
			name: "httpError into httpError",
			err:  &httpError{Message: "woops", Code: 418},
			want: &httpError{Message: "woops", Code: 418},
		},
		{
			name: "error into httpError",
			err:  errors.New("woops"),
			want: &httpError{Message: "Internal Server Error", Code: 500},
		},
		{
			name: "nil into httpError",
			err:  nil,
			want: &httpError{Message: "Internal Server Error", Code: 500},
		},
		{
			name: "httpError into httpError (keep inner error)",
			err:  &httpError{Message: "woops", Code: 418, inner: errors.New("woops")},
			want: &httpError{Message: "woops", Code: 418, inner: errors.New("woops")},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := httpErrorOf(testcase.err)

			if testcase.want.Code != got.Code {
				t.Errorf("incorrect conversion: httpError.Code, want %d, got %d", testcase.want.Code, got.Code)
			}

			if testcase.want.Message != got.Message {
				t.Errorf("incorrect conversion: httpError.Message, want %s, got %s", testcase.want.Message, got.Message)
			}

			if testcase.want.inner != nil && got.inner != nil && testcase.want.inner == got.inner {
				t.Errorf("incorrect conversion: httpError.inner, want %+v, got %+v", testcase.want.inner, got.inner)
			}
		})
	}
}
