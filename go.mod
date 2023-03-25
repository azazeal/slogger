module github.com/azazeal/slogger

go 1.19

require (
	github.com/stretchr/testify v1.8.2
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.0.0-20230322163611-c05651b9ef2a // Published accidentally.
)
