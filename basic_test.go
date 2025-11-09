package tr_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

func ExampleTr_basic() {
	// echo "hello" | tr 'a-z' 'A-Z'
	gloo.MustRun(
		Tr(Set1("a-z"), Set2("A-Z"), strings.NewReader("hello")),
	)
	// Output:
	// HELLO
}
