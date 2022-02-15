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
				"value": true
			},
			{
				"id": 3,
				"parent_id": 1,
				"key": "sayHello",
				"operator": "eq",
				"value": false
			},
			{
				"id": 4,
				"parent_id": 3,
				"Name": "Goodbye"
			},
			{
				"id": 5,
				"parent_id": 2,
				"key": "gender",
				"operator": "eq",
				"value": "F"
			},
			{
				"id": 6,
				"parent_id": 5,
				"Name": "Hello Miss"
			},
			{
				"id": 7,
				"parent_id": 2,
				"value": "fallback"
			},
			{
				"id": 8,
				"parent_id": 7,
				"Name": "Hello"
			},
			{
				"id": 9,
				"parent_id": 2,
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
				"Name": "Hello Sir"
			},
			{
				"id": 12,
				"parent_id": 9,
				"key": "age",
				"operator": "lte",
				"value": 60,
				"Headers": {
					"year": 89
				}
			},
			{
				"id": 13,
				"parent_id": 12,
				"Name": "Hello dude",
				"Content": "mura",
				"Headers": {
					"name": "kalle"
				}
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

	fmt.Println(t)
	// output: Hello dude
	//                                ╭────╮
	//                                │root│
	//                                ╰──┬─╯
	//                         ╭─────────┴─────────────────────────╮
	//                ╭────────┴───────╮                  ╭────────┴────────╮
	//                │sayHello eq true│                  │sayHello eq false│
	//                ╰────────┬───────╯                  ╰────────┬────────╯
	//      ╭──────────────────┴┬──────────────────╮               │
	//╭─────┴─────╮       ╭─────┴─────╮       ╭────┴───╮       ╭───┴───╮
	//│gender eq F│       │gender eq M│       │fallback│       │Goodbye│
	//╰─────┬─────╯       ╰─────┬─────╯       ╰────┬───╯       ╰───────╯
	//      │             ╭─────┴──────╮           │
	//╭─────┴────╮  ╭─────┴────╮  ╭────┴────╮   ╭──┴──╮
	//│Hello Miss│  │age lte 60│  │age gt 60│   │Hello│
	//╰──────────╯  ╰─────┬────╯  ╰────┬────╯   ╰─────╯
	//                    │            │
	//              ╭─────┴────╮  ╭────┴────╮
	//              │Hello dude│  │Hello Sir│
	//              ╰──────────╯  ╰─────────╯
}
