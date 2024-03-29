[![Build Status](https://github.com/azazeal/slogger/actions/workflows/build.yml/badge.svg)](https://github.com/azazeal/slogger/actions/workflows/build.yml)
[![Coverage Report](https://coveralls.io/repos/github/azazeal/slogger/badge.svg?branch=master)](https://coveralls.io/github/azazeal/slogger?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/azazeal/slogger.svg)](https://pkg.go.dev/github.com/azazeal/slogger)

# slogger

Package `slogger` implements supporting functionality related to the
[`log/slog`](https://pkg.go.dev/log/slog) package.

For more details, you may review the package documentation [here](https://pkg.go.dev/github.com/azazeal/slogger).

## Usage

```go
package main

import (
	"context"

	"github.com/azazeal/slogger"
)

func main() {
	// grab a reference to a slog.Logger configured by the $LOG_LEVEL & $LOG_FORMAT environment
	// variables.
	logger := slogger.FromEnv()

	// store that reference to a top-level Context that propagates throughout the program
	ctx := slogger.NewContext(context.Background(), logger)

	// pass the Context, and therefore the slog.Logger, around; log as needed
	doSomething(ctx)
	doSomethingElse(ctx)
}

func doSomething(ctx context.Context) {
	logger := slogger.FromContext(ctx)

	logger.Warn("did something")
}

func doSomethingElse(ctx context.Context) {
	logger := slogger.FromContext(ctx)

	logger.Info("did something else")
}
```
