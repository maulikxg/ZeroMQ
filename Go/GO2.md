
---
# **Flow Control in Go – Advanced Learning**

---

## **1️⃣ `for` Loop in Depth**

### 🔹 **How Go Compiles and Executes a `for` Loop**
When you write:
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
This is internally converted into:
```go
i := 0
for {
    if !(i < 5) {
        break
    }
    fmt.Println(i)
    i++
}
```
**Takeaway:**
- The compiler **inlines** the loop and generates machine-level conditional jumps.
- `break` is automatically inserted when the condition fails.

### **🛑 Edge Case: What if `i` is Changed Inside the Loop?**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
    i += 2  // 🚨 Modifying loop variable manually
}
```
**Output:**
```
0
3
```
**Issue:**
- The `i++` from the loop header **still executes**, leading to an **unexpected jump** in values.

**Fix:** Avoid modifying the loop variable inside the body unless absolutely necessary.

---

## **2️⃣ `for` Continued (Using `continue`)**

### **🔹 Theory**
- `continue` **skips the rest of the iteration** and jumps to the next cycle.
- It does **not reset the loop variable**.

### **Edge Case: Misuse with `i++`**
```go
for i := 0; i < 5; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```
**Output:**
```
1
3
```
**Why?**
- `continue` **skips `fmt.Println(i)`**, but `i++` still happens.
- This does **not create an infinite loop** but **can lead to unexpected missing outputs**.

---

## **3️⃣ `for` as Go’s `while` Loop**
- If you **omit** initialization and post statements, the loop acts as a `while` loop.

```go
x := 5
for x > 0 {
    fmt.Println(x)
    x--
}
```
🔹 **Internally, Go optimizes this by using a single jump instruction to evaluate the condition.**

### **Edge Case: Unintended Infinite Loop**
```go
x := 5
for x > 0 {
    fmt.Println(x)
}
```
💥 **Issue:** `x--` is missing → The condition `x > 0` never changes → **Infinite loop!**

---

## **4️⃣ `Forever` (Infinite Loop)**
```go
for {
    fmt.Println("Running forever...")
}
```
💡 **Common Use Cases:**  
✅ Background goroutines  
✅ Long-running network services  
✅ Event-driven programs

### **🛑 Edge Case: Resource Consumption**
Infinite loops can lead to **100% CPU usage** if not handled properly.  
🔹 **Fix:** Add a `time.Sleep` inside the loop.
```go
for {
    fmt.Println("Running forever...")
    time.Sleep(time.Second) // Prevents high CPU usage
}
```

---

## **5️⃣ `if` Statement in Go**

### **🔹 Theory**
- Go **doesn’t need parentheses** around conditions (`if (x > 5) {}` ❌).
- `if` statements **can declare variables inside them**.

### **Edge Case: Variable Scope**
```go
if x := 10; x > 5 {
    fmt.Println(x)
}
fmt.Println(x) // ❌ Compile error: x is not defined
```
🔹 The variable `x` **only exists inside the `if` block**.

---

## **6️⃣ `if` with a Short Statement**

### **🔹 Why Use It?**
- Makes conditions **concise**.
- Reduces **boilerplate code**.

```go
if err := someFunction(); err != nil {
    fmt.Println("Error:", err)
}
```
🔹 The variable `err` is **only available inside the `if` block**.

### **Edge Case: Reusing the Variable**
```go
if x := 10; x > 5 {
    fmt.Println(x)
}
fmt.Println(x) // ❌ Error: x is undefined
```
**Fix:** Declare `x` before the `if` statement.

---

## **7️⃣ `switch` in Go**

### **🔹 How Go’s `switch` Works**
- **No automatic fall-through** (unlike C).
- **Evaluates top-down** and exits on the first match.

```go
switch x := 2; x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two") // ✅ First match, stops here
case 3:
    fmt.Println("Three")
default:
    fmt.Println("Unknown")
}
```
💡 **No need for `break` statements like in C!**

### **🛑 Edge Case: Handling Multiple Cases**
```go
switch x {
case 1, 2, 3:
    fmt.Println("1, 2, or 3")
default:
    fmt.Println("Something else")
}
```
🔹 Multiple cases **reduce redundancy**.

---

## **8️⃣ `switch` Without a Condition**

- Works **like multiple `if` statements**.
- Evaluates **first true condition**.

```go
switch {
case x > 10:
    fmt.Println("Greater than 10")
case x > 5:
    fmt.Println("Greater than 5") // ✅ Stops here if x > 5
}
```
🔹 **More readable than nested `if-else` statements**.

---

## **9️⃣ `defer` in Go**

### **🔹 How `defer` Works**
- **Executes functions at the end** of the surrounding function.
- Useful for **resource cleanup**.

```go
func main() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
}
```
**Output:**
```
This runs first
This runs last
```
🔹 **Great for closing files, databases, etc.**

### **🛑 Edge Case: Deferred Function Argument Evaluation**
```go
func main() {
    x := 5
    defer fmt.Println("Deferred:", x)
    x = 10
}
```
**Output:**
```
Deferred: 5
```
💡 **Why?**
- The argument (`x`) **is evaluated immediately**, but execution is delayed.

---

## **1️⃣0️⃣ Stacking `defer` Calls**

- **LIFO (Last In, First Out) execution**.

```go
func main() {
    defer fmt.Println("One")
    defer fmt.Println("Two")
    defer fmt.Println("Three")
}
```
**Output:**
```
Three
Two
One
```
🔹 **Last `defer` statement executes first!**

---

# **Final Summary**
| Concept | Key Takeaway |
|---------|-------------|
| `for` Loop | The only loop in Go, works like `while`. |
| `continue` | Skips the current iteration. |
| Infinite Loop | `for {}` runs forever, must be stopped. |
| `switch` | Stops at the first match, no fall-through. |
| `defer` | Executes at the end, works in LIFO order. |

---
