# Go Learning Labs 🚀

```
This repository contains my solutions for the weekly labs as I learn the Go programming language.

## Repository Structure

The repository is organized by week, with each lab task in its own sub-directory.

    goLabs/
    ├── README.md
    ├── week1/
    │   ├── 01-fizzbuzz/
    │   └── ...
    ├── week2/
    │   └── ...
    └── ...

---

## 🎯 Week 1: Go Fundamentals & Data Structures

**Topics Covered:** Go Modules, variables, basic types, control flow, arrays, slices, maps, `range` keyword.

### 💻 Labs

- [ ] **FizzBuzz Program:** A Go program that correctly prints numbers 1-100, replacing multiples of 3 with "Fizz", 5 with "Buzz", and both with "FizzBuzz".
- [ ] **Find the Maximum Value in a Slice:** A function `FindMax(numbers []int) int` that correctly returns the largest integer from a given slice.
- [ ] **Implement a Simple Stack:** A `Stack` type implemented using a slice with `Push(item int)`, `Pop() (int, error)`, and `Peek() (int, error)` methods. Popping or peeking from an empty stack returns an error.
- [ ] **Word Frequency Counter:** A program that takes a string and prints the frequency of each word using a map (case-insensitive).

---

## 🎯 Week 2: Structuring Code & Data

**Topics Covered:** Structs, pointers, advanced functions (variadic, multiple returns), packages, `PascalCase` vs `camelCase`.

### 💻 Labs

- [ ] **Define a `User` Struct:** Define a `User` struct with fields `ID (string)`, `FirstName (string)`, `LastName (string)`, and `IsActive (bool)`.
- [ ] **Create a Pointer Receiver Method:** Create a method `Deactivate()` with a pointer receiver `(*User)` that sets the user's `IsActive` field to `false`.
- [ ] **Build a `calculator` Package:** Create a new directory and package named `calculator` containing exported functions `Add`, `Subtract`, `Multiply`, and `Divide`.
- [ ] **Use the Custom Package:** The `main` package successfully imports the `calculator` package and uses its functions to perform and print calculations.

---

## 🎯 Week 3: Interfaces, Methods & Idiomatic Go

**Topics Covered:** Methods, interfaces, implicit implementation, the `error` type, JSON marshalling/unmarshalling.

### 💻 Labs

- [ ] **Define a `Shape` Interface:** Define a `Shape` interface with one method: `Area() float64`.
- [ ] **Implement the `Shape` Interface:** Create `Circle` and `Rectangle` structs, each with methods that correctly implement the `Shape` interface.
- [ ] **Demonstrate Polymorphism:** A function `PrintTotalArea(shapes ...Shape)` calculates and prints the total area of any combination of `Circle` and `Rectangle` shapes passed to it.
- [ ] **Implement Error Handling:** Modify the `calculator.Divide` function to return `(float64, error)`, returning a non-nil error for division by zero.
- [ ] **Mini-Project 1: CLI Todo List Manager:** A program that runs from the CLI (`add`, `list`, `complete`), uses a `Task` struct, and persists the list to `tasks.json` on every change.

---

## 🎯 Week 4: Concurrency Fundamentals

**Topics Covered:** Goroutines, `sync.WaitGroup`, race conditions, `sync.Mutex`, `sync/atomic`.

### 💻 Labs

- [ ] **Concurrent URL Downloader:** A program takes a list of URLs, downloads their contents concurrently using goroutines, and uses a `WaitGroup` to wait for all downloads to finish.
- [ ] **Demonstrate and Fix a Race Condition:** Write a program where 1000 goroutines increment a shared counter. First, show the incorrect final value (the race condition), then fix it using a `sync.Mutex` so it correctly prints `1000`.

---

## 🎯 Week 5: Advanced Concurrency & Communication

**Topics Covered:** Channels (buffered/unbuffered), `select` statement, `range` over channels, worker pools.

### 💻 Labs

- [ ] **Producer-Consumer with a Buffered Channel:** One goroutine (producer) sends 10 numbers to a buffered channel. A second goroutine (consumer) reads and prints them without deadlocking.
- [ ] **`select` with a Timeout:** A goroutine sends a message to a channel after a 3-second delay. The `main` goroutine's `select` statement waits for the message but will time out after 2 seconds.
- [ ] **Implement a Worker Pool:** Implement a worker pool with 3 goroutines. Create a "jobs" channel to send 10 tasks, which the workers receive and process.

---

## 🎯 Week 6: Testing & Project Structure

**Topics Covered:** `testing` package, unit tests, table-driven tests, mocking with interfaces, benchmarking.

### 💻 Labs

- [ ] **Write a Table-Driven Test:** Write a table-driven test `TestDivide` for the `calculator.Divide` function. Include test cases for a valid division and division by zero.
- [ ] **Test Logic with a Mock:** Write a business logic function that uses a `DataStore` interface. Write a test that injects a mock `DataStore` to verify the logic without any real I/O.
```
