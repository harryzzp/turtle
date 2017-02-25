package freejson

import (
	"encoding/json"
	"fmt"
)

func FreeType(theJson string) interface{} {

	var anything interface{}
	json.Unmarshal([]byte(theJson), &anything)
	var ft interface{}
	switch v := anything.(type) {
	case float64:
		fmt.Printf("NUMBER: %f\n", v)
		ft = v
	case string:
		fmt.Printf("STRING: %s\n", v)
		ft = v
	case map[string]interface{}:
		fmt.Printf("STRING MAP: %s\n", v)
		ft = v
	case []interface{}:
		fmt.Printf("STRING arrays: %s\n", v)
		ft = v
	default:
		panic("I don't know how to handle this!\n")
	}
	return ft
}
