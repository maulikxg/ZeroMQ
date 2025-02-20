
### **1. Pointers in Go**
A **pointer** is a variable that stores the memory address of another variable. Pointers allow efficient memory usage and can help in modifying function arguments.

#### **Basic Pointer Declaration**
```go
var p *int  // Declares a pointer to an int
```

#### **Assigning and Using Pointers**
```go
x := 42
p := &x  // p stores the memory address of x
fmt.Println(*p) // Dereferencing: prints 42
*p = 100  // Modifies x through p
fmt.Println(x) // Prints 100
```

#### **Edge Cases**
1. **Nil Pointers:**  
   If a pointer is not initialized, it holds a `nil` value, which can cause runtime errors.
   ```go
   var p *int
   fmt.Println(p) // Prints <nil>
   *p = 10 // Causes panic: invalid memory address dereference
   ```

2. **Pointer Arithmetic (Not Supported in Go)**  
   Unlike C, Go doesn’t support pointer arithmetic.
   ```go
   p++ // Compile error!
   ```

---

### **2. Structs in Go**
A **struct** is a composite data type that groups related data.

#### **Defining and Initializing a Struct**
```go
type Person struct {
    Name string
    Age  int
}

var p1 = Person{"Alice", 30}
p2 := Person{Name: "Bob"} // Age defaults to 0
```

#### **Edge Cases**
1. **Struct Comparisons**
    - Structs can be compared if they don’t contain slices, maps, or functions.
   ```go
   type A struct {
       x int
   }
   a1 := A{10}
   a2 := A{10}
   fmt.Println(a1 == a2) // true
   ```

2. **Embedding (Inheritance in Go)**
    - Go doesn’t have classes, but **embedding** provides struct inheritance-like behavior.
   ```go
   type Employee struct {
       Person
       Salary int
   }
   e := Employee{Person: Person{Name: "Charlie", Age: 28}, Salary: 5000}
   fmt.Println(e.Name) // Accessing embedded struct field
   ```

---

### **3. Struct Fields**
Each struct field has a name and type.

#### **Exported vs. Unexported Fields**
- Fields that start with an **uppercase letter** are **exported** (public).
- Fields with a **lowercase letter** are **unexported** (private).

```go
type Example struct {
    PublicField  int
    privateField string
}
```

#### **Edge Cases**
1. **JSON Encoding Issues**
    - Unexported fields won’t be included in JSON output.
   ```go
   type Data struct {
       Visible   string
       hidden    string
   }
   d := Data{"show", "hide"}
   jsonData, _ := json.Marshal(d)
   fmt.Println(string(jsonData)) // {"Visible":"show"}
   ```

---

### **4. Pointers to Structs**
Structs can be used with pointers to avoid copying large data.

```go
type Car struct {
    Brand string
}

c := Car{"Tesla"}
pc := &c
pc.Brand = "BMW" // Modifies original struct
fmt.Println(c.Brand) // BMW
```

---

### **5. Struct Literals**
Structs can be initialized in different ways:

```go
p1 := Person{"Alice", 30}           // Ordered initialization
p2 := Person{Name: "Bob"}           // Named initialization
p3 := &Person{Name: "Charlie", Age: 25} // Pointer to struct
```

Edge Case:
- Omitting values sets fields to their zero values.

---

### **6. Arrays in Go**
An **array** is a fixed-size sequence of elements of the same type.

```go
var arr [3]int // [0, 0, 0]
arr[0] = 10
fmt.Println(arr) // [10, 0, 0]
```

#### **Edge Cases**
1. **Array size is part of the type**
   ```go
   var a [3]int
   var b [4]int
   fmt.Println(a == b) // Compilation error: different types
   ```

2. **Passing Arrays to Functions (Copy Issue)**  
   Arrays are passed **by value**, causing copying.
   ```go
   func modify(arr [3]int) {
       arr[0] = 100
   }
   arr := [3]int{1, 2, 3}
   modify(arr)
   fmt.Println(arr[0]) // Still 1 (not modified)
   ```

---

### **7. Slices in Go**
A **slice** is a dynamic view of an array.

```go
s := []int{1, 2, 3}
s = append(s, 4)
fmt.Println(s) // [1, 2, 3, 4]
```

#### **Edge Cases**
1. **Appending Beyond Capacity**
   ```go
   s := make([]int, 2, 2)
   s = append(s, 100) // Allocates a new underlying array
   ```

2. **Slicing Beyond Length Causes Panic**
   ```go
   arr := []int{1, 2, 3}
   fmt.Println(arr[:4]) // Causes runtime panic
   ```

---

