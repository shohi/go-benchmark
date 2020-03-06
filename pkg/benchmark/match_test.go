package benchmark

import (
	"regexp"
	"strings"
	"testing"

	"github.com/gobwas/glob"
)

// TODO: performance comparison
// regular express vs strings.Contains vs glob
var results = make([]bool, 0, 10)

func benchMatch(b *testing.B, substr string, fn func(s string) bool) {
	var result bool
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = fn(substr)
	}

	_ = result
	// results = append(results, result)
}

func benchMatch_Contains(b *testing.B, substr string) {
	var result bool

	str := "xxxx-http-client-yyyy"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = strings.Contains(str, substr)
	}

	results = append(results, result)
}

func benchMatch_Glob(b *testing.B, substr string) {
	g, _ := glob.Compile("*http-client*")

	var result bool
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = g.Match(substr)
	}

	results = append(results, result)
}

func benchMatch_Re(b *testing.B, substr string) {
	re := regexp.MustCompile(".*http-client.*")

	var result bool
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = re.MatchString(substr)
	}

	results = append(results, result)
}

func BenchmarkMatch(b *testing.B) {
	substr := "http-client"

	str := "xxxx-http-client-yyyy"
	g, _ := glob.Compile("*http-client*")
	re := regexp.MustCompile(".*http-client.*")

	cases := []struct {
		name    string
		matchFn func(string) bool
	}{
		{"Contains", func(s string) bool { return strings.Contains(str, s) }},
		{"Glob", func(s string) bool { return g.Match(s) }},
		{"Regexp", func(s string) bool { return re.MatchString(s) }},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			benchMatch(b, substr, c.matchFn)
		})
	}

	// fmt.Printf("result ===> %v\n", results)
}
