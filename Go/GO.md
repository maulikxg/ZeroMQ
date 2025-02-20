
---

### **1. Packages in Go**
#### **What is a Package?**
- A **package** is a collection of Go source files in the same directory that are compiled together.
- Every Go program is made up of packages.
- The `main` package is special—it defines an executable program.

#### **Types of Packages**
1. **Executable Package (`main`)**
    - Any package named `main` can be compiled to produce an executable file.
    - It must contain a `main()` function.

2. **Library Packages**
    - These are reusable code libraries that can be imported.
    - Examples: `fmt`, `math`, `net/http`, `encoding/json`.

#### **Example 1: A Simple Main Package**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
- This program has the **`main`** package and prints a message.

#### **Example 2: Creating a Custom Package**
1. **Create a directory `mathutil` and a file `mathutil.go`**
```go
// File: mathutil/mathutil.go
package mathutil

// Add function - exported
func Add(a, b int) int {
    return a + b
}

// multiply function - unexported (lowercase first letter)
func multiply(a, b int) int {
    return a * b
}
```

2. **Use this package in `main.go`**
```go
package main

import (
    "fmt"
    "myproject/mathutil" // Import the custom package
)

func main() {
    result := mathutil.Add(10, 5)
    fmt.Println("Addition:", result)
}
```
- `Add` is an **exported function** (first letter capitalized).
- `multiply` is **not exported** (cannot be accessed outside `mathutil`).

#### **Edge Cases**
- **Importing Unused Packages**
    - Go throws an error for unused imports.
    - Use `_` (blank identifier) to import but not use a package.
  ```go
  import _ "fmt"
  ```
- **Name Conflicts**
    - If two packages have the same name, alias them:
  ```go
  import m "mathutil"
  fmt.Println(m.Add(5, 5))
  ```

---

### **2. Imports in Go**
#### **Importing Standard & Custom Packages**
- Standard Library:
  ```go
  import "fmt"
  ```
- Multiple Imports (Grouped)
  ```go
  import (
      "fmt"
      "math"
  )
  ```
- Importing Custom Packages
    - The package must be in `GOPATH` or module-based.

#### **Example: Importing a Package with Alias**
```go
import m "math"

func main() {
    fmt.Println(m.Sqrt(25))
}
```

#### **Edge Cases**
- **Import Loops**
    - Import cycles are not allowed. Example:
      ```go
      package A
      import "B"
  
      package B
      import "A"
      ```
        - This will throw an **import cycle error**.

---


## **3. Exported Names in Go**
### **What Are Exported and Unexported Names?**
- In Go, a **name (variable, function, struct, or method)** is:
    - **Exported** if it **starts with an uppercase letter**.
    - **Unexported** if it **starts with a lowercase letter**.

- Exported names can be accessed from other packages.
- Unexported names are private to the package.

### **Example: Exported vs. Unexported**
```go
package mypackage

// Exported function (Accessible outside the package)
func PublicFunction() {
    fmt.Println("I am an exported function")
}

// Unexported function (Not accessible outside this package)
func privateFunction() {
    fmt.Println("I am a private function")
}
```

Now, if we try to call `privateFunction()` from another package, it will result in a **compilation error**.

### **Example: Using an Exported Name**
```go
package main

import (
    "fmt"
    "mypackage"
)

func main() {
    mypackage.PublicFunction() // ✅ Works
    mypackage.privateFunction() // ❌ ERROR: cannot refer to unexported function
}
```

### **Edge Cases**
1. **Exporting Structs but Keeping Fields Unexported**
   ```go
   package person

   type Person struct {
       Name  string // Exported
       age   int    // Unexported
   }
   ```
    - `Name` is accessible from other packages, but `age` is not.

2. **Workaround for Unexported Fields: Getter/Setter**
   ```go
   package person

   type Person struct {
       name string
   }

   func (p *Person) GetName() string {
       return p.name
   }
   ```
    - Now we can indirectly access `name` using `GetName()`.

---

## **4. Functions in Go**
### **Basic Function Syntax**
```go
func add(a int, b int) int {
    return a + b
}
```
- Functions can have **parameters and return values**.

### **Example: Multiple Return Values**
```go
func divide(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}
```
```go
q, r := divide(10, 3)
fmt.Println("Quotient:", q, "Remainder:", r)
```

