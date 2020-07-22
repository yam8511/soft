package soft

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

// TFormat use go:text/template to format
//
// Example1:
//
//     soft.TFormat("Hello {.a}, my name is {.b}", map[string]interface{}{"a": "World", "b": "Zuolar"})
//
// Example2:
//
//     soft.TFormat("My name is {.Name}, {.Age} years old. Again, My Name is {.Name}", struct {
//         Name string
//         Age  int
//     }{
//         Name: "Zuolar",
//         Age:  9999,
//     })
func TFormat(txt string, payload interface{}) string {
	tpl, err := template.New(txt).Delims("{", "}").Parse(txt)
	if err != nil {
		return txt
	}

	w := bytes.NewBufferString("")
	err = tpl.Execute(w, payload)
	if err != nil {
		return txt
	}
	return w.String()
}

// MFormat use key:value to format
//
// Example1:
//
//     soft.MFormat("Hello {a}, my name is {b}", map[string]interface{}{"a": "World", "b": "Zuolar"})
//
// Example2:
//
//     soft.MFormat("My name is {name}, {age} years old. Again, My Name is {name}", map[string]interface{}{"name": "Zuolar", "age": 9999})
func MFormat(txt string, payload map[string]interface{}) string {
	for k, v := range payload {
		txt = strings.ReplaceAll(txt, "{"+k+"}", fmt.Sprint(v))
	}

	return txt
}

// SFormat use index: 0,1,2... to format
//
// Example1:
//
//     soft.SFormat("Hello {0}, my name is {1}", "World", "Zuolar")
//
// Example2:
//
//     soft.SFormat("My name is {0}, {1} years old. Again, My Name is {0}", "Zuolar", 9999)
func SFormat(txt string, payload ...interface{}) string {
	for i, v := range payload {
		txt = strings.ReplaceAll(txt, "{"+strconv.Itoa(i)+"}", fmt.Sprint(v))
	}

	return txt
}
