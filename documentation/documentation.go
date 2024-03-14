package documentation

// func LoadJSON(path string) string {
// 	jsonBytes, err := os.ReadFile(path)

// 	// jsonBytes, err := os.ReadFile("documentation/users.json")
// 	if err != nil {
// 		fmt.Println("Error reading JSON file:", err)
// 		return ""
// 	}
// 	return string(jsonBytes)
// }

// func LoadSwagger() {
// 	userTemplate := LoadJSON("documentation/users.json")
// 	var UsersSwaggerInfo = &swag.Spec{
// 		InfoInstanceName: "swagger",
// 		SwaggerTemplate:  userTemplate,
// 	}
// 	swag.Register(UsersSwaggerInfo.InstanceName(), UsersSwaggerInfo)
// }
