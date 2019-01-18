package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	js, err := simplejson.NewJson([]byte(`{
    "test": {
      "string_array": ["asdf", "ghjk", "zxcv"],
      "array": [1, "2", 3],
      "arraywithsubs": [{"subkeyone": 1},
      {"subkeytwo": 2, "subkeythree": 3}],
      "int": 10,
      "float": 5.150,
      "bignum": 9223372036854775807,
      "string": "simplejson",
      "bool": true
    }
  }`))
	if err != nil {
		panic("json format error")
	}

	s, err := js.Get("test").Get("string").String()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
