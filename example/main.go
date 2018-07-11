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
				"name": "root"
			},
			{
				"id": 2,
				"parent_id": 1,
				"key": "sayHello",
				"operator": "eq",
				"value": false
			},
			{
				"id": 3,
				"name": "GoodBye",
				"parent_id": 2,
				"value": "fallback"
			},
			{
				"id": 4,
				"parent_id": 1,
				"key": "sayHello",
				"operator": "eq",
				"value": true
			},
			{
				"id": 5,
				"parent_id": 4,
				"key": "gender",
				"operator": "eq",
				"value": "F"
			},
			{
				"id": 6,
				"name": "Hello Miss",
				"parent_id": 5,
				"value": "fallback"
			},
			{
				"id": 7,
				"parent_id": 4,
				"value": "fallback"
			},
			{
				"id": 8,
				"name": "Hello",
				"parent_id": 7,
				"value": "fallback"
			},
			{
				"id": 9,
				"parent_id": 4,
				"key": "gender",
				"operator": "eq",
				"value": "M"
			},
			{
				"id": 10,
				"parent_id": 9,
				"key": "age",
				"operator": "gt",
				"value": 60
			},
			{
				"id": 11,
				"parent_id": 10,
				"name": "Hello Sir",
				"value": "fallback"
			},
			{
				"id": 12,
				"parent_id": 9,
				"key": "age",
				"operator": "lte",
				"value": 60
			},
			{
				"id": 13,
				"parent_id": 12,
				"name": "Hello dude",
				"value": "fallback"
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
