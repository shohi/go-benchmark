// -*- vim: set ft=gotexttmpl: -*-
package benchmark

// Code generated by go generate; DO NOT EDIT.

import (
	"testing"
    radix "github.com/armon/go-radix"
    iradix "github.com/hashicorp/go-immutable-radix"
)

var keys = []string{
{{- range $i := until . -}}
	{{ $key := (printf "key-%v" .) }}
	"{{ $key }}",
{{- end }}
}

var vals = []string{
{{- range $i := until . -}}
	{{ $val := (printf "vals-%v" .) }}
	"{{ $val }}",
{{- end }}
}

var strMap = make(map[string]string)
var intMap = make(map[int]string)
var intArr = make([]string, 0)
var rdx = radix.New()
var irdx = iradix.New()


func init() {
	for k := 0; k < len(keys); k++ {
		strMap[keys[k]] = vals[k]
		intMap[k] = vals[k]
		intArr = append(intArr, vals[k])
        rdx.Insert(keys[k], vals[k])
        irdx, _, _ = irdx.Insert([]byte(keys[k]), vals[k])
	}
}

func getValUsingSwitchByInt(key int) string {
	switch key {
{{- range $i := until . -}}
	{{ $val := (printf "val-%v" .) }}
	case {{ . }}:
		return "{{ $val }}"
{{- end }}
	default:
		return "unknown"
	}
}

func getValUsingSwitchByString(key string) string {
    switch key {
{{- range $i := until . -}}
	{{ $key := (printf "key-%v" .) }}
	{{ $val := (printf "val-%v" .) }}
	case "{{ $key }}":
		return "{{ $val }}"
{{- end }}
	default:
		return "unknown"
	}
}
func getValUsingMapByInt(key int) string {
	return intMap[key]
}

func getValUsingSliceByInt(key int) string {
	return intArr[key]
}

func getValUsingMapByString(key string) string {
	return strMap[key]
}

func getValUsingMapByRedix(key string) string {
    if val, ok := rdx.Get(key); ok {
        return val.(string)
    }

	return ""
}

func getValUsingMapByiRedix(key string) string {
    if val, ok := irdx.Get([]byte(key)); ok {
        return val.(string)
    }

	return ""
}

func BenchmarkSwitchString(b *testing.B) {
	cases := []struct {
		name string
		fn   func(string) string
	}{
		{"FromSwitch", getValUsingSwitchByString},
		{"FromMap", getValUsingMapByString},
        {"FromRedix", getValUsingMapByRedix},
        {"FromiRedix", getValUsingMapByiRedix},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			var result string
			for i := 0; i < b.N; i++ {
				key := keys[i%len(keys)]
				result = c.fn(key)
			}
			_ = result
		})
	}
}

func BenchmarkSwitchInt(b *testing.B) {
	cases := []struct {
		name string
		fn   func(int) string
	}{
		{"FromSwith", getValUsingSwitchByInt},
		{"FromMap", getValUsingMapByInt},
		{"FromSlice", getValUsingSliceByInt},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			var result string
			for i := 0; i < b.N; i++ {
				result = c.fn(i % len(keys))
			}
			_ = result
		})
	}
}