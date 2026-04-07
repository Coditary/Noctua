package main

import (
	"fmt"
	"os"

	coreerrors "github.com/Coditary/Noctua/internal/core/errors"
	"github.com/Coditary/Noctua/internal/platform/bootstrap"
)

func main() {
	app, err := bootstrap.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(coreerrors.ExitCode(err))
	}

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(coreerrors.ExitCode(err))
	}
}
