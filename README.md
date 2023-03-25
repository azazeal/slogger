[![Build Status](https://github.com/azazeal/slogger/actions/workflows/build.yml/badge.svg)](https://github.com/azazeal/slogger/actions/workflows/build.yml)
[![Coverage Report](https://coveralls.io/repos/github/azazeal/slogger/badge.svg?branch=master)](https://coveralls.io/github/azazeal/slogger?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/azazeal/slogger.svg)](https://pkg.go.dev/github.com/azazeal/slogger)

# slogger

Package `slogger` implements supporting functionality related to the
[`exp/x/slog`](https://pkg.go.dev/golang.org/x/exp/slog) package.

More specifically, it implements functionality that used to exist in `exp/x/slog`, like the
[`NewContext`](https://pkg.go.dev/github.com/azazeal/slogger#NewContext) and
[`FromContext`](https://pkg.go.dev/github.com/azazeal/slogger#FromContext) functions, as well as new helpers like the
`FromEnv` function.

For more details, you may review the package documentation [here](https://pkg.go.dev/github.com/azazeal/slogger).