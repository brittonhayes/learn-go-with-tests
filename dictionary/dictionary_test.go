package dictionary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	searchTests := []struct {
		name       string
		dictionary Dictionary
		key        string
		expected   string
		wantErr    bool
	}{
		{"Find key in dictionary", Dictionary{"test": "this is just a test"}, "test", "this is just a test", false},
		{"Return error missing", Dictionary{"hidden": "can't find me"}, "shown", "this is just a test", true},
	}

	for _, tt := range searchTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.dictionary.Search(tt.key)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAdd(t *testing.T) {
	addTests := []struct {
		name       string
		dictionary Dictionary
		word       string
		definition string
	}{
		{"Find key in dictionary", Dictionary{"test": "this is just a test"}, "new", "a new entry"},
	}

	for _, tt := range addTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dictionary.Add(tt.word, tt.definition)
			actual, err := tt.dictionary.Search(tt.word)
			if err != nil {
				t.Fatalf("did not expect error but received %s", err.Error())
			}

			assert.Equal(t, tt.definition, actual)
		})
	}
}
