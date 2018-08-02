package client

import (
	"text/tabwriter"
	"os"
	"fmt"
)

func listUserRoutesAndCallbacksInATable() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "#", "Route", "Callback")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "--", "-----", "--------")

	i := 1
	for callback, route := range routes.All() {
		fmt.Fprintf(w, "\n %d\t%s \t%s\t", i, route, callback)
		i += 1
	}
}
