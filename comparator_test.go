package dtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var containstt = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "Contains should not support others type than string as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      "123",
		v2:      &Tree{Value: 123},
		message: "Contains should not support others type than string as Tree Value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "abcdefghijkl",
		v2:      &Tree{Value: "mnopqrstu"},
		message: "Contains should return false if v1 does not contains v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "abcdefghijkl",
		v2:      &Tree{Value: "def"},
		message: "Contains should return true if v1 contains v2",
		result:  true,
		err:     nil,
	},
}

// TestContains test contains feature
func TestContains(t *testing.T) {
	for _, tt := range containstt {
		// Act
		result, err := contains(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var counttt = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "Count should not support others type than []interface{} as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      []interface{}{1, 2, 3},
		v2:      &Tree{Value: "not float64"},
		message: "Count should not support others type than float64 as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      []interface{}{1, 2, 3},
		v2:      &Tree{Value: 5.0},
		message: "Count should return false if len(v1) != v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      []interface{}{1, 2, 3},
		v2:      &Tree{Value: 3.0},
		message: "Count should return false if len(v1) == v2",
		result:  true,
		err:     nil,
	},
}

// TestCount test Count feature
func TestCount(t *testing.T) {
	for _, tt := range counttt {
		// Act
		result, err := count(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var gttt = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "gt should not support others type than float64 and string as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 123},
		message: "gt should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: 123},
		message: "gt should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "e",
		v2:      &Tree{Value: "g"},
		message: "gt (string) should return false if v1 < v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "g"},
		message: "gt (string) should return false if v1 == v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "t",
		v2:      &Tree{Value: "g"},
		message: "gt (string) should return true if v1 > v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      9.0,
		v2:      &Tree{Value: 10.0},
		message: "gt (float46) should return false if v1 < v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 10.0},
		message: "gt (float46) should return false if v1 == v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      15.0,
		v2:      &Tree{Value: 10.0},
		message: "gt (float46) should return true if v1 > v2",
		result:  true,
		err:     nil,
	},
}

// TestGt test GreatherThan feature
func TestGt(t *testing.T) {
	for _, tt := range gttt {
		// Act
		result, err := gt(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var lttt = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "lt should not support others type than float64 and string as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 123},
		message: "lt should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: 123},
		message: "lt should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "e"},
		message: "lt (string) should return false if v1 > v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "g"},
		message: "lt (string) should return false if v1 == v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "t"},
		message: "lt (string) should return true if v1 < v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 9.0},
		message: "lt (float64) should return false if v1 > v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 10.0},
		message: "lt (float64) should return false if v1 == v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 15.0},
		message: "lt (float64) should return true if v1 < v2",
		result:  true,
		err:     nil,
	},
}

// TestLt test LessThan feature
func TestLt(t *testing.T) {
	for _, tt := range lttt {
		// Act
		result, err := lt(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var gtett = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "gte should not support others type than float64 and string as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 123},
		message: "gte should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: 123},
		message: "gte should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "e",
		v2:      &Tree{Value: "g"},
		message: "gte (string) should return false if v1 < v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "g"},
		message: "gte (string) should return true if v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      "t",
		v2:      &Tree{Value: "g"},
		message: "gte (string) should return true if v1 > v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      9.0,
		v2:      &Tree{Value: 10.0},
		message: "gte (float64) should return false if v1 < v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 10.0},
		message: "gte (float64) should return false if v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      15.0,
		v2:      &Tree{Value: 10.0},
		message: "gte (float64) should return true if v1 > v2",
		result:  true,
		err:     nil,
	},
}

// TestGte test GreatherThan Or Equal feature
func TestGte(t *testing.T) {
	for _, tt := range gtett {
		// Act
		result, err := gte(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var ltett = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "lte should not support others type than float64 and string as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 123},
		message: "lte should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: 123},
		message: "lte should not support others type than float64 and string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "e"},
		message: "lte (string) should return false if v1 > v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "g"},
		message: "lte (string) should return true if v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      "g",
		v2:      &Tree{Value: "t"},
		message: "lte (string) should return true if v1 < v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 9.0},
		message: "lte (float64) should return false if v1 > v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 10.0},
		message: "lte (float64) should return true if v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      10.0,
		v2:      &Tree{Value: 15.0},
		message: "lte (float64) should return true if v1 < v2",
		result:  true,
		err:     nil,
	},
}

// TestLte test LessThan Or Equal feature
func TestLte(t *testing.T) {
	for _, tt := range ltett {
		// Act
		result, err := lte(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var eqtt = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "eq should not support others type than  string, float64, bool, []interface{} (interface{} being a string or float64) as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 123},
		message: "eq should not support others type than float64, bool, string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "123.0",
		v2:      &Tree{Value: 123.0},
		message: "eq should not support others type than float64, bool, string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      true,
		v2:      &Tree{Value: 123.0},
		message: "eq should not support others type than float64, bool, string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "123.0",
		v2:      &Tree{Value: 123.0},
		message: "eq should not support others type than float64, bool, string as Tree value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 122.0},
		message: "eq (float64) should return false when v1 != v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      123.0,
		v2:      &Tree{Value: 123.0},
		message: "eq (float64) should return true when v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: "b"},
		message: "eq (string) should return false when v1 != v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: "a"},
		message: "eq (string) should return true when v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      true,
		v2:      &Tree{Value: false},
		message: "eq (bool) should return false when v1 != v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      true,
		v2:      &Tree{Value: true},
		message: "eq (bool) should return true when v1 == v2",
		result:  true,
		err:     nil,
	},
	{
		v1:      []interface{}{true, false},
		v2:      &Tree{Value: 122.0},
		message: "eq should not support others type than []interface{} (interface{} being a string or float64) as request",
		result:  false,
		err:     nil,
	},
	{
		v1:      []interface{}{123.0, 456.0},
		v2:      &Tree{Value: true},
		message: "eq ([]interface{} => float64) should return false when v2 type not the same as v1 element",
		result:  false,
		err:     nil,
	},
	{
		v1:      []interface{}{123.0, 456.0},
		v2:      &Tree{Value: 122.0},
		message: "eq ([]interface{} => float64) should return false when v2 are not in v1",
		result:  false,
		err:     nil,
	},
	{
		v1:      []interface{}{123.0, 456.0},
		v2:      &Tree{Value: 456.0},
		message: "eq ([]interface{} => float64) should return true when v2 are in v1",
		result:  true,
		err:     nil,
	},
	{
		v1:      []interface{}{"a", "b"},
		v2:      &Tree{Value: true},
		message: "eq ([]interface{} => string) should return false when v2 type not the same as v1 element",
		result:  false,
		err:     nil,
	},
	{
		v1:      []interface{}{"a", "b"},
		v2:      &Tree{Value: "c"},
		message: "eq ([]interface{} => string) should return false when v2 are not in v1",
		result:  false,
		err:     nil,
	},
	{
		v1:      []interface{}{"a", "b"},
		v2:      &Tree{Value: "a"},
		message: "eq ([]interface{} => string) should return true when v2 are in v1",
		result:  true,
		err:     nil,
	},
	{
		v1:      []interface{}{"a", "b"},
		v2:      &Tree{Value: "c"},
		message: "eq ([]interface{} => string) should return false when v2 are not in v1",
		result:  false,
		err:     nil,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: []interface{}{1, 2}},
		message: "eq (TreeValue []interface{} => int) without string in the tree value should return nil",
		result:  false,
		err:     nil,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: []interface{}{"1", "2"}},
		message: "eq (TreeValue []interface{} => string) not finding string should return nil",
		result:  false,
		err:     nil,
	},
	{
		v1:      "a",
		v2:      &Tree{Value: []interface{}{"a", "b"}},
		message: "eq (TreeValue []interface{} => string) finding string should return The node",
		result:  true,
		err:     nil,
	},
	{
		v1:      1.0,
		v2:      &Tree{Value: []interface{}{1, 2}},
		message: "eq (TreeValue []interface{} => int) without float64 in the tree value should return nil",
		result:  false,
		err:     nil,
	},
	{
		v1:      1.0,
		v2:      &Tree{Value: []interface{}{3.0, 4.0}},
		message: "eq (TreeValue []interface{} => float64) not finding float64 should return nil",
		result:  false,
		err:     nil,
	},
	{
		v1:      1.0,
		v2:      &Tree{Value: []interface{}{1.0, 2.0}},
		message: "eq (TreeValue []interface{} => float64) finding float64 should return The node",
		result:  true,
		err:     nil,
	},
}

// TestEq test Equal feature
func TestEq(t *testing.T) {
	for _, tt := range eqtt {
		// Act
		result, err := eq(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var comparett = []struct {
	op      string
	v2      *Tree
	err     error
	result  bool
	message string
}{
	{
		v2:      &Tree{Value: "fallback"},
		message: "Compare should always return true, when it is the fallback node",
		result:  true,
		err:     nil,
	},
	{
		v2:      &Tree{Value: 123},
		message: "Compare should always return false, when the operator is not supported",
		op:      "abc",
		result:  false,
		err:     ErrOperator,
	},
}

// TestCompare test the Compare function
func TestCompare(t *testing.T) {
	for _, tt := range comparett {
		// Act
		result, err := compare(nil, "", tt.op, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

var regexptt = []struct {
	v1     interface{}
	v2     *Tree
	err    error
	result bool

	message string
}{
	{
		v1:      123,
		message: "Contains should not support others type than string as request",
		result:  false,
		err:     ErrNotSupportedType,
	},
	{
		v1:      "123",
		v2:      &Tree{Value: 123},
		message: "Contains should not support others type than string as Tree Value",
		result:  false,
		err:     ErrBadType,
	},
	{
		v1:      "abcdefghijkl",
		v2:      &Tree{Value: "[0-9]+"},
		message: "Contains should return false if v1 does not match v2",
		result:  false,
		err:     nil,
	},
	{
		v1:      "abcdefgh45jkl",
		v2:      &Tree{Value: "[0-9]+"},
		message: "Contains should return true if v1 contains v2",
		result:  true,
		err:     nil,
	},
}

// TestRegex test regular expression feature
func TestRegex(t *testing.T) {
	for _, tt := range containstt {
		// Act
		result, err := regex(tt.v1, tt.v2)

		// Assert
		assert.Equal(t, tt.err, err, tt.message)
		assert.Equal(t, tt.result, (result != nil), tt.message)
	}
}

// TestPercentage_Without_Parent_Node_Should_Return_Nil should return nil, if no parent
func TestPercentage_Without_Parent_Node_Should_Return_Nil(t *testing.T) {
	//Arrange
	percentTree := &Tree{}
	//Act
	result, err := percentage(nil, percentTree)

	//Assert
	assert.Equal(t, err, ErrNoParentNode, "percentage should return an error, if no parents")
	assert.Nil(t, result, "percentage should return nil, if no parent")
}

// TestPercentage_Without_Brother_Node_Should_Return_ItSelf should return itself, if no brother
func TestPercentage_Without_Brother_Node_Should_Return_ItSelf(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{})
	//Act
	result, err := percentage(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "percentage should not return an error, if there is no brothers")
	assert.Equal(t, rootTree.GetChild()[0], result, "percentage should return the node itself, when there is no brothers")
}

// TestPercentage_Without_FloatValue_Should_Return_Nil should return nil, if the value is no parsable to float64
func TestPercentage_Without_FloatValue_Should_Return_Nil(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		Operator: "%",
		Value:    123,
	})
	rootTree.AddNode(&Tree{
		Operator: "%",
		Value:    123,
	})
	//Act
	result, err := percentage(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "percentage should not return an error, if the value is no parsable to float64")
	assert.Nil(t, result, "percentage should return nil, if the value is no parsable to float64")
}

// TestPercentage_With_FallBack_Should_Return_Fallback fallback should be returned if it is defined and there is no others choice
func TestPercentage_With_FallBack_Should_Return_Fallback(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		Operator: "%",
		Value:    123,
	})
	rootTree.AddNode(&Tree{
		Value: FallbackType,
	})
	//Act
	result, err := percentage(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "percentage should not return an error, if fallback is defined")
	assert.Equal(t, rootTree.GetChild()[1], result, "percentage should return falback, if fallback is defined, and there is no others choice")
}

// TestPercentage_Should_Return_A_Node should return a node, if all is ok
func TestPercentage_Should_Return_A_Node(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		Operator: "%",
		Value:    50.0,
	})
	rootTree.AddNode(&Tree{
		Operator: "%",
		Value:    50.0,
	})
	//Act
	result, err := percentage(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "percentage should not return an error, if all is ok")
	assert.NotNil(t, result, "percentage should return a node if all is ok")
}
