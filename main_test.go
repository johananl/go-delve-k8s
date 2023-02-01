package main

import "testing"

func TestGenRandomString(t *testing.T) {
	testCases := []struct {
		in      int
		wantErr bool
	}{
		{
			in: 1,
		},
		{
			in: 2,
		},
		{
			in: -1,
		},
		{
			in:      0,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		_, err := genRandomString(tc.in)

		if err != nil && !tc.wantErr {
			t.Fatalf("unexpected error: %v", err)
		}

		if err == nil && tc.wantErr {
			t.Fatal("expected an error")
		}
	}
}
