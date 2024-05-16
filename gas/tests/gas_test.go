package tests

import (
	"main/gas"
	"os"
	"testing"
)

var gasObj *gas.GAS

func TestMain(m *testing.M) {
	gasObj = new(gas.GAS)

	code := m.Run()
	os.Exit(code)
}

func Test_insert(t *testing.T) {
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

func Test_retrieve(t *testing.T) {
	gasObj.AddResult("retrieve", "true")

	rr := gasObj.GetResults("retrieve")
	if len(rr.Results) < 1 {
		t.Error("Should have found the key")
	}

	for x := range rr.Results {
		if rr.Results[x].Word == "retrieve" {
			for y := range rr.Results[x].Values {
				if rr.Results[x].Values[y] == "true" {
					return
				}
			}
		}
	}

	t.Error("Should have found the value: retrieve")
}
