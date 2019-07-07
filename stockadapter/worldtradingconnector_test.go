package stockadapter

import (
	"testing"
)

func TestDoRequest(t *testing.T) {
	cases := []struct {
		Symbol        string
		ExpectedError bool
	}{
		{"AAPL",
			false},
		{
			"MSFT",
			false,
		},
		{
			"AMEX",
			true,
		},
	}
	for _, tc := range cases {
		data, err := doRequest(tc.Symbol)
		if err != nil && tc.ExpectedError {
			t.Log("got error ", err)
		} else if err != nil && tc.ExpectedError == false {
			t.Fatal("failed", err)
		}
		if err == nil && len(data) > 0 {
			t.Log("working fine ")
		}
	}
}
