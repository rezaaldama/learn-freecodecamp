// --------------------------------------------------------------------------------
// Concept							| Go (go mod)					| MERN (npm/yarn)										|
// --------------------------------------------------------------------------------
// Project Config				| go.mod							| package.json											|
// Lock File						| go.sum							| package-lock.json (or yarn.lock)	|
// Dependency Folder		| vendor/ (optional)	| node_modules/ (mandatory)					|
// Install Command			| go mod tidy					| npm install												|
// Initialization				| go mod init					| npm init													|
// --------------------------------------------------------------------------------

// 1. Initiate project
// mkdir project_name
// cd project_name
// go mod init project_name

// 2. Write the code
// code main.go

package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

func init() {
	fmt.Println("Write initial configuration here")
}

// Golangci-lint is a fast, parallel-running aggregate linter for Go.
// golangci-lint run main.go

// 3. Resolve external dependencies
// go mod tidy

// 4. Create vendor folder for external dependencies
// go mod vendor

// 6. Build the executable
// go build -mod=vendor -o hello_world.exe .

// 7. Run the executable
// ./hello_world.exe
