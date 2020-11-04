package data

import (
    "fmt"
    "testing"

    "github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
    tr, err := NewRates(hclog.Default()
    
    if err !=nil {
        t.fatal(err)

    }
    fmt.Sprintf("%#v", tr.rates)
}