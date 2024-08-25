# Data Types

1. **String:** Strings are used to represent text and are surrounded by `("")` or `('')`.
   
2. **Number:** There are two types:
   - **Integer:** It can be `uint` (unsigned integer) or `int`.
   - **Float:** It can be either `float32` or `float64`.

3. **Booleans:** It has two values: `true` or `false`.

4. **Arrays:** A collection of elements of similar data types.

5. **Slices:** A container like an array but more flexible.

6. **Maps:** Key-value pairs where keys and values can be strings, numbers, or other data types.

# Dynamically vs Statically Typed Languages

- **Statically typed languages:** Throw an error at compile time when using the wrong data type. Examples: C++, Java.
- **Dynamically typed languages:** Do not enforce a strict data type system. Examples: Python, JavaScript.
- **Go:** Go has a concept of data types that can be explicitly declared by the programmer or inferred by the compiler.

# Variables

- Variables are used to store data.

**Syntax:**
```go
var <variable name> <data type> = <value>
