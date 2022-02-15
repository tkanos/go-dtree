package main

import (
	"fmt"

	"github.com/tkanos/go-dtree"
)

func main() {
	data := []dtree.Tree{
		{
			ID:       1,
			Name:     "1",
			ParentID: 0,
			Value:    true,
			Operator: "eq",
			Key:      "1",
			Legacy: map[string][]interface{}{
				"cost": {10},
				"test": {"testValue"},
			},
		},
		{
			ID:       2,
			Name:     "2",
			ParentID: 1,
			Value:    int64(12),
			Operator: "eq",
			Key:      "2",
			Legacy: map[string][]interface{}{
				"cost": {20},
			},
		},
		{
			ID:       3,
			Name:     "3",
			ParentID: 1,
			Value:    12,
			Operator: "eq",
			Key:      "3",
			Legacy: map[string][]interface{}{
				"cost": {50},
			},
		},
	}

	tree := dtree.CreateTree(data)

	request := map[string]interface{}{
		"1": true,
		"2": int64(12),
	}

	resolve, err := tree.Resolve(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(tree)

	// output: legacy map[cost:[20 10] test:[testValue]]
	fmt.Println("legacy", resolve.Legacy)
}
