package tr_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

func ExampleTr_squeeze() {
	// echo "hello    world" | tr -s ' '
	gloo.MustRun(
		Tr(Squeeze, Set1(" "), strings.NewReader("hello    world")),
	)
	// Output:
	// hello world
}
