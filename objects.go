/* Copyright 2021 Kilobit Labs Inc. */

package objected

import _ "fmt"
import "errors"

import "strings"

type Value interface{}

type Values []interface{}

type ValueMapFunc func(i int, val Value) Value

func (vals Values) Map(f ValueMapFunc) Values {

	results := Values{}

	for i, val := range vals {
		result := f(i, val)

		if result != nil {
			results = append(results, result)
		}
	}

	return results
}

// Gets values from a collection of Objects.
//
func (vals Values) GetValues(query string) Values {

	results := vals.Map(ValueMapFunc(func(i int, val Value) Value {

		var obj Object

		switch v := val.(type) {

		case map[string]interface{}:
			obj = Object(v)

		case Object:
			obj = v

		default:
			return nil
		}

		res, ok := obj.Get(query)
		if !ok {
			return nil
		}

		return res
	}))

	return results
}

type Object map[string]interface{}

// Recursively looks up key segments in Objects using '.' as a
// separator.
//
func (obj Object) Get(query string) (Value, bool) {

	if strings.HasSuffix(query, ".") {
		return nil, false
	}

	key := query
	rest := ""

	i := strings.Index(query, ".")
	if i != -1 {
		key = query[:i]
		rest = query[i+1:]
	}

	val, ok := obj[key]
	if !ok {
		return nil, false
	}

	if rest == "" {
		return val, true
	}

	switch v := val.(type) {

	case map[string]interface{}:
		return Object(v).Get(rest)

	case Object:
		return v.Get(rest)

	default:
		return nil, false
	}
}

// Gets a value and coerces it into a string.
//
// Returns "" if not found or not coercible.
//
func (obj Object) GetString(query string) string {

	result := ""

	val, ok := obj.Get(query)
	if ok {
		str, ok := val.(string)
		if ok {
			result = str
		}
	}

	return result
}

var ErrorNotFound = errors.New("Not found.")
var ErrorNotANumber error = errors.New("Not a number.")

// Gets a value and coerces it into a number.
//
func (obj Object) GetNumber(query string) (float64, error) {

	result := float64(0)
	var err error = nil

	val, ok := obj.Get(query)
	if ok {
		switch v := val.(type) {
		case int:
			result = float64(v)
		case float64:
			result = v
		default:
			err = ErrorNotANumber
		}
	} else {

		err = ErrorNotFound
	}

	return result, err
}

func (obj Object) Keys() []string {
	keys := []string{}

	for key, _ := range obj {
		keys = append(keys, key)
	}

	return keys
}

func (obj Object) Values() Values {

	vals := Values{}
	for _, val := range obj {
		vals = append(vals, val)
	}

	return vals
}
