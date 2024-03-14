package main

import (
	"github.com/hammer-code/lms-be/cmd"
)


// func init() {

// 	docTemplate := Load()
// 	var SwaggerInfo = &swag.Spec{
// 		InfoInstanceName: "swagger",
// 		SwaggerTemplate:  docTemplate,
// 	}
// 	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
// }
func main() {
	cmd.Execute()
}
