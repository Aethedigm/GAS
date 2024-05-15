package tests

import (
	"main/gas"
	"testing"
)

func Test_insert(t *testing.T) {
	gasObj := new(gas.GAS)

	gasObj.AddResult("key", "value")
	res := gasObj.GetResults("key")

	for x := range res.Results {
		if res.Results[x].Word == "key" {
			for y := range res.Results[x].Values {
				if res.Results[x].Values[y] == "value" {
					return
				} else {
					t.Error("Non-inserted value found")
				}
			}
		} else {
			t.Error("Non-inserted value found")
		}
	}

	t.Error("Should have found the key and value")
}
