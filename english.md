# Comprehensive Guide to Learning the Go Programming Language (Golang)

[![Go](https://img.shields.io/badge/Go-1.2x-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)

## Language [English🇬🇧](README.md) & [Indonesia🇮🇩](README.md)
Welcome to the comprehensive learning guide for the Go Programming Language (often called Golang)! This document aims to be a complete resource for anyone looking to learn Go, from the basics to more advanced topics. Go is a compiled, statically-typed programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is known for its simplicity, efficiency, strong built-in concurrency support, and great tooling.

This repository (or document) serves as a roadmap and collection of notes/explanations for your Go learning journey.

## Why Learn Go?

Before diving in, let's understand why Go has become a popular choice for many developers and companies:

1.  **Simplicity:** Go has a clean and minimalist syntax with a relatively small number of keywords (around 25). This makes it easier to learn and read compared to many other languages. Its focus is on code readability and maintainability.
2.  **Built-in Concurrency:** Go was designed with concurrency in mind from the start. Goroutines (concurrently running functions/methods) and Channels (communication mechanisms between goroutines) make it easy to write programs that can handle many tasks simultaneously without the complexity of traditional threads and locking.
3.  **High Performance:** As a language compiled to native machine code, Go offers excellent performance, approaching C/C++ in many cases, while maintaining the development ease of dynamic languages. Its garbage collector is also designed for low latency.
4.  **Fast Compilation:** Go's compilation times are remarkably fast, significantly improving developer productivity during the development cycle (write-compile-test).
5.  **Static Typing:** Go is a statically-typed language, meaning variable types are checked at compile time. This helps catch many errors early in the development process and improves program reliability.
6.  **Rich Standard Library & Tooling:** Go comes with an extensive and robust standard library, covering functionality for networking, cryptography, data processing (like JSON, XML), and much more. The Go command-line tools (`go build`, `go test`, `go fmt`, `go vet`, etc.) are powerful and well-integrated.
7.  **Growing Ecosystem:** The Go community is active and continually growing. Many large open-source projects (like Docker, Kubernetes, Terraform, Prometheus) are written in Go, demonstrating its capability for building large-scale systems and infrastructure.
8.  **Single Binary:** The Go build process produces a single static binary file that includes all necessary dependencies (unless using Cgo). This greatly simplifies the deployment process – just copy and run the binary.

## Prerequisites

To follow this guide effectively, you are expected to have:

*   **Basic Understanding of Programming Concepts:** Familiarity with variables, data types, control structures (if/else, loops), and functions from another programming language will be very helpful.
*   **Comfort with the Command Line/Terminal:** Most Go tooling is run via the terminal.
*   **Text Editor or IDE:** You'll need a place to write code. Popular choices include Visual Studio Code (with the Go extension), GoLand (paid IDE), Vim, or Emacs with Go plugins.

## Table of Contents

*   [Why Learn Go?](#why-learn-go)
*   [Prerequisites](#prerequisites)
*   [1. Installation and Environment Setup](#1-installation-and-environment-setup)
    *   [Downloading and Installing Go](#downloading-and-installing-go)
    *   [Verifying Installation](#verifying-installation)
    *   [Environment Variables (`GOROOT`, `GOPATH`, `GOBIN`)](#environment-variables-goroot-gopath-gobin)
    *   [Introduction to Go Modules](#introduction-to-go-modules)
    *   [Editor/IDE Setup](#editoride-setup)
*   [2. Your First Go Program: "Hello, World!"](#2-your-first-go-program-hello-world)
    *   [Basic Program Structure](#basic-go-program-structure)
    *   [Writing the Code](#writing-the-code)
    *   [Running the Code (`go run`)](#running-the-code-go-run)
    *   [Compiling the Code (`go build`)](#compiling-the-code-go-build)
*   [3. Go Language Basics](#3-go-language-basics)
    *   [Packages and Imports](#packages-and-imports)
    *   [Variables](#variables)
        *   [`var` Declaration](#var-declaration)
        *   [Short Declaration (`:=`)](#short-declaration-)
        *   [Zero Values](#zero-values)
        *   [Variable Scope](#variable-scope)
    *   [Constants](#constants)
    *   [Basic Data Types](#basic-data-types)
        *   [Boolean (`bool`)](#boolean-bool)
        *   [Numeric (`int`, `float64`, `complex128`, etc.)](#numeric-int-float64-complex128-etc)
        *   [String (`string`)](#string-string)
        *   [Rune and Byte](#rune-and-byte)
    *   [Composite Data Types](#composite-data-types)
        *   [Array](#array)
        *   [Slice](#slice)
        *   [Map](#map)
        *   [Struct](#struct)
    *   [Control Flow Structures](#control-flow-structures)
        *   [Branching (`if`, `else`)](#branching-if-else-if-else)
        *   [Branching (`switch`)](#branching-switch)
        *   [Looping (`for`)](#looping-for)
        *   [`break` and `continue`](#break-and-continue)
        *   [`defer`](#defer)
*   [4. Functions](#4-functions)
    *   [Basic Declaration](#basic-function-declaration)
    *   [Parameters and Return Values](#parameters-and-return-values)
    *   [Multiple Return Values](#multiple-return-values)
    *   [Named Return Values](#named-return-values)
    *   [Variadic Functions](#variadic-functions)
    *   [Functions as First-Class Values](#functions-as-first-class-values)
    *   [Anonymous Functions and Closures](#anonymous-functions-and-closures)
*   [5. Pointers](#5-pointers)
    *   [Pointer Concepts](#what-are-pointers)
    *   [Operators `&` and `*`](#operators--address-of-and--dereference)
    *   [When to Use Pointers?](#when-to-use-pointers)
    *   [Pointers to Structs](#pointers-to-structs)
*   [6. Structs and Methods](#6-structs-and-methods)
    *   [Defining Structs](#defining-structs-recap)
    *   [Creating Struct Instances](#creating-struct-instances) <!-- Implicitly covered -->
    *   [Accessing Fields](#accessing-struct-fields) <!-- Implicitly covered -->
    *   [Methods](#methods)
    *   [Pointer vs. Value Receivers](#pointer-receiver-vs-value-receiver)
    *   [Embedding (Composition)](#embedding-composition)
*   [7. Interfaces](#7-interfaces)
    *   [Interface Concepts](#interface-concepts)
    *   [Defining Interfaces](#defining-interfaces)
    *   [Implicit Implementation](#implicit-implementation)
    *   [The Empty Interface (`interface{}`)](#the-empty-interface-interface)
    *   [Type Assertion and Type Switch](#type-assertion-and-type-switch)
    *   [Example: The `error` Interface](#common-example-the-error-interface)
*   [8. Error Handling](#8-error-handling)
    *   [Error Conventions in Go](#error-conventions-in-go)
    *   [Creating Errors (`errors.New`, `fmt.Errorf`)](#creating-errors-errorsnew-fmterrorf)
    *   [Wrapping Errors (Go 1.13+)](#wrapping-errors-error-wrapping---go-113)
    *   [`panic` and `recover`](#panic-and-recover)
*   [9. Concurrency](#9-concurrency)
    *   [Goroutines](#goroutines)
        *   [Starting Goroutines (`go`)](#starting-goroutines-go-keyword)
        *   [Synchronization (`sync.WaitGroup`)](#synchronization-with-syncwaitgroup)
    *   [Channels](#channels)
        *   [Creating Channels (`make`)](#creating-channels-make)
        *   [Sending and Receiving](#sending-and-receiving-data)
        *   [Blocking Operations](#blocking-operations)
        *   [Buffered vs. Unbuffered](#buffered-vs-unbuffered-channels)
        *   [Closing Channels (`close`)](#closing-channels-close)
        *   [Iterating over Channels (`range`)](#iterating-over-channels-with-range)
    *   [`select` Statement](#select-statement)
    *   [Common Concurrency Patterns](#common-concurrency-patterns)
    *   [`sync` Primitives (Mutex, RWMutex, etc.)](#sync-package-mutex-rwmutex-etc)
*   [10. Modules & Packages](#10-modules--packages)
    *   [Project Directory Structure](#project-directory-structure)
    *   [Initializing a Module (`go mod init`)](#initializing-a-module-go-mod-init)
    *   [`go.mod` and `go.sum` Files](#gomod-and-gosum-files)
    *   [Dependency Management (`go get`, `go mod tidy`)](#managing-dependencies-go-get-go-mod-tidy)
    *   [Dependency Versions](#dependency-versions)
    *   [Vendoring (`go mod vendor`)](#vendoring-dependencies-go-mod-vendor)
    *   [Creating Library Packages](#creating-library-packages)
*   [11. Testing in Go](#11-testing-in-go)
    *   [Testing Basics (`testing`, `_test.go`)](#testing-basics-_testgo)
    *   [Writing Tests (`TestXxx`, `*testing.T`)](#writing-tests-testxxx-testingt)
    *   [Running Tests (`go test`)](#running-tests-go-test)
    *   [Test Coverage](#test-coverage)
    *   [Table-Driven Tests](#table-driven-tests)
    *   [Benchmarking (`BenchmarkXxx`, `*testing.B`)](#benchmarking-benchmarkxxx-testingb)
    *   [Example Tests (`ExampleXxx`)](#example-tests-examplexxx)
    *   [Test Fixtures (Helpers, `TestMain`)](#test-fixtures-setup--teardown)
*   [12. Standard Library Highlights](#12-standard-library-highlights)
    *   [`fmt`](#fmt)
    *   [`net/http`](#nethttp)
    *   [`encoding/json`](#encodingjson)
    *   [`io`, `bufio`](#io-bufio)
    *   [`os`](#os)
    *   [`time`](#time)
    *   [`strings`](#strings)
    *   [`strconv`](#strconv)
    *   [`context`](#context)
    *   [`sync`, `sync/atomic`](#sync-syncatomic)
    *   [`database/sql`](#databasesql)
    *   [`log`](#log)
    *   [`flag`](#flag)
    *   [`regexp`](#regexp)
    *   [`testing`](#testing)
    *   *(And many more...)*
*   [13. Building and Deploying Go Applications](#13-building-and-deploying-go-applications)
    *   [Building Executables (`go build`)](#building-executables-go-build)
    *   [Cross-Compilation (`GOOS`, `GOARCH`)](#cross-compilation-goos-goarch)
    *   [Reducing Binary Size (`ldflags`)](#reducing-binary-size-ldflags)
    *   [Deployment Strategies (Copy, Docker, Serverless, PaaS)](#deployment-strategies-copy-docker-serverless-paas)
    *   [Application Configuration (Env, Flags, Files)](#application-configuration-env-flags-files)
*   [14. Advanced Go Topics](#14-advanced-go-topics)
    *   [Reflection (`reflect`)](#reflection-reflect-package)
    *   [Build Constraints / Build Tags](#build-constraints--build-tags)
    *   [Cgo (C Interoperability)](#cgo-c-interoperability)
    *   [Profiling and Optimization (`pprof`)](#profiling-and-optimization-pprof)
    *   [Generics (Go 1.18+)](#generics-go-118)
*   [15. Ecosystem & Tooling](#15-ecosystem--tooling)
    *   [Other `go` commands (`doc`, `vet`, `fmt`, `work`, etc.)](#other-go-commands-doc-vet-fmt-work-etc)
    *   [Linters (`staticcheck`, `golangci-lint`)](#linters-staticcheck-golangci-lint)
    *   [Debuggers (`delve`)](#debuggers-delve)
    *   [Popular Web Frameworks (Gin, Echo, Chi, Fiber)](#popular-web-frameworks-gin-echo-chi-fiber)
    *   [ORMs and Database Libraries (GORM, sqlx, SQLBoiler)](#orms-and-database-libraries-gorm-sqlx-sqlboiler)
    *   [CLI Libraries (Cobra, Viper)](#cli-libraries-cobra-viper)
    *   [gRPC](#grpc)
    *   [Structured Logging (Zerolog, Zap)](#structured-logging)
*   [16. Best Practices & Style Guide](#16-best-practices--style-guide)
    *   [Effective Go](#effective-go)
    *   [Naming](#naming-packages-variables-functions-etc)
    *   [Formatting (`gofmt`/`goimports`)](#formatting-gofmtgoimports)
    *   [Error Handling](#error-handling-1) <!-- Added -1 to differentiate from section 8 -->
    *   [Interface Size](#interface-size)
    *   [Composition over Inheritance](#composition-over-inheritance)
    *   [Avoiding Global State](#avoiding-global-state)
*   [Contributing](#contributing)
*   [Additional Resources](#additional-resources)
*   [License](#license)

---

## 1. Installation and Environment Setup

The first step is to install Go on your system and configure the development environment.

### Downloading and Installing Go

Visit the official Go download page: [https://golang.org/dl/](https://golang.org/dl/)

Choose the installer package suitable for your operating system (Windows, macOS, Linux). Follow the installation instructions:

*   **Windows:** Use the MSI installer. It will set the path automatically.
*   **macOS:** Use the PKG installer or install via Homebrew (`brew install go`).
*   **Linux:** Download the `.tar.gz` archive, extract it to `/usr/local` (or another preferred location), and add `/usr/local/go/bin` to your `PATH` environment variable.
    ```bash
    # Example for Linux (adjust version)
    wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin # Add this to ~/.profile or ~/.bashrc
    source ~/.profile # or source ~/.bashrc
    ```

### Verifying Installation

Open a new terminal or command prompt and run the following command:

```bash
go version
```

You should see output displaying the installed Go version, e.g., `go version go1.21.0 linux/amd64`.

Also run:

```bash
go env
```

This command will display environment variables relevant to Go, such as `GOROOT` and `GOPATH`.

### Environment Variables (`GOROOT`, `GOPATH`, `GOBIN`)

*   **`GOROOT`**: The location of your Go installation. Usually set automatically by the installer. You generally *do not* need to change this.
*   **`GOPATH`**: (Less central since Go Modules, but still relevant) Specifies the location of your *workspace*. By default, this is usually `$HOME/go` (Linux/macOS) or `%USERPROFILE%\go` (Windows). `GOPATH` contains directories:
    *   `src/`: Where your project source code used to reside (less common for new projects with Modules).
    *   `pkg/`: Contains compiled package objects.
    *   `bin/`: Contains compiled executable binaries (`go install`).
*   **`GOBIN`**: (Optional) Specifies the directory where binaries installed with `go install` will be placed. If unset, binaries go into `$GOPATH/bin`. Ensure the `bin` directory (either `$GOPATH/bin` or `$GOBIN`) is in your system `PATH` to run installed binaries from anywhere.

### Introduction to Go Modules

Since Go 1.11, **Go Modules** are the standard way to manage dependencies. You no longer *have* to place your code inside `GOPATH`. Modules allow you to explicitly define project dependencies in a `go.mod` file. We'll cover this in more detail later. For now, know that you can create Go projects in *any* directory on your system.

### Editor/IDE Setup

Configure your text editor or IDE for Go development:

*   **VS Code:** Install the official `Go` extension from Microsoft. It provides features like IntelliSense, linting, debugging, code navigation, automatic formatting with `gofmt` on save, etc.
*   **GoLand:** A commercial IDE from JetBrains dedicated to Go. Very powerful and feature-rich.
*   **Vim/Emacs/Sublime Text:** Many plugins are available to support Go development (e.g., `vim-go`, `go-mode.el`).

Ensure essential features like automatic formatting (`gofmt` or `goimports`) and linting are enabled to maintain code consistency and quality.

---

## 2. Your First Go Program: "Hello, World!"

Let's create the simplest Go program to ensure everything works.

### Basic Go Program Structure

Every executable Go program must have:

1.  A `package main` declaration.
2.  A `func main()` function, which is the program's entry point.
3.  (Usually) `import` statements to use functionality from other packages.

### Writing the Code

1.  Create a new directory for your first project *outside* `GOPATH` (if you're using Go Modules). For example:
    ```bash
    mkdir ~/hello-go
    cd ~/hello-go
    ```
2.  Initialize a Go module (optional but recommended for real projects):
    ```bash
    go mod init example.com/hello-go
    # Replace example.com/hello-go with your unique module path
    # This will create a go.mod file
    ```
3.  Create a new file named `main.go` (or another name ending in `.go`).
4.  Write the following code inside `main.go`:

    ```go
    // Main package declaration, required for executable programs
    package main

    // Import the "fmt" package, which contains functions for formatted I/O (like printing)
    import "fmt"

    // The main function, the program's entry point
    func main() {
        // Call the Println function from the fmt package to print text to the console
        fmt.Println("Hello, Go World!")
    }
    ```

### Running the Code (`go run`)

The `go run` command compiles and runs Go source files directly without creating a permanent binary file.

```bash
go run main.go
```

Expected output:

```
Hello, Go World!
```

### Compiling the Code (`go build`)

The `go build` command compiles the Go source code and its dependencies into an executable binary file.

```bash
go build
# Or: go build main.go
```

This command will create a binary file in the current directory. The binary's name will be the same as the directory name (if `go build` is run without file arguments) or the filename (without the `.go` extension, if a filename is given). If you used `go mod init example.com/hello-go`, the binary will likely be named `hello-go` (based on the last part of the module path).

To run it:

*   **Linux/macOS:** `./hello-go`
*   **Windows:** `hello-go.exe`

The output will be the same: `Hello, Go World!`

The advantage of `go build` is producing a self-contained binary that can be distributed and run without needing the source code or a Go installation on the target machine (as long as the architecture and OS match).

---

## 3. Go Language Basics

Let's learn the fundamental elements of Go.

### Packages and Imports

Go organizes code into *packages*. Every Go file starts with a `package` declaration.

*   `package main`: Signifies this package will compile into an executable program.
*   Other packages (e.g., `package mylib`): Signify this package is a library usable by other packages. The package name is usually the same as the directory name containing the `.go` files.
*   `import` statement: Used to access functionality from other packages.
    ```go
    import (
        "fmt"        // Standard package for I/O
        "math/rand" // Standard package for random numbers
        "os"         // Standard package for OS interactions

        "github.com/gin-gonic/gin" // Example import of an external package (via Go Modules)
        myAlias "example.com/mylib/internal/utils" // Import with an alias
        _ "github.com/lib/pq" // "Blank" import (underscore): only runs the package's init() function(s), without using the package name directly. Useful for side effects like registering database drivers.
    )
    ```
*   Convention: Package names are lowercase. If an import path is long, you can provide an *alias*.
*   Visibility: Names (functions, types, variables, constants) starting with a **capital letter** are Exported and can be accessed from other packages. Names starting with a **lowercase letter** are unexported and only accessible within the same package.

### Variables

Variables are used to store data. Go is a *statically typed* language, so the variable's type must be known at compile time.

#### `var` Declaration

The most explicit form. Can be used at the package level or inside functions.

```go
package main

import "fmt"

// Variable declaration at package level
var globalMessage string = "This is a global message"
var isGlobal bool // Will be initialized with its zero value (false)

func main() {
    // Variable declaration at function level
    var i int        // Declaration without initialization (gets zero value: 0)
    var f float64 = 3.14 // Declaration with initialization
    var s string     // Zero value: "" (empty string)
    var b bool = true

    fmt.Println(i, f, s, b) // Output: 0 3.14  true

    // Declaring multiple variables at once
    var x, y int = 10, 20
    var (
        firstName    string = "Budi"
        lastName     string = "Santoso"
        age          int    = 30
    )
    fmt.Println(x, y)           // Output: 10 20
    fmt.Println(firstName, lastName, age) // Output: Budi Santoso 30
    fmt.Println(globalMessage) // Output: This is a global message
}
```

#### Short Declaration (`:=`)

A more concise way to declare *and* initialize variables. **Can only be used inside functions.** Go will *infer* the data type from the assigned value.

```go
func main() {
    // Short declaration
    message := "Hello from short declaration!" // Type inferred as string
    count := 100                         // Type inferred as int
    isValid := true                      // Type inferred as bool
    pi := 3.14159                       // Type inferred as float64

    fmt.Println(message, count, isValid, pi)

    // Short declaration for multiple variables
    host, port := "localhost", 8080 // host: string, port: int
    fmt.Println(host, port)

    // Important: `:=` is only for new declarations. To assign a new value to an existing variable, use `=`
    count = 101 // Assigning a new value (assignment)
    // count := 102 // This would error if count is already declared in the same scope
    fmt.Println(count)

    // Can be used if at least one variable on the left side is new
    host, err := "127.0.0.1", error(nil) // ok if err doesn't exist yet
    fmt.Println(host, err)
}
```

#### Zero Values

If a variable is declared without an explicit initial value, Go initializes it with its *zero value* based on its type:

*   `0` for numeric types (int, float, etc.)
*   `false` for boolean type (`bool`)
*   `""` (empty string) for `string` type
*   `nil` for pointers, interfaces, slices, maps, channels, and function types.

#### Variable Scope

*   **Package Level:** Variables declared outside any function are accessible to all functions within that package (and exported if their name starts with a capital letter).
*   **Function Level:** Variables declared inside a function (including parameters) are only accessible within that function.
*   **Block Level:** Variables declared inside a block (e.g., inside an `if` or `for`) are only accessible within that block.

```go
package main

import "fmt"

var packageVar = "I am at package level"

func main() {
	funcVar := "I am at function level"
	fmt.Println(packageVar) // OK
	fmt.Println(funcVar)   // OK

	if true {
		blockVar := "I am at if block level"
		fmt.Println(blockVar) // OK
		fmt.Println(funcVar)  // OK (variables from outer scope still accessible)

        // Shadowing: Declaring a variable with the same name in an inner scope
        funcVar := "I am funcVar inside if"
        fmt.Println(funcVar) // Output: I am funcVar inside if
	}

    fmt.Println(funcVar)   // Output: I am at function level (the original one)
	// fmt.Println(blockVar) // Error: undefined: blockVar (outside if block scope)
}
```

### Constants

Constants are values fixed at compile time and cannot be changed during runtime. Declared using the `const` keyword.

```go
package main

import "fmt"

// Package level constants
const Pi float64 = 3.14159265359
const AppName = "Cool Application" // Type inferred (string)

// Constants can be typed or untyped
const typedInt int = 100
const untypedInt = 200 // More flexible, can be used in contexts requiring different numeric types

// iota: sequential number generator for constants
const (
    Sunday = iota // 0
    Monday        // 1 (auto increments)
    Tuesday       // 2
    Wednesday     // 3
    Thursday      // 4
    Friday        // 5
    Saturday      // 6
)

const (
    // Can be used in expressions
    _  = iota             // Ignore 0
    KB = 1 << (10 * iota) // 1 << (10*1) = 1024 (Kilobyte)
    MB = 1 << (10 * iota) // 1 << (10*2) = 1048576 (Megabyte)
    GB = 1 << (10 * iota) // 1 << (10*3) (Gigabyte)
    TB = 1 << (10 * iota) // 1 << (10*4) (Terabyte)
)

func main() {
	const localConst = "This is a local constant"

	fmt.Println(Pi)
	fmt.Println(AppName)
	fmt.Println(localConst)

	fmt.Println("Day:", Monday, Saturday) // Output: Day: 1 6

	fmt.Printf("1 GB = %d Bytes\n", GB) // Output: 1 GB = 1073741824 Bytes

	// typedInt + untypedInt // OK, untyped will adapt
	// var myFloat float32 = typedInt // Error: cannot use typedInt (type int) as float32
	var myFloat float32 = untypedInt // OK: untypedInt can become float32
	fmt.Println(myFloat) // Output: 200
}
```

Constants must have a value determinable at compile time.

### Basic Data Types

#### Boolean (`bool`)

Value `true` or `false`. Zero value: `false`.

```go
var isActive bool = true
var isEnabled = false // Type inferred
```

#### Numeric (`int`, `float64`, `complex128`, etc.)

*   **Integer:**
    *   `int`: Platform-dependent size (32 or 64 bits). Use this for general integers.
    *   `int8`, `int16`, `int32`, `int64`: Specific bit sizes, signed.
    *   `uint`: Platform-dependent size, unsigned.
    *   `uint8` (alias: `byte`), `uint16`, `uint32`, `uint64`: Specific bit sizes, unsigned.
    *   `uintptr`: An unsigned integer large enough to store the bits of a pointer value.
    Zero value: `0`.
*   **Floating-Point:**
    *   `float32`: IEEE-754 single-precision.
    *   `float64`: IEEE-754 double-precision. Use this for general floats.
    Zero value: `0.0`.
*   **Complex:**
    *   `complex64`: Real and imaginary parts are `float32`.
    *   `complex128`: Real and imaginary parts are `float64`.
    Zero value: `(0+0i)`.

```go
var age int = 25
var temperature float64 = 26.5
var complexValue complex128 = 2 + 3i
var dataByte byte = 'A' // Alias for uint8
```

**Important:** Go does not perform implicit numeric type conversions. You must perform explicit conversions.

```go
var a int = 10
var b float64 = 3.5
// var c = a + b // Error: invalid operation: a + b (mismatched types int and float64)
var c = float64(a) + b // OK: Convert int to float64
fmt.Println(c) // Output: 13.5
```

#### String (`string`)

Represents a sequence of bytes, usually containing UTF-8 encoded text. Strings in Go are **immutable** (cannot be changed after creation). Zero value: `""`.

```go
var greeting string = "Hello"
var name = "Go Developer"

// Concatenation
var message = greeting + ", " + name + "!"
fmt.Println(message) // Output: Hello, Go Developer!

// Raw string literal
var multiLine = `This is a string
that can span
multiple lines.
Interpolation does not work here: ${var}`
fmt.Println(multiLine)

// Accessing bytes (not Unicode characters)
fmt.Println(greeting[0]) // Output: 72 (ASCII/UTF-8 code for 'H')

// String length (number of bytes)
fmt.Println(len(greeting)) // Output: 5

// To iterate over Unicode characters, use `range`
for index, runeValue := range "Go语言" {
    // Note: %c prints the character, %U prints the Unicode code point
    fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", index, runeValue, runeValue)
}
/* Output:
Index: 0, Rune: G, Unicode: U+0047
Index: 1, Rune: o, Unicode: U+006F
Index: 2, Rune: 语, Unicode: U+8BED
Index: 5, Rune: 言, Unicode: U+8A00  (Note the index jumps because Chinese characters require 3 bytes)
*/
```

#### Rune and Byte

*   `rune`: Alias for `int32`. Represents a Unicode *code point*. Used when working with individual characters in a string. Rune literals are written with single quotes: `'a'`, `'好'`.
*   `byte`: Alias for `uint8`. Represents a single byte of data. Useful for binary data or ASCII strings. Byte literals are the same as ASCII rune literals: `'A'`, `'\n'`.

### Composite Data Types

Types built from basic or other composite types.

#### Array

A sequence of elements with a **fixed size** and the same element type. The size is part of the array's type. Less flexible than slices, so used less often directly. Zero value: an array with zero values for each element.

```go
// Integer array with 5 elements
var numbers [5]int
numbers[0] = 10
numbers[1] = 20
fmt.Println(numbers)        // Output: [10 20 0 0 0]
fmt.Println(len(numbers))   // Output: 5

// Declaration with literal
primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes) // Output: [2 3 5 7 11 13]

// Let the compiler count the size (...)
names := [...]string{"Alice", "Bob", "Charlie"}
fmt.Println(len(names)) // Output: 3
fmt.Println(names)      // Output: [Alice Bob Charlie]

// Multi-dimensional array
var matrix [2][3]int
matrix[0] = [3]int{1, 2, 3}
matrix[1] = [3]int{4, 5, 6}
fmt.Println(matrix) // Output: [[1 2 3] [4 5 6]]
```

#### Slice

A **dynamic and flexible** view into the elements of an *underlying array*. Slices are far more common than arrays in Go. A slice consists of three components:
1.  A pointer to the first element of the array visible to the slice.
2.  Length: The number of elements in the slice.
3.  Capacity: The number of elements in the underlying array from the slice's first element to the end of the array.

Zero value: `nil`. A `nil` slice has length 0 and capacity 0.

```go
// Slice of strings (underlying array created automatically)
var fruits []string // nil slice
fmt.Println(fruits == nil) // Output: true

// Create a slice with make(type, length, capacity)
// Capacity is optional, defaults to length
numbers := make([]int, 3, 5) // len=3, cap=5
fmt.Println(numbers)       // Output: [0 0 0]
fmt.Println("Len:", len(numbers), "Cap:", cap(numbers)) // Output: Len: 3 Cap: 5

// Create a slice with a literal (like array, but without size)
colors := []string{"Red", "Green", "Blue"}
fmt.Println(colors)        // Output: [Red Green Blue]
fmt.Println("Len:", len(colors), "Cap:", cap(colors)) // Output: Len: 3 Cap: 3

// Slicing: Create a new slice from an existing array or slice
// slice[low:high] -> elements from index low up to (but not including) high
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4] // Elements at index 1, 2, 3
fmt.Println(s)            // Output: [3 5 7]
fmt.Println("Len:", len(s), "Cap:", cap(s)) // Output: Len: 3 Cap: 5 (from index 1 to the end of primes)

s2 := colors[0:2]
fmt.Println(s2) // Output: [Red Green]

// Defaults: low=0, high=len(slice)
s3 := colors[:2] // Same as colors[0:2]
s4 := colors[1:] // Same as colors[1:len(colors)]
s5 := colors[:]  // Slice copy (shares the same underlying array)

// Add elements to a slice with `append`
// append can grow the underlying array if capacity is insufficient
numbers = append(numbers, 10)
fmt.Println(numbers)        // Output: [0 0 0 10]
fmt.Println("Len:", len(numbers), "Cap:", cap(numbers)) // Output: Len: 4 Cap: 5

numbers = append(numbers, 20, 30) // Capacity exceeded
fmt.Println(numbers)        // Output: [0 0 0 10 20 30]
// New capacity will be allocated (usually doubled)
fmt.Println("Len:", len(numbers), "Cap:", cap(numbers)) // Output: Len: 6 Cap: 10 (or another value depending on allocation strategy)

// Combine two slices
moreColors := []string{"Yellow", "Purple"}
allColors := append(colors, moreColors...) // Note the ...
fmt.Println(allColors) // Output: [Red Green Blue Yellow Purple]

// Iterate over a slice with `range`
for index, value := range allColors {
    fmt.Printf("Index: %d, Color: %s\n", index, value)
}
// If only value is needed: for _, value := range allColors { ... }
// If only index is needed: for index := range allColors { ... }

// Multi-dimensional slice
board := [][]string{
    {"_", "_", "_"},
    {"_", "_", "_"},
    {"_", "_", "_"},
}
board[0][0] = "X"
board[1][1] = "O"
fmt.Println(board) // Output: [[X _ _] [_ O _] [_ _ _]]
```

**Important:** Because slices are *views* into an underlying array, modifying elements through one slice will be visible to other slices referencing the same array. `append` *may* return a new slice with a new underlying array if the capacity is exceeded.

#### Map

An unordered collection of **key-value** pairs. Keys must be of a *comparable* data type (can be compared with `==`, like `string`, `int`, `bool`, `pointer`, `struct` without slice/map/function fields). Values can be of any type. Zero value: `nil`. A `nil` map cannot be written to; it must be initialized with `make` or a literal.

```go
// Map with string keys and int values
var scores map[string]int // nil map
// scores["Alice"] = 100 // Runtime error: assignment to entry in nil map

// Initialize with make
ages := make(map[string]int)
ages["Alice"] = 30
ages["Bob"] = 25
fmt.Println(ages)        // Output: map[Alice:30 Bob:25] (order not guaranteed)
fmt.Println(ages["Alice"]) // Output: 30

// Initialize with a literal
populations := map[string]int{
    "Jakarta":  10_000_000,
    "Surabaya": 3_000_000,
    "Bandung":  2_500_000, // Trailing comma required if brace is on a new line
}
fmt.Println(populations)

// Accessing elements
fmt.Println(populations["Jakarta"]) // Output: 10000000

// Accessing a non-existent key returns the zero value for the value type
fmt.Println(populations["Medan"]) // Output: 0 (zero value for int)

// Checking for key existence (comma ok idiom)
pop, ok := populations["Medan"]
if ok {
    fmt.Println("Medan Population:", pop)
} else {
    fmt.Println("Medan data not found.") // This will be printed
}

val, exists := ages["Charlie"]
fmt.Println("Charlie exists?", exists, "Age:", val) // Output: Charlie exists? false Age: 0

// Updating a value
ages["Alice"] = 31
fmt.Println(ages)

// Deleting an element
delete(ages, "Bob")
fmt.Println(ages) // Output: map[Alice:31]

// Map length (number of key-value pairs)
fmt.Println(len(populations)) // Output: 3

// Iterate over a map with `range` (order not guaranteed!)
for key, value := range populations {
    fmt.Printf("City: %s, Population: %d\n", key, value)
}
// If only key is needed: for key := range populations { ... }
// If only value is needed: for _, value := range populations { ... }
```

#### Struct

A composite data type that groups *fields* (member data) of potentially different types under a single name. Useful for representing real-world entities or structured data. Zero value: a struct with zero values for each of its fields.

```go
package main

import "fmt"

// Define a struct type named Person
type Person struct {
    FirstName string // Exported field (capital letter)
    LastName  string
    Age       int
    isMarried bool // Unexported field (lowercase letter)
}

// Define another struct type
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Structs can contain other structs (composition)
type Contact struct {
	Email       string
	Phone       string
	HomeAddress Address // Field of type Address
}

func main() {
    // Create an instance (variable) of the Person struct
    var p1 Person
    p1.FirstName = "Andi"
    p1.LastName = "Wijaya"
    p1.Age = 28
    p1.isMarried = false
    fmt.Println(p1) // Output: {Andi Wijaya 28 false}
    fmt.Println("First Name:", p1.FirstName) // Output: First Name: Andi

    // Create an instance with a struct literal
    p2 := Person{
        FirstName: "Siti",
        LastName:  "Aminah",
        Age:       32,
        isMarried: true, // Note: can access unexported fields within the same package
    }
    fmt.Println(p2) // Output: {Siti Aminah 32 true}

    // Struct literal without field names (order must match definition) - less recommended
    p3 := Person{"Rudi", "Hartono", 40, true}
    fmt.Println(p3)

    // Struct literal filling only some fields (rest get zero values)
    p4 := Person{FirstName: "Dewi"}
    fmt.Println(p4) // Output: {Dewi  0 false}

    // Create a Contact instance with an Address
    c1 := Contact{
        Email: "andi.w@example.com",
        Phone: "0812345678",
        HomeAddress: Address{
            Street: "Jl. Merdeka No. 10",
            City:   "Jakarta",
            ZipCode:"10110",
        },
    }
    fmt.Println(c1)
    fmt.Println("City:", c1.HomeAddress.City) // Output: City: Jakarta

    // Pointer to a struct
    p5 := &Person{"Bambang", "Pamungkas", 42, true} // p5 is *Person
    fmt.Println(p5.FirstName) // Go automatically dereferences: (*p5).FirstName
    // Same as: fmt.Println((*p5).FirstName)

    // Anonymous struct (struct without a type name, defined inline)
    point := struct {
        X int
        Y int
    }{
        X: 10,
        Y: 20,
    }
    fmt.Println(point) // Output: {10 20}
}
```

### Control Flow Structures

Manage the order of statement execution in a program.

#### Branching (`if`, `else if`, `else`)

Execute blocks of code based on boolean conditions.

```go
package main

import "fmt"

func main() {
    score := 75
    grade := ""

    if score >= 90 {
        grade = "A"
    } else if score >= 80 {
        grade = "B"
    } else if score >= 70 {
        grade = "C"
    } else if score >= 60 {
        grade = "D"
    } else {
        grade = "E"
    }
    fmt.Println("Score:", score, "Grade:", grade) // Output: Score: 75 Grade: C

    // If with a short statement (variable declaration within the if)
    // The 'num' variable is only scoped to the if-else if-else block
    if num := 10; num%2 == 0 {
        fmt.Printf("%d is even\n", num) // Output: 10 is even
    } else {
        fmt.Printf("%d is odd\n", num)
    }
    // fmt.Println(num) // Error: undefined: num

    // Conditions don't need parentheses `()`
    x := 5
    if x > 0 { // if (x > 0) { } -> not idiomatic
        fmt.Println("x is positive")
    }
}
```

#### Branching (`switch`)

A cleaner alternative to long `if-else if-else` chains, especially when comparing a single value against multiple possibilities.

```go
package main

import (
	"fmt"
	"runtime" // Needed for the OS switch example
)

func main() {
    day := "Monday"
    activity := ""

    switch day {
    case "Monday":
        activity = "Start of week meeting"
        // No automatic fallthrough (no `break` needed)
    case "Tuesday", "Wednesday", "Thursday": // Multiple cases
        activity = "Routine work"
    case "Friday":
        activity = "Weekly review & weekend prep"
    case "Saturday", "Sunday":
        activity = "Weekend!"
    default: // Optional
        activity = "Invalid day"
    }
    fmt.Printf("Day %s: %s\n", day, activity) // Output: Day Monday: Start of week meeting

    // Switch without an expression (like if-else if-else)
    score := 85
    grade := ""
    switch { // Equivalent to: switch true { ... }
    case score >= 90:
        grade = "A"
    case score >= 80:
        grade = "B"
    default:
        grade = "C or lower"
    }
    fmt.Println("Grade:", grade) // Output: Grade: B

    // Switch with a short statement
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("macOS")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("%s\n", os)
    }

    // Fallthrough: explicitly continue to the next case
    num := 1
    switch num {
    case 1:
        fmt.Println("One")
        fallthrough // Continue to case 2
    case 2:
        fmt.Println("Two (or fallthrough from 1)")
    case 3:
        fmt.Println("Three")
    }
    /* Output:
    One
    Two (or fallthrough from 1)
    */

    // Type Switch (covered more in Interfaces section)
    var i interface{} = "hello"
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    } // Output: String: hello
}
```

#### Looping (`for`)

Go has only one looping construct: `for`. But it can be used in various ways.

```go
package main

import "fmt"

func main() {
    // 1. C-style form (init; condition; post)
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println("Sum (C-style):", sum) // Output: Sum (C-style): 45

    // 2. "While" form (condition only)
    // Init and post statements are optional
    n := 1
    for n < 100 { // Same as while (n < 100) in other languages
        n *= 2
    }
    fmt.Println("n (while-style):", n) // Output: n (while-style): 128

    // 3. "Infinite loop" form (no condition)
    /*
       count := 0
       for { // Infinite loop
           fmt.Println("Looping...")
           count++
           if count > 2 {
               break // Must have a way to exit
           }
       }
    */

    // 4. `for range` form (for iterating over slices, arrays, maps, strings, channels)
    // Slice
    items := []string{"apple", "banana", "cherry"}
    fmt.Println("Iterating Slice:")
    for index, value := range items {
        fmt.Printf("  Index: %d, Value: %s\n", index, value)
    }

    // Map (order not guaranteed)
    capitals := map[string]string{"Indonesia": "Jakarta", "Japan": "Tokyo"}
    fmt.Println("\nIterating Map:")
    for country, capital := range capitals {
        fmt.Printf("  Capital of %s is %s\n", country, capital)
    }

    // String (iterating Unicode runes/characters)
    fmt.Println("\nIterating String:")
    for i, r := range "Go€" {
        fmt.Printf("  Byte index: %d, Rune: %c\n", i, r)
    }
    /* Output:
      Byte index: 0, Rune: G
      Byte index: 1, Rune: o
      Byte index: 2, Rune: €
    */
}
```

#### `break` and `continue`

*   `break`: Exits the innermost `for` loop or `switch` statement.
*   `continue`: Proceeds to the next iteration of the innermost `for` loop.

```go
package main

import "fmt"

func main() {
    // Example break
    fmt.Println("Example break:")
    for i := 0; i < 10; i++ {
        if i == 5 {
            fmt.Println("  Stopping at i=5")
            break // Exit the for loop
        }
        fmt.Printf("  i = %d\n", i)
    }

    // Example continue
    fmt.Println("\nExample continue (skip even numbers):")
    for j := 0; j < 6; j++ {
        if j%2 == 0 {
            continue // Skip the rest of the block, proceed to the next iteration (j++)
        }
        fmt.Printf("  j = %d (odd)\n", j)
    }
    /* Output:
      j = 1 (odd)
      j = 3 (odd)
      j = 5 (odd)
    */

    // Labels with break/continue (rarely used, can make code harder to read)
    fmt.Println("\nExample break with label:")
OuterLoop: // Label
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            fmt.Printf("  x=%d, y=%d\n", x, y)
            if x == 1 && y == 1 {
                fmt.Println("    Breaking out of OuterLoop")
                break OuterLoop // Exit the loop labeled OuterLoop
            }
        }
    }
}
```

#### `defer`

Schedules a function (or method) call to be executed *just before* the function containing the `defer` statement finishes (either by reaching the end, a `return`, or a `panic`). Deferred calls are executed in **Last-In, First-Out (LIFO)** order.

Extremely useful for cleaning up resources like closing files, unlocking mutexes, or closing database connections.

```go
package main

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	fmt.Println("Opening file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err // defer won't run if error occurs before defer statement
	}
	// Schedule file.Close() to run just before readFile finishes
	// This will always run, whether an error occurs after this or not
	defer fmt.Println("Scheduling file close...") // This call runs second
	defer file.Close() // This call runs LAST (LIFO)
	defer fmt.Println("Second defer call") // This call runs FIRST

	fmt.Println("Reading file contents...")
	// ... (file reading logic) ...
	// Suppose an error occurs during reading
	// if someCondition { return fmt.Errorf("error during read") } -> defers still run

	fmt.Println("Finished reading file (before return/end of function)")
	return nil // defers run before returning nil
}

func main() {
	fmt.Println("--- Starting main ---")
	// Create dummy file for example
	dummyFile := "dummy.txt"
	f, _ := os.Create(dummyFile)
	f.WriteString("test data")
	f.Close()

	err := readFile(dummyFile)
	if err != nil {
		fmt.Println("Main received error:", err)
	} else {
		fmt.Println("Main: readFile finished without error")
	}

	// Clean up dummy file
	os.Remove(dummyFile)

	fmt.Println("\n--- Defer Order Example ---")
	fmt.Println("One")
	defer fmt.Println("Four (first defer)")
	fmt.Println("Two")
	defer fmt.Println("Three (second defer)")
	fmt.Println("--- Finishing main ---")
    // Deferred output appears after "--- Finishing main ---"
    // Three (second defer)
    // Four (first defer)
}

/* Example readFile Output:
Opening file: dummy.txt
Reading file contents...
Finished reading file (before return/end of function)
Second defer call
Scheduling file close...
(file.Close() runs here)
Main: readFile finished without error
*/
```
**Important:** Arguments to deferred functions are evaluated *when the `defer` statement is executed*, not when the deferred call runs.

```go
func exampleDeferArgs() {
	i := 0
	// The value of 'i' (0) is captured when defer is called
	defer fmt.Println("Value of i when defer executes:", i)
	i++
	fmt.Println("Value of i before return:", i) // Output: 1
    // The deferred call will print: Value of i when defer executes: 0
}
```

---

## 4. Functions

Functions are reusable blocks of code that perform a specific task.

### Basic Function Declaration

```go
package main

import "fmt"

// Function with no parameters and no return value
func sayHello() {
    fmt.Println("Hello!")
}

// The main function is a basic example
func main() {
    sayHello() // Calling the function
}
```

### Parameters and Return Values

Functions can accept parameters (input) and return values (output). The data types of parameters and return values must be declared.

```go
package main

import "fmt"

// Function with one parameter (name string) and one return value (string)
func greet(name string) string {
    message := "Hello, " + name + "!"
    return message // Return a string value
}

// Function with two parameters (a, b int) and one return value (int)
func add(a int, b int) int {
    return a + b
}

// If consecutive parameters have the same type,
// the type can be written only for the last one
func multiply(x, y int, factor float64) float64 {
    return float64(x*y) * factor
}

func main() {
    greeting := greet("Go User")
    fmt.Println(greeting) // Output: Hello, Go User!

    sum := add(5, 3)
    fmt.Println("5 + 3 =", sum) // Output: 5 + 3 = 8

    result := multiply(4, 5, 1.5)
    fmt.Println("4 * 5 * 1.5 =", result) // Output: 4 * 5 * 1.5 = 30.0
}
```

### Multiple Return Values

Functions in Go can return more than one value. This is often used to return a result and an error status simultaneously.

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Function returns two values: int and error
func divide(numerator, denominator int) (int, error) {
    if denominator == 0 {
        // Return zero value for int (0) and an error
        return 0, errors.New("cannot divide by zero")
    }
    // Return result and nil for error (indicating no error)
    return numerator / denominator, nil
}

// Function returns string and bool
func checkLength(s string) (string, bool) {
	if len(s) > 10 {
		return "Too long", false
	}
	return "Length OK", true
}

func main() {
	// Capture both return values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result) // Output: 10 / 2 = 5
	}

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: cannot divide by zero
	} else {
		fmt.Println("5 / 0 =", result)
	}

    // Ignore one return value using the blank identifier (_)
    statusMsg, _ := checkLength("short")
    fmt.Println("Status:", statusMsg) // Output: Status: Length OK

    _, isValid := checkLength("a very very very long string")
    if !isValid {
        fmt.Println("String is not valid!") // Output: String is not valid!
    }

	// Example from standard library (strconv.Atoi)
	numStr := "123"
	numInt, convErr := strconv.Atoi(numStr)
	if convErr != nil {
		fmt.Println("Conversion failed:", convErr)
	} else {
		fmt.Println("Conversion result:", numInt) // Output: Conversion result: 123
	}

	numStr = "abc"
	numInt, convErr = strconv.Atoi(numStr)
	if convErr != nil {
		fmt.Println("Conversion failed:", convErr) // Output: Conversion failed: strconv.Atoi: parsing "abc": invalid syntax
	} else {
		fmt.Println("Conversion result:", numInt)
	}
}
```

### Named Return Values

You can name the return values in the function signature. These variables are declared and initialized to their zero values at the start of the function. A `return` statement without arguments (`naked return`) will return the current values of these named variables.

While possible, using `naked return` is generally discouraged in longer functions as it can reduce readability. It's often better to return values explicitly.

```go
package main

import "fmt"

// Using named return values (result int, success bool)
func subtract(a, b int) (result int, success bool) {
	// 'result' and 'success' are already declared (int=0, bool=false)
	if a < b {
		success = false
		result = 0 // Explicit is better, though not required with naked return
		return result, success // Explicit return is clearer
	}

	result = a - b
	success = true
	// return // Naked return: returns the current values of 'result' and 'success'
    return result, success // More recommended
}

func main() {
	res, ok := subtract(10, 3)
	fmt.Printf("Result: %d, Success: %t\n", res, ok) // Output: Result: 7, Success: true

	res, ok = subtract(5, 8)
	fmt.Printf("Result: %d, Success: %t\n", res, ok) // Output: Result: 0, Success: false
}
```

### Variadic Functions

Functions that can accept a *variable* number of arguments for their final parameter. The variadic parameter is denoted by `...` before the type. Inside the function, this parameter is treated as a slice of that type.

```go
package main

import "fmt"

// The sumNumbers function accepts any number of integers (zero or more)
func sumNumbers(label string, numbers ...int) int {
	fmt.Printf("Received for '%s': %v (type: %T)\n", label, numbers, numbers)
	total := 0
	// 'numbers' is an []int slice
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	total1 := sumNumbers("Set 1") // Zero arguments for numbers
	fmt.Println("Total 1:", total1) // Output: Total 1: 0

	total2 := sumNumbers("Set 2", 1, 2, 3)
	fmt.Println("Total 2:", total2) // Output: Total 2: 6

	total3 := sumNumbers("Set 3", 10, 20, 30, 40, 50)
	fmt.Println("Total 3:", total3) // Output: Total 3: 150

	// Passing a slice as a variadic argument
	nums := []int{5, 10, 15}
	total4 := sumNumbers("Set 4", nums...) // Use ... when calling with a slice
	fmt.Println("Total 4:", total4) // Output: Total 4: 30
}
/* Detailed Output:
Received for 'Set 1': [] (type: []int)
Total 1: 0
Received for 'Set 2': [1 2 3] (type: []int)
Total 2: 6
Received for 'Set 3': [10 20 30 40 50] (type: []int)
Total 3: 150
Received for 'Set 4': [5 10 15] (type: []int)
Total 4: 30
*/
```

### Functions as First-Class Values

In Go, functions are *first-class citizens*. This means functions can be:

*   Assigned to variables.
*   Passed as arguments to other functions.
*   Returned as values from other functions.

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// Function that accepts another function as a parameter
// operation is a parameter of type 'func(int, int) int'
func calculate(x, y int, operation func(int, int) int) int {
	fmt.Printf("Running operation on %d and %d\n", x, y)
	return operation(x, y)
}

// Function that returns another function
func createMultiplier(factor int) func(int) int {
	// An anonymous function is returned
	return func(n int) int {
		return n * factor
	}
}

func main() {
	// Assign a function to a variable
	var op func(int, int) int
	op = add
	result := op(10, 5)
	fmt.Println("Result op (add):", result) // Output: Result op (add): 15

	op = subtract
	result = op(10, 5)
	fmt.Println("Result op (subtract):", result) // Output: Result op (subtract): 5

	// Pass a function as an argument
	sumResult := calculate(20, 7, add)
	fmt.Println("Result calculate (add):", sumResult) // Output: Running operation on 20 and 7 \n Result calculate (add): 27

	diffResult := calculate(20, 7, subtract)
	fmt.Println("Result calculate (subtract):", diffResult) // Output: Running operation on 20 and 7 \n Result calculate (subtract): 13

    // Pass an anonymous function directly
    multResult := calculate(5, 6, func(a, b int) int { return a * b })
    fmt.Println("Result calculate (anonymous multiply):", multResult) // Output: Running operation on 5 and 6 \n Result calculate (anonymous multiply): 30

	// Get a function from another function
	double := createMultiplier(2) // double is now func(int) int
	triple := createMultiplier(3) // triple is now func(int) int

	fmt.Println("Double 5:", double(5)) // Output: Double 5: 10
	fmt.Println("Triple 5:", triple(5)) // Output: Triple 5: 15
}
```

### Anonymous Functions and Closures

*   **Anonymous Function:** A function without a name. Useful for creating small inline functions or for closures.
*   **Closure:** A function that "remembers" the lexical environment in which it was defined, even if executed outside that scope. This means an anonymous function can access and modify variables from its parent function.

```go
package main

import "fmt"

// This function returns a closure
func counter() func() int {
	count := 0 // This variable is "captured" by the closure
	// The anonymous function is returned
	return func() int {
		count++ // Accesses and modifies 'count' from the outer scope
		return count
	}
}

func main() {
	// Example of an anonymous function called immediately
	func(message string) {
		fmt.Println("Anonymous message:", message)
	}("Hello directly!") // Output: Anonymous message: Hello directly!

	// Create counter closures
	c1 := counter()
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 1
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 2
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 3

	// Each call to counter() creates a new closure instance with its own separate 'count' state
	c2 := counter()
	fmt.Println("Counter 2:", c2()) // Output: Counter 2: 1
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 4 (c1 still has its own state)
}
```

---

## 5. Pointers

Pointers are often considered complex in languages like C/C++, but in Go, their use is more controlled and safer. A pointer stores the *memory address* of a value, not the value itself.

### What are Pointers?

Imagine computer memory as a series of numbered boxes (addresses). When you declare a regular variable, like `x := 10`, Go allocates a memory box, gives it an address, and stores the value `10` inside it.

A *pointer* is a variable whose *content* is the *memory address* of another variable. So, instead of storing `10`, a pointer to `x` would store the address of the box where `10` is kept.

### Operators `&` (Address Of) and `*` (Dereference)

Go uses two main operators for working with pointers:

1.  **`&` (Address Of Operator):** Placed before a variable, the `&` operator returns the *memory address* of that variable. The result is a pointer. If `v` is a variable, then `&v` is a pointer to `v`. The type of `&v` is `*T`, where `T` is the type of `v`.
2.  **`*` (Dereference Operator):** Placed before a *pointer* variable, the `*` operator accesses the *value* stored at the memory address pointed to by the pointer. If `p` is a pointer, then `*p` is the value that `p` points to.

```go
package main

import "fmt"

func main() {
    // Regular variable
    x := 100
    fmt.Printf("Value of x: %d, Address of x: %p\n", x, &x) // %p is the format verb for addresses

    // Declare a pointer to an int
    var p *int // p is a pointer to int. Its zero value is nil.
    fmt.Printf("Value of p (before assignment): %v\n", p) // Output: <nil>

    // Assign the address of x to pointer p
    p = &x
    fmt.Printf("Value of p (address of x): %p\n", p)

    // Access the value pointed to by p (dereferencing)
    fmt.Printf("Value pointed to by p (*p): %d\n", *p) // Output: 100

    // Change the value of x through pointer p
    *p = 200 // Change the value at the address p points to
    fmt.Printf("Value of x after change via p: %d\n", x) // Output: 200

    // Pointers also have their own address
    fmt.Printf("Address of pointer p: %p\n", &p)

    // Pointer to a pointer
    var pp **int // pp is a pointer to a pointer to an int
    pp = &p
    fmt.Printf("Value of pp (address of p): %p\n", pp)
    fmt.Printf("Value pointed to by p (*p) via pp (**pp): %d\n", **pp) // Dereference twice

    // `new()` function: another way to create pointers
    // new(T) allocates memory for type T, initializes it with T's zero value,
    // and returns a pointer to that memory (*T).
    ptrStr := new(string) // ptrStr is *string, pointing to "" (zero value of string)
    fmt.Printf("Value of ptrStr: %p, Value *ptrStr: '%s'\n", ptrStr, *ptrStr)
    *ptrStr = "Hello from new()"
    fmt.Printf("Value *ptrStr after change: '%s'\n", *ptrStr)
}

/* Example Output (Addresses will vary):
Value of x: 100, Address of x: 0xc000018030
Value of p (before assignment): <nil>
Value of p (address of x): 0xc000018030
Value pointed to by p (*p): 100
Value of x after change via p: 200
Address of pointer p: 0xc000006028
Value of pp (address of p): 0xc000006028
Value pointed to by p (*p) via pp (**pp): 200
Value of ptrStr: 0xc000044210, Value *ptrStr: ''
Value *ptrStr after change: 'Hello from new()'
*/
```

**`nil` Pointer:** A pointer that does not point to any memory address has the value `nil`. Attempting to dereference a `nil` pointer will cause a *runtime panic*.

```go
var pNil *int
// fmt.Println(*pNil) // This will PANIC: runtime error: invalid memory address or nil pointer dereference
if pNil != nil {
    fmt.Println(*pNil) // Safe because of the check
} else {
    fmt.Println("pNil is nil")
}
```

### When to Use Pointers?

1.  **Allowing Functions to Modify Arguments:** In Go, function arguments are passed *by value* (a copy). If you want a function to modify the original variable passed as an argument, you must pass a *pointer* to that variable.

    ```go
    package main

    import "fmt"

    // Accepts int (by value)
    func incrementValue(val int) {
        val++ // Only modifies the local copy
        fmt.Printf("  Value inside incrementValue: %d\n", val)
    }

    // Accepts pointer to int (*int)
    func incrementPointer(ptr *int) {
        *ptr++ // Modifies the value at the address the pointer points to
        fmt.Printf("  Value inside incrementPointer (*ptr): %d\n", *ptr)
    }

    func main() {
        num := 10
        fmt.Printf("Value of num before incrementValue: %d\n", num)
        incrementValue(num)
        fmt.Printf("Value of num after incrementValue: %d\n", num) // Still 10

        fmt.Printf("\nValue of num before incrementPointer: %d\n", num)
        incrementPointer(&num) // Pass the address of num
        fmt.Printf("Value of num after incrementPointer: %d\n", num) // Becomes 11
    }
    /* Output:
    Value of num before incrementValue: 10
      Value inside incrementValue: 11
    Value of num after incrementValue: 10

    Value of num before incrementPointer: 10
      Value inside incrementPointer (*ptr): 11
    Value of num after incrementPointer: 11
    */
    ```

2.  **Efficiency When Passing Large Data (Structs):** Passing a large struct as a function argument involves copying the entire struct. Passing a pointer to the struct only copies the memory address (typically 8 bytes on a 64-bit system), which can be much more efficient.

3.  **Signaling "Absence" or Optional Values:** A pointer can be `nil`, which can be used to indicate that a value is absent or not set, especially for data types that don't have a natural `nil` representation (like `int`, whose zero value `0` might be a valid value).

    ```go
    package main
    import "fmt"

    type Config struct {
    	Timeout *int // Pointer to int, can be nil if not set
    	Retries int    // Regular int, 0 is its zero value
    }

    func main() {
    	defaultTimeout := 30
    	cfg1 := Config{ Timeout: &defaultTimeout, Retries: 3 }
    	cfg2 := Config{ Retries: 5 } // Timeout not set (will be nil)

    	if cfg1.Timeout != nil { fmt.Printf("Cfg1 Timeout: %d\n", *cfg1.Timeout) } // Cfg1 Timeout: 30
    	if cfg2.Timeout != nil { fmt.Printf("Cfg2 Timeout: %d\n", *cfg2.Timeout) } else { fmt.Println("Cfg2 Timeout: default") } // Cfg2 Timeout: default
    }
    ```

4.  **Method Receivers (Pointer Receiver):** Pointers are frequently used as receivers for methods to allow the method to modify the struct's state (discussed in the Methods section).

### Pointers to Structs

Working with pointers to structs is very common in Go. Go provides *syntactic sugar* so you don't need to explicitly dereference to access struct fields via a pointer.

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // p is a pointer to v (*Vertex)

	// Accessing fields via pointer
	// Standard way (explicit dereference):
	fmt.Println((*p).X) // Output: 1

	// Go way (implicit dereference - more common and readable):
	fmt.Println(p.X) // Output: 1 (Go automatically dereferences p)

	// Modifying fields via pointer
	p.X = 100
	fmt.Println(v.X) // Output: 100 (original v's value changed)
	fmt.Println(v)   // Output: {100 2}
}
```

**Difference between `new(T)` and `&T{}`:**

*   `p := new(T)`: Allocates memory, initializes with `T`'s zero value, returns a pointer `*T`.
*   `p := &T{...}`: Allocates memory, initializes with the provided literal values, returns a pointer `*T`. This is more common as you usually want to initialize fields when creating a struct.

```go
p1 := new(Vertex)     // p1 is *Vertex pointing to {0 0}
p2 := &Vertex{}       // p2 is *Vertex pointing to {0 0} (same as new)
p3 := &Vertex{X: 1}   // p3 is *Vertex pointing to {1 0}
p4 := &Vertex{1, 2}   // p4 is *Vertex pointing to {1 2}
fmt.Println(p1, p2, p3, p4)
```

---

## 6. Structs and Methods

We've covered struct basics. Now let's dive into *embedding* (composition) and *methods*.

### Defining Structs (Recap)

Structs are collections of named fields.

```go
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
}
```

### Embedding (Composition)

Go does not support traditional inheritance like in other OOP languages. Instead, Go favors **composition through embedding**. You can embed one struct within another to "borrow" its fields and methods.

To embed, declare a field within the struct without giving it a name, only its type.

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method for Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

type Contact struct {
	Email string
	Phone string
}

// Manager struct embeds Person and Contact
type Manager struct {
	Person        // Embed Person (no field name, just the type)
	Contact       // Embed Contact
	Department string // Field specific to Manager
	Level      int
}

// Manager can also have its own methods
func (m Manager) DelegateTask() {
	fmt.Printf("%s (%s) is delegating tasks.\n", m.Name, m.Department)
}

func main() {
	// Create a Manager instance
	m := Manager{
		Person: Person{ // Initialize the embedded struct
			Name: "Budi Gunawan",
			Age:  45,
		},
		Contact: Contact{
			Email: "budi.g@company.com",
			Phone: "0811-111-222",
		},
		Department: "Technology",
		Level:      5,
	}

	// Access fields from the embedded struct (Promoted Fields)
	// Fields from Person and Contact can be accessed directly as if they belong to Manager
	fmt.Println("Name:", m.Name)       // Output: Name: Budi Gunawan (from Person)
	fmt.Println("Email:", m.Email)     // Output: Email: budi.g@company.com (from Contact)
	fmt.Println("Department:", m.Department) // Output: Department: Technology (from Manager)

	// Access the embedded struct explicitly (if needed, e.g., name conflict)
	fmt.Println("Age (explicit):", m.Person.Age) // Output: Age (explicit): 45

	// Call methods from the embedded struct (Promoted Methods)
	m.Greet() // Output: Hello, my name is Budi Gunawan and I am 45 years old. (Greet method from Person)

	// Call Manager's own method
	m.DelegateTask() // Output: Budi Gunawan (Technology) is delegating tasks.
}
```

**Field/Method Promotion:** When a struct is embedded, the fields and methods of the embedded struct are "promoted" to the containing struct. This means you can access them directly as if they were defined on the outer struct. If there's a name conflict (the outer struct has a field/method with the same name), you must access the embedded version explicitly (e.g., `m.Person.Name`).

Composition is more flexible than classical inheritance as it allows you to build functionality by combining independent parts.

### Methods

A method is a **function associated with a particular type**. This type is called the *receiver*. Methods allow you to define behavior for your data types (especially structs).

Method declaration is similar to function declaration, but includes an extra *receiver* parameter before the method name.

```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Method with a Point receiver (value receiver)
// `p` is a copy of the Point instance calling this method
func (p Point) DistanceFromOrigin() float64 {
	// Changes to p here won't affect the original Point
	// p.X = 100
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Method with a *Point receiver (pointer receiver)
// `p` is a pointer to the Point instance calling this method
// This method can modify the original Point instance.
func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
	fmt.Printf("  Inside Scale: Point becomes %v\n", *p)
}

func main() {
	pt1 := Point{3, 4}
	fmt.Printf("Point pt1: %v\n", pt1)

	// Call method with value receiver
	dist := pt1.DistanceFromOrigin()
	fmt.Printf("Distance of pt1 from origin: %.2f\n", dist) // Output: 5.00
	fmt.Printf("Point pt1 after DistanceFromOrigin: %v\n", pt1) // Remains {3 4}

	fmt.Println("---")

	// Call method with pointer receiver
	// Go automatically converts pt1 to &pt1 when calling the Scale method
	pt1.Scale(2)
	// or explicitly: (&pt1).Scale(2)
	fmt.Printf("Point pt1 after Scale(2): %v\n", pt1) // Becomes {6 8}

	fmt.Println("---")

	// Call method with pointer receiver on a pointer
	pt2 := &Point{1, 1}
	fmt.Printf("Point pt2 (pointer): %v\n", pt2)
	pt2.Scale(5) // Can be called directly on the pointer
	fmt.Printf("Point pt2 after Scale(5): %v\n", pt2) // Becomes &{5 5}

	// Call method with value receiver on a pointer
	dist2 := pt2.DistanceFromOrigin() // Go automatically dereferences pt2 to *pt2
	// or explicitly: (*pt2).DistanceFromOrigin()
	fmt.Printf("Distance of pt2 from origin: %.2f\n", dist2) // Output: 7.07 (sqrt(5^2 + 5^2))
}
```

### Pointer Receiver vs. Value Receiver

The choice between a pointer receiver (`*T`) and a value receiver (`T`) is important:

1.  **Value Receiver (`func (r T) MethodName()`)**
    *   The method operates on a **copy** of the receiver value.
    *   Cannot modify the original receiver value.
    *   Safe to use concurrently without locking (as it works on a copy).
    *   Suitable if the method doesn't need to change the receiver's state or if the receiver is a small/immutable data type (like basic types or small structs).

2.  **Pointer Receiver (`func (r *T) MethodName()`)**
    *   The method operates on the **original receiver value** via a pointer.
    *   **Can modify** the original receiver value.
    *   More efficient for large structs as it avoids copying the entire struct.
    *   Allows the method to handle `nil` receivers (requires an `if r == nil` check).
    *   **Required** if the method needs to change the receiver's state.
    *   If you modify state, you might need to consider concurrency (e.g., using mutexes) if multiple goroutines could call methods on the same instance.

**General Convention:**

*   If *any* method for type `T` needs a pointer receiver (for modification or efficiency), then **all** methods for type `T` should probably use pointer receivers for consistency.
*   When in doubt, use a pointer receiver.

Go automatically handles the conversion between values and pointers when calling methods. If `v` is of type `T`, `v.PointerMethod()` will be interpreted as `(&v).PointerMethod()`. If `p` is of type `*T`, `p.ValueMethod()` will be interpreted as `(*p).ValueMethod()`.

---

## 7. Interfaces

Interfaces in Go provide a way to specify *behavior* rather than a specific data structure. An interface defines a set of *method signatures*. Any data type that implements *all* the methods in that signature set is automatically considered to satisfy the interface.

### Interface Concepts

Think of an interface as a *contract*. The contract says, "Anyone wanting to be considered type X must be able to do A, B, and C." The interface doesn't care *how* A, B, and C are done, only *that* they can be done (i.e., they have methods with matching signatures).

### Defining Interfaces

Interfaces are defined using the `type` and `interface` keywords, followed by a list of method signatures.

```go
package main

import (
	"fmt"
	"math"
)

// The Shape interface defines behavior for geometric shapes
// Any type wanting to be a Shape MUST have Area() float64
// and Perimeter() float64 methods.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Concrete type: Rectangle
type Rectangle struct {
	Width, Height float64
}

// Rectangle implements Shape (it has Area and Perimeter methods)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Concrete type: Circle
type Circle struct {
	Radius float64
}

// Circle implements Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that works with the Shape interface
// This function doesn't care if s is a Rectangle or Circle,
// as long as s fulfills the Shape contract (has Area and Perimeter).
func PrintShapeInfo(s Shape) {
	fmt.Printf("Type: %T\n", s) // %T shows the concrete type
	fmt.Printf("  Area: %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// rect and circ are concrete types, but both satisfy the Shape interface
	fmt.Println("Rectangle Info:")
	PrintShapeInfo(rect)

	fmt.Println("\nCircle Info:")
	PrintShapeInfo(circ)

	// We can create a slice of Shapes
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
		rect, // Can include existing variables
		circ,
	}

	fmt.Println("\nInfo from Shapes Slice:")
	totalArea := 0.0
	for _, s := range shapes {
		PrintShapeInfo(s)
		totalArea += s.Area()
	}
	fmt.Printf("\nTotal Area of all shapes: %.2f\n", totalArea)
}
```

### Implicit Implementation

This is a key feature of Go: **interface implementation is implicit**. You don't need to explicitly declare `type Rectangle implements Shape`. As long as `Rectangle` has all the methods defined in `Shape` (with the same signatures), Go automatically considers `Rectangle` to implement `Shape`.

This promotes *decoupling*. The package defining the interface doesn't need to know about the packages implementing it, and vice versa.

### The Empty Interface (`interface{}`)

The empty interface is an interface that has no methods at all: `interface{}`. Since there are no methods to implement, **all data types** in Go automatically satisfy the empty interface.

This makes `interface{}` useful for handling values of unknown or varying types. A variable of type `interface{}` can hold a value of any type. It's similar to `Object` in Java or `any` in TypeScript, but should be used carefully as it bypasses static type safety.

```go
package main

import "fmt"

// Function that accepts any type
func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	var data interface{} // Variable of empty interface type

	data = 10 // Store an int
	describe(data) // Output: Value: 10, Type: int

	data = "Hello Go" // Store a string
	describe(data) // Output: Value: Hello Go, Type: string

	data = true // Store a bool
	describe(data) // Output: Value: true, Type: bool

	data = struct{ Name string }{"Sample"} // Store an anonymous struct
	describe(data) // Output: Value: {Sample}, Type: struct { Name string }

	// Slice of empty interfaces
	mixed := []interface{}{1, "two", false, 3.14, nil}
	for _, v := range mixed {
		describe(v)
	}
}
```

### Type Assertion and Type Switch

When you have a value in an interface variable (especially `interface{}`), you often need to find out the underlying *concrete* type or convert it back to the concrete type to use its specific fields or methods. This is done with *type assertion* or a *type switch*.

**Type Assertion:** Takes an interface value and extracts a value of a specific concrete type. Syntax: `value, ok := i.(T)`.

*   `i`: The interface variable.
*   `T`: The concrete type you *expect*.
*   `value`: Will contain the value of `i` converted to type `T` if successful.
*   `ok`: A boolean, `true` if `i` indeed holds type `T`, `false` otherwise.

If you perform the assertion `value := i.(T)` without the `ok`, and `i` does not hold type `T`, the program will *panic*. Using the `value, ok` form is safer.

```go
package main

import "fmt"

func process(i interface{}) {
	fmt.Printf("Processing: %v (%T)\n", i, i)

	// Type Assertion to string (safe form)
	strVal, ok := i.(string)
	if ok {
		fmt.Printf("  It's a string! Length: %d\n", len(strVal))
		return // Done processing
	}

	// Type Assertion to int (safe form)
	intVal, ok := i.(int)
	if ok {
		fmt.Printf("  It's an integer! Squared value: %d\n", intVal*intVal)
		return
	}

	// If you are sure of the type (be careful, can panic)
	// floatVal := i.(float64) // Will panic if i is not float64
	// fmt.Printf("  It's a float: %f\n", floatVal)

	fmt.Println("  Type not recognized or handled.")
}

func main() {
	process("Hello World")
	process(42)
	process(true) // This type is not handled by the ifs above
}
/* Output:
Processing: Hello World (string)
  It's a string! Length: 11
Processing: 42 (int)
  It's an integer! Squared value: 1764
Processing: true (bool)
  Type not recognized or handled.
*/
```

**Type Switch:** A cleaner and more idiomatic way to perform several type assertions in sequence, similar to a regular `switch` but on types.

```go
package main

import "fmt"

func describeWithTypeSwitch(i interface{}) {
	fmt.Printf("Describing: %v (%T) -> ", i, i)
	switch v := i.(type) { // 'v' will have the concrete type inside each case
	case string:
		fmt.Printf("String with length %d\n", len(v))
	case int:
		fmt.Printf("Integer, value %d\n", v)
	case bool:
		fmt.Printf("Boolean, value %t\n", v)
	case float64:
		fmt.Printf("Float64, value %f\n", v)
	case nil:
		fmt.Println("A nil value")
	default:
		// v here still has the type interface{} (or the original type of i)
		fmt.Printf("Another type: %T\n", v)
	}
}

func main() {
	describeWithTypeSwitch("Go")
	describeWithTypeSwitch(123)
	describeWithTypeSwitch(false)
	describeWithTypeSwitch(3.14)
	describeWithTypeSwitch(nil)
	describeWithTypeSwitch([]int{1, 2})
}
/* Output:
Describing: Go (string) -> String with length 2
Describing: 123 (int) -> Integer, value 123
Describing: false (bool) -> Boolean, value false
Describing: 3.14 (float64) -> Float64, value 3.140000
Describing: <nil> (<nil>) -> A nil value
Describing: [1 2] ([]int) -> Another type: []int
*/
```

### Common Example: The `error` Interface

One of the most important and common interfaces in Go is the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

Any type with an `Error() string` method is considered to satisfy the `error` interface. This is Go's standard way of representing and handling error conditions. We'll discuss this more in the Error Handling section.

**Interface Principles in Go:**

*   **Accept interfaces, return structs:** Functions should accept interface types as parameters if they only need specific *behavior*, not a concrete type. This increases flexibility and testability. Functions should generally return concrete types (structs) so the caller has full information about what was returned.
*   **Smaller interfaces are better:** Define interfaces with only the methods that are strictly necessary. The `io.Reader` and `io.Writer` interfaces are great examples:
    ```go
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    type Writer interface {
        Write(p []byte) (n int, err error)
    }
    ```
    Many types (Files, network connections, buffers) can implement these simple interfaces, allowing code that works with `Reader` or `Writer` to be highly reusable.

---

## 8. Error Handling

Go takes a unique approach to error handling that emphasizes *explicit handling* rather than exception mechanisms (like try-catch) found in many other languages.

### Error Conventions in Go

1.  **Errors are Ordinary Values:** Errors in Go are represented by values that satisfy the built-in `error` interface. This means errors can be passed around, checked, logged, etc., just like any other value.
2.  **Functions Return Errors:** Functions that can fail typically return an `error` as their *last* return value. If the operation succeeds, the returned `error` value is `nil`. If it fails, the `error` value contains information about the failure.
3.  **Explicit Error Checking:** The caller of a function is responsible for *checking* the returned `error` value, usually immediately after the function call, using `if err != nil`.

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Example 1: Opening a file
	filePath := "non_existent_file.txt"
	file, err := os.Open(filePath) // os.Open returns (*File, error)
	if err != nil {
		// Handle the error (e.g., log and exit, or try something else)
		fmt.Printf("Failed to open file '%s': %v\n", filePath, err)
		// os.Exit(1) // Might want to exit if the error is fatal
	} else {
		// If no error (err == nil), proceed using 'file'
		fmt.Printf("Successfully opened file: %s\n", file.Name())
		// Don't forget to close the file when done
		defer file.Close()
		// ... do something with the file ...
	}

	fmt.Println("---")

	// Example 2: Converting string to integer
	numStr := "123a"
	num, err := strconv.Atoi(numStr) // strconv.Atoi returns (int, error)
	if err != nil {
		fmt.Printf("Failed to convert '%s' to int: %v\n", numStr, err)
	} else {
		fmt.Printf("Conversion result: %d\n", num)
	}

	numStr = "456"
	num, err = strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Failed to convert '%s' to int: %v\n", numStr, err)
	} else {
		fmt.Printf("Conversion result: %d\n", num)
	}
}

/* Output:
Failed to open file 'non_existent_file.txt': open non_existent_file.txt: no such file or directory
---
Failed to convert '123a' to int: strconv.Atoi: parsing "123a": invalid syntax
Conversion result: 456
*/
```

The `if err != nil { return ..., err }` pattern is very common in Go for propagating errors up to higher-level callers.

### Creating Errors (`errors.New`, `fmt.Errorf`)

You can create your own error values using:

1.  **`errors.New(message string)`:** Creates a simple error with a static text message.

    ```go
    import "errors"

    func validateInput(input string) error {
        if input == "" {
            return errors.New("input cannot be empty")
        }
        return nil
    }
    ```

2.  **`fmt.Errorf(format string, a ...interface{})`:** Creates an error with a formatted message, similar to `fmt.Printf`. Useful for including dynamic details in the error message.

    ```go
    package main
    import "fmt"

    func connectToDB(host string, port int) error {
        if port <= 0 || port > 65535 {
            // Include the port value in the error message
            return fmt.Errorf("invalid database port: %d", port)
        }
        // ... connection logic ...
        fmt.Printf("Attempting connection to %s:%d...\n", host, port)
        // Suppose connection fails
        return fmt.Errorf("failed to connect to %s:%d (host not found)", host, port)
    }

    func main() {
        err := connectToDB("localhost", 80000)
        if err != nil { fmt.Println("Error:", err) } // Error: invalid database port: 80000

        err = connectToDB("db.example.com", 5432)
        if err != nil { fmt.Println("Error:", err) }
        /* Output (if connection fails):
        Attempting connection to db.example.com:5432...
        Error: failed to connect to db.example.com:5432 (host not found)
        */
    }
    ```

### Wrapping Errors (Error Wrapping - Go 1.13+)

Sometimes, you want to add context to an error received from another function without losing the original error information. Go 1.13 introduced *error wrapping*.

*   Use the `%w` format verb in `fmt.Errorf` to wrap an error.
*   Use `errors.Is(err, target error)` to check if an error in the *wrapping chain* matches a specific `target` error. Useful for checking specific error kinds (e.g., `os.ErrNotExist`).
*   Use `errors.As(err, target interface{})` to check if an error in the wrapping chain matches a specific *type*, and if so, assign that error to `target`. Useful if you have custom error types and want to access their specific fields.

```go
package main

import (
	"errors"
	"fmt"
	"os" // For os.ErrNotExist and custom error example
)

// Custom error type
type ConfigError struct {
	FileName string
	Field    string
	Err      error // Wraps the original error (if any)
}

// Implement the error interface
func (e *ConfigError) Error() string {
	msg := fmt.Sprintf("configuration error in file '%s'", e.FileName)
	if e.Field != "" {
		msg += fmt.Sprintf(", field '%s'", e.Field)
	}
	// Note: Don't include e.Err in the main message directly
	// Use Unwrap() for that purpose
	return msg
}

// Implement Unwrap() to support errors.Is and errors.As
func (e *ConfigError) Unwrap() error {
	return e.Err // Return the wrapped error
}


func loadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// Wrap the error from os.Open with additional context
		// Use %w for wrapping
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	defer file.Close()

	// ... parsing logic ...
	// Suppose a field is missing
	missingField := "database_url"
	if missingField != "" { // Simulate parsing error
		// Parsing error, doesn't wrap the OS error directly
		parseErr := &ConfigError{FileName: filename, Field: missingField}
		// Wrap the custom parseErr
		return fmt.Errorf("failed to parse configuration: %w", parseErr)
	}

	return nil
}

func main() {
    fmt.Println("--- Testing file not found ---")
	err := loadConfig("non_existent_config.yaml")
	if err != nil {
		fmt.Println("Main error:", err) // Main error: failed to load configuration: open non_existent_config.yaml: no such file or directory

		// Check if the error was caused by file not found
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  Detail: Configuration file not found.")
		}

		// Check if the error is of type ConfigError
		var configErr *ConfigError
		if errors.As(err, &configErr) {
			fmt.Printf("  Detail: Error on field '%s' in file '%s'\n", configErr.Field, configErr.FileName)
            // Check if the ConfigError wraps another error
            wrapped := configErr.Unwrap()
            if wrapped != nil {
                fmt.Println("    Wrapped error:", wrapped)
            }
		} else {
            fmt.Println("  Error is not a ConfigError.")
            // We can still get the wrapped error directly
            wrapped := errors.Unwrap(err)
            if wrapped != nil {
                 fmt.Println("  Wrapped OS error:", wrapped) // Shows the os.ErrNotExist details
            }
        }
	}

    fmt.Println("\n--- Testing parsing error ---")
	// Simulate parsing error
    // Create dummy file first
    dummyConf := "config.yaml"
    f, _ := os.Create(dummyConf)
    f.Close()

	err = loadConfig(dummyConf)
    if err != nil {
        fmt.Println("Main error (parsing):", err) // Main error (parsing): failed to parse configuration: configuration error in file 'config.yaml', field 'database_url'

		if errors.Is(err, os.ErrNotExist) { // Won't match
			fmt.Println("  Detail: Configuration file not found.")
		}

		var configErr *ConfigError
		if errors.As(err, &configErr) { // Will match
			fmt.Printf("  Detail: Error on field '%s' in file '%s'\n", configErr.Field, configErr.FileName)
            wrapped := errors.Unwrap(err) // Get the wrapped error (which is configErr itself in this case)
            fmt.Println("    Wrapped error (via errors.Unwrap):", wrapped) // Output: configuration error...
            wrapped = configErr.Unwrap() // Check if ConfigError wraps *another* error
            if wrapped != nil {
                 fmt.Println("    Inner wrapped error (via method Unwrap):", wrapped)
            } else {
                 fmt.Println("    Inner wrapped error (via method Unwrap): <nil>") // Output: <nil> because this ConfigError didn't wrap another error
            }
		}
    }
    os.Remove(dummyConf) // Clean up dummy file
}
```

### `panic` and `recover`

Go has `panic` and `recover` mechanisms similar to exceptions in other languages, but their use is **strongly discouraged** for normal error handling.

*   **`panic(v interface{})`**: Stops the normal execution flow of the current function. Execution of `defer`red functions still runs. If the `panic` is not handled by a `recover` in the same goroutine, the program crashes and prints a stack trace.
*   **`recover()`**: Only useful inside a `defer`red function. Stops the `panic` process and returns the value passed to `panic`. If the goroutine is not panicking, `recover` returns `nil`.

**When might `panic` be used?**

1.  **Unrecoverable Errors:** Situations where the program absolutely cannot continue (e.g., internal inconsistencies, critical memory allocation failure). Example: Indexing a slice out of bounds (`slice[10]` when slice length is 5) causes a built-in panic.
2.  **Initialization Failures:** If the program cannot start because initial setup failed (e.g., cannot read a critical configuration file at startup).
3.  **(Rarely) Within Libraries:** Some libraries might use panic/recover internally to simplify error handling across many layers, but they should recover the panic at the public API boundary and return it as a regular `error` value.

**Example Usage of `recover` (use with caution):**

```go
package main

import "fmt"

func mightPanic(shouldPanic bool) {
	// Defer function to potentially recover from a panic
	defer func() {
		// recover() only works inside defer
		if r := recover(); r != nil {
			fmt.Println("PANIC DETECTED (in recover):", r)
			// Could log the error here, or perform additional cleanup
		}
	}() // Call the anonymous deferred function

	fmt.Println("Before potential panic...")
	if shouldPanic {
		panic("Something really bad happened!")
	}
	// This code won't execute if panic occurs
	fmt.Println("After potential panic (won't be reached if panic)")
}

func main() {
	fmt.Println("--- Calling mightPanic(false) ---")
	mightPanic(false)
	fmt.Println("mightPanic(false) finished.")

	fmt.Println("\n--- Calling mightPanic(true) ---")
	mightPanic(true)
	// Execution RETURNS here after recover in mightPanic
	fmt.Println("mightPanic(true) finished (after recover).")
}
/* Output:
--- Calling mightPanic(false) ---
Before potential panic...
After potential panic (won't be reached if panic)
mightPanic(false) finished.

--- Calling mightPanic(true) ---
Before potential panic...
PANIC DETECTED (in recover): Something really bad happened!
mightPanic(true) finished (after recover).
*/
```

**Error Handling Conclusion:** Prioritize using `error` values and `if err != nil` checks for all anticipated error conditions. Use error wrapping (`%w`, `errors.Is`, `errors.As`) to add context and check for specific errors. Use `panic` only for truly exceptional, unrecoverable situations.

---

## 9. Concurrency

Concurrency is one of Go's primary strengths. Go provides simple yet powerful built-in features for writing programs that can do many things simultaneously: **Goroutines** and **Channels**.

**Concurrency vs. Parallelism:**

*   **Concurrency:** Dealing with many tasks *as if* they run at the same time. It's about *structuring* a program to manage multiple things at once (e.g., handling many web requests). Doesn't necessarily run on multiple CPU cores.
*   **Parallelism:** Executing many tasks *physically* at the same time, usually leveraging multiple CPU cores.

Go makes it easy to write concurrent code, and the Go runtime will automatically try to run goroutines in parallel on available CPU cores.

### Goroutines

A goroutine is a **function or method that runs concurrently** with other functions/methods. Think of it like a very lightweight *thread*, managed by the Go runtime, not directly by the operating system. Creating a goroutine is much cheaper than creating an OS thread. Thousands or even millions of goroutines can run within a single process.

#### Starting Goroutines (`go` keyword)

To run a function as a goroutine, simply add the `go` keyword before the function call.

```go
package main

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// time.Sleep simulates work or waiting for I/O
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Message from '%s': %s - iteration %d\n", s, s, i)
	}
	fmt.Printf("'%s' finished.\n", s)
}

func main() {
	fmt.Println("Starting main goroutine.")

	// Run say("Hello") as a new goroutine
	go say("Hello", 5)

	// Run say("World") as a new goroutine
	go say("World", 3)

	// The main() function itself runs in a goroutine (the main goroutine).
	// If the main goroutine finishes, the entire program exits,
	// even if other goroutines haven't completed.
	fmt.Println("Main goroutine waiting briefly...")

	// We need a way to wait for other goroutines to finish.
	// A simple (BUT NOT RECOMMENDED for production) way: sleep briefly.
	time.Sleep(1 * time.Second) // Give other goroutines time to run

	fmt.Println("Main goroutine finished.")
	// The "Hello" goroutine might not be finished here if sleep is too short.
}

/* Example Output (Order can vary!):
Starting main goroutine.
Main goroutine waiting briefly...
Message from 'World': World - iteration 0
Message from 'Hello': Hello - iteration 0
Message from 'World': World - iteration 1
Message from 'Hello': Hello - iteration 1
Message from 'World': World - iteration 2
'World' finished.
Message from 'Hello': Hello - iteration 2
Message from 'Hello': Hello - iteration 3
Message from 'Hello': Hello - iteration 4
'Hello' finished.
Main goroutine finished.
*/
```

**Problem:** In the example above, we used `time.Sleep` in `main` to wait for other goroutines. This is a bad approach because we don't know exactly how long the other goroutines will take. The correct way is to use synchronization mechanisms like `sync.WaitGroup` or channels.

#### Synchronization with `sync.WaitGroup`

`sync.WaitGroup` is a common way to wait for a collection of goroutines to finish.

1.  Create a `WaitGroup`: `var wg sync.WaitGroup`
2.  Before starting each goroutine you want to wait for, call `wg.Add(1)`.
3.  Inside each goroutine, call `defer wg.Done()` to signal completion (decrements the `WaitGroup` counter).
4.  In `main` (or the waiting goroutine), call `wg.Wait()` to block execution until the `WaitGroup` counter becomes zero.

```go
package main

import (
	"fmt"
	"sync" // Import the sync package
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Ensure Done() is called when the worker finishes, even if it panics
	defer wg.Done()

	fmt.Printf("Worker %d: Starting\n", id)
	// Simulate work
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d: Finished\n", id)
}

func main() {
	// Declare a WaitGroup
	var wg sync.WaitGroup

	numWorkers := 5
	fmt.Printf("Starting %d workers...\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Increment the WaitGroup counter BEFORE starting the goroutine
		wg.Add(1)
		// Run the worker as a goroutine
		// Need to pass 'i' as an argument so each goroutine gets the correct value of i
		// Otherwise, all goroutines would use the last value of 'i' from the loop (closure issue)
		go worker(i, &wg)
	}

	fmt.Println("Main: Waiting for all workers to finish...")
	// Wait until the WaitGroup counter becomes 0
	wg.Wait()

	fmt.Println("Main: All workers have finished.")
}

/* Example Output (Order of worker completion can vary):
Starting 5 workers...
Main: Waiting for all workers to finish...
Worker 1: Starting
Worker 2: Starting
Worker 3: Starting
Worker 4: Starting
Worker 5: Starting
Worker 1: Finished
Worker 2: Finished
Worker 3: Finished
Worker 4: Finished
Worker 5: Finished
Main: All workers have finished.
*/
```

### Channels

Channels are **typed communication pipes** that allow goroutines to send and receive values to/from each other. Channels provide a safe way to share data between goroutines without requiring explicit locking (though mutexes are still needed in other cases).

**Key Principle:** *"Do not communicate by sharing memory; instead, share memory by communicating."*

#### Creating Channels (`make`)

Channels are created using the `make` function.

```go
// Create an unbuffered channel for strings
chStr := make(chan string)

// Create a buffered channel for ints with a capacity of 3
chIntBuffered := make(chan int, 3)
```

#### Sending and Receiving Data

*   **Sending:** Use the `<-` operator to send a value into a channel: `channel <- value`
*   **Receiving:** Use the `<-` operator to receive a value from a channel: `value := <-channel` or `<-channel` (if the value isn't needed).

Send and receive operations are **blocking** by default (for unbuffered channels, and for buffered channels when the buffer is full/empty).

#### Blocking Operations

*   **Sending to an unbuffered channel:** *Blocks* the sending goroutine until another goroutine is *ready to receive* from that channel.
*   **Receiving from an unbuffered channel:** *Blocks* the receiving goroutine until another goroutine *sends* a value to that channel.
*   **Sending to a buffered channel:** Does *not* block if the buffer is *not full*. Blocks if the buffer is *full*, until space becomes available (due to a receiver).
*   **Receiving from a buffered channel:** Does *not* block if the buffer is *not empty*. Blocks if the buffer is *empty*, until a new value is sent.

This blocking is a fundamental synchronization mechanism in Go.

```go
package main

import (
	"fmt"
	"time"
)

// Function that sends a message to a channel
func sendMessage(ch chan string, msg string) {
	fmt.Printf("Sending: '%s'\n", msg)
	time.Sleep(500 * time.Millisecond) // Simulate time
	ch <- msg // Send to channel (will block if unbuffered and no receiver ready)
	fmt.Printf("Sent: '%s'\n", msg)
}

// Function that receives a message from a channel
func receiveMessage(ch chan string) {
	fmt.Println("Waiting for message...")
	receivedMsg := <-ch // Receive from channel (will block until sender sends)
	fmt.Printf("Received: '%s'\n", receivedMsg)
}

func main() {
	// Create an unbuffered channel
	messageChannel := make(chan string)

	// Run sender and receiver as goroutines
	go sendMessage(messageChannel, "Hello Channel!")
	go receiveMessage(messageChannel)

	// Give goroutines time to run
	// (In a real app, use WaitGroup or other sync mechanism)
	time.Sleep(2 * time.Second)
	fmt.Println("Main finished.")
}

/* Example Output (Order might vary slightly, but blocking will occur):
Waiting for message...
Sending: 'Hello Channel!'
(After 500ms, sender sends to channel, waiting receiver can receive)
Sent: 'Hello Channel!'
Received: 'Hello Channel!'
Main finished.
*/
```

#### Buffered vs. Unbuffered Channels

*   **Unbuffered (Capacity 0):** `make(chan T)`. Requires sender and receiver to be ready simultaneously (direct synchronization).
*   **Buffered (Capacity > 0):** `make(chan T, capacity)`. Sender can send values into the buffer without waiting for a receiver, as long as the buffer isn't full. Receiver can receive values from the buffer without waiting for a sender, as long as the buffer isn't empty. Useful for decoupling or smoothing out bursts of data production/consumption.

```go
package main

import "fmt"

func main() {
	// Buffered channel with capacity 2
	bufferedChan := make(chan int, 2)

	fmt.Println("Sending 1 to buffer...")
	bufferedChan <- 1 // Doesn't block (buffer 1/2)
	fmt.Println("Sending 2 to buffer...")
	bufferedChan <- 2 // Doesn't block (buffer 2/2)

	// bufferedChan <- 3 // This would BLOCK because the buffer is full

	fmt.Println("Receiving from buffer...")
	val1 := <-bufferedChan // Doesn't block, receives 1 (buffer 1/2)
	fmt.Printf("Received: %d\n", val1)

	fmt.Println("Receiving from buffer...")
	val2 := <-bufferedChan // Doesn't block, receives 2 (buffer 0/2)
	fmt.Printf("Received: %d\n", val2)

	// val3 := <-bufferedChan // This would BLOCK because the buffer is empty
}
```

#### Closing Channels (`close`)

*   A channel can be closed by the **sender** (and only the sender) to indicate that no more values will be sent. Use `close(channel)`.
*   Attempting to send on a closed channel causes a *panic*.
*   Receiving from a closed channel immediately returns the *zero value* for the channel's type.
*   You can use the two-value receive form `v, ok := <-ch` to check if a channel is closed. `ok` will be `false` if the channel is closed and there are no values left in the buffer.

#### Iterating over Channels with `range`

You can use `for range` to receive values from a channel repeatedly until the channel is *closed*.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Producer: Sending %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Producer: Finished sending, closing channel.")
	close(ch) // Close the channel after all values are sent
}

func consume(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Consumer %d: Starting\n", id)
	// The for range loop will automatically stop when 'ch' is closed
	for value := range ch {
		fmt.Printf("Consumer %d: Received %d\n", id, value)
		time.Sleep(200 * time.Millisecond) // Simulate processing
	}
	fmt.Printf("Consumer %d: Channel closed, finished.\n", id)
}

func main() {
	dataChan := make(chan int, 3) // Buffered channel
	var wg sync.WaitGroup

	// Start the producer
	go produce(dataChan, 5)

	// Start multiple consumers
	numConsumers := 2
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consume(i, dataChan, &wg)
	}

	// Wait for all consumers to finish
	wg.Wait()
	fmt.Println("Main: All consumers finished.")
}
```
### `select` Statement

The `select` statement allows a goroutine to wait on **multiple channel operations** simultaneously. `select` will block until one of its `case`s (a channel operation) is ready to execute (is non-blocking). If multiple `case`s are ready, `select` will choose one pseudo-randomly.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1 sends to ch1 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	// Goroutine 2 sends to ch2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	fmt.Println("Waiting for message from ch1 or ch2...")

	// Loop twice to receive from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
			// Other cases can be added here (e.g., timeout channel)
			// case <-time.After(3 * time.Second):
			//     fmt.Println("Timeout waiting for message!")
			//     return
			// default: // If no case is ready, default will run (non-blocking select)
			//     fmt.Println("No message ready, trying again later...")
			//     time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("Finished receiving two messages.")
}

/* Output:
Waiting for message from ch1 or ch2...
Received: Message from channel 1
Received: Message from channel 2
Finished receiving two messages.
*/
```

**`select` with `default`:** Makes the `select` non-blocking. If no channel is ready, the `default` case will be executed.
**`select` with timeout:** Uses `time.After(duration)` in a `case` to limit the waiting time.

### Common Concurrency Patterns

*   **Worker Pools:** A number of goroutines (workers) process tasks from an input channel and send results to an output channel.
*   **Fan-out, Fan-in:** One goroutine distributes work to many goroutines (fan-out), then another goroutine collects results from those many goroutines (fan-in).
*   **Rate Limiting:** Controls how often an operation can be performed, often using tickers or buffered channels.
*   **Pipeline:** A series of processing stages, where the output of one stage becomes the input for the next, connected by channels.

### `sync` Package (Mutex, RWMutex, etc.)

Although channels are the idiomatic way for communication, sometimes you still need to protect access to *shared memory* directly, especially for complex state or when performance is critical. The `sync` package provides traditional synchronization primitives:

*   **`sync.Mutex` (Mutual Exclusion Lock):** Allows only *one* goroutine to access the protected data at a time.
    *   `mutex.Lock()`: Locks the mutex (blocks if already locked).
    *   `mutex.Unlock()`: Unlocks the mutex. Common pattern: `defer mutex.Unlock()`.
*   **`sync.RWMutex` (Reader/Writer Mutex):** Allows *many* goroutines to read data concurrently, but only *one* goroutine can write at a time (and no readers while a writer is active). More optimal if data is read more often than written.
    *   `rwMutex.RLock()`, `rwMutex.RUnlock()`: Lock/unlock for readers.
    *   `rwMutex.Lock()`, `rwMutex.Unlock()`: Lock/unlock for writers.
*   **`sync.Once`:** Guarantees a function is executed *exactly once*, even if called from multiple goroutines. Useful for singleton initialization or one-time setup.
    *   `once.Do(func() { ... })`
*   **`sync.Pool`:** A temporary cache for reusable objects, reducing pressure on the garbage collector. Useful for objects that are frequently created and discarded (e.g., buffers).

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter safe for concurrent access using Mutex
type SafeCounter struct {
	mu    sync.Mutex // Mutex to protect access to 'value'
	value int
}

// Method to increment counter (needs Lock/Unlock)
func (c *SafeCounter) Increment() {
	c.mu.Lock()   // Lock before accessing/modifying value
	defer c.mu.Unlock() // Ensure Unlock is called when the function finishes
	c.value++
}

// Method to get counter value (needs Lock/Unlock)
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	numIncrements := 1000

	wg.Add(numIncrements)
	for i := 0; i < numIncrements; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter.Value()) // Should be 1000
}
```

With `go mod tidy`, `go.mod` and `go.sum` will be updated automatically to reflect the dependencies actually used in your code.

#### 10.3. Working with Versions
Go Modules use **Semantic Versioning (SemVer)**. You can specify a specific version, the latest version, or even a specific commit hash when adding dependencies:
```bash
# Get specific version
go get github.com/gin-gonic/gin@v1.7.4

# Get latest version (same major version)
go get github.com/gin-gonic/gin@latest

# Get specific commit hash (less common)
go get github.com/gin-gonic/gin@abcdef123456
```
Go will try to find compatible versions for all your dependencies.

#### 10.4. Vendor Dependencies (`go mod vendor`)
The `go mod vendor` command copies all your project's dependencies into a `vendor` directory at the project root. If a `vendor` directory exists, the `go build` command will use packages from there by default, instead of downloading them from the module cache.
```bash
go mod vendor
```
**When to use vendor?**
*   When you need fully *offline* and *reproducible* builds without relying on proxies or external sources.
*   For strict security or license audits, where you need an exact copy of the dependency code.
*   In CI/CD build environments that might have limited network access.

Although Go Modules with the local cache and `go.sum` are already very reliable, vendoring remains a valid option for specific use cases.

#### 10.5. Private Dependencies
If your project depends on private repositories (e.g., on GitHub Enterprise or a private GitLab instance), you need to configure Go to know about them. Use the `GOPRIVATE` environment variable:
```bash
export GOPRIVATE=*.corp.example.com,github.com/yourorganization/*
```
This tells the Go tool not to use public proxies (like `proxy.golang.org`) for matching paths and to use Git (or another appropriate protocol) directly.

---

### 11. Testing in Go

Go has *first-class* support for testing integrated directly into the toolchain and standard library (`testing` package). Go's approach to testing is simple, conventional, and effective.

#### 11.1. Testing Basics (`*_test.go`)
*   Test files must end with `_test.go` (e.g., `calculator_test.go` to test `calculator.go`).
*   Test files reside in the *same package* as the code being tested.
*   Test functions start with `Test` followed by a descriptive name (with a capital first letter), and accept a single argument `*testing.T`.
    ```go
    package calculator

    import "testing"

    func TestAdd(t *testing.T) {
        result := Add(2, 3)
        expected := 5
        if result != expected {
            // Report failure without stopping other tests
            t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
        }
    }

    func TestSubtract(t *testing.T) {
         // ... test logic ...
         if /* failure condition */ {
            // Report failure AND stop execution of this specific test
            t.Fatalf("Subtract(...) failed: expected %v, got %v", expected, actual)
         }
         // Code after t.Fatal or t.Fatalf will not be executed
    }

    func TestMultiply(t *testing.T) {
        // You can also print logs during testing (only visible if test fails or -v flag is used)
        t.Log("Testing multiplication...")
        // ... test logic ...
        if /* failure condition */ {
             t.Error("Multiplication failed")
        }
    }
    ```
*   `t.Error`/`t.Errorf`: Logs the failure and continues execution of the test (and other tests).
*   `t.Fatal`/`t.Fatalf`: Logs the failure and immediately stops execution of the *current test function* (other tests in the same file or package will still run). Use this if the test cannot proceed after the failure.
*   `t.Log`/`t.Logf`: Prints information (useful for debugging). Output is only shown if the test fails or if the `-v` (verbose) flag is used when running `go test`.

#### 11.2. Running Tests
Use the `go test` command:
```bash
# Run all tests in the current directory
go test

# Run all tests in the current directory and subdirectories
go test ./...

# Run tests with verbose output (showing test names and logs)
go test -v

# Run tests whose names match a regular expression
go test -run TestAdd # Only run TestAdd
go test -run ^TestAdd$ # Only TestAdd (exact match)
go test -run /Sub. # Run TestSubtract or other tests containing "Sub"

# Run tests and calculate coverage
go test -cover

# Generate a coverage profile for further analysis
go test -coverprofile=coverage.out

# View the coverage report in HTML format in the browser
go tool cover -html=coverage.out
```

#### 11.3. Table-Driven Tests
This is a very common pattern in Go for efficiently testing various inputs and outputs for the same function.
```go
package calculator

import "testing"

func TestAddTableDriven(t *testing.T) {
    // Define the structure for each test case
    testCases := []struct {
        name     string // Descriptive name for the test case
        a, b     int
        expected int
    }{
        {"Positive", 2, 3, 5},
        {"Negative", -1, -5, -6},
        {"Zero", 0, 10, 10},
        {"Positive Negative", 5, -3, 2},
    }

    // Iterate through each test case
    for _, tc := range testCases {
        // Run each case as a separate sub-test
        t.Run(tc.name, func(st *testing.T) { // Use st *testing.T for sub-test
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                st.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```
Using `t.Run` allows for better isolation between test cases and clearer output when failures occur.

#### 11.4. Benchmark Tests
For measuring code performance.
*   Benchmark functions start with `Benchmark` and accept `*testing.B`.
*   The code to be measured is placed inside a `for i := 0; i < b.N; i++` loop. `b.N` is adjusted automatically by the testing framework to achieve stable measurements.
```go
package mypackage

import "testing"

func BenchmarkMyFunction(b *testing.B) {
    // Perform expensive setup outside the loop (if any)
    setupData := prepareExpensiveData()

    b.ResetTimer() // Reset the timer after setup

    for i := 0; i < b.N; i++ {
        // Code whose performance is to be measured
        MyFunction(setupData)
    }

    // b.StopTimer() // Can be used if there is teardown after the loop
}
```
Running benchmarks:
```bash
# Run all benchmarks in the current directory
go test -bench=.

# Run benchmarks matching a regex
go test -bench=MyFunction

# Run benchmarks and display memory allocations
go test -bench=. -benchmem
```

#### 11.5. Example Tests
Used for documentation and ensuring code examples work as expected.
*   Example functions start with `Example`.
*   The output printed to `stdout` is compared with the `// Output:` comment at the end of the function.
```go
package greetings

import "fmt"

func ExampleSayHello() {
	message := SayHello("Gopher")
	fmt.Println(message)
	// Output: Hello, Gopher! Welcome!
}

func ExampleSayGoodbye() {
    fmt.Println(SayGoodbye("Alice"))
    fmt.Println(SayGoodbye("Bob"))
    // Unordered output:
    // Goodbye, Alice. See you!
    // Goodbye, Bob. See you!
} // Use "Unordered output:" if the order is not guaranteed
```
Example tests are run with `go test` and are also displayed by `godoc`.

#### 11.6. Test Fixtures (Setup & Teardown)
For complex setup/teardown (e.g., database connections, test servers):
*   **Helper Functions:** The most common way is to create helper functions that perform setup and return resources along with a teardown function.
*   **`TestMain`:** A special function `func TestMain(m *testing.M)` can be defined per test package. This allows for global setup before all tests in the package run and teardown afterward.
    ```go
    package mypackage_test // Often in a _test package for black-box testing

    import (
        "fmt"
        "os"
        "testing"
    )

    func TestMain(m *testing.M) {
        fmt.Println("Performing global setup...")
        // Setup (e.g., start database docker container)

        // Run all tests in the package
        exitCode := m.Run()

        fmt.Println("Performing global teardown...")
        // Teardown (e.g., stop & remove docker container)

        // Exit with the exit code from m.Run()
        os.Exit(exitCode)
    }

    func TestSomething(t *testing.T) {
        t.Log("Running TestSomething...")
        // Test logic...
    }

    func TestAnother(t *testing.T) {
         t.Log("Running TestAnother...")
        // Test logic...
    }
    ```

---

### 12. Standard Library

One of Go's greatest strengths is its rich, comprehensive, and well-designed standard library. Often, you can accomplish many tasks without needing external dependencies. It's highly recommended to familiarize yourself with it.

Some key packages (this is not an exhaustive list):

*   **`fmt`**: Functions for formatted I/O (similar to C's `printf` and `scanf`), like `Println`, `Printf`, `Scanf`.
*   **`io`**: Basic primitives for I/O, including the fundamental `io.Reader` and `io.Writer` interfaces.
*   **`os`**: Functions for interacting with the operating system (files, directories, environment variables, processes), like `os.Open`, `os.Create`, `os.ReadFile`, `os.Args`, `os.Getenv`.
*   **`bufio`**: Implementation of buffered I/O for `io.Reader` and `io.Writer`, improving efficiency (e.g., `bufio.NewReader`, `bufio.NewScanner`).
*   **`strings`**: Utility functions for string manipulation (split, join, contains, replace, trim, etc.).
*   **`strconv`**: Conversion between strings and basic types (int, float, bool) and vice versa.
*   **`encoding/json`**: Encoding and decoding JSON data (`json.Marshal`, `json.Unmarshal`).
*   **`encoding/xml`**, **`encoding/csv`**, **`encoding/base64`**: For other data formats.
*   **`net/http`**: A very powerful and widely used implementation of HTTP clients and servers. Building web servers or APIs in Go is very easy with this package.
*   **`net`**: Basic package for networking (TCP/IP, UDP, domain name resolution).
*   **`sync`**: Basic concurrency primitives like `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Pool`.
*   **`sync/atomic`**: Low-level atomic operations for basic types, useful for concurrency without mutexes in specific cases.
*   **`context`**: Mechanism for managing cancellation, deadlines, and request-scoped data across API calls and goroutines. Crucial in server applications and concurrency.
*   **`time`**: Functions for measuring and displaying time (`time.Now`, `time.Parse`, `time.Format`, `time.Duration`, `time.Ticker`, `time.Timer`).
*   **`regexp`**: Regular expression implementation.
*   **`sort`**: Sorting algorithms for slices and custom data types (requires implementing the `sort.Interface` interface).
*   **`flag`**: Command-line argument parsing.
*   **`log`**: Simple logging. For serious applications, third-party libraries (like `zerolog`, `zap`) are often preferred for their structured logging features.
*   **`testing`**: Support for unit tests, benchmark tests, and example tests (as discussed in the previous section).
*   **`reflect`**: Implementation of runtime reflection. Allows dynamic inspection and manipulation of Go objects. Very powerful, but use with caution as it can be complex, slower, and bypasses type safety.
*   **`unsafe`**: Contains operations that violate Go's type safety (like pointer arithmetic). Only use if you *really* know what you're doing, typically for low-level performance optimization or C interoperability.
*   **`crypto/...`**: Packages for cryptographic functions (hashing like MD5, SHA1, SHA256; encryption AES, RSA; etc.).
*   **`database/sql`**: Generic interface for interacting with SQL databases. You need a specific database driver (e.g., `github.com/lib/pq` for PostgreSQL, `github.com/go-sql-driver/mysql` for MySQL).
*   **`text/template`** and **`html/template`**: Server-side templating systems. `html/template` automatically performs escaping to prevent XSS, making it safer for HTML output.

**How to Explore:**
*   Use the official documentation: [pkg.go.dev](https://pkg.go.dev/)
*   Use the `godoc` command:
    ```bash
    # View documentation for the fmt package in the terminal
    godoc fmt

    # View documentation for the Println function from the fmt package
    godoc fmt Println

    # Run a local documentation server at http://localhost:6060
    godoc -http=:6060
    ```

---

### 13. Building and Deploying Go Applications

Go is designed to make the build and deployment process simple.

#### 13.1. Building an Executable (`go build`)
The `go build` command compiles Go packages and their dependencies into a single executable binary.
```bash
# Be in your project directory (containing main.go)
# Build an executable with the default name (directory name or main.go)
go build

# Build an executable with a specific output name
go build -o myapp

# Build for another directory
go build ./cmd/mycliprogram
```
By default (especially on Linux without Cgo), Go produces a **statically linked binary**. This means the binary contains all the code needed to run, including the Go runtime and all dependencies (already compiled). You don't need to install the Go runtime or other libraries on the target server. Just copy the binary!

#### 13.2. Cross-Compilation
Go makes it easy to compile for different operating systems and architectures than your development machine. This is done by setting the `GOOS` (Target Operating System) and `GOARCH` (Target Architecture) environment variables before running `go build`.

Some common targets:
*   **Linux (64-bit):** `GOOS=linux GOARCH=amd64 go build -o myapp-linux`
*   **Windows (64-bit):** `GOOS=windows GOARCH=amd64 go build -o myapp.exe`
*   **macOS (64-bit, Intel):** `GOOS=darwin GOARCH=amd64 go build -o myapp-mac-intel`
*   **macOS (64-bit, Apple Silicon):** `GOOS=darwin GOARCH=arm64 go build -o myapp-mac-arm`
*   **Linux (ARM 64-bit):** `GOOS=linux GOARCH=arm64 go build -o myapp-linux-arm64` (common for Raspberry Pi 4, AWS Graviton servers)

You can see all supported combinations with: `go tool dist list`

#### 13.3. Reducing Binary Size
Go binaries can be relatively large due to static linking. For production deployments, you might want to reduce their size:
```bash
# Use ldflags to strip debug symbols and DWARF table
# -s: Omit the symbol table
# -w: Omit the DWARF debugging information
go build -ldflags="-s -w" -o myapp-small
```
This can significantly reduce the file size without affecting functionality (but makes debugging harder if a crash occurs).

#### 13.4. Deployment Strategies
Since Go produces a single binary, deployment can be very simple:

1.  **Direct Copy (SCP/Rsync):** The easiest way for a single server or a few servers. Build the binary (cross-compile if necessary), then copy it to the server and run it. You might need to manage the process using `systemd`, `supervisor`, or `screen`/`tmux`.
2.  **Container (Docker):** The most popular approach today. It fits perfectly with Go because static binaries allow for *very small* Docker images. Use a *multi-stage build* in your `Dockerfile`:
    ```dockerfile
    # ---- Builder Stage ----
    # Use the official Go image as a builder
    FROM golang:1.21-alpine AS builder

    # Set the working directory inside the container
    WORKDIR /app

    # Copy go.mod and go.sum first for layer caching
    COPY go.mod go.sum ./
    # Download dependencies
    RUN go mod download

    # Copy the entire source code
    COPY . .

    # Build the Go application, disable CGO, build statically, strip debug info
    RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /myapp ./cmd/myapp

    # ---- Final Stage ----
    # Use a very small base image (scratch or alpine)
    # scratch is an empty image, the most minimal
    FROM scratch
    # Or use alpine if you need a shell or basic tools / CA certificates
    # FROM alpine:latest
    # RUN apk --no-cache add ca-certificates

    # Set the working directory
    WORKDIR /

    # Copy only the built executable from the builder stage
    COPY --from=builder /myapp /myapp
    # If using alpine and need certificates (e.g., for outgoing HTTPS)
    # COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

    # Port the application will expose (if it's a web server)
    # EXPOSE 8080

    # Command to run the application when the container starts
    ENTRYPOINT ["/myapp"]
    ```
    The `scratch` image results in the smallest container but lacks a shell or basic filesystem. `alpine` is slightly larger but provides a useful minimal Linux environment.
3.  **Serverless (AWS Lambda, Google Cloud Functions, Azure Functions):** Go is a great choice for serverless due to its fast startup times and small deployment size (especially with stripped binaries and minimal Docker images). Cloud platforms usually provide native Go runtimes or allow you to bring your own (e.g., via containers).
4.  **Platform as a Service (PaaS - Heroku, Google App Engine):** These platforms often have built-in support for Go, simplifying the build and deployment process further.
5.  **Kubernetes:** If using Docker, the next step for large-scale orchestration is Kubernetes. You'll define Deployments, Services, etc., to manage your Go containers.

#### 13.5. Application Configuration
Deployed applications need configuration. Common approaches in Go:
*   **Environment Variables:** Easy to set in various environments (local, Docker, Kubernetes, PaaS). Use `os.Getenv`. Often considered best practice for many configurations.
*   **Command-Line Flags:** Useful for CLI tools or options at startup. Use the `flag` package.
*   **Configuration Files:** For more complex configurations. Common formats are JSON, YAML, or TOML. You'll need external libraries (e.g., `viper`) or manual parsing (`encoding/json`, etc.).

---

### 14. Next Steps & Advanced Topics

After mastering these basics, here are some areas for further exploration:

*   **Advanced Concurrency Patterns:** Learn patterns like Pipelines, Fan-in/Fan-out, Cancellation Propagation with `context`, Rate Limiting, Circuit Breakers.
*   **Reflection (`reflect`)**: Understand when and how to use reflection safely and effectively, and its performance implications.
*   **Cgo:** Learn how to call C code from Go and vice versa. Useful for leveraging existing C libraries or for critical performance optimizations. Understand the overhead and complexity it introduces.
*   **Profiling and Optimization:** Use Go's built-in tools (`pprof`) to analyze CPU, memory, blocking, and goroutines. Learn how to identify bottlenecks and optimize your code.
    *   `runtime/pprof`: For profiling programs from within the code.
    *   `net/http/pprof`: For profiling live web applications.
    *   `go tool pprof`: The tool for analyzing profile data.
*   **Generics (Go 1.18+):** Learn how to use type parameters to write more flexible and reusable code without sacrificing type safety. Understand when generics are appropriate.
*   **Web Frameworks:** Explore popular frameworks like [Gin](https://github.com/gin-gonic/gin), [Echo](https://github.com/labstack/echo), [Chi](https://github.com/go-chi/chi), or [Fiber](https://github.com/gofiber/fiber) to speed up API and web application development. Understand the trade-offs between using a framework versus the standard `net/http` package.
*   **Database/ORM:** Dive deeper into `database/sql` and its drivers. Learn helper libraries like `sqlx` or ORMs (Object-Relational Mappers) like [GORM](https://gorm.io/) or [SQLBoiler](https://github.com/volatiletech/sqlboiler). Understand the pros and cons of ORMs.
*   **gRPC:** A high-performance RPC (Remote Procedure Call) framework popular for microservice communication. Go has excellent gRPC support.
*   **Structured Logging:** Use libraries like [Zerolog](https://github.com/rs/zerolog) or [Zap](https://github.com/uber-go/zap) for logging that is more efficient and easily parsed by log aggregation systems.
*   **Go Internals:** If interested, learn about the Go runtime, garbage collector, goroutine scheduler, and how Go works under the hood.

---

### 15. Contributing

If you want to contribute to this project (e.g., fix typos, add examples, suggest new sections), please follow these standard steps:

1.  **Fork** this repository.
2.  Create a **new branch** for your feature or fix (`git checkout -b feature/feature-name` or `fix/bug-description`).
3.  Make your changes. Add explanations, code examples, or relevant fixes.
4.  If you add code, make sure to add corresponding **tests** (if applicable).
5.  Ensure your code follows **idiomatic Go style** and passes linters (if the project uses them).
6.  **Commit** your changes with clear messages (`git commit -m "feat: Add explanation about context"`).
7.  **Push** your branch to your fork (`git push origin feature/feature-name`).
8.  Open a **Pull Request (PR)** from your branch in your fork to the `main` (or `master`) branch of the original repository.
9.  Explain your changes in the PR description.

If you find issues or have suggestions, feel free to open an **Issue** in this repository.

---

### 16. MIT License

Happy Go Learning! This language is powerful, efficient, and fun to use. Don't hesitate to experiment, read others' code, and contribute to the community. The best journey in learning a programming language is by building something! Happy coding! 🚀
