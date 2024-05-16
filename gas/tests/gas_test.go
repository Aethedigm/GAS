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

func cleanup(t *testing.T) {
	t.Cleanup(func() {
		gasObj = new(gas.GAS)
	})
}

func Test_retrieve_will_ignore_special_characters(t *testing.T) {
	gasObj.AddResult("main", "val")
	res := gasObj.GetResults("ma~in")

	for x := range res.Results {
		if res.Results[x].Word == "main" {
			return
		} else {
			t.Error("Should have found 'main'")
		}
	}

	t.Error("Should have found 'main'")
	cleanup(t)
}

func Test_insert_will_not_add_special_characters(t *testing.T) {
	gasObj.AddResult("specia!l", "a&b")
	res := gasObj.GetResults("special")

	for x := range res.Results {
		if res.Results[x].Word == "special" {
			for y := range res.Results[x].Values {
				if res.Results[x].Values[y] == "ab" {
					return
				} else {
					t.Error("Expected 'ab' got ", res.Results[x].Values[y])
				}
			}
		} else {
			t.Error("Non-inserted value found")
		}
	}

	t.Error("Should have found 'special'")
	cleanup(t)
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
	cleanup(t)
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
	cleanup(t)
}
