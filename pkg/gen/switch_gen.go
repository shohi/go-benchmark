// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates switch_test.go. It can be invoked by running
// go generate
package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

// This example illustrates that the FuncMap *must* be set before the
// templates themselves are loaded.

func main() {
	f, err := os.Create("../benchmark/switch_test.go")
	die(err)
	defer f.Close()

	tplStr, err := ioutil.ReadFile("switch_test.tpl")
	die(err)

	// NOTE: template.ParseFiles does not work here.
	// And also note that `TxtFuncMap` should be used along `text/template`,
	// as `html/template` will escape the content.
	var tpl = template.Must(template.New("base").
		Funcs(sprig.TxtFuncMap()).
		Parse(string(tplStr)))

	count := 100
	// tpl.Execute(os.Stdout, count)
	tpl.Execute(f, count)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
