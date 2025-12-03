## Go Architect Progress Tracker

**Role:** You are the User's Principal Mentor.
**Context:** This file is the user's "Ground Truth" of known concepts.
**Rules:**

1. **Personalization:** Scan for `[x]`. Build analogies _only_ using `[x]` concepts.
2. **Taxonomy:**
   - **L1 (Junior):** Syntax & Mechanics ("I can make it work")
   - **L2 (Senior):** Production & Robustness ("I can ship it safely")
   - **L3 (Architect):** Internals & Decisions ("I make the right trade-offs")

---

## Structs, Methods, & Composition

### L1: Define data models and attach logic.

- [x] **Syntax:** I can define a `struct` with typed fields and instantiate it using both struct literals and the `var` keyword.
- [x] **Pointers:** I understand the difference between a value variable (`t`) and a pointer variable (`&t`), and I can dereference a pointer (`*p`) to access the underlying value.
- [x] **Receivers:** I can declare methods on types. I know the syntax difference between a Value Receiver (`func (s MyStruct)`) and a Pointer Receiver (`func (s *MyStruct)`).
- [x] **Visibility:** I can control field and method access (public vs. private) across packages using capitalization.
- [x] **Embedding:** I can use anonymous fields (struct embedding) to "borrow" fields and methods from one struct into another.
- [x] **Constructors:** I can implement "Constructor" functions (e.g., `NewUser()`) to initialize structs with default values, since Go lacks built-in constructors.

---

## Slice vs. Array

### L1: Use collections to store data without crashing.

- [x] **Syntax:** I can declare a fixed-size array (`var a [5]int`) versus a dynamic slice (`var s []int`), and I know that `[5]int` and `[4]int` are completely different types.
- [x] **Semantics:** I understand that assigning an array copies **all** its data (Value Semantics), whereas assigning a slice only copies the "Header" (Pointer Semantics).
- [x] **Manipulation:** I can use `append()` to grow a slice and `len()`/`cap()` to inspect it.
- [x] **Construction:** I can use `make([]T, len, cap)` to initialize a slice with specific dimensions.
- [x] **Slicing:** I can create a "view" into an existing array or slice using the `[low:high]` syntax (e.g., `s[1:4]`).
- [x] **Standard Lib:** I use the `slices` package (Go 1.21+) for common operations (Sorting, Binary Search) instead of writing my own loops.

---

## Maps & Hash Collisions

### L1: Store and retrieve key-value pairs correctly.

- [x] **Syntax:** I can initialize maps using `make(map[K]V)` vs literal syntax, and I know that writing to a `nil` map causes a panic.
- [x] **Access:** I can retrieve values and strictly use the "Comma-Ok" idiom (`val, ok := m[key]`) to distinguish between a "missing key" and a "zero-value" entry.
- [x] **Manipulation:** I can add entries, use the `delete(m, key)` function, and iterate over maps using `range`.
- [x] **Iteration Order:** I understand that map iteration order is **randomized** intentionally by the runtime and I never write code that relies on the order of keys.
- [x] **Key Constraints:** I know that map keys must be **comparable** (types that support `==`), meaning I cannot use Slices, Maps, or Functions as keys.

---

## Interfaces

### L1: Polymorphic code using shared behaviors.

- [x] **Syntax:** I can define an `interface` with method signatures and implement it implicitly by defining those methods on a concrete type.
- [x] **Polymorphism:** I can write a function that accepts an interface type (e.g., `func Speak(s Speaker)`) and pass different concrete structs that satisfy it.
- [x] **Type Assertion:** I can extract the concrete value from an interface using `val, ok := i.(ConcreteType)` and handle the failure case safely.
- [x] **Type Switches:** I can use a `switch v := i.(type)` block to handle multiple possible concrete types stored within an interface variable.
- [x] **The Empty Interface:** I understand `interface{}` (or `any` in Go 1.18+) represents "zero methods," meaning _all_ types satisfy it.

---

## Garbage Collection (GC)

### L1: Memory is managed automatically but not infinite.

- [x] **Concept:** I understand the difference between **Stack** (fast, automatic cleanup per function) and **Heap** (slower, GC-managed).
- [x] **Lifecycle:** I understand that variables created in a function "live" until they are no longer referenced, at which point the GC reclaims them.
- [x] **Leaks:** I understand that "Memory Leaks" in Go are usually logical errors (e.g., forgetting to remove an item from a long-lived Map or Slice) rather than unmanaged memory.

---

## Defer

### L1: Clean up resources safely.

- [x] **Syntax:** I can register a function call to be executed when the surrounding function returns using the `defer` keyword.
- [x] **LIFO Order:** I understand that multiple deferred calls execute in **Last-In, First-Out** (stack) order. (e.g., `defer A; defer B` executes B, then A).
- [x] **Resource Management:** I immediately place `defer resource.Close()` after successfully acquiring a resource (File, Network Connection, Mutex) to prevent leaks.
- [x] **Scope Rule:** I understand that `defer` is bound to the **Function**, not the `{ block }`. Calling `defer` inside an `if` block or `for` loop schedules it for the _end of the function_, not the end of the block.

---

## Error Handling vs. Panics

### L1: Handle errors explicitly; do not crash.

- [x] **Philosophy:** I understand that Go does not have `try/catch`. I accept that handling errors is a core part of the code logic, not an afterthought.
- [x] **Syntax:** I can implement the standard check: `if err != nil { return err }`.
- [x] **Creation:** I can create simple error values using `errors.New("text")` and formatted errors using `fmt.Errorf("code: %d", code)`.
- [x] **Panic:** I understand that `panic()` stops the ordinary flow of control and begins **Stack Unwinding**. I know it is reserved for unrecoverable states (e.g., startup configuration missing).
- [x] **Recover:** I know the syntax to `recover()` from a panic using a `defer` statement.

