module gosysadmin

go 1.21.0

replace github.com/gosysadmin/basics v0.0.1 => ../basics

require (
	github.com/gosysadmin/basics v0.0.1
	github.com/stretchr/testify v1.9.0
	github.com/valyala/fastjson v1.6.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
