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
				"cost": []interface{}{10},
				"test": []interface{}{
					"testValue",
				},
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
				"cost": []interface{}{20},
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

	// output: legacy map[cost:[20 10]]
	fmt.Println("legacy", resolve.Legacy)
}
