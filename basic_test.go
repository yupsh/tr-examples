package tr_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

func ExampleTr_basic() {
	// echo "hello" | tr 'a-z' 'A-Z'
	yup.MustRun(
		Tr(Set1("a-z"), Set2("A-Z"), strings.NewReader("hello")),
	)
	// Output:
	// HELLO
}