### **8. Slices Are Like References to Arrays**
Modifying a slice modifies the underlying array.

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[:2]
s[0] = 100
fmt.Println(arr) // [100, 2, 3, 4, 5]
```

---

### **9. Slice Literals & Defaults**
Slice literals:
```go
s := []int{10, 20, 30}
```

Default slicing:
```go
arr := []int{1, 2, 3, 4}
fmt.Println(arr[:])  // Full slice
fmt.Println(arr[1:]) // From index 1 to end
fmt.Println(arr[:2]) // From start to index 2
```

---

### **10. Nil Slices**
A nil slice has a `nil` reference but behaves like an empty slice.

```go
var s []int
fmt.Println(s == nil) // true
```

---

### **11. Creating a Slice with `make`**
```go
s := make([]int, 5, 10) // Length 5, capacity 10
fmt.Println(len(s), cap(s)) // 5 10
```

---

### **12. Slices of Slices**
Slices can contain slices.

```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
}
fmt.Println(matrix[1][2]) // 6
```

---

### **13. Appending to a Slice**
```go
s := []int{1, 2}
s = append(s, 3, 4, 5)
fmt.Println(s) // [1, 2, 3, 4, 5]
```

Edge Case:  
Appending large data may allocate a new array.

---

## **1. Range in Go**
The `range` keyword in Go is used to iterate over different data structures like slices, arrays, maps, and strings.

### **Iterating Over a Slice**
```go
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Println("Index:", index, "Value:", value)
}
```
- The **index** and **value** of each element are returned.

### **Ignoring Index or Value**
```go
for _, value := range nums { // Ignoring index
    fmt.Println(value)
}
```
```go
for index := range nums { // Ignoring value
    fmt.Println(index)
}
```

### **Edge Cases**
1. **Modifying Elements in a `range` Loop Doesn't Affect the Slice**
   ```go
   for _, v := range nums {
       v = v * 2 // Doesn't modify the original slice
   }
   fmt.Println(nums) // Output: [10, 20, 30]
   ```
    - This happens because `v` is a copy of the slice element.

2. **Using a Pointer in `range` to Modify the Slice**
   ```go
   for i := range nums {
       nums[i] *= 2 // This modifies the original slice
   }
   fmt.Println(nums) // Output: [20, 40, 60]
   ```

---

## **2. Maps in Go**
A **map** is a built-in data structure that stores key-value pairs.

### **Declaring and Initializing Maps**
```go
var myMap map[string]int // Uninitialized (nil)
myMap = make(map[string]int) // Initialized

students := map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(students["Alice"]) // Output: 25
```

---

## **3. Map Literals**
Using map literals:
```go
grades := map[string]int{
    "Math":    90,
    "Science": 85,
    "English": 88,
}
```
- Keys are unique; duplicate keys overwrite values.

---

## **4. Mutating Maps (Adding, Updating, Deleting)**
### **Adding & Updating Elements**
```go
grades["History"] = 80  // Adds a new key
grades["Math"] = 95  // Updates existing key
```

### **Deleting Elements**
```go
delete(grades, "Science")
```

### **Checking if a Key Exists**
```go
value, exists := grades["Math"]
if exists {
    fmt.Println("Math grade:", value)
} else {
    fmt.Println("Math grade not found")
}
```

### **Edge Cases**
1. **Accessing a Non-Existing Key Returns the Zero Value**
   ```go
   fmt.Println(grades["Geography"]) // Outputs: 0 (default int)
   ```

2. **Nil Maps**
   ```go
   var m map[string]int
   m["test"] = 100 // Causes panic: assignment to entry in nil map
   ```

---

## **5. Function Values**
Functions in Go are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from functions.

### **Assigning a Function to a Variable**
```go
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4)) // 7
```

### **Passing Functions as Arguments**
```go
func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}
result := operate(5, 2, add)
fmt.Println(result) // 7
```

### **Returning Functions from Functions**
```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
fmt.Println(double(5)) // 10
```

---

## **6. Function Closures**
Closures are **functions that capture variables from their surrounding scope**.

### **Example of Closure Capturing a Variable**
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

### **Edge Case**
- The variable `count` is preserved across function calls, unlike normal function variables.

---

## **7. Fibonacci Closure Exercise**
### **Generating Fibonacci Numbers Using a Closure**
```go
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        temp := a
        a, b = b, a+b
        return temp
    }
}

fib := fibonacci()
fmt.Println(fib()) // 0
fmt.Println(fib()) // 1
fmt.Println(fib()) // 1
fmt.Println(fib()) // 2
fmt.Println(fib()) // 3
```
- Each call **remembers** previous values.

---
