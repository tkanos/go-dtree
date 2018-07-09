package main

import (
	"fmt"
	"os"

	"github.com/tkanos/go-dtree"
)

func main() {
	jsonTree := []byte(`[
			{
				"id": 1,
				"name": "sayHello"
			},
			{
				"id": 2,
				"name": "GoodBye",
				"parent_id": 1,
				"operator": "eq",
				"value": false
			},
			{
				"id": 3,
				"name": "gender",
				"parent_id": 1,
				"operator": "eq",
				"value": true
			},
			{
				"id": 4,
				"name": "Hello Miss",
				"parent_id": 3,
				"operator": "eq",
				"value": "F"
			},
			{
				"id": 5,
				"name": "Hello",
				"parent_id": 3,
				"value": "fallback"
			},
			{
				"id": 6,
				"name": "age",
				"parent_id": 3,
				"operator": "eq",
				"value": "M"
			},
			{
				"id": 7,
				"name": "Hello Sir",
				"parent_id": 6,
				"operator": "gt",
				"value": 60
			},
			{
				"id": 8,
				"name": "Hello dude",
				"parent_id": 6,
				"operator": "lte",
				"value": 60
			}
		]`)

	t, err := dtree.LoadTree(jsonTree)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	request := make(map[string]interface{})
	request["sayHello"] = true
	request["gender"] = "M"
	request["age"] = 35.0 //does not use int, the engine only support float (if you want  do a PR to include int, it's up to you)

	/*request := []byte(`{
			"sayHello": false,
			"gender":   "M",
			"age": 35
		}`)

	v, _ := t.ResolveJSON(request)
	*/

	v, _ := t.Resolve(request)

	fmt.Println(v.Name)
	// output : Hello dude
}
