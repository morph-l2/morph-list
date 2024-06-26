package main

import (
	"fmt"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

const (
	SchemaFilePath       = "./src/token.schema.json"
	DevNetTokenListPath  = "./src/devnet/tokenList.json"
	TestNetTokenListPath = "./src/testnet/tokenList.json"
	HoleskyTokenListPath = "./src/holesky/tokenList.json"
	QAnetTokenListPath   = "./src/qanet/tokenList.json"
)

var CheckPath = [...]string{
	DevNetTokenListPath,
	TestNetTokenListPath,
	HoleskyTokenListPath,
	QAnetTokenListPath,
}

func main() {
	// get time now
	currentTime := time.Now().UTC().Format(time.RFC3339)
	_ = currentTime
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + SchemaFilePath)
	for i := 0; i < len(CheckPath); i++ {
		tokenListPath := CheckPath[i]
		documentLoader := gojsonschema.NewReferenceLoader("file://" + tokenListPath)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			panic(err.Error())
		}

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

}
