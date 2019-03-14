
[![CircleCI](https://circleci.com/gh/tkanos/go-dtree.svg?style=svg)](https://circleci.com/gh/tkanos/go-dtree)
[![Go Report Card](https://goreportcard.com/badge/github.com/tkanos/go-dtree)](https://goreportcard.com/report/github.com/tkanos/go-dtree)
[![Coverage Status](https://coveralls.io/repos/github/tkanos/go-dtree/badge.svg?branch=master)](https://coveralls.io/github/tkanos/go-dtree?branch=master)

<img src="./docs/images/dtree-logo-name.png" height=251>


## GO-DTREE - Golang package for resolve decisions Tree

DTree allow you to define an Decision Tree in json

```json
{
				"id": 1,
				"name": "Root"
			},
			{
				"id": 2,
				"name": "IsFirstTree",
                "parent_id": 1,
                "key":"isFirstTree",
				"operator": "eq",
				"value": true
			},
			{
				"id": 3,
				"name": "IsNotTheFirstTree",
                "parent_id": 1,
                "key":"isFirstTree",
				"operator": "eq",
				"value": false
            },
            {
				"id": 4,
				"name": "Welcome",
                "parent_id": 2,
            },
            {
				"id": 5,
				"name": "Congrats",
                "parent_id": 3,
			}

```

loaded 

```golang
  tree, err :=dtree.LoadTree([]byte(jsonTree))
```

it will create :

<img src="./docs/images/first-tree.png" height=60%>

If you want to programmaticaly build your tree, you can also use the CreateTree Method.

```golang
var myTree []Tree
// append your nodes on myTree and then
tree :=dtree.CreateTree(myTree)
```

Then we can resolve the decision Tree by passing another json, representing the needed value.  

```golang
    request := []byte(`{
		"isFirstTree":     true
	}`)

    node, _ := t.ResolveJSON(request)

    fmt.Println(node.Name)
    // Output: Welcome
```

you can also define it programmatically,

```golang
    request := make(map[string]interface{}) 
    request["isFirstTree"] = true

    node, _ := t.Resolve(request)

    fmt.Println(node.Name)
    // Output: Welcome
```

but in this case be careful, to don't use int (not supported), only floats.

This one was a simple decision Tree. You can build more complexe with more nodes, with others operators than only equal.

## Available Operators :
| operator       | description                                                                         |
| -------------- | ----------------------------------------------------------------------------------- |
| eq (or ==)     | equality (for string, bool, numbers, arrays)                                        |
| ne (or !=)     | not equal (for string, bool, numbers, arrays)                                       |
| gt (or >)      | gt (for string, numbers)                                                            |
| lt (or <)      | lt (for string, numbers)                                                            |
| gte (or >=)    | gte (for string, numbers)                                                           |
| lte (or <=)    | lte (for string, numbers)                                                           |
| contains       | does the string (defined on the value of the Tree) is contained on the json request |
| count          | count (only for arrays)                                                             |
| regexp         | do a regexp (only for string)                                                       |
| percent (or %) | do a random selection based on percentages                                          |

You can also define your own operators 

## string comparition

All the strings are case compared (without ToLower before the comparition)

```
s1 == s2
```

if you are not sure about the data, and you want to do ToLower comparition, if you two choice:
- Write your own operator
- Lower your tree data and json request, before to call Resolve.


## Custom operators 

You can define your own custom operators (on the example I do a len of an array pass on request, and i check if it matches the Value of the path of the node on the Tree (it is the code of the already implemented operator "count"))

```golang
f := func(t *TreeOptions) {
    t.Operators = make(map[string]dtree.Operator)
    t.Operators["len"] = func(requests map[string]interface{}, node *Tree) (*Tree, error) {
        if v1, ok := requests[node.Key]; ok {
            switch t1 := v1.(type) {
            case []interface{}:
                if t2, ok := node.Value.(float64); ok {
                    if len(t1) == int(t2) {
                        return node, nil
                    }
                    return nil, nil
                }

                return nil, ErrBadType
            default:
                return nil, ErrNotSupportedType
            }
        }
        return nil, nil
    }
}
```
Of course if you have really good operators that you want to add to DTREE, does not hesitate to do a PR.

## Options :

By default if dtree cannot resolve one node (because bad parameters), we consider this node as false, and it continues.
We can set the option StopIfConvertingError at true, on this case dtree will stop once it found an parsing error.

```golang
f := func(t *TreeOptions) {
    t.StopIfConvertingError = true
}
```

We can also define a fallback value. It means on this case, that if all others path are in false, it goes to this one.

```json
"value": "fallback"
```

We can also set an order, to define the order of the evaluation (but of course fallback will always be the last (even if you don't say so))

```json
"order": "1"
```

The Node Tree as a parameter content, it's a interface{}, that allow you to put whatever you want.

## Context :

You can give a context to the Tree, is mostly used for debugging, like this you will be able to know what are the path that your request takes inside the tree.

```golang
t.WithContext(context.Background())

v, _ := t.Resolve(request)

sliceOfString := dtree.GetNodePathFromContext(t.Context())
fmt.Println(sliceOfString)

// sliceofstring contains is a slice where each string is a node on the format `id : key value operator expectedvalue`
// example : 3 : productid 1234 gt 1230
```


