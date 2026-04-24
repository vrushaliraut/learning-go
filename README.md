# Learning Go

A learning project to understand Go programming language fundamentals.

## Structure

All chapters are now **modular** - each concept is in its own runnable file!

### Chapter 01: Setup Go Basics
- `types_declaration.go` - String and type basics
- `boolean_declaration.go` - Boolean values
- `integer_types.go` - Integer operations and types
- `floating_point.go` - Float types and precision
- `strings_and_runes.go` - Strings, runes, and conversions
- `constants.go` - Constants declaration
- `variable_declarations.go` - Variable declaration patterns
- `zero_values.go` - Zero value concepts
- `type_conversion.go` - Type conversion and casting
- `networklimitation.go` - Network requests

### Chapter 02: Types and Variable Declaration
- `explicit_types.go` - Explicit type declarations
- `type_inference.go` - Type inference
- `multiple_declarations.go` - Multiple variable declarations
- `constants_declaration.go` - Constants

### Chapter 03: Composite Types
- `main.go` - Slices and maps examples

### Chapter 05: Functions
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

## Running the Code

**✅ Run individual concepts:**

```bash
# Chapter 01 examples
go run ./ch01setupgo/types_declaration.go
go run ./ch01setupgo/boolean_declaration.go
go run ./ch01setupgo/integer_types.go
go run ./ch01setupgo/floating_point.go
go run ./ch01setupgo/strings_and_runes.go
go run ./ch01setupgo/constants.go
go run ./ch01setupgo/variable_declarations.go
go run ./ch01setupgo/zero_values.go
go run ./ch01setupgo/type_conversion.go
go run ./ch01setupgo/networklimitation.go

# Chapter 02 examples
go run ./ch02_types_declaration/explicit_types.go
go run ./ch02_types_declaration/type_inference.go
go run ./ch02_types_declaration/multiple_declarations.go
go run ./ch02_types_declaration/constants_declaration.go

# Chapter 05 examples (same pattern)
go run ./ch05_functions/calculator.go
go run ./ch05_functions/closures.go
# ... etc
```

**Navigate to a chapter and run a specific file:**

```bash
cd ch01setupgo
go run types_declaration.go
go run boolean_declaration.go
# ... etc
```

## Important Notes

⚠️ **All chapters are modular** - Each Go file has its own `main()` function. You must run individual files, not entire directories.

✅ **Correct:** `go run ./ch01setupgo/types_declaration.go`  
❌ **Incorrect:** `go run ./ch01setupgo` (will error with multiple main functions)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
