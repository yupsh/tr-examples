package tr_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

func ExampleTr_delete() {
	// echo "hello123world" | tr -d '0-9'
	gloo.MustRun(
		Tr(Delete, Set1("0-9"), strings.NewReader("hello123world")),
	)
	// Output:
	// helloworld
}
