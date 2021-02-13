// A test program for the cli package.
package main

import (
	"fmt"

	"src.elv.sh/pkg/cli"
	"src.elv.sh/pkg/cli/mode"
	"src.elv.sh/pkg/cli/term"
	"src.elv.sh/pkg/cli/tk"
)

func main() {
	app := cli.NewApp(cli.AppSpec{})
	w := mode.NewNavigation(app, mode.NavigationSpec{
		Bindings: tk.MapBindings{
			term.K('x'): func(tk.Widget) { app.CommitCode() },
		},
	})
	app.SetAddon(w, false)

	code, err := app.ReadCode()
	fmt.Println("code:", code)
	if err != nil {
		fmt.Println("err", err)
	}
}
