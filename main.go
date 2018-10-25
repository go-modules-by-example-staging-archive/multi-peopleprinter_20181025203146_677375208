package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	majBranch "github.com/go-modules-by-example-staging/goinfo-maj-branch/designers"
	majBranch2 "github.com/go-modules-by-example-staging/goinfo-maj-branch/v2/designers"

	majSubdir "github.com/go-modules-by-example-staging/goinfo-maj-subdir/designers"
	majSubdir2 "github.com/go-modules-by-example-staging/goinfo-maj-subdir/v2/designers"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	w := func(format string, args ...interface{}) {
		fmt.Fprintf(tw, format, args...)
	}

	w("github.com/go-modules-by-example-staging/goinfo-maj-branch/designers.Names(): %v\n", majBranch.Names())
	w("github.com/go-modules-by-example-staging/goinfo-maj-branch/v2/designers.FullNames(): %v\n", majBranch2.FullNames())

	w("github.com/go-modules-by-example-staging/goinfo-maj-subdir/designers.Names(): %v\n", majSubdir.Names())
	w("github.com/go-modules-by-example-staging/goinfo-maj-subdir/v2/designers.FullNames(): %v\n", majSubdir2.FullNames())

	tw.Flush()
}
