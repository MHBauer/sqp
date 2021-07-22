package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

func isDifferent(a, b map[string]interface{}) bool {

	for k, v := range a {
		bv, exists := b[k]
		//fmt.Println(k, v, bv)
		if !exists {
			return true
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			bvv, ok := bv.(map[string]interface{})
			if !ok { // type mismatch, different
				return true
			}
			if isDifferent(vv, bvv) {
				return true
			}
		default:
			if v != bv {
				return true
			}
		}

	}

	return false
}
