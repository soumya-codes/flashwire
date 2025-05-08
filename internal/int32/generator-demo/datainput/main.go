package main

import (
	"github.com/soumya-codes/flashwire/internal/int32/codegen"
	"log"
)

// To run this code, from the terminal, go to directory where this file is located and run "go run main.go"
func main() {
	gen, err := codegen.NewGenerator("DataInput", "internal/int32/generator-demo/datainput", "main", "internal/int32/codegen/templates")
	if err != nil {
		log.Fatalf("Failed to create generator: %v", err)
	}
	err = gen.Generate("internal/int32/generator-demo/datainput/datainput.go")
	if err != nil {
		log.Fatalf("Failed to generate code: %v", err)
	}
	log.Println("✅ Code generation successful!")
}
