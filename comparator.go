package dtree

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// FallbackType the fallback value (by default = "fallback"), can be overrided
var FallbackType = "fallback"

// ErrOperator : unknow operator
var ErrOperator = errors.New("unknow operator")

// ErrBadType : Types between request and Tree are different, so we are unable to compare them
var ErrBadType = errors.New("types are different")

// ErrNotSupportedType : Type in request are not supported
var ErrNotSupportedType = errors.New("type not supported")

// ErrNoNode : No treeValue was sent
var ErrNoNode = errors.New("Node is nil")

// ErrNoParentNode : Node has no parent
var ErrNoParentNode = errors.New("Node has no parent")

func compare(request interface{}, op string, treeValue *Tree, operators ...map[string]func(interface{}, *Tree) (*Tree, error)) (*Tree, error) {

	if treeValue == nil {
		return nil, ErrNoNode
	}
	// Check if it is a fallback value
	if v, ok := treeValue.Value.(string); ok && v == FallbackType {
		return treeValue, nil
	}

	if operators != nil {
		for _, operator := range operators {
			if f, ok := operator[op]; ok {
				return f(request, treeValue)
			}
		}
	}

	switch op {
	case "eq", "==":
		return eq(request, treeValue)
	case "ne", "!=":
		b, err := eq(request, treeValue)
		if b == nil {
			return treeValue, err
		}
		return nil, err
	case "gt", ">":
		return gt(request, treeValue)
	case "lt", "<":
		return lt(request, treeValue)
	case "gte", ">=":
		return gte(request, treeValue)
	case "lte", "<=":
		return lte(request, treeValue)
	case "contains":
		return contains(request, treeValue)
	case "count":
		return count(request, treeValue)
	case "regexp":
		return regex(request, treeValue)
	case "percent", "%":
		return percentage(request, treeValue)
	default:
		return nil, ErrOperator
	}
}

// eq check if v1 == v2 (only for string, float64, bool, []interface{} (interface{} being a string or float64))
func eq(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case float64:
		switch t2 := v2.Value.(type) {
		case float64:
			if t1 == t2 {
				return v2, nil
			}
			return nil, nil
		case []interface{}:
			for _, v := range t2 {
				if t2, ok := v.(float64); ok {
					if t1 == t2 {
						return v2, nil
					}
				}
			}
			return nil, nil
		default:
			return nil, ErrBadType
		}
	case string:
		switch t2 := v2.Value.(type) {
		case string:
			if t1 == t2 {
				return v2, nil
			}
			return nil, nil
		case []interface{}:
			for _, v := range t2 {
				if t2, ok := v.(string); ok {
					if t1 == t2 {
						return v2, nil
					}
				}
			}
			return nil, nil
		default:
			return nil, ErrBadType
		}
	case bool:
		if t2, ok := v2.Value.(bool); ok {
			if t1 == t2 {
				return v2, nil
			}
			return nil, nil
		}

		return nil, ErrBadType
	case []interface{}:
		for _, v := range v1.([]interface{}) {
			switch tv := v.(type) {
			case float64:
				if t2, ok := v2.Value.(float64); ok {
					if tv == t2 {
						return v2, nil
					}
				}
			case string:
				if t2, ok := v2.Value.(string); ok {
					if tv == t2 {
						return v2, nil
					}
				}
			}
		}
		return nil, nil
	default:
		return nil, ErrNotSupportedType
	}
}

// gt check if v1 > v2 (only for float64 and string)
func gt(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case float64:
		if t2, ok := v2.Value.(float64); ok {
			if t1 > t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	case string:
		if t2, ok := v2.Value.(string); ok {
			if t1 > t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// lt check if v1 < v2 (only for float64 and string)
func lt(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case float64:
		if t2, ok := v2.Value.(float64); ok {
			if t1 < t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	case string:
		if t2, ok := v2.Value.(string); ok {
			if t1 < t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// gte check if v1 >= v2 (only for float64 and string)
func gte(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case float64:
		if t2, ok := v2.Value.(float64); ok {
			if t1 >= t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	case string:
		if t2, ok := v2.Value.(string); ok {
			if t1 >= t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// lte check if v1 <= v2 (only for float64 and string)
func lte(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case float64:
		if t2, ok := v2.Value.(float64); ok {
			if t1 <= t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	case string:
		if t2, ok := v2.Value.(string); ok {
			if t1 <= t2 {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// contains check if a string v1 contains a string v2
func contains(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case string:
		if t2, ok := v2.Value.(string); ok {
			if strings.Contains(t1, t2) {
				return v2, nil
			}
			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// count check if the length of a slice v1 == (int)v2
func count(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case []interface{}:
		if t2, ok := v2.Value.(float64); ok {
			if len(t1) == int(t2) {
				return v2, nil
			}
			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// v check if the regexp pattern provided by v2 match v1 string
func regex(v1 interface{}, v2 *Tree) (*Tree, error) {
	switch t1 := v1.(type) {
	case string:
		if t2, ok := v2.Value.(string); ok {
			matched, _ := regexp.MatchString(t2, t1)
			if matched {
				return v2, nil
			}

			return nil, nil
		}

		return nil, ErrBadType
	default:
		return nil, ErrNotSupportedType
	}
}

// percentage rolls the dice, to know if it falls on one of the bucket of the percents node.
func percentage(v1 interface{}, v2 *Tree) (*Tree, error) {
	if v2.Parent == nil {
		return nil, ErrNoParentNode
	}

	brothersNode := v2.Parent.GetChild()

	if brothersNode == nil || len(brothersNode) == 1 {
		return v2, nil
	}
	var fallbackNode *Tree
	rand.Seed(int64(time.Now().Nanosecond()))
	var percent = rand.Float64() * 100.0
	var total float64

	for _, node := range brothersNode {
		if node.Operator == "%" || node.Operator == "percent" {
			if tn, ok := node.Value.(float64); ok {
				max := total + tn
				if percent <= max {
					return node, nil
				}
				total = total + tn
			}
		}

		// search if it exist a fallback node
		if tn, ok := node.Value.(string); ok && tn == FallbackType {
			fallbackNode = node
		}
	}

	// check for fallback
	if fallbackNode != nil {
		return fallbackNode, nil
	}

	return nil, nil
}