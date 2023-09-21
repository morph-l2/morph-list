package main

import (
	"fmt"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	schemaLoader := gojsonschema.NewReferenceLoader("file://./src/schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://./src/list.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	// get time now
	currentTime := time.Now().UTC().Format(time.RFC3339)
	_ = currentTime

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		panic("check ERR")
	}
}