### **Named Return Values**
```go
func rectangleDimensions(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return
}
```
- Named return values allow us to **omit the `return` values** explicitly.

### **Variadic Functions**
- Functions that accept **variable numbers of arguments**:
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```
```go
fmt.Println(sum(1, 2, 3, 4, 5)) // ✅ Works with multiple arguments
```

### **Function as a Parameter**
```go
func applyFunction(f func(int, int) int, a int, b int) int {
    return f(a, b)
}

func multiply(x, y int) int {
    return x * y
}

fmt.Println(applyFunction(multiply, 3, 4)) // ✅ Output: 12
```

### **Edge Cases**
1. **Returning a Function**
   ```go
   func multiplier(factor int) func(int) int {
       return func(x int) int {
           return x * factor
       }
   }

   double := multiplier(2)
   fmt.Println(double(5)) // ✅ Output: 10
   ```
    - We can **return functions** from functions.

2. **Closure Capturing Variables**
   ```go
   func counter() func() int {
       count := 0
       return func() int {
           count++
           return count
       }
   }

   next := counter()
   fmt.Println(next()) // ✅ Output: 1
   fmt.Println(next()) // ✅ Output: 2
   ```
    - Each `counter()` call creates a **new closure** with its own state.

    
## 🔹 **Where is `count` Stored?**
Normally, when a function exits, its stack is cleaned up, but **closures work differently**.

✅ The `count` variable is **not stored in the stack**, but in the **heap**.

- When `counter()` runs, `count` is **allocated on the heap** (instead of the stack).
- The anonymous function **captures** `count`, which keeps it alive even after `counter()` finishes.
- This means `count` persists between function calls.

---

## 🔹 **How Closures Work in Memory**
1. `counter()` is called → `count := 0` is **created on the heap**.
2. The anonymous function is returned but **remembers** the `count` variable.
3. `next()` calls the function and updates `count`, which persists in memory.

---

## 🔹 **Why is This Useful?**
Closures are useful for:
- **Maintaining state** between function calls without using global variables.
- **Encapsulating logic** inside functions.
- **Creating factory functions** (like `counter()`).

---

## 🔹 **Key Takeaways**
- Closures **"trap"** variables from their surrounding scope and keep them alive.
- Variables captured by a closure are **stored on the heap**, not the stack.
- Each call to `counter()` creates a **new instance** of `count`, meaning different counters don’t interfere.

---

### 🔥 **Example with Multiple Closures**
```go
counter1 := counter()
counter2 := counter()