---

## Type Aliases

### L1: Distinguish New Type vs Renamed Type.

- [x] **Syntax:** I know the syntactic difference between a Type Definition (`type MyInt int`) and a Type Alias (`type MyInt = int`).
- [x] **Identity:** I understand that an Alias is **identical** to the underlying type. I can assign a variable of type `MyInt` (alias) to a variable of type `int` without casting.
- [x] **New Types:** I understand that a Type Definition creates a distinct type. I know I _must_ explicitly cast (`T(val)`) to convert between the new type and the underlying type.
- [x] **Built-ins:** I recognize that `byte` is an alias for `uint8` and `rune` is an alias for `int32` in the standard library.

---

## Generics (Type Parameters)

### L1: Read/write basic generic functions.

- [x] **Syntax:** I can declare a generic function using square brackets: `func Map[T any](s []T, f func(T) T) []T`.
- [x] **Constraints:** I understand `any` (alias for `interface{}`) and `comparable` (types that support `==`). I know I cannot use `==` on a generic type `T` unless it is constrained by `comparable`.
- [x] **Instantiation:** I can call generic functions both with explicit type arguments (`Fn[int](5)`) and by relying on **Type Inference** (`Fn(5)`).
- [x] **Generic Structs:** I can define data structures that hold generic values, such as `type Box[T any] struct { Val T }`.

---

## Testing (Unit & Suites)

### L1: Verify code works using standard lib.

- [x] **Syntax:** I can create a test file ending in `_test.go` and define functions with the signature `func TestName(t *testing.T)`.
- [x] **Assertion Mechanics:** I understand the difference between `t.Error/t.Fail` (continue execution) and `t.Fatal/t.FailNow` (stop current test immediately).
- [x] **Table-Driven Tests:** I can implement the **Table-Driven Test** pattern: defining a slice of structs (inputs/outputs) and iterating over them. I accept this as the _mandatory_ Go testing idiom.
- [x] **Running Tests:** I can use `go test ./...` to run all tests in a module and `go test -v` to see verbose output.
- [x] **Code Coverage:** I can run `go test -cover` to see what percentage of my code is exercised by tests.

---

## Test-Driven Development (TDD)

### L1: Basic test verification.

- [x] **The Cycle:** I can execute the "Red-Green-Refactor" loop: Write a failing test first, write the minimal code to pass it, then clean up.
- [x] **Assertions:** I understand that Go has no built-in `assert` library. I can write manual checks: `if got != want { t.Errorf(...) }`.

---

## Channels & Concurrency Safety

### L1: Transmit data; identify blockers.

- [x] **Syntax:** I can declare, initialize (`make`), and distinguish between **unbuffered** and **buffered** channels.
- [x] **Blocking Mechanics:** I understand that sending to a full channel or receiving from an empty unbuffered channel blocks execution, and I can explain why.
- [x] **Directionality:** I can correctly use the arrow syntax (`<-ch` vs `ch<-`) to send and receive values.
- [x] **Range & Close:** I can iterate over a channel using `for range` and understand that the loop terminates only when the channel is closed.
- [x] **Basic Deadlock:** I can recognize the runtime panic `"fatal error: all goroutines are asleep - deadlock!"` and explain why it happens (circular dependency or no sender/receiver).
- [x] **Race Concept:** I can define a "Race Condition" in my own words: _Two threads accessing the same memory, at least one is a write, and no ordering is enforced._

---

## Goroutines & The Scheduler

### L1: Spawn concurrent tasks and wait.

- [x] **Syntax:** I can start a concurrent execution using the `go` keyword (`go myFunction()`).
- [x] **Lifecycle:** I understand that if the `main()` function returns, the program exits immediately, killing all other running goroutines without warning.
- [x] **Synchronization:** I can use `sync.WaitGroup` (`Add`, `Done`, `Wait`) to ensure the main program waits for goroutines to complete.
- [x] **Closures:** I can wrap logic in anonymous functions (`go func() { ... }()`) to execute quick tasks without defining named functions.
- [x] **No Return:** I understand that a goroutine cannot "return" a value to its caller like a normal function; it must communicate results via Channels or Shared Memory.

---

## Context

### L1: Pass context; set basic timeouts.

- [x] **The First Arg:** I follow the convention of passing `ctx context.Context` as the **first argument** to functions, never embedding it in a struct (with rare exceptions).
- [x] **Roots:** I know when to use `context.Background()` (main, init, tests) vs. `context.TODO()` (when I'm unsure or refactoring).
- [x] **Deadlines:** I can use `context.WithTimeout(parent, duration)` and `context.WithDeadline(parent, time)` to ensure operations don't hang forever.
- [x] **Cancellation:** I can use `context.WithCancel(parent)` to create a `cancel` function, and I understand that calling it stops the context's children.
- [x] **Cleanup:** I strictly adhere to the rule: **Always call the `cancel` function** (often via `defer`) returned by `WithTimeout` or `WithCancel` to release resources, even if the operation succeeds.

---

## Synchronization Primitives (Sync & Atomic)

### L1: Prevent data races.

- [x] **Mutex Syntax:** I can use `sync.Mutex` to protect a shared map or slice. I strictly follow the pattern: `Lock()`, perform operation, `Unlock()`.
- [x] **Read/Write Split:** I understand that `sync.RWMutex` allows multiple readers (`RLock`) but only one writer (`Lock`), and I use it for read-heavy workloads.
- [x] **Atomic Counters:** I can use `atomic.AddInt64` and `atomic.LoadInt64` to maintain thread-safe counters without the overhead of a full Mutex.
- [x] **Race Detector:** I verify my synchronization using `go test -race` and treat any output as a critical bug.
