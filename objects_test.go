/* Copyright 2021 Kilobit Labs Inc. */

package objected_test

import _ "fmt"
import _ "errors"

import o "kilobit.ca/go/objected"

import "kilobit.ca/go/tested/assert"
import "testing"

func TestObjectTest(t *testing.T) {
	assert.Expect(t, true, true, "Failed Sanity Check.")
}

func TestGet(t *testing.T) {

	tests := []struct {
		obj    o.Object
		path   string
		result bool
		exp    interface{}
	}{

		{o.Object{"foo": "bar"}, "foo", true, "bar"},
		{o.Object{"foo": "bar"}, "nope", false, nil},
		{o.Object{"foo": "bar", "bing": "bang"}, "foo", true, "bar"},
		{o.Object{"foo": "bar", "bing": "bang"}, "bing", true, "bang"},
		{o.Object{"foo": o.Object{"bar": "bing"}}, "foo.bar", true, "bing"},
		{o.Object{"foo": map[string]interface{}{"bar": "bing"}}, "foo.bar", true, "bing"},

		{o.Object{"foo": map[string]interface{}{"bar": "bing"}}, ".foo.bar", false, nil},
		{o.Object{"foo": map[string]interface{}{"bar": "bing"}}, "foo.bar.", false, nil},

		{o.Object{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"bing": "bang",
				},
			},
		}, "foo.bar.bing", true, "bang"},

		{o.Object{
			"foo": o.Object{
				"bar": map[string]interface{}{
					"bing": "bang",
				},
			},
		}, "foo.bar.bing", true, "bang"},

		{o.Object{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"bing": "bang",
				},
			},
		}, "foo.bor.bing", false, nil},

		{o.Object{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"bing": "bang",
				},
			},
		}, "foo.bar.bong", false, nil},
	}

	for _, test := range tests {

		act, ok := test.obj.Get(test.path)
		assert.Expect(t, test.result, ok, test.obj, test.path, test.result, test.exp)
		assert.Expect(t, test.exp, act, test.obj, test.path, test.result, test.exp)
	}
}

func TestGetValues(t *testing.T) {

	tests := []struct {
		vals  o.Values
		query string
		exps  o.Values
	}{

		{
			o.Values{},
			"foo",
			o.Values{},
		},

		{
			o.Values{
				o.Object{"foo": "bar"},
				o.Object{"foo": "bing"},
			},
			"foo",
			o.Values{"bar", "bing"},
		},

		{
			o.Values{
				o.Object{"foo": "bar"},
				map[string]interface{}{"foo": "bing"},
			},
			"foo",
			o.Values{"bar", "bing"},
		},

		{
			o.Values{
				map[string]interface{}{"foo": "bar"},
				map[string]interface{}{"foo": "bing"},
			},
			"foo",
			o.Values{"bar", "bing"},
		},

		{
			o.Values{
				map[string]interface{}{"foo": "bar"},
				map[string]interface{}{"fo": "bing"},
			},
			"foo",
			o.Values{"bar"},
		},

		{
			o.Values{
				map[string]interface{}{"foo": "bar"},
				map[string]interface{}{"foo": "bing"},
			},
			"f",
			o.Values{},
		},
	}

	for _, test := range tests {
		acts := test.vals.GetValues(test.query)
		assert.ExpectDeep(t, test.exps, acts, test.vals, test.query, test.exps)
	}
}

func TestGetNumber(t *testing.T) {

	tests := []struct {
		obj   o.Object
		query string
		exp   float64
		err   error
	}{
		{o.Object{"foo": 42}, "foo", 42, nil},
	}

	for _, test := range tests {

		result, err := test.obj.GetNumber(test.query)
		assert.Expect(t, test.err, err, test.obj, test.query, test.exp, test.err)
		assert.Expect(t, test.exp, result, test.obj, test.query, test.exp, test.err)
	}
}