fmt.Println(counter1()) // 1
fmt.Println(counter1()) // 2
fmt.Println(counter2()) // 1 (Separate instance)
fmt.Println(counter2()) // 2
```
Since `counter1` and `counter2` are **separate closures**, they maintain **independent** `count` variables.


---

## **5. Variables in Go**
### **Declaring Variables**
```go
var name string = "GoLang"
var age int = 25
var isDeveloper bool = true
```

### **Short Variable Declaration (`:=`)**
```go
count := 42  // Automatically infers type (int)
message := "Hello, Go!" // Automatically infers type (string)
```
- **Only works inside functions**.

### **Edge Cases**
1. **Cannot Use `:=` Outside a Function**
   ```go
   count := 100  // ❌ ERROR: "count := 100" outside a function
   ```
    - Use `var` for global variables.

2. **Redeclaring Variables in the Same Scope**
   ```go
   count := 5
   count := 10 // ❌ ERROR: Cannot redeclare variable
   ```

---

## **6. Constants in Go**
### **Declaring Constants**
```go
const Pi = 3.14
const Greeting = "Hello, Go!"
```
- Constants **cannot be changed** once declared.

### **Typed vs. Untyped Constants**
```go
const x int = 10   // Typed constant
const y = 20       // Untyped constant
```
- **Typed constants** have a fixed type (`int`, `string`, etc.).
- **Untyped constants** can **adapt their type** in different contexts.

### **Constant Expressions**
```go
const a = 5 * 2  // Computed at compile-time
```

### **Enumerations with `iota`**
```go
const (
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
)
```
- `iota` is an **auto-incrementing constant generator**.

### **Edge Cases**
1. **Cannot Assign to a Constant**
   ```go
   Pi = 3.1415 // ❌ ERROR: Cannot assign to Pi
   ```

2. **Using `iota` in Bitwise Flags**
   ```go
   const (
       Read  = 1 << iota  // 1 (0001)
       Write              // 2 (0010)
       Execute            // 4 (0100)
   )
   ```

---




## **7. Basic Types in Go**
Go has the following basic types:
### **Integer Types**
| Type | Size | Range |
|------|------|--------------------------------|
| `int8` | 8-bit | -128 to 127 |
| `int16` | 16-bit | -32,768 to 32,767 |
| `int32` | 32-bit | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64-bit | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `uint8` | 8-bit | 0 to 255 |
| `uint16` | 16-bit | 0 to 65,535 |
| `uint32` | 32-bit | 0 to 4,294,967,295 |
| `uint64` | 64-bit | 0 to 18,446,744,073,709,551,615 |

- `int` and `uint` are **platform-dependent**:
    - 32-bit systems: `int` is `int32`, `uint` is `uint32`
    - 64-bit systems: `int` is `int64`, `uint` is `uint64`

### **Floating-Point Types**
| Type | Size | Precision |
|------|------|-----------|
| `float32` | 32-bit | ~6 decimal places |
| `float64` | 64-bit | ~15 decimal places |

```go
var f1 float32 = 3.14159
var f2 float64 = 2.718281828459
```

### **Complex Numbers**
- Go supports **complex numbers** with `complex64` and `complex128`.
```go
c1 := complex(3, 4) // 3 + 4i
fmt.Println(real(c1), imag(c1)) // Output: 3 4
```

### **Boolean Type**
- `bool` (true/false)
```go
var isActive bool = true
```

### **String Type**
- Immutable sequence of bytes.
- Can contain Unicode characters.
```go
s := "Hello, 世界" 
fmt.Println(len(s)) // Counts bytes, not characters
```

### **Edge Cases**
1. **Integer Overflow**
   ```go
   var x uint8 = 255
   x++ // Wraps around to 0
   fmt.Println(x) // Output: 0
   ```
2. **Floating-Point Precision Issues**
   ```go
   fmt.Println(0.1 + 0.2 == 0.3) // Output: false (Due to floating-point imprecision)
   ```
3. **Strings Are Immutable**
   ```go
   s := "hello"
   s[0] = 'H' // ❌ ERROR: Strings are immutable
   ```

---

## **8. Zero Values in Go**
Every variable in Go has a **default "zero value"** when not explicitly initialized.

| Type | Zero Value |
|------|-----------|
| `int`, `float` | `0`, `0.0` |
| `bool` | `false` |
| `string` | `""` (empty string) |
| `pointer`, `slice`, `map`, `channel`, `interface`, `function` | `nil` |

```go
var a int       // 0
var b float64   // 0.0
var c bool      // false
var d string    // ""
var e []int     // nil
```

### **Edge Cases**
1. **Zero Value for Structs**
   ```go
   type Person struct {
       Name string
       Age  int
   }
   var p Person
   fmt.Println(p) // Output: {"" 0}
   ```
2. **Zero Value for Pointers**
   ```go
   var ptr *int
   fmt.Println(ptr) // Output: nil
   ```

---

## **9. Type Conversions in Go**
Go **does not allow implicit type conversion**.

### **Explicit Type Conversion**
```go
var x int = 10
var y float64 = float64(x) // Convert int → float64
```

### **String to Integer**
```go
import "strconv"

