# Go Concurrency Learning & Problem Solving Patterns

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Topic](https://img.shields.io/badge/Topic-Concurrency_&_Parallelism-FF6F00?style=for-the-badge)](https://golang.org/doc/effective_go#concurrency)
[![Status](https://img.shields.io/badge/Status-Fully_Tested_&_Passing-4CAF50?style=for-the-badge)](https://golang.org/doc/code#Testing)

Welcome to the **Go Concurrency Learning Project**! This repository is a hands-on, practical guide designed to help you master concurrency in Go. It contains real-life problem-solving patterns, step-by-step challenges, and production-safe strategies for managing asynchronous execution.

---

## 🎯 Project Objectives & Purpose

1. **Practical Concurrency Patterns**: Real-world simulations (like a pizzeria order pipeline and a multi-threaded bank ledger) that mirror scenarios found in production services.
2. **Safe & Defensive Programming**: Direct guidance on avoiding the most common concurrent programming pitfalls—such as race conditions, goroutine leaks, deadlocks, and premature application termination.
3. **Elevate Your Go Skills**: Structured exercises to transition from sequential thinking to designing concurrent, message-driven applications using Go's native Communicating Sequential Processes (CSP) primitives.

---

## 📂 Project Structure & Deep Dive

The workspace is organized into five primary learning modules, each focusing on specific Go concurrency primitives.

### 1. [go routines](file:///Users/koustavdas/projects/go-concurrency/go%20routines)
* **Goal**: Coordinate asynchronous tasks and prevent early application termination.
* **Core Primitives**: `sync.WaitGroup`
* **Real-World Analogy**: Running multiple asynchronous jobs that must all complete before producing a final report or exiting.
* **Key Files**:
  - [main.go](file:///Users/koustavdas/projects/go-concurrency/go%20routines/main.go): Implements goroutines that concurrently modify and update message buffers, coordinated safely using a WaitGroup.
  - [main_test.go](file:///Users/koustavdas/projects/go-concurrency/go%20routines/main_test.go): Comprehensive unit tests verifying concurrent message updating, output printing, and the execution flow of the main function.

### 2. [chans, mutex](file:///Users/koustavdas/projects/go-concurrency/chans,%20mutex)
* **Goal**: Implement thread-safe communication and protect shared state.
* **Core Primitives**: `chan`, `select`, `sync.Mutex`
* **Real-World Scenarios**:
  - **The Pizzeria Simulator (Producer-Consumer Pattern)**:
    - Leverages channels (`chan`) to pass pizza orders asynchronously from a worker pipeline to a receiver.
    - Demonstrates safe channel closure and clean-up using a dedicated quit channel structure to prevent goroutine leaks.
  - **The Bank Balancer (Shared Memory Protection)**:
    - Simulates multiple concurrent income streams updating a single account ledger variable.
    - Solves critical race conditions by locking/unlocking the critical section using a mutual exclusion lock (`sync.Mutex`).
* **Key Files**:
  - [main.go](file:///Users/koustavdas/projects/go-concurrency/chans,%20mutex/main.go): Implements the complete pizzeria loop and the mutex-protected bank ledger.
  - [main_test.go](file:///Users/koustavdas/projects/go-concurrency/chans,%20mutex/main_test.go): Test suite that captures standard output pipes to assert execution correctness, validating bank operations.

### 3. [channel-select](file:///Users/koustavdas/projects/go-concurrency/channel-select)
* **Goal**: Understand channel buffering, directional channels, and multi-channel selection.
* **Core Primitives**: `chan`, `select`, directional channels (`<-chan`, `chan<-`)
* **Real-World Scenarios**:
  - **Buffered Channels**: Demonstrates how a buffered channel allows the sender to continue sending messages without blocking until the buffer capacity is reached.
  - **Ping-Pong Communication**: Demonstrates how two goroutines can coordinate and exchange data using directional channels (`<-chan` and `chan<-`) for shout processing.
  - **Multi-Server Select**: Simulates receiving responses from multiple asynchronous servers running at different speeds, handling messages dynamically using the `select` statement.
* **Key Files**:
  - [buffered.go](file:///Users/koustavdas/projects/go-concurrency/channel-select/buffered.go): Showcases sender blocking behavior on full vs empty buffered channels via [buffered](file:///Users/koustavdas/projects/go-concurrency/channel-select/buffered.go#L18).
  - [pingPong.go](file:///Users/koustavdas/projects/go-concurrency/channel-select/pingPong.go): Shows type-safe directional channels communicating interactively via [PingPong](file:///Users/koustavdas/projects/go-concurrency/channel-select/pingPong.go#L18).
  - [selectCh.go](file:///Users/koustavdas/projects/go-concurrency/channel-select/selectCh.go): Uses `select` to handle asynchronous channel events dynamically via [selectCh](file:///Users/koustavdas/projects/go-concurrency/channel-select/selectCh.go#L25).
  - [main.go](file:///Users/koustavdas/projects/go-concurrency/channel-select/main.go): Entry point running the buffered channel simulation.

### 4. [dining-philosophers](file:///Users/koustavdas/projects/go-concurrency/dining-philosophers)
* **Goal**: Prevent resource contention and avoid deadlocks in a shared-resource environment.
* **Core Primitives**: `sync.Mutex`, `sync.WaitGroup`
* **Real-World Analogy**: The classic Dining Philosophers problem where five [Philosopher](file:///Users/koustavdas/projects/go-concurrency/dining-philosophers/dining.go#L10) entities must share five forks. Demonstrates deadlock prevention by ordering how resources (mutexes/forks) are acquired (locking the lower-indexed fork first).
* **Key Files**:
  - [dining.go](file:///Users/koustavdas/projects/go-concurrency/dining-philosophers/dining.go): Implements the deadlock-free dining loop using [dine](file:///Users/koustavdas/projects/go-concurrency/dining-philosophers/dining.go#L43) and [diningPhilosophers](file:///Users/koustavdas/projects/go-concurrency/dining-philosophers/dining.go#L63).
  - [dining_test.go](file:///Users/koustavdas/projects/go-concurrency/dining-philosophers/dining_test.go): Verifies that all philosophers successfully seat, eat, and leave without deadlock.

### 5. [sleeping-barber](file:///Users/koustavdas/projects/go-concurrency/sleeping-barber)
* **Goal**: Implement complex resource scheduling and synchronization with dynamic participant arrival and departure.
* **Core Primitives**: `chan`, `select`, `time.After`
* **Real-World Analogy**: The Sleeping Barber problem simulating a [BarberShop](file:///Users/koustavdas/projects/go-concurrency/sleeping-barber/barberShop.go#L9) with a limited waiting room. Barbers sleep when no clients are present and are awoken when a client arrives. Shows how to handle shop closing, stopping client generation, and sending barbers home safely.
* **Key Files**:
  - [barber.go](file:///Users/koustavdas/projects/go-concurrency/sleeping-barber/barber.go): Entry point that initializes the shop, manages opening/closing time, and spawns client arrivals concurrently.
  - [barberShop.go](file:///Users/koustavdas/projects/go-concurrency/sleeping-barber/barberShop.go): Implements the barber behavior, including waiting room capacity checks (via non-blocking select channel writes) and barber napping/waking logic.

---

## 🛠️ How to Run & Test

Navigate to each module directory to run the implementations and their tests.

### Running Module 1 (Goroutines & WaitGroups)
```bash
# Navigate to directory
cd "go routines"

# Run the program
go run main.go

# Run tests
go test -v
```

### Running Module 2 (Channels & Mutexes)
```bash
# Navigate to directory
cd "chans, mutex"

# Run the program
go run main.go

# Run tests
go test -v
```

### Running Module 3 (Channel Select)
```bash
# Navigate to directory
cd "channel-select"

# Run the program
go run .
```

### Running Module 4 (Dining Philosophers)
```bash
# Navigate to directory
cd "dining-philosophers"

# Run the program
go run dining.go

# Run tests
go test -v
```

### Running Module 5 (Sleeping Barber)
```bash
# Navigate to directory
cd "sleeping-barber"

# Run the program
go run .
```

### 🔍 Crucial: Running with the Race Detector
Go has a built-in race detector that spots concurrent memory access issues. Run your tests with the `-race` flag to guarantee your code is free of race conditions:
```bash
go test -race -v ./...
```

---

## 💡 Key Concurrency Concepts Explained

### 🚀 Goroutines (`go func()`)
Goroutines are multiplexed onto a small number of OS threads. They are incredibly lightweight, costing only a few kilobytes of memory to start, allowing you to run hundreds of thousands of them concurrently.

### 🛑 WaitGroups (`sync.WaitGroup`)
When you spawn a goroutine, the main application loop will not wait for it to complete. A `WaitGroup` maintains a counter:
- `Add(n)` increases the counter.
- `Done()` (usually called via `defer`) decreases the counter.
- `Wait()` blocks until the counter reaches zero.

### 🛡️ Mutexes (`sync.Mutex`)
When multiple goroutines try to read and write to the same memory address concurrently, a **data race** occurs, leading to unpredictable crashes and state corruption. A `Mutex` ensures only one goroutine can execute a critical section at a time:
```go
balance.Lock()
// Critical Section: Only one goroutine can execute this at a time
bankBalance += income.amount
balance.Unlock()
```

### 📣 Channels (`chan`)
Channels are typed conduits that allow goroutines to communicate and synchronize without explicit locks.
- **Unbuffered Channels**: Block the sender until the receiver is ready, ensuring synchronization.
- **Buffered Channels**: Allow senders to queue messages without blocking up to the buffer capacity.
- **Directional Channels**: Limit channels to receive-only (`<-chan`) or send-only (`chan<-`) behavior to enforce compile-time safety and cleaner API design.
- **Select Statement**: Allows a goroutine to wait on multiple communication operations (e.g., waiting for new data or a quit signal simultaneously) or implement non-blocking attempts (using a `default` case).

---

## 📜 Go Concurrency Golden Rules

> [!IMPORTANT]
> **Share Memory by Communicating, Not Communicating by Sharing Memory**
> Instead of using locks to coordinate access to shared state, prefer passing ownership of data through channels.

> [!WARNING]
> **Avoid Goroutine Leaks**
> Never start a goroutine unless you know how and when it will terminate. If a goroutine is blocked forever trying to write to a channel that no one is reading, it leaks memory. Use quit channels or a `context.Context` to signal termination.

> [!TIP]
> **Keep Critical Sections Small**
> When using `sync.Mutex`, keep the locked section of code as brief as possible. Holding a lock while performing heavy I/O or network requests will completely destroy concurrency benefits.

