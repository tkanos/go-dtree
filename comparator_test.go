package dtree

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type want struct {
	result bool
	err    error
}

type tt struct {
	name string
	v1   interface{}
	v2   *Tree
	err  error
	want map[string]want
}

var tts = []tt{
	{
		name: "string v1 < v2",
		v1:   "e",
		v2:   &Tree{Value: "g"},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: true,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "string v1 > v2",
		v1:   "b",
		v2:   &Tree{Value: "a"},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: false,
				err:    nil,
			},
			"gt": {
				result: true,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "string v1 = v2",
		v1:   "e",
		v2:   &Tree{Value: "e"},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "int v1 < v2",
		v1:   1,
		v2:   &Tree{Value: 2},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: true,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "int v1 > v2",
		v1:   2,
		v2:   &Tree{Value: 1},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: false,
				err:    nil,
			},
			"gt": {
				result: true,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "int v1 = v2",
		v1:   2,
		v2:   &Tree{Value: 2},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "int64 v1 < v2",
		v1:   int64(1),
		v2:   &Tree{Value: int64(2)},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: true,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "int64 v1 > v2",
		v1:   int64(2),
		v2:   &Tree{Value: int64(1)},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: false,
				err:    nil,
			},
			"gt": {
				result: true,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "int64 v1 = v2",
		v1:   int64(2),
		v2:   &Tree{Value: int64(2)},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "float64 v1 < v2",
		v1:   float64(1),
		v2:   &Tree{Value: float64(2)},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: true,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "float64 v1 > v2",
		v1:   float64(2),
		v2:   &Tree{Value: float64(1)},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: false,
				err:    nil,
			},
			"gt": {
				result: true,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "float64 v1 = v2",
		v1:   float64(2),
		v2:   &Tree{Value: float64(2)},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
			"lt": {
				result: false,
				err:    nil,
			},
			"lte": {
				result: true,
				err:    nil,
			},
			"gt": {
				result: false,
				err:    nil,
			},
			"gte": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "([]string) v1 == v2",
		v1:   []string{"a", "b"},
		v2:   &Tree{Value: []interface{}{"a"}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(string) []interface{v1,...} == []interface{v2,...}",
		v1:   []interface{}{"g", "a"},
		v2:   &Tree{Value: []interface{}{"g"}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(string) []interface{v1,...} != []interface{v2,...}",
		v1:   []interface{}{"g", "a"},
		v2:   &Tree{Value: []interface{}{"b"}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(string) []interface{v1,...} == v2",
		v1:   []interface{}{"g", "a"},
		v2:   &Tree{Value: "g"},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(string) v1 == []interface{v1,...}",
		v1:   "g",
		v2:   &Tree{Value: []interface{}{"g", "a"}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "string []interface{v1,...} != v2",
		v1:   []interface{}{"b", "a"},
		v2:   &Tree{Value: "g"},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "string v1 != []interface{v1,...}",
		v1:   "g",
		v2:   &Tree{Value: []interface{}{"b", "a"}},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(float64) v1 != v2",
		v1:   9.0,
		v2:   &Tree{Value: 10.0},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(float64) v1 == v2",
		v1:   10.0,
		v2:   &Tree{Value: 10.0},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(float64) []interface{v1,...} == v2",
		v1:   []interface{}{9.0, 8.8},
		v2:   &Tree{Value: 8.8},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(float64) v1 == []interface{v1,...}",
		v1:   float64(9.0),
		v2:   &Tree{Value: []interface{}{9.0, 8.8}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(float64) []interface{v1,...} != v2",
		v1:   []interface{}{9.0, 8.8},
		v2:   &Tree{Value: 7.7},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(float64) v1 != []interface{v1,...}",
		v1:   float64(7.7),
		v2:   &Tree{Value: []interface{}{9.0, 8.8}},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(int64) []interface{v1,...} == v2",
		v1:   []interface{}{int64(9), int64(8)},
		v2:   &Tree{Value: int64(8)},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(int64) v1 == []interface{v1,...}",
		v1:   int64(8),
		v2:   &Tree{Value: []interface{}{int64(9), int64(8)}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(int64) []interface{v1,...} != v2",
		v1:   []interface{}{int64(9), int64(8)},
		v2:   &Tree{Value: int64(7)},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(int64) v1 != []interface{v1,...}",
		v1:   int64(7),
		v2:   &Tree{Value: []interface{}{int64(9), int64(8)}},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(int) []interface{v1,...} == v2",
		v1:   []interface{}{9, 8},
		v2:   &Tree{Value: 8},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(int) v1 == []interface{v1,...}",
		v1:   8,
		v2:   &Tree{Value: []interface{}{9, 8}},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(int) []interface{v1,...} != v2",
		v1:   []interface{}{9, 8},
		v2:   &Tree{Value: 7},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(int) v1 != []interface{v1,...}",
		v1:   7,
		v2:   &Tree{Value: []interface{}{9, 8}},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(bool) v1 != v2",
		v1:   false,
		v2:   &Tree{Value: true},
		want: map[string]want{
			"eq": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "(bool) v1 == v2",
		v1:   true,
		v2:   &Tree{Value: true},
		want: map[string]want{
			"eq": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(compare) fallback",
		v2:   &Tree{Value: "fallback"},
		want: map[string]want{
			"compare": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "(compare) unsupported operator",
		v2:   &Tree{Value: 123, Operator: "abc"},
		want: map[string]want{
			"compare": {
				result: false,
				err:    ErrOperator,
			},
		},
	},
	{
		v1: "abcdefghijkl",
		v2: &Tree{Value: "[0-9]+"},
		want: map[string]want{
			"regex": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		v1: "abcdefgh45jkl",
		v2: &Tree{Value: "[0-9]+"},
		want: map[string]want{
			"regex": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "Not supported type (int)",
		v1:   123,
		want: map[string]want{
			"regex": {
				result: false,
				err:    ErrNotSupportedType,
			},
		},
	},
	{
		name: "forbidden int type in tree value",
		v1:   "123",
		v2:   &Tree{Value: 123},
		want: map[string]want{
			"regex": {
				result: false,
				err:    ErrBadType,
			},
		},
	},
	{
		name: "!(v1 U v2)",
		v1:   "abcdef",
		v2:   &Tree{Value: "fed"},
		want: map[string]want{
			"contains": {
				result: false,
				err:    nil,
			},
		},
	},
	{
		name: "v1 U v2",
		v1:   "abcdef",
		v2:   &Tree{Value: "def"},
		want: map[string]want{
			"contains": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "int U {}",
		v1:   123,
		v2:   &Tree{},
		want: map[string]want{
			"contains": {
				result: false,
				err:    ErrNotSupportedType,
			},
		},
	},
	{
		name: "string U int",
		v1:   "123",
		v2:   &Tree{Value: 123},
		want: map[string]want{
			"contains": {
				result: false,
				err:    ErrBadType,
			},
		},
	},
	{
		name: "v1 not allowed",
		v1:   123,
		want: map[string]want{
			"count": {
				result: false,
				err:    ErrNotSupportedType,
			},
		},
	},
	{
		name: "int",
		v1:   []int{1, 2, 3},
		v2:   &Tree{Value: 3},
		want: map[string]want{
			"count": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "int64",
		v1:   []int64{1, 2, 3},
		v2:   &Tree{Value: int64(3)},
		want: map[string]want{
			"count": {
				result: true,
				err:    nil,
			},
		},
	},
	{
		name: "float64",
		v1:   []float64{1.1, 2.2, 3.3},
		v2:   &Tree{Value: float64(3.0)},
		want: map[string]want{
			"count": {
				result: true,
				err:    nil,
			},
		},
	},
}

func TestComparator(t *testing.T) {
	for _, tt := range tts {
		for wantKey, want := range tt.want {
			t.Run(fmt.Sprintf("%s+%s", tt.name, wantKey), func(t *testing.T) {
				var (
					got *Tree
					err error
				)
				switch wantKey {
				case "eq":
					got, err = eq(tt.v1, tt.v2)
				case "gt":
					got, err = gt(tt.v1, tt.v2)
				case "gte":
					got, err = gte(tt.v1, tt.v2)
				case "lt":
					got, err = lt(tt.v1, tt.v2)
				case "lte":
					got, err = lte(tt.v1, tt.v2)
				case "compare":
					got, err = compare(nil, nil, tt.v2, nil)
				case "regex":
					got, err = regex(tt.v1, tt.v2)
				case "contains":
					got, err = contains(tt.v1, tt.v2)
				case "count":
					got, err = count(tt.v1, tt.v2)
				default:
					t.Logf("No matching test for key %q", wantKey)
				}
				assert.Equal(t, want.err, err)
				assert.Equal(t, want.result, (got != nil))
			})
		}
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
	assert.Equal(t, rootTree.GetChild()[1], result, "percentage should return fallback, if fallback is defined, and there is no others choice")
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

func TestCRC32(t *testing.T) {
	num1 := crc32Num("entity1", "salt1", 1000)
	num2 := crc32Num("entity2", "salt1", 1000)
	num3 := crc32Num("entity1", "salt1", 1000)
	assert.Equal(t, num1, num3)
	assert.NotEqual(t, num1, num2)
}

// TestAbTest_Without_Parent_Node_Should_Return_Nil should return nil, if no parent
func TestAbTest_Without_Parent_Node_Should_Return_Nil(t *testing.T) {
	//Arrange
	percentTree := &Tree{}
	//Act
	result, err := abTest(nil, percentTree)

	//Assert
	assert.Equal(t, err, ErrNoParentNode, "A/B Test should return an error, if no parents")
	assert.Nil(t, result, "A/B Test should return nil, if no parent")
}

// func TestAbTest_Without_Brother_Node_Should_Return_ItSelf should return itself, if no brother
func TestAbTest_Without_Brother_Node_Should_Return_ItSelf(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{})
	//Act
	result, err := abTest(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "A/B Test should not return an error, if there is no brothers")
	assert.Equal(t, rootTree.GetChild()[0], result, "A/B Test should return the node itself, when there is no brothers")
}

// TestAbTest_Without_FloatValue_Should_Return_Nil should return nil, if the value is no parsable to float64
func TestAbTest_Without_FloatValue_Should_Return_Nil(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		Operator: "ab",
		Value:    123,
	})
	rootTree.AddNode(&Tree{
		Operator: "ab",
		Value:    123,
	})
	//Act
	result, err := abTest(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "A/B Test should not return an error, if the value is no parsable to float64")
	assert.Nil(t, result, "A/B Test should return nil, if the value is no parsable to float64")
}

// TestAbTest_With_FallBack_Should_Return_Fallback fallback should be returned if it is defined and there is no others choice
func TestAbTest_With_FallBack_Should_Return_Fallback(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		Operator: "ab",
		Value:    123,
	})
	rootTree.AddNode(&Tree{
		Value: FallbackType,
	})
	//Act
	result, err := abTest(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err, "A/B Test should not return an error, if fallback is defined")
	assert.Equal(t, rootTree.GetChild()[1], result, "A/B Test should return fallback, if fallback is defined, and there is no others choice")
}

// TestAbTest_Should_Return_A_Node should return a node, if all is ok
func TestAbTest_Should_Return_A_Node_Like_For_precentage(t *testing.T) {
	if os.Getenv("TEST_SKIP") != "" {
		t.Skip("Skipping random test, that won't work on CI process")
	}
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		ID:       123,
		Operator: "ab",
		Value:    50.0,
	})
	rootTree.AddNode(&Tree{
		ID:       456,
		Operator: "ab",
		Value:    50.0,
	})
	//Act
	result1, err1 := abTest(nil, rootTree.GetChild()[0])
	result2, err2 := abTest(nil, rootTree.GetChild()[0])
	result3, err3 := abTest(nil, rootTree.GetChild()[0])
	result4, err4 := abTest(nil, rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err1, "A/B Test should not return an error, if all is ok")
	assert.NoError(t, err2, "A/B Test should not return an error, if all is ok")
	assert.NoError(t, err3, "A/B Test should not return an error, if all is ok")
	assert.NoError(t, err4, "A/B Test should not return an error, if all is ok")
	assert.True(t, result1.ID != result2.ID || result2.ID != result3.ID || result3.ID != result4.ID, "A/B Test should return 4 different node if all is ok")
}

// TestAbTest_Should_Return_A_Node should return a node, if all is ok
func TestAbTest_Should_Return_The_Same_Node_By_UserId(t *testing.T) {
	//Arrange
	rootTree := &Tree{}
	rootTree.AddNode(&Tree{
		ID:       123,
		Operator: "ab",
		Value:    50.0,
	})
	rootTree.AddNode(&Tree{
		ID:       456,
		Operator: "ab",
		Value:    50.0,
	})
	//Act
	result1, err1 := abTest("Felipe", rootTree.GetChild()[0])
	result2, _ := abTest("Felipe", rootTree.GetChild()[0])
	result3, _ := abTest("Felipe", rootTree.GetChild()[0])
	result4, _ := abTest("Felipe", rootTree.GetChild()[0])

	result1_1, err2 := abTest("Another", rootTree.GetChild()[0])
	result2_1, _ := abTest("Another", rootTree.GetChild()[0])
	result3_1, _ := abTest("Another", rootTree.GetChild()[0])
	result4_1, _ := abTest("Another", rootTree.GetChild()[0])

	//Assert
	assert.NoError(t, err1, "A/B Test should not return an error, if all is ok")
	assert.NoError(t, err2, "A/B Test should not return an error, if all is ok")
	assert.True(t, result1.ID == result2.ID && result2.ID == result3.ID && result3.ID == result4.ID, "A/B Test should return 4 same node if all is ok")
	assert.NotEqual(t, result1.ID, result1_1.ID, "A/B Test should return 2 different node if different usersId")
	assert.True(t, result1_1.ID == result2_1.ID && result2_1.ID == result3_1.ID && result3_1.ID == result4_1.ID, "A/B Test should return 4 same node if all is ok")
}