s := "123"
num, err := strconv.Atoi(s) // "123" → 123
```

### **Edge Cases**
1. **Loss of Precision**
   ```go
   var f float64 = 3.99
   var i int = int(f) // Drops decimal part (3, not rounded)
   ```
2. **Converting Between Unsigned & Signed**
   ```go
   var u uint = 255
   var i int = int(u) // Works, but negative numbers may cause issues
   ```

---

## **10. Type Inference in Go**
Go automatically infers the type when `:=` is used.
```go
a := 10       // int
b := 3.14     // float64
c := "hello"  // string
```
- This **avoids explicit type declarations**.

### **Edge Cases**
1. **Cannot Reassign with Different Type**
   ```go
   a := 42
   a = "hello" // ❌ ERROR: Cannot assign string to int
   ```
2. **Mixing Integer & Floating-Point**
   ```go
   var x int = 10
   var y float64 = 3.14
   sum := x + y // ❌ ERROR: mismatched types
   ```

---

## **11. Constants in Go**
Constants are **immutable** and declared using `const`.

```go
const Pi = 3.14159
const Greeting = "Hello, Go!"
```

### **Typed vs. Untyped Constants**
```go
const x int = 10 // Typed constant
const y = 20     // Untyped constant
```

### **iota: Auto-Incrementing Constants**
```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```
- `iota` automatically increments for each line.

### **Edge Cases**
1. **Cannot Modify Constants**
   ```go
   Pi = 3.14 // ❌ ERROR: Cannot assign to Pi
   ```
2. **Skipping Values with `iota`**
   ```go
   const (
       Read  = 1 << iota // 1 (0001)
       Write             // 2 (0010)
       Execute           // 4 (0100)
   )
   ```

---

## **Next Steps**
So far, we’ve covered:
✅ **Basic Types**  
✅ **Zero Values**  
✅ **Type Conversions**  
✅ **Type Inference**  
✅ **Constants & iota**

🔹 Next, I’ll cover:
- **Functions Continued (Higher-Order Functions, Closures)**
- **Multiple Return Values**
- **More Complex Edge Cases**

Would you like me to continue in this structured way, or do you have any specific topics you want more depth on? 🚀

Great! Let's continue with in-depth explanations, real-world scenarios, and edge cases for the next topics. 🚀

---

# **12. Functions Continued (Higher-Order Functions, Closures, and Edge Cases)**

## **12.1 Higher-Order Functions**
Go **treats functions as first-class citizens**, meaning functions can be:  
✅ Passed as arguments  
✅ Returned from other functions  
✅ Assigned to variables

### **Example 1: Function as a Parameter**
```go
package main

import "fmt"

func applyOperation(a int, b int, op func(int, int) int) int {
    return op(a, b)
}

func add(x, y int) int { return x + y }

func main() {
    result := applyOperation(5, 10, add) 
    fmt.Println(result) // Output: 15
}
```
🔹 `applyOperation` takes a function as an argument and applies it to numbers.

### **Example 2: Function Returning a Function**
```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    fmt.Println(double(5)) // Output: 10
}
```
🔹 `multiplier` returns a function that multiplies numbers by a given factor.

### **Edge Cases**
1. **Passing `nil` as a function argument**
   ```go
   var fn func(int) int
   fmt.Println(fn(5)) // ❌ PANIC: calling nil function
   ```
   ✅ Always check for `nil` before calling function variables.

2. **Returning Function Pointers Instead of Values**
   ```go
   func getOperator() func(int, int) int {
       return add // Function pointer
   }
   ```

---

# **13. Multiple Return Values**
Go allows functions to **return multiple values**, commonly used for:  
✅ Returning **results & errors**  
✅ Avoiding **global variables**

### **Example 1: Returning Two Values**
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: 5
    }
}
```
🔹 If `b == 0`, it returns an **error instead of panicking**.

### **Example 2: Ignoring Unwanted Return Values**
```go
quotient, _ := divide(9, 3) // Ignoring the error
```

### **Edge Cases**
1. **Forgetting to Handle Errors**
   ```go
   res, err := divide(10, 0) // ❌ Error ignored, may cause runtime issues
   ```
   ✅ Always handle returned errors!

2. **Using Named Return Values**
   ```go
   func info() (name string, age int) {
       name = "Alice"
       age = 25
       return // Implicitly returns (name, age)
   }
   ```
   🔹 Named return values can **improve readability** but may cause confusion.

---

# **14. Variables with Initializers**
Go supports **declaring and initializing** variables in a single step.

```go
var x = 100 // Type inferred as int
y := "Hello" // Short variable declaration
```

🔹 When to use `var` vs `:=`:  
| `var` | `:=` |
|-------|------|
| Used in **global scope** | Used in **function scope** |
| Explicit type declaration possible | Implicit type inference |
| Can be declared without initialization | Requires initialization |

### **Example: Multiple Initializations**
```go
var a, b, c = 1, 2, "Go"
```

