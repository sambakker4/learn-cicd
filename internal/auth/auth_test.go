package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		header string
		str    string
		err    error
	} {
		{
			header: "",
			str: "",
			err: ErrNoAuthHeaderIncluded,
		},
		{
			header: "ApiKey asdf",
			str: "asdf",
			err: nil,
		},
		{
			header: "asdf",
			str: "",
			err: MalformedHeader,
		},
		{
			header: "asdf sdfasf",
			str: "",
			err: MalformedHeader,
		},
	}

	for _, testcase := range cases {
		header := http.Header{}
		header.Set("Authorization", testcase.header)
		apiKey, err := GetAPIKey(header)
		if apiKey != testcase.str {
			t.Errorf("Expected output of %v but got: %v", testcase.str, apiKey)
		}
		
		if err != testcase.err {
			t.Errorf("Expected error of %v but got: %v", testcase.err, err)
		}

		
		if err == nil && testcase.err != nil {
			t.Errorf("Expected an error but got none")
		}

		if err != nil && testcase.err == nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	}
}
