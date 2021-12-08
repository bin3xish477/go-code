package main

/*
$ go test -bench=. -benchmem
*/

import (
        "bytes"
        "strings"
        "testing"
)

func BenchmarkSprintf(b *testing.B) {
        var s string
        for n := 0; n < b.N; n++ {
                s = fmt.Sprintf("%s%s", s, "a")
        }
}

func BenchmarkBytesBuffer(b *testing.B) {
        var buf bytes.Buffer
        for n := 0; n < b.N; n++ {
                buf.WriteString("a")
        }
}

func BenchmarkStringsBuilders(b *testing.B) {
        var builder strings.Builder
        for n := 0; n < b.N; n++ {
                builder.WriteString("a")
        }
}
