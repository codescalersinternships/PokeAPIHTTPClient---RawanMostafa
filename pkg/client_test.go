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
		t.Errorf("Error! type mismatch, wanted: %t got: %t", reflect.TypeOf(obj1), reflect.TypeOf(obj2))
	}
	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("Error! values mismatch, wanted: %v got: %v", obj1, obj2)
	}
}

func parseJson(t *testing.T, filePath string) (pokemon Pokemon) {
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
	expectedPokemon:=parseJson(t,"../testdata/pokemon1.json")

	timeout := time.Second
	c := NewClient("/pokemon", timeout)
	gotPokemon, err := c.Pokemon(1)
	if err != nil {
		t.Error(err)
	}
	assertEquality(t,expectedPokemon,gotPokemon)
}
