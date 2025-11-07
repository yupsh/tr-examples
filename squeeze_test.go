package tr_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

func ExampleTr_squeeze() {
	// echo "hello    world" | tr -s ' '
	yup.MustRun(
		Tr(Squeeze, Set1(" "), strings.NewReader("hello    world")),
	)
	// Output:
	// hello world
}

