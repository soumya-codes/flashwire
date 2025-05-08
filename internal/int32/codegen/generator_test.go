package codegen

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	tests := []struct {
		name         string
		source       string
		typeName     string
		expectError  bool
		expectedSnip string // some small string to check in generated output
	}{
		{
			name: "SimpleStructWithInt32",
			source: `
package generated

type DataInput struct {
	Foo int32
	Bar int32
}
`,
			typeName:     "DataInput",
			expectError:  false,
			expectedSnip: "WriteInt32(d.Foo)", // basic snippet presence check
		},
		{
			name: "StructWithMixedTypes",
			source: `
package generated

type MixedInput struct {
	Foo int32
	Name string
	Age int32
}
`,
			typeName:     "MixedInput",
			expectError:  false,
			expectedSnip: "WriteInt32(d.Foo)",
		},
		{
			name: "StructWithNoInt32Fields",
			source: `
package generated

type EmptyInput struct {
	Name string
}
`,
			typeName:    "EmptyInput",
			expectError: true,
		},
		{
			name: "StructNotFound",
			source: `
package generated

type SomethingElse struct {
	ID int32
}
`,
			typeName:    "NotExisting",
			expectError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			inputFile := filepath.Join(tmpDir, "input.go")
			outputDir := filepath.Join(tmpDir, "generated")

			err := os.WriteFile(inputFile, []byte(tc.source), 0644)
			if err != nil {
				t.Fatalf("Failed to write temp input file: %v", err)
			}

			// 💥 Patch here: create output dir
			err = os.MkdirAll(outputDir, 0755)
			if err != nil {
				t.Fatalf("Failed to create output dir: %v", err)
			}

			gen, err := NewGenerator(tc.typeName, outputDir, "codegen", templatesPath())
			if err != nil {
				t.Fatalf("Failed to create generator: %v", err)
			}

			err = gen.Generate(inputFile)
			if tc.expectError {
				if err == nil {
					t.Fatalf("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check generated file exists
			generatedFile := filepath.Join(outputDir, tc.typeName+"_gen.go")
			data, err := os.ReadFile(generatedFile)
			if err != nil {
				t.Fatalf("Failed to read generated file: %v", err)
			}
			if tc.expectedSnip != "" && !contains(string(data), tc.expectedSnip) {
				t.Errorf("Expected snippet %q not found in generated output", tc.expectedSnip)
			}
		})
	}
}

func templatesPath() string {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	return filepath.Join(baseDir, "templates")
}

// helper to check snippet presence
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (len(substr) == 0 || s[len(s)-len(substr):] == substr || s[:len(substr)] == substr || contains(s[1:], substr))
}
