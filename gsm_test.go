package gsm

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoad(t *testing.T) {

	tests := map[string]struct {
		key       int
		wantValue string
		wantOK    bool
	}{
		"successful load":   {key: 1, wantValue: "one", wantOK: true},
		"unsuccessful load": {key: 2, wantValue: "", wantOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")

			value, ok := gsm.Load(tc.key)

			diff := cmp.Diff(tc.wantValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestClear(t *testing.T) {

	tests := map[string]struct {
		key       int
		wantValue string
		wantOK    bool
	}{
		"unsuccessful load": {key: 1, wantValue: "", wantOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")

			gsm.Clear()

			value, ok := gsm.Load(tc.key)

			diff := cmp.Diff(tc.wantValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestLoadOrStore(t *testing.T) {

	tests := map[string]struct {
		key                    int
		value                  string
		wantActual             string
		wantLoaded             bool
		afterLoadOrStoreValue  string
		afterLoadOrStoreLoadOK bool
	}{
		"existing key": {key: 1, value: "notOne", wantActual: "one", wantLoaded: true, afterLoadOrStoreValue: "one", afterLoadOrStoreLoadOK: true},
		"new key":      {key: 2, value: "two", wantActual: "", wantLoaded: false, afterLoadOrStoreValue: "two", afterLoadOrStoreLoadOK: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")

			actual, loaded := gsm.LoadOrStore(tc.key, tc.value)

			diff := cmp.Diff(tc.wantActual, actual)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantLoaded, loaded)
			if diff != "" {
				t.Fatal(diff)
			}

			afterLoadOrStoreValue, afterLoadOrStoreLoadOK := gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterLoadOrStoreValue, afterLoadOrStoreValue)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterLoadOrStoreLoadOK, afterLoadOrStoreLoadOK)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}

}

func TestLoadAndDelete(t *testing.T) {

	tests := map[string]struct {
		key               int
		wantValue         string
		wantLoaded        bool
		afterDeleteValue  string
		afterDeleteLoadOK bool
	}{
		"existing key": {key: 1, wantValue: "one", wantLoaded: true, afterDeleteValue: "", afterDeleteLoadOK: false},
		"unknown key":  {key: 2, wantValue: "", wantLoaded: false, afterDeleteValue: "", afterDeleteLoadOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")

			value, loaded := gsm.LoadAndDelete(tc.key)

			diff := cmp.Diff(tc.wantValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantLoaded, loaded)
			if diff != "" {
				t.Fatal(diff)
			}

			afterDeleteValue, afterDeleteLoadOK := gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterDeleteValue, afterDeleteValue)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterDeleteLoadOK, afterDeleteLoadOK)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}

}

func TestDelete(t *testing.T) {

	tests := map[string]struct {
		key               int
		initialValue      string
		initialLoadOK     bool
		afterDeleteValue  string
		afterDeleteLoadOK bool
	}{
		"first load": {key: 1, initialValue: "one", initialLoadOK: true, afterDeleteValue: "", afterDeleteLoadOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")

			value, ok := gsm.Load(tc.key)

			diff := cmp.Diff(tc.initialValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.initialLoadOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}

			gsm.Delete(tc.key)

			afterDeleteValue, afterDeleteLoadOK := gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterDeleteValue, afterDeleteValue)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterDeleteLoadOK, afterDeleteLoadOK)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestSwap(t *testing.T) {

	tests := map[string]struct {
		key             int
		value           string
		wantPrevious    string
		wantLoaded      bool
		afterSwapValue  string
		afterSwapLoadOK bool
	}{
		"existing key": {key: 1, value: "updatedOne", wantPrevious: "one", wantLoaded: true, afterSwapValue: "updatedOne", afterSwapLoadOK: true},
		"unknown key":  {key: 2, value: "updatedTwo", wantPrevious: "", wantLoaded: false, afterSwapValue: "updatedTwo", afterSwapLoadOK: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")

			previous, loaded := gsm.Swap(tc.key, tc.value)

			diff := cmp.Diff(tc.wantPrevious, previous)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantLoaded, loaded)
			if diff != "" {
				t.Fatal(diff)
			}

			afterSwapValue, afterSwapLoaded := gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterSwapValue, afterSwapValue)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterSwapLoadOK, afterSwapLoaded)
			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

func TestCompareAndSwap(t *testing.T) {

	tests := map[string]struct {
		key             int
		oldValue        string
		newValue        string
		wantSwapped     bool
		afterSwapValue  string
		afterSwapLoadOK bool
	}{
		"existing key and value":              {key: 1, oldValue: "one", newValue: "updatedOne", wantSwapped: true, afterSwapValue: "updatedOne", afterSwapLoadOK: true},
		"existing key not equal value":        {key: 1, oldValue: "badone", newValue: "updatedOne", wantSwapped: false, afterSwapValue: "one", afterSwapLoadOK: true},
		"not existing key not existing value": {key: 2, oldValue: "test", newValue: "updatedTest", wantSwapped: false, afterSwapValue: "", afterSwapLoadOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			// Store a new value at beginning of each test
			gsm.Store(1, "one")

			swapped := gsm.CompareAndSwap(tc.key, tc.oldValue, tc.newValue)

			diff := cmp.Diff(tc.wantSwapped, swapped)
			if diff != "" {
				t.Fatal(diff)
			}

			afterSwapValue, afterSwapLoaded := gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterSwapValue, afterSwapValue)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterSwapLoadOK, afterSwapLoaded)
			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

func TestCompareAndDelete(t *testing.T) {

	tests := map[string]struct {
		key               int
		oldValue          string
		wantDeleted       bool
		afterDeleteValue  string
		afterDeleteLoadOK bool
	}{
		"existing key and value":           {key: 1, oldValue: "one", wantDeleted: true, afterDeleteValue: "", afterDeleteLoadOK: false},
		"existing key and non-equal value": {key: 1, oldValue: "badone", wantDeleted: false, afterDeleteValue: "one", afterDeleteLoadOK: true},
		"non existing key":                 {key: 2, oldValue: "", wantDeleted: false, afterDeleteValue: "", afterDeleteLoadOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			// Store a new value at beginning of each test
			gsm.Store(1, "one")

			deleted := gsm.CompareAndDelete(tc.key, tc.oldValue)

			diff := cmp.Diff(tc.wantDeleted, deleted)
			if diff != "" {
				t.Fatal(diff)
			}

			afterDeleteValue, afterDeleteLoadOK := gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterDeleteValue, afterDeleteValue)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterDeleteLoadOK, afterDeleteLoadOK)
			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

func TestRange(t *testing.T) {

	tests := map[string]struct {
		name         string
		wantRangeKVs map[int]string
	}{
		"range": {
			wantRangeKVs: map[int]string{
				1: "one",
				2: "two",
				3: "three",
				4: "four",
				5: "five",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")
			gsm.Store(2, "two")
			gsm.Store(3, "three")
			gsm.Store(4, "four")
			gsm.Store(5, "five")

			rangeKVs := make(map[int]string)

			for key, value := range gsm.Range() {
				rangeKVs[key] = value
			}

			diff := cmp.Diff(tc.wantRangeKVs, rangeKVs)
			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

func TestKeys(t *testing.T) {

	tests := map[string]struct {
		name           string
		wantSortedKeys []int
	}{
		"range": {wantSortedKeys: []int{1, 2, 3, 4, 5}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")
			gsm.Store(2, "two")
			gsm.Store(3, "three")
			gsm.Store(4, "four")
			gsm.Store(5, "five")

			var sortedKeys []int

			for key := range gsm.Keys() {
				sortedKeys = append(sortedKeys, key)
			}

			slices.Sort(sortedKeys)

			diff := cmp.Diff(tc.wantSortedKeys, sortedKeys)
			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

func TestValues(t *testing.T) {

	wantSortedValues := []string{"one", "two", "three", "four", "five"}
	slices.Sort(wantSortedValues)

	tests := map[string]struct {
		name             string
		wantSortedValues []string
	}{
		"range": {wantSortedValues: wantSortedValues},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var gsm GenericSyncMap[int, string]

			gsm.Store(1, "one")
			gsm.Store(2, "two")
			gsm.Store(3, "three")
			gsm.Store(4, "four")
			gsm.Store(5, "five")

			var sortedValues []string

			for key := range gsm.Values() {
				sortedValues = append(sortedValues, key)
			}

			slices.Sort(sortedValues)

			diff := cmp.Diff(tc.wantSortedValues, sortedValues)
			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}
