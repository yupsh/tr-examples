package tr_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tr"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleTr_fromFile_basic() {
	// cat testdata/text.txt | tr 'a-z' 'A-Z'
	yup.MustRun(
		Tr(Set1("a-z"), Set2("A-Z"), yup.File("testdata/text.txt")),
	)
	// Output:
	// HELLO
}

