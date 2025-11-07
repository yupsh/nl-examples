package nl_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/nl"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleNl_fromFile_basic() {
	// nl testdata/text.txt
	yup.MustRun(
		Nl(yup.File("testdata/text.txt")),
	)
	// Output:
	//      1	First line
	//      2	Second line
	//      3	Third line
}

