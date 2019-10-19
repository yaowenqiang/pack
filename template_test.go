package pack

import (
	"testing"
	"bytes"
	"text/template"
)

// go test -bench .

func BenchmarkExample(b *testing.B) {
	temp, _ := template.New("").Parse("Hello, Oo")	
	b.ResetTimer()

	var buf bytes.Buffer

	for i := 0; i < b.N; i ++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}
}
