package sqlgen

import (
	"testing"
)

func TestColumnString(t *testing.T) {
	var s, e string

	column := Column{"role.name"}

	s = column.Compile(defaultTemplate)
	e = `"role"."name"`

	if s != e {
		t.Fatalf("Got: %s, Expecting: %s", s, e)
	}
}

func TestColumnAs(t *testing.T) {
	var s, e string

	column := Column{"role.name as foo"}

	s = column.Compile(defaultTemplate)
	e = `"role"."name" AS "foo"`

	if s != e {
		t.Fatalf("Got: %s, Expecting: %s", s, e)
	}
}