### **Edge Cases**
1. **Re-declaring Variables in the Same Scope**
   ```go
   x := 10
   x := 20 // ❌ ERROR: no new variables
   ```
   ✅ Use `=` instead if modifying an existing variable.

---

# **15. Short Variable Declarations (`:=`)**
✅ Used inside functions  
✅ No need for `var` keyword

### **Example: Declaring Multiple Variables**
```go
name, age := "Alice", 30
fmt.Println(name, age) // Output: Alice 30
```

### **Edge Cases**
1. **Cannot Use `:=` Outside Functions**
   ```go
   package main
   name := "Alice" // ❌ ERROR: syntax error
   ```
   ✅ Use `var` instead for package-level variables.

2. **Partial Redeclaration**
   ```go
   a, b := 1, "Hello"
   a, c := 2, 3.14 // ✅ Allowed: `a` is reassigned, `c` is new
   ```

---

# **16. Zero Values (Revisited with Structs & Slices)**
We covered zero values earlier, but let's see more **practical cases**.

### **Example: Zero Values for Structs**
```go
type Person struct {
    Name string
    Age  int
}

func main() {
    var p Person
    fmt.Println(p) // Output: {"" 0}
}
```

### **Example: Zero Values for Slices**
```go
var numbers []int
fmt.Println(numbers == nil) // Output: true
```
🔹 Slices default to `nil` until initialized.

### **Edge Cases**
1. **Empty Structs Take Up Memory**
   ```go
   type Empty struct{}
   fmt.Println(Empty{}) // Output: {}
   ```
   ✅ **Use empty structs** (`struct{}`) when you need a **zero-size placeholder**.

---

# **17. Type Conversions (Advanced Cases)**
### **Case 1: Integer to String**
```go
import "strconv"

num := 42
str := strconv.Itoa(num) // "42"
```

### **Case 2: String to Float**
```go
f, err := strconv.ParseFloat("3.14", 64)
```

### **Edge Cases**
1. **Converting Large Integers May Lose Precision**
   ```go
   var big int64 = 9223372036854775807
   var small int32 = int32(big) // ❌ Possible overflow
   ```

---

# **18. Constants (Advanced Cases with iota)**
### **Case 1: Skipping Values**
```go
const (
    First  = iota + 1 // 1
    _                 // Skipped value (2)
    Third             // 3
)
```

### **Case 2: Bitwise Flags**
```go
const (
    Read   = 1 << iota // 1 (0001)
    Write              // 2 (0010)
    Execute            // 4 (0100)
)
```
🔹 Used in file permission systems (`chmod` in Unix).

---


# **1️⃣ Type Inference in More Depth**

Go has a **strongly typed** but **type-inferred** system. This means:
- **Explicit type declaration is optional** (when using `:=`)
- **The compiler automatically determines types**

---

## **🔹 Basic Type Inference**

### **Example 1: Implicit Type Inference with `:=`**
```go
package main

import "fmt"

func main() {
    x := 10   // Inferred as int
    y := 3.14 // Inferred as float64
    z := "Go" // Inferred as string

    fmt.Printf("Type of x: %T\n", x) // Output: int
    fmt.Printf("Type of y: %T\n", y) // Output: float64
    fmt.Printf("Type of z: %T\n", z) // Output: string
}
```
🔹 Go **infers types** based on the **right-hand side value**.

---

## **🔹 Inference with Constants**
Go **treats constants differently**—they are **untyped until assigned**.

### **Example 2: Constants Adopting Different Types**
```go
const value = 100 // Untyped constant

var a int = value        // value acts as int
var b float64 = value    // value acts as float64
var c complex128 = value // value acts as complex128

fmt.Printf("Types: %T %T %T\n", a, b, c) 
// Output: int float64 complex128
```
🔹 **Constants are more flexible** in type inference than variables.

---

## **🔹 Inference in Expressions**
When combining different types, Go ensures **precision is not lost**.

### **Example 3: Mixing Integer and Float**
```go
x := 5
y := 2.5
result := x * y  // ❌ ERROR: mismatched types int and float64
```
✅ Solution: Convert `x` explicitly.
```go
result := float64(x) * y // ✅ Works, result is float64
```
---

