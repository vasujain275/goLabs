# Go Professional Backend Labs 🚀

This repository contains my solutions for the weekly labs and mini-projects focused on mastering Go for professional backend development, based on my internship preparation roadmap.

---

## 🎯 Week 1: Go Fundamentals & Structured Code

**Topics Covered:** Go Modules, variables, control flow, arrays, slices, maps, `range`, structs, pointers, advanced functions (variadic, multiple returns), packages, `PascalCase` vs `camelCase`.

### 💻 Labs & Projects

- [x] **Task: Implement a Word Frequency Counter.**
  - **DoD:** A program takes a string and prints the frequency of each word using a map. The count is case-insensitive and ignores punctuation.
- [x] **Task: Define a `User` Model.**
  - **DoD:** A `User` struct is defined with fields `ID (string)`, `FirstName (string)`, `LastName (string)`, `Email (string)`, and `IsActive (bool)`. A function `NewUser(firstName, lastName, email string) (*User, error)` is created that validates input (e.g., non-empty names, valid email format) and returns a new user pointer or an error.
- [x] **Task: Create a Pointer Receiver Method.**
  - **DoD:** A method `Deactivate()` with a pointer receiver `(*User)` is created. When called, it sets the user's `IsActive` field to `false`. Another method `FullName() string` with a value receiver returns the user's full name.
- [x] **Task: Build a `utils` Package.**
  - **DoD:** A new directory and package named `utils` is created. It contains exported functions for common tasks, like `GenerateUUID()` and `IsValidEmail(email string) bool`.
- [x] **Task: Read "Effective Go" - Data Sections.**
  - **DoD:** You have read and taken notes on the "Data" sections of Effective Go, including slices, maps, and allocation.

---

## 🎯 Week 2: Interfaces, Errors & Web Services with `net/http`

**Topics Covered:** Methods, interfaces, implicit implementation, the `error` type, custom errors, JSON marshalling/unmarshalling, `net/http` package.

### 💻 Labs & Projects

- [x] **Task: Define a `DataStore` Interface.**
  - **DoD:** A `DataStore` interface is defined with methods like `GetUser(id string) (*User, error)` and `CreateUser(user *User) error`.
- [x] **Task: Implement the `DataStore` Interface with a `MemoryStore`.**
  - **DoD:** A `MemoryStore` struct (using a map) is created that correctly implements the `DataStore` interface. The methods handle cases like a user not being found by returning a custom error.
- [x] **Task: Build a Basic HTTP Server.**
  - **DoD:** A server running on port 8080 uses `http.ServeMux` to handle routes. `GET /health` returns a `200 OK` with a JSON body `{"status": "up"}`.
- [x] **Task: Create a JSON API Endpoint.**
  - **DoD:** A `userHandler` struct is created that depends on your `DataStore` interface. It handles `POST /users` by decoding a JSON request body into a `User` struct and using the `DataStore` to save it. It returns a `201 Created` status.
- [ ] **Mini-Project: `go-shortener` - A URL Shortening Web Service**
  - **DoD:** A web service built with only the `net/http` package. `POST /shorten` accepts a JSON body `{"url": "long_url"}` and returns `{"short_code": "xyz123"}`. `GET /{short_code}` redirects the user to the original long URL with a `302 Found` status. Data is stored in-memory using your `MemoryStore` implementation. All endpoints have structured logging.

---

## 🎯 Week 3: Mastering Concurrency

**Topics Covered:** Goroutines, `sync.WaitGroup`, race conditions, `sync.Mutex`, `sync.RWMutex`, channels (buffered/unbuffered), `select` statement, worker pools.

### 💻 Labs & Projects

- [ ] **Task: Concurrent Map Writes (The Wrong Way and The Right Way).**
  - **DoD:** First, write a program where multiple goroutines write to a shared map, demonstrating the race condition with `go run -race`. Second, fix it using `sync.RWMutex`.
- [ ] **Task: Implement a Worker Pool.**
  - **DoD:** A worker pool with 4 goroutines is implemented to process a series of "jobs". Goroutines receive tasks from a `jobs` channel and send results to a `results` channel. The main function waits for all results before exiting.
- [ ] **Task: `select` with Cancellation Context.**
  - **DoD:** A function takes a `context.Context` and a list of URLs. It launches a goroutine for each URL fetch. It uses a `select` statement to return the first successful result while immediately cancelling all other in-flight requests using the context.
- [ ] **Mini-Project: `go-grep` - A Concurrent Log File Searcher**
  - **DoD:** A CLI tool `go-grep <pattern> <file1> [file2...]`. It searches for the pattern in multiple files concurrently. Each file is processed by a separate goroutine. Results are sent over a channel to a single "printer" goroutine to avoid interleaved output. A `sync.WaitGroup` ensures the program waits for all files to be processed.

---

## 🎯 Week 4: Professional Testing, Tooling & Project Structure

**Topics Covered:** `testing` package, unit tests, table-driven tests, mocking with interfaces, benchmarking with `testing.B`, profiling with `pprof`.

### 💻 Labs & Projects

- [ ] **Task: Write a Table-Driven Test.**
  - **DoD:** A table-driven test is written for a `utils.IsValidEmail` function, covering multiple valid and invalid email formats.
- [ ] **Task: Test the `go-shortener` API Handler with a Mock.**
  - **DoD:** A unit test for your `POST /shorten` handler is written. It uses a mock `DataStore` to simulate success and failure cases without needing a real web server. It also uses `net/http/httptest` to check the HTTP status code and response body.
- [ ] **Task: Write a Benchmark Test.**
  - **DoD:** A benchmark test `BenchmarkShortenURL` is created for your core URL shortening logic to measure its performance.
- [ ] **Task: Profile a CPU-intensive function with `pprof`.**
  - **DoD:** You write a function that performs a computationally expensive task (e.g., calculating many Fibonacci numbers). You use `pprof` to generate a CPU profile and can identify the bottleneck.
- [ ] **Mini-Project: `go-chat` - A Simple TCP Chat Server**
  - **DoD:** A TCP server that accepts multiple client connections. Each connection is handled in its own goroutine. Messages from one client are broadcast to all others. Channels are used to manage connections and messages safely. The project is structured logically with `server`, `client`, and `main` packages and has at least 50% test coverage.
