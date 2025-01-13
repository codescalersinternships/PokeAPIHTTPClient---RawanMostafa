package pokemon

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func assertEquality(t *testing.T, obj1 any, obj2 any) {
	t.Helper()
	if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
		t.Errorf("Error! type mismatch, wanted: %t \n\n got: %t", reflect.TypeOf(obj1), reflect.TypeOf(obj2))
	}
	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("Error! values mismatch, wanted: %v \n\n got: %v", obj1, obj2)
	}
}

func parseResource(t *testing.T, filePath string) (resource Resource) {
	t.Helper()
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("error opening json file %v", err)
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&resource); err != nil {
		t.Errorf("error parsing json file %v", err)
	}

	return
}

func parsePokemon(t *testing.T, filePath string) (pokemon Pokemon) {
	t.Helper()
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("error opening json file %v", err)
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&pokemon); err != nil {
		t.Errorf("error parsing json file %v", err)
	}

	return
}

func TestPokemon(t *testing.T) {
	expectedPokemon := parsePokemon(t, "../testdata/pokemon1.json")

	c := NewClient("/pokemon", time.Second)
	gotPokemon, err := c.Pokemon(1)
	if err != nil {
		t.Error(err)
	}
	assertEquality(t, expectedPokemon, gotPokemon)
}

func TestResource(t *testing.T) {

	testcases := []struct {
		name             string
		expectedJsonPath string
		endpoint         string
		options          []int
	}{
		{
			name:             "resources of ability, no options",
			expectedJsonPath: "../testdata/resource_ability.json",
			endpoint:         "/ability",
			options:          make([]int, 0),
		},
		{
			name:             "resources of ability, offset=5, limit=5",
			expectedJsonPath: "../testdata/resource_ability_offset_limit.json",
			endpoint:         "/ability",
			options:          []int{5, 5},
		},
		{
			name:             "resources of ability, offset=5",
			expectedJsonPath: "../testdata/resource_ability_offset.json",
			endpoint:         "/ability",
			options:          []int{5},
		},
		{
			name:             "resources of machine, no options",
			expectedJsonPath: "../testdata/resource_machine.json",
			endpoint:         "/machine",
			options:          make([]int, 0),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			expectedResource := parseResource(t, testcase.expectedJsonPath)

			c := NewClient(testcase.endpoint, time.Second)
			gotResource, err := c.Resource(testcase.options...)
			if err != nil {
				t.Error(err)
			}
			assertEquality(t, expectedResource, gotResource)
		})
	}

}
