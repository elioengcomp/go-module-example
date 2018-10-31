package example

import (
	"github.com/kr/pretty"
)

const version = "v3.0.0"

func Exec() string {
	
	return pretty.Sprintf("This is go-module-example project %s", version)
}
