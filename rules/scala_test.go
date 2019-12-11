package rules

import (
	"testing"
)

func TestParseScala(t *testing.T) {
	file := []byte(`
  package foo.bar

  import x.y.z.Foo
  import a.b._
  import p.{x => a}
  import p.{x, y}

  object test extends Application {
    println("test")
  }
  `)

	result := parseScala(file)

	if len(result["foo.bar"]) != 4 {
		t.Error("expected package 'foo.bar', got", result, result)
	}

	expectListsEq(t, []string{"x.y.z.Foo", "a.b._", "p.{x => a}", "p.{x, y}"}, result["foo.bar"])
}
