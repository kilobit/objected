ObjectEd
========

A simple Object access library that let's you work with arbitrary
objects in Golang more like you would in JavaScript.

Especially useful when working with the encoding/json package.

Status: In-Development

```
type MyObj struct {
	Object
}

func (m MyObj) Two() string {

	vals, _ := m.Get("onetwothree")
	
	return string(vals[1])
}

func main() {

	bs := []byte(`{"message": "Hello World!", "onetwothree": ["one", "two", "three"]}`)

	m := MyObj{}
	json.Unmarshal(bs, &m)
	
	fmt.Println(m.GetString("message"))
	
	vals, _ := m.Get("onetwothree")
	
	for _, val := range vals {
		fmt.Println(val)
	}
	
	fmt.Println(m.Two())
}

```

Features
--------

- Walk Object hierarchies with dotted queries, e.g. "foo.bar.bing"
- Access data as it is in it's Unmarshaled form.
- Handles arbitrary values, collections of values and objects.
- Get values from a list of object with queries.
- Useful on its own or augment Objects with your own types and methods.

Future:
- Provide graceful shortcuts for numeric data.

Installation
------------

```
go get kilobit.ca/go/objected
go test -v
```

Building
--------

```
go get kilobit.ca/go/objected
go build
```

Contribute
----------

Please help improve this project!

Submit pull requests via github.com/kilobit/objected

Support
-------

Please submit issues through github.com/kilobit/objected

License
-------

See LICENSE.

--  
Created: Feb 9, 2021  
By: Christian Saunders <cps@kilobit.ca>  
Copyright 2021 Kilobit Labs Inc.
