package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	tr, err := NewRates(hclog.Default(), "USD")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\nRates %+v\n", tr.rates)

	// c := 1
	// for k, _ := range tr.rates {
	// 	fmt.Printf("%s=%d;\n", k, c)
	// 	c++
	// }
}
