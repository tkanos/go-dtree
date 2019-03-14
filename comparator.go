package dtree

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

// FallbackType the fallback value (by default = "fallback"), can be overrided
var FallbackType = "fallback"

// ErrOperator : unknow operator
var ErrOperator = errors.New("unknow operator")

// ErrBadType : Types between request and Tree are different, so we are unable to compare them
var ErrBadType = errors.New("types are different")

// ErrNotSupportedType : Type in request are not supported
var ErrNotSupportedType = errors.New("type not supported")

// ErrNoNode : No Node was sent
var ErrNoNode = errors.New("Node is nil")

// ErrNoParentNode : Node has no parent
var ErrNoParentNode = errors.New("Node has no parent")

func compare(requests map[string]interface{}, jsonValue interface{}, node *Tree, config *TreeOptions) (*Tree, error) {

	if node == nil {
		return nil, ErrNoNode
	}

	// Check if it is a fallback value
	if v, ok := node.Value.(string); (ok && v == FallbackType) || len(node.Operator) == 0 {
		return node, nil
	}

	// if the operators override an existing one defined on isExistingOperator, we check first the opeators else we do it in the default
	if config != nil && config.OverrideExistingOperator {
		if r, err := runOperators(requests, node, config); err != ErrOperator {
			return r, err
		}
	}

	switch node.Operator {
	case "eq", "==":
		return eq(jsonValue, node)
	case "ne", "!=":
		b, err := eq(jsonValue, node)
		if b == nil {
			return node, err
		}
		return nil, err
	case "gt", ">":
		return gt(jsonValue, node)
	case "lt", "<":
		return lt(jsonValue, node)
	case "gte", ">=":
		return gte(jsonValue, node)
	case "lte", "<=":
		return lte(jsonValue, node)
	case "contains":
		return contains(jsonValue, node)
	case "count":
		return count(jsonValue, node)
	case "regexp":
		return regex(jsonValue, node)
	case "percent", "%":
		return percentage(jsonValue, node)
	default:
		if config != nil && config.OverrideExistingOperator == false {
			if r, err := runOperators(requests, node, config); err != ErrOperator {
				return r, err
			}
		}
		return nil, ErrOperator
	}
}

func runOperators(requests map[string]interface{}, node *Tree, config *TreeOptions) (*Tree, error) {
	if config != nil && config.Operators != nil {
		for k, f := range config.Operators {
			if k == node.Operator {
				return f(requests, node)
			}
		}
	}
	return nil, ErrOperator
}

func isExistingOperator(name string) bool {
	switch name {
	case "eq", "==", "ne", "!=", "gt", ">", "lt", "<", "gte", ">=", "lte", "<=", "contains", "count", "regexp", "percent", "%":
		return true
	default:
		return false
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
				if t2, ok := v2.Value.([]interface{}); ok {
					for _, vs := range t2 {
						if t2, ok := vs.(float64); ok {
							if tv == t2 {
								return v2, nil
							}
						}
					}
				}
			case string:
				if t2, ok := v2.Value.(string); ok {
					if tv == t2 {
						return v2, nil
					}
				}
				if t2, ok := v2.Value.([]interface{}); ok {
					for _, vs := range t2 {
						if t2, ok := vs.(string); ok {
							if tv == t2 {
								return v2, nil
							}
						}
					}
				}
			}
		}
		return nil, nil
	case []string:
		for _, v := range v1.([]interface{}) {
			switch tv := v.(type) {
			case string:
				if t2, ok := v2.Value.(string); ok {
					if tv == t2 {
						return v2, nil
					}
				}
			}
		}
		return nil, nil
	case []float64:
		for _, v := range v1.([]interface{}) {
			switch tv := v.(type) {
			case float64:
				if t2, ok := v2.Value.(float64); ok {
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
	if v2.GetParent() == nil {
		return nil, ErrNoParentNode
	}

	brothersNode := v2.GetParent().GetChild()

	if brothersNode == nil || len(brothersNode) == 1 {
		return v2, nil
	}
	var fallbackNode *Tree
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
