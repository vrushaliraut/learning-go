# Learning Go

A learning project to understand Go programming language fundamentals.

## Structure

- **ch01setupgo** - Chapter 01: Setup Go basics
- **ch02_types_declaration** - Chapter 02: Types and Variable Declaration
- **ch03_composite_types** - Chapter 03: Composite Types (Slices, Maps)
- **ch05_functions** - Chapter 05: Functions (Modular - each concept is independently runnable)

## Running the Code

To run any chapter:

```bash
# Run Chapter 01
go run ./ch01setupgo

# Run Chapter 02
go run ./ch02_types_declaration

# Run Chapter 03
go run ./ch03_composite_types

# Run Chapter 05 (modular structure)
go run ./ch05_functions/calculator.go
go run ./ch05_functions/anonymous_function.go
go run ./ch05_functions/closures.go
# ... and so on for each concept
```

Or navigate to a chapter directory and run:

```bash
go run .
```

### Chapter 05: Functions - Modular Structure

Chapter 05 has been restructured to be **modular**. Each Go file is independently runnable and demonstrates a specific function concept:

**⚠️ IMPORTANT:** You cannot run the entire `ch05_functions` directory anymore because each file has its own `main()` function. Run individual files instead.

**✅ Correct Usage:**
```bash
# Run individual concepts
go run ./ch05_functions/calculator.go
go run ./ch05_functions/anonymous_function.go
go run ./ch05_functions/closures.go
go run ./ch05_functions/function_concept.go
# ... etc
```

**❌ This will NOT work:**
```bash
go run ./ch05_functions  # Error: multiple main functions
```

**Available runnable files:**
- `calculator.go` - Calculator with error handling
- `anonymous_function.go` - Anonymous functions
- `closures.go` - Closures and variable capture
- `filelengthwithdefer.go` - File operations with defer
- `function_concept.go` - Basic function concepts
- `function_map.go` - Functions as map values
- `functionvalues.go` - Functions as first-class values
- `multiplereturnvalue.go` - Multiple return values
- `named_and_option_param.go` - Named and optional parameters
- `package_level_anonymous.go` - Package-level anonymous functions
- `simulatedowhile.go` - Simulating do-while loops
- `singlevariabletrap.go` - Single variable assignment trap
- `variadic_parameter.go` - Variadic parameters

Each file can be run independently: `go run ./ch05_functions/filename.go`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
