# **ZeroMQ Thread Safety: Context vs. Sockets**

## **1️⃣ Process vs. Thread (General Understanding)**

| Feature          | **Process** 🏢 | **Thread** 🧵 |
|-----------------|--------------|-------------|
| **Definition**  | A separate execution unit with its own memory space. | A lightweight execution unit within a process, sharing memory. |
| **Memory Space** | Has its own **isolated** memory space. | Shares memory **with other threads** of the same process. |
| **Communication** | Uses **IPC (Inter-Process Communication)** mechanisms like message queues, pipes, or ZeroMQ sockets. | Uses **shared memory**, making communication faster but riskier (race conditions). |
| **Safety** | No data corruption since processes **don't share memory**. | Needs synchronization (mutex, locks) to prevent **race conditions**. |
| **Efficiency** | More overhead (context switching, more system resources). | More efficient, lightweight, and faster context switching. |
| **ZeroMQ Handling** | Can **safely** use the same socket instance in different processes. | **Cannot** share the same socket across threads. |

---

## **2️⃣ Why is ZeroMQ Context Thread-Safe but Sockets Are Not?**

### **✅ ZeroMQ Context (`zmq.Context`) is Thread-Safe**
- **Manages sockets, memory, I/O threads, and message queues.**
- **Multiple threads can safely create sockets from a shared context.**
- **Internally synchronized** to prevent race conditions.

### **❌ ZeroMQ Sockets (`zmq.Socket`) are NOT Thread-Safe**
- **Handles message queues and I/O operations.**
- **Modifying the same socket from multiple threads can cause data corruption.**
- **Each thread must create its own socket.**

---

## **3️⃣ Real-Life Analogy**

### **🔹 ZeroMQ Context (Thread-Safe) → The Restaurant's Kitchen**
- The kitchen (context) manages all food orders.
- Multiple waiters (threads) can place orders safely.
- The kitchen only organizes orders but does not serve food directly.

✅ **Why is it thread-safe?**
- The kitchen only **manages** orders, ensuring no conflict between waiters.

### **🔹 ZeroMQ Sockets (Not Thread-Safe) → The Serving Trays**
- Each waiter (thread) needs **their own tray** (socket) to serve food.
- If multiple waiters share a single tray, food may get mixed up or dropped.

❌ **Why is it NOT thread-safe?**
- Multiple people using the same tray at the same time will cause **disasters**.
- Each waiter (thread) should have **their own tray (socket)** to prevent chaos.

---

## **4️⃣ Key Rules for Using ZeroMQ in Multi-Threaded Applications**

| **Rule** | **Explanation** |
|---------|--------------|
| **1. Create One Context Per Application** ✅ | A `zmq.Context` is lightweight and should be shared across all threads. |
| **2. Each Thread Gets Its Own Socket** ✅ | Since sockets are not thread-safe, each thread should create its own socket. |
| **3. Use `inproc://` for Inter-Thread Communication** ✅ | This allows ZeroMQ sockets in different threads to communicate safely. |
| **4. Do NOT Share Sockets Between Threads** ❌ | This causes data corruption and crashes. |

---

## **5️⃣ Example: Process vs. Thread in ZeroMQ**

### **✅ Multi-Process Example (Safe to Share Sockets via IPC)**
```go
package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"os/exec"
	"time"
)

func worker() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	defer socket.Close()
	socket.Bind("ipc:///tmp/zmq-socket")

	for {
		msg, _ := socket.Recv(0)
		fmt.Println("Worker received:", msg)
		socket.Send("Reply from worker", 0)
	}
}

func main() {
	cmd := exec.Command(os.Args[0], "worker")
	cmd.Start()
	time.Sleep(time.Second)

	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	defer socket.Close()
	socket.Connect("ipc:///tmp/zmq-socket")

	socket.Send("Hello Worker!", 0)
	reply, _ := socket.Recv(0)
	fmt.Println("Main process received:", reply)

	cmd.Process.Kill()
}
```
✅ **Works because each process has a separate memory space.**

---

## **ZMQ and Go Routines**

## `zmq.socket` is **NOT** Thread Safe.

Hence Passing the sockets through separate Go routine will cause anomalies. 
Sharing a single socket across goroutines without proper synchronization can lead to race conditions and undefined behavior.

### **BAD CODE**
```go
func myfn(socket *zmq.Socket) {
    // Using socket
}

func main() {
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.PUB) // ANY type of socket

    go myfn(socket) //BAD IDEA
}
```

### Instead, Pass the `zmq.Context` across the Routines, Instantiate and Utilize `sockets` under the same routine, & Synchronize execution using `Channels`.

### **GOOD CODE**
```go
func myfn(context *zmq.Context) // Pass channel variables if sync needed
{
    socket, _ := context.NewSocket(zmq.PUB)
    defer socket.Close()
    // Utilize and end, exclusively for this routine.
}

func main() {
    context, _ := zmq.NewContext()
    defer context.Term()

    go myfn(context)
}
```

