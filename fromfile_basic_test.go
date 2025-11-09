package nl_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/nl"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleNl_fromFile_basic() {
	// nl testdata/text.txt
	gloo.MustRun(
		Nl(gloo.File("testdata/text.txt")),
	)
	// Output:
	//      1	First line
	//      2	Second line
	//      3	Third line
}