## **🔹 Edge Cases**
### **Case 1: Unexpected Integer Overflow**
```go
var smallInt int8 = 127
var bigInt = smallInt + 1 // ❌ ERROR: type mismatch
```
✅ Solution: Explicit conversion
```go
var bigInt = int16(smallInt) + 1
```

### **Case 2: Default Float Type**
```go
x := 3.14 
fmt.Printf("%T\n", x) // Output: float64 (not float32)
```
✅ If you need `float32`, **explicitly define it**:
```go
var y float32 = 3.14
```
---

## **🔹 Advanced Type Inference with Interfaces**
If a function returns **interface{}**, type inference is dynamic.

### **Example 4: Detecting Inferred Type at Runtime**
```go
func returnUnknown() interface{} {
    return 42 // Could be anything
}

func main() {
    value := returnUnknown()

    switch v := value.(type) {
    case int:
        fmt.Println("Got an int:", v)
    case string:
        fmt.Println("Got a string:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```
🔹 **Type assertions** allow checking inferred types at runtime.

---

# **2️⃣ Complex Scenarios with Go Types**

Go’s type system is simple but powerful. Let’s explore:  
✅ **Aliased Types**  
✅ **Empty Structs (`struct{}`) for Memory Optimization**  
✅ **Function Types**  
✅ **Custom Types with Methods**

---

## **🔹 1. Aliased Types**
Go allows **type aliases**, which are **distinct but share the same underlying type**.

### **Example 1: Defining an Aliased Type**
```go
type Age int

func main() {
    var a Age = 25
    fmt.Printf("Type: %T, Value: %d\n", a, a) 
    // Output: main.Age, Value: 25
}
```
🔹 `Age` is **not the same as `int`**, so implicit conversion **isn’t allowed**:
```go
var b int = a  // ❌ ERROR: cannot use Age as int
```
✅ Solution: Explicit conversion:
```go
var b int = int(a) // ✅ Works
```

---

## **🔹 2. Empty Structs (`struct{}`) for Memory Optimization**
Go allows **zero-sized types** using `struct{}`.

### **Example 2: Using `struct{}` to Save Memory**
```go
type Empty struct{}

var x Empty
fmt.Println(x) // Output: {}
```
🔹 The **size of `struct{}` is always 0 bytes**.

✅ **Common Use Case: Efficient Maps for Set Implementation**
```go
var seen = make(map[string]struct{})
seen["hello"] = struct{}{}
```
🔹 **Saves memory** compared to `map[string]bool`.

---

## **🔹 3. Function Types (`func()`)**
Go allows **functions to be stored as values**.

### **Example 3: Assigning Functions to Variables**
```go
type MathFunc func(int, int) int

func add(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }

func main() {
    var operation MathFunc = add
    fmt.Println(operation(5, 3)) // Output: 8

    operation = multiply
    fmt.Println(operation(5, 3)) // Output: 15
}
```
🔹 Functions **can be stored in variables and switched dynamically**.

---

## **🔹 4. Custom Types with Methods**
You can define **methods on custom types**.

### **Example 4: Adding Methods to a Custom Type**
```go
type Temperature float64

func (t Temperature) CelsiusToFahrenheit() float64 {
    return float64(t)*1.8 + 32
}

func main() {
    temp := Temperature(30)
    fmt.Println(temp.CelsiusToFahrenheit()) // Output: 86
}
```
🔹 **Methods allow attaching behavior** to custom types.

---

## **🔹 Edge Cases**
### **Case 1: Assigning Structs with Different Field Orders**
```go
type A struct {
    Name string
    Age  int
}

type B struct {
    Age  int
    Name string
}

var a A
var b B = a  // ❌ ERROR: Different struct layout
```
✅ **Structs must match exactly in both field types and order.**

---

# **🔹 Summary**
| Concept | Key Takeaways |
|---------|--------------|
| **Type Inference** | Implicitly assigns types based on values. |
| **Constant Inference** | Constants remain **untyped** until assigned. |
| **Aliased Types** | Custom types improve clarity but need explicit conversion. |
| **Empty Structs** | `struct{}` is a **zero-byte type** used for optimization. |
| **Function Types** | Functions can be assigned, passed, and returned dynamically. |
| **Custom Types with Methods** | Attach behaviors to custom types using methods. |

---
