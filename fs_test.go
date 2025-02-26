package embedvendor_test

import (
	"reflect"
	"testing"

	"github.com/jsteenb2/embedvendor"
)

func TestWalk(t *testing.T) {
	got := embedvendor.WalkFS()

	want := []string{
		"testdata",
		"testdata/foo",
		"testdata/foo/notvendor",
		"testdata/foo/notvendor/baz",
		"testdata/foo/notvendor/baz/file.json",
		"testdata/foo/vendor",
		"testdata/foo/vendor/bar",
		"testdata/foo/vendor/bar/msg.txt",
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("values do not match:\n\twant:\t%v\n\tgot:\t%v", want, got)
	}
}
