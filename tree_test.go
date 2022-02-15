package dtree

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var treeTest = []byte(`[
	{
		"id": 1,
		"name": "root"
	},
	{
		"id": 2,
		"parent_id": 1,
		"key": "isTest",
		"operator": "eq",
		"value": false

	},
	{
		"id": 3,
		"name": "Never Reach",
		"parent_id": 2,
		"value": "fallback"
	},
	{
		"id": 4,
		"parent_id": 1,
		"key": "isTest",
		"operator": "eq",
		"value": true
	},
	{
		"id": 5,
		"parent_id": 4,
		"operator": "gt",
		"key": "count",
		"value": 10,
		"order":1
	},
	{
		"id": 6,
		"name": "FinalNode 2",
		"parent_id": 5,
		"value": "fallback"
	},
	{
		"id": 7,
		"parent_id": 4,
		"operator": "lt",
		"key": "count",
		"value": 10,
		"order":2
	},
	{
		"id": 8,
		"name": "FinalNode 1",
		"parent_id": 7,
		"value": "fallback"
	},
	{
		"id": 9,
		"name": "FinalNode 3",
		"parent_id": 4,
		"value": "fallback"
	}
]`)

func TestTree_SimpleTest(t *testing.T) {
	// Arrange

	//Load Tree
	tr, err := LoadTree(treeTest)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Load request
	jsonRequest := []byte(`{
		"isTest":  true,
		"count":   15
	}`)

	//Act
	result, err := tr.ResolveJSON(jsonRequest)

	//Assert
	assert.NoError(t, err, "Resolve should not have errors")
	assert.Equal(t, "FinalNode 2", result.Name)
}

func TestTree_SimpleTest_With_Error_Config(t *testing.T) {
	// Arrange

	//Load Tree
	tr, err := LoadTree(treeTest)
	if err != nil {
		t.Fail()
	}

	// Load request
	jsonRequest := []byte(`{
		"isTest":  true,
		"count":   "15"
	}`)

	f := func(t *TreeOptions) {
		t.StopIfConvertingError = true
	}

	//Act
	result, err := tr.ResolveJSON(jsonRequest, f)
	//Assert
	assert.Error(t, err, "Resolve should not return an error when the type of the request is the not the same as the one defined on tree")
	assert.Equal(t, "isTest", result.Key)
}

func TestTree_SimpleTest_Without_Error_Config(t *testing.T) {
	// Arrange

	//Load Tree
	tr, err := LoadTree(treeTest)
	if err != nil {
		t.Fail()
	}

	// Load request
	jsonRequest := []byte(`{
		"isTest":  true,
		"count":   "15"
	}`)

	f := func(t *TreeOptions) {
		t.StopIfConvertingError = false
	}

	//Act
	result, err := tr.ResolveJSON(jsonRequest, f)

	//Assert
	assert.NoError(t, err, "Resolve should not return an error even if the type of the request is the not the same as the one defined on tree")
	assert.Equal(t, "FinalNode 3", result.Name)
}

func TestTree_SimpleTest_With_Bad_Json(t *testing.T) {
	//Act
	_, err := LoadTree([]byte("not a json"))

	//Assert
	assert.Error(t, err, "LoadTree should return an error if the json is malformed")
}

func TestTree_SimpleTest_Resolving_Bad_Json(t *testing.T) {
	// Arrange

	//Load Tree
	tr, err := LoadTree(treeTest)
	if err != nil {
		t.Fail()
	}

	// Load request
	jsonRequest := []byte("Obviously not a json")

	//Act
	_, err = tr.ResolveJSON(jsonRequest)

	//Assert
	assert.Error(t, err, "Resolve should return an error  if the jsonrequest is malformed")
}

func ExampleLoadTree() {
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
			"value": 60
		},
		{
			"id": 13,
			"parent_id": 12,
			"Name": "Hello dude"
		}
	]`)

	t, err := LoadTree(jsonTree)
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

func TestTreeLegacy(t *testing.T) {
	tree := []Tree{
		{
			ID:   1,
			Name: "root",
		},
		{
			ID:       2,
			ParentID: 0,
			Name:     "1",
			Legacy: map[string][]interface{}{
				"cost": {20},
			},
		},
		{
			ID:       3,
			ParentID: 2,
			Name:     "2",
			Value:    true,
			Operator: "eq",
			Key:      "2",
			Legacy:   map[string][]interface{}{"cost": {50}},
		},
		{
			ID:       4,
			ParentID: 3,
			Name:     "4",
			Value:    true,
			Operator: "eq",
			Key:      "4",
			Legacy:   map[string][]interface{}{"cost": {40}},
		},
		{
			ID:       5,
			ParentID: 2,
			Name:     "5",
			Value:    true,
			Operator: "eq",
			Key:      "5",
			Legacy:   map[string][]interface{}{"cost": {100}},
		},
	}
	newTree := CreateTree(tree)

	fmt.Println(newTree)

	request := make(map[string]interface{})
	request["2"] = true
	request["3"] = false

	want := map[string][]interface{}{
		"cost": {50, 20},
	}

	v, err := newTree.Resolve(request)
	assert.NoError(t, err)

	assert.Equal(t, want, v.Legacy)

}
