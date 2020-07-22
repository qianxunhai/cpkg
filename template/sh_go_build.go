package template
var Sh_go_build = `#!/bin/bash
export GO111MODULE="on"
export GOPROXY=https://goproxy.cn
{{.GoRoot}}/go mod tidy
{{.GoRoot}}/go build -o {{.Name}} main.go`