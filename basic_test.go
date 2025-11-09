package nl_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/nl"
)

func ExampleNl_basic() {
	// echo "First line\nSecond line\nThird line" | nl
	gloo.MustRun(
		Nl(strings.NewReader("First line\nSecond line\nThird line")),
	)
	// Output:
	//      1	First line
	//      2	Second line
	//      3	Third line
}
