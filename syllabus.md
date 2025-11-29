# GO ARCHITECT SYLLABUS & PROGRESS TRACKER

**Objective:** Zero-to-Mastery Roadmap
**Format:** Interactive Knowledge Tracking

---

## 🧠 Knowledge Taxonomy (The Standard)

This syllabus requires mastery at three distinct levels of abstraction.

| Level  | Role Equivalent           | Goal                             | Criteria for Mastery                                                                                                                                                                                     |
| :----- | :------------------------ | :------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **L1** | **Junior / Foundation**   | _"I can make it work."_          | Focus on **Syntax & Mechanics**. You can write code that compiles and performs basic operations using standard keywords.                                                                                 |
| **L2** | **Senior / Professional** | _"I can ship it safely."_        | Focus on **Production & Robustness**. You use idiomatic patterns, handle errors gracefully, ensure concurrency safety, and write testable code.                                                          |
| **L3** | **Principal / Architect** | _"I make the right trade-offs."_ | Focus on **Internals & Decisions**. You understand the Runtime (GC, Scheduler, Memory) deeply enough to justify architectural trade-offs (e.g., Latency vs. Throughput) for specific system constraints. |

---

## 🤖 INSTRUCTIONS FOR AI ASSISTANTS

**You are the User's Principal Mentor.** This file represents the user's **Ground Truth** of known concepts.

**1. Personalize based on `[x]` (The Knowledge Boundary):**

- **Scan the file:** Before answering, identify which concepts are marked `[x]` (Known) and `[ ]` (Unknown).
- **Analogy Strategy:** When explaining a new concept, build analogies _only_ using concepts the user has already mastered (`[x]`).
- **Safety Guard:** Do **not** use advanced jargon or concepts from unchecked `[ ]` items to explain basic topics. This causes confusion.

**2. The Audit Standard (L3):**

- If the user claims mastery of an **L3** topic, do not accept surface-level answers.
- **Challenge them:** Ask for the _trade-off_. (e.g., "Why use a value receiver here if it increases stack usage? How does that impact the GC?").
- The goal of L3 is to verify **Decision Making Capabilities**, not just trivia recall.

---

# Syllabus

## Structs, Methods, & Composition

_“Go is not a classic OO language; we do not build hierarchies, we compose behaviors. We trade the rigid comfort of inheritance for the flexible reality of composition.”_

This module covers the "Object-Oriented" features of Go: how to define data (**Structs**), how to refer to it (**Pointers**), how to give it behavior (**Receivers**), and how to build complex systems from simple parts (**Composition**).

---

### L1: The Foundation (Junior)

_Goal: I can define data models and attach logic to them._

- [ ] **[Syntax]:** I can define a `struct` with typed fields and instantiate it using both struct literals and the `var` keyword.
- [ ] **[Pointers]:** I understand the difference between a value variable (`t`) and a pointer variable (`&t`), and I can dereference a pointer (`*p`) to access the underlying value.
- [ ] **[Receivers]:** I can declare methods on types. I know the syntax difference between a Value Receiver (`func (s MyStruct)`) and a Pointer Receiver (`func (s *MyStruct)`).
- [ ] **[Visibility]:** I can control field and method access (public vs. private) across packages using capitalization.
- [ ] **[Embedding]:** I can use anonymous fields (struct embedding) to "borrow" fields and methods from one struct into another.
- [ ] **[Constructors]:** I can implement "Constructor" functions (e.g., `NewUser()`) to initialize structs with default values, since Go lacks built-in constructors.

### L2: The Professional (Senior)

_Goal: I write idiomatic code that is concurrency-safe and API-stable._

- [ ] **[Semantics]:** I can decide between Value Receivers and Pointer Receivers based on semantics (Mutation vs. Immutability) rather than just "performance."
- [ ] **[Nil Safety]:** I can prevent `nil` pointer dereference panics by using idiomatic guard clauses and meaningful zero-values.
- [ ] **[Interface Compliance]:** I understand how Receiver types affect Interface satisfaction (e.g., a pointer receiver method means the _value_ type does not implement the interface).
- [ ] **[Tags]:** I can utilize Struct Tags (e.g., `` `json:"name"` ``) to control serialization and understand how reflection utilizes them.
- [ ] **[Composition vs. Inheritance]:** I avoid deep embedding chains. I can refactor "Is-A" relationships (Inheritance thinking) into "Has-A" relationships (Composition thinking) to prevent fragile base class problems.
- [ ] **[Promoted Methods]:** I can handle namespace collisions when multiple embedded structs have methods with the same name, and I understand how method promotion works.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I optimize memory layout and GC pressure; I design for cache locality._

- [ ] **[Memory Layout]:** I can reorder struct fields to minimize memory footprint by optimizing for **padding and alignment** (reducing wasted bytes between fields).
- [ ] **[GC Pressure Strategy]:** I can recognize when to use Value Semantics (passing copies) instead of Pointers to reduce the **Garbage Collector's scanning set**, justifying the CPU cost of copying against the latency savings in GC pauses.
- [ ] **[Escape Analysis]:** I can predict when a variable will "escape to the heap" versus stay on the stack. I use this to design APIs that minimize heap allocations (e.g., avoiding returning pointers to local variables in hot paths).
- [ ] **[Cache Locality]:** I can choose between a "Slice of Structs" vs. a "Struct of Slices" (SoA vs. AoS) based on access patterns to maximize CPU cache hits.
- [ ] **[API Stability]:** I can decide when to Embed a type in a public API versus using a Named Field. I justify this based on the risk of leaking internal implementation details (Promoted Methods) into the public interface.
- [ ] **[False Sharing]:** I can identify and mitigate "False Sharing" in concurrent structs by using padding to ensure frequently modified atomic counters reside on different cache lines.

---

## Slice vs. Array

_“Arrays are the hardware reality (fixed, contiguous, value types). Slices are the software abstraction (dynamic, flexible, reference types). Mastery lies in knowing when to strip away the abstraction.”_

This module distinguishes between the rigid, stack-friendly **Array** and the dynamic, heap-backed **Slice**. While 95% of Go code uses slices, the "Elite" 5% knows that slices are just syntactic sugar over arrays—and sometimes, that sugar is too expensive.

---

### L1: The Foundation (Junior)

_Goal: I can use collections to store data without crashing._

- [ ] **[Syntax]:** I can declare a fixed-size array (`var a [5]int`) versus a dynamic slice (`var s []int`), and I know that `[5]int` and `[4]int` are completely different types.
- [ ] **[Semantics]:** I understand that assigning an array copies **all** its data (Value Semantics), whereas assigning a slice only copies the "Header" (Pointer Semantics).
- [ ] **[Manipulation]:** I can use `append()` to grow a slice and `len()`/`cap()` to inspect it.
- [ ] **[Construction]:** I can use `make([]T, len, cap)` to initialize a slice with specific dimensions.
- [ ] **[Slicing]:** I can create a "view" into an existing array or slice using the `[low:high]` syntax (e.g., `s[1:4]`).

### L2: The Professional (Senior)

_Goal: I optimize for allocation efficiency and prevent memory leaks._

- [ ] **[Pre-allocation]:** I can prevent performance-killing reallocation loops by pre-calculating capacity: `make([]T, 0, expectedSize)`.
- [ ] **[Growth Algorithm]:** I understand that `append` doubles capacity for small slices (up to ~256 elements) but grows by ~1.25x for larger ones, and I can explain how this affects memory spikes.
- [ ] **[Memory Leak Prevention]:** I can identify the "Zombie Array" problem: keeping a tiny slice of a massive backing array keeps the _entire_ array in memory. I fix this using `copy()` to a fresh, smaller slice.
- [ ] **[Copying]:** I use the built-in `copy(dst, src)` function correctly, knowing it only copies the minimum of `len(dst)` and `len(src)`.
- [ ] **[Variadic Expansion]:** I can explode a slice into arguments using the ellipsis operator (`func(s...)`).

### L3: The Architect (Mastery & Trade-offs)

_Goal: I manipulate memory layout for SIMD, Cache Locality, and Zero-Allocation._

- [ ] **[Internals]:** I can visualize the **Slice Header** (`struct { ptr, len, cap }`) and explain why passing a slice to a function allows mutation of elements but _not_ the slice header itself (e.g., `append` inside a function doesn't update the caller's slice length).
- [ ] **[Bounds Check Elimination]:** I can write code that proves safety to the compiler (e.g., `_ = b[3]` before accessing indices 0-3), forcing it to remove expensive runtime bounds checks.
- [ ] **[Stack Allocation Strategy]:** I choose small Arrays over Slices in hot paths to force **Stack Allocation**, avoiding GC pressure entirely, whereas Slices almost always force their backing array to the Heap.
- [ ] **[Zero-Copy Casts]:** I can safely cast a Slice to an Array Pointer using Go 1.17+ syntax (`(*[N]T)(s)`) to regain array performance characteristics (fixed size invariants) without copying memory.
- [ ] **[Cache Line Optimization]:** I can use Arrays with padding (e.g., `[64]byte`) to prevent **False Sharing** between goroutines modifying adjacent data, whereas slices make alignment guarantees difficult.
- [ ] **[SIMD & Vectorization]:** I understand that the compiler is more likely to auto-vectorize loops over Arrays (static bounds) than Slices (dynamic bounds), and I choose Arrays for heavy number-crunching kernels.

---

## Maps & Hash Collisions

_“A map is a hash table, not a magic box. It trades memory for O(1) access. The Architect knows when that trade is too expensive and when the hash function itself becomes an attack vector.”_

This module covers the **Hash Table** implementation in Go. While L1/L2 focus on usage, L3 dives into the `hmap` struct, bucket mechanics, and how Go mitigates Hash Flooding attacks (Collisions).

---

### L1: The Foundation (Junior)

_Goal: I can store and retrieve key-value pairs correctly._

- [ ] **[Syntax]:** I can initialize maps using `make(map[K]V)` vs literal syntax, and I know that writing to a `nil` map causes a panic.
- [ ] **[Access]:** I can retrieve values and strictly use the "Comma-Ok" idiom (`val, ok := m[key]`) to distinguish between a "missing key" and a "zero-value" entry.
- [ ] **[Manipulation]:** I can add entries, use the `delete(m, key)` function, and iterate over maps using `range`.
- [ ] **[Iteration Order]:** I understand that map iteration order is **randomized** intentionally by the runtime and I never write code that relies on the order of keys.
- [ ] **[Key Constraints]:** I know that map keys must be **comparable** (types that support `==`), meaning I cannot use Slices, Maps, or Functions as keys.

### L2: The Professional (Senior)

_Goal: I write concurrency-safe code and optimize for common access patterns._

- [ ] **[Concurrency Safety]:** I know that standard maps are **not thread-safe** for concurrent read/write. I can prevent "concurrent map writes" panics using `sync.RWMutex`.
- [ ] **[Sets]:** I can implement a "Set" data structure using `map[KeyType]struct{}` to minimize memory usage (0-byte values).
- [ ] **[Pre-sizing]:** I can optimize performance by providing a capacity hint `make(map[K]V, size)` to prevent expensive runtime resizing and rehashing during initialization.
- [ ] **[Deep Equality]:** I can handle complex keys (like structs) correctly, ensuring that fields used for hashing don't change after insertion.
- [ ] **[Addressability]:** I understand that map elements are **not addressable**. I know I cannot get a pointer to a map value (`&m[k]`) because the map might grow and move memory, invalidating the pointer.
- [ ] **[Clearing]:** I can efficiently clear a map using the compiler-optimized idiom (Go 1.21+ `clear(m)` or looping delete) without re-allocating.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I understand buckets, evacuation, and security risks._

- [ ] **[Bucket Topology]:** I can explain the `hmap` internals: how data is stored in **Buckets** (holds 8 keys/values) and how **Overflow Buckets** are linked when collisions occur. I understand that Go uses **Chaining** (via overflow buckets) rather than Open Addressing.
- [ ] **[TopHash]:** I understand how Go uses the "Top Byte" (8 bits) of the hash to quickly select a bucket or check for presence before doing a full key comparison (expensive).
- [ ] **[Evacuation & Growth]:** I can explain "Incremental Evacuation." When a map grows (Load Factor > 6.5), it doesn't copy everything at once (latency spike). It moves buckets gradually as we access them.
- [ ] **[Hash Flooding (DoS)]:** I can explain **Algorithmic Complexity Attacks**. I know that Go seeds the hash function randomly at startup to prevent attackers from predicting hash collisions and forcing $O(1)$ lookups into $O(N)$ CPU spikes.
- [ ] **[Memory Layout]:** I can explain why `map[int]int` puts keys and values together in memory ( `K K K ... V V V` ) inside the bucket to eliminate padding, rather than alternating `K V K V`.
- [ ] **[Small N Optimization]:** I can identify when _not_ to use a map. For very small datasets ($N < 20$), I can justify using a Slice and linear search because CPU cache locality makes it faster than the hashing overhead.

---

## Interfaces

_“Interfaces are the spine of Go software design. In other languages, interfaces are contracts you sign upfront. In Go, they are discoveries you make after the fact. We define behavior, not hierarchy.”_

This module covers Go's implicit interface system. It moves from basic polymorphism to the internal machinery of dynamic dispatch, ensuring you know not just _how_ to abstract, but _what it costs_ to abstract.

---

### L1: The Foundation (Junior)

_Goal: I can write polymorphic code using shared behaviors._

- [ ] **[Syntax]:** I can define an `interface` with method signatures and implement it implicitly by defining those methods on a concrete type (no `implements` keyword).
- [ ] **[Polymorphism]:** I can write a function that accepts an interface type (e.g., `func Speak(s Speaker)`) and pass different concrete structs that satisfy it.
- [ ] **[Type Assertion]:** I can extract the concrete value from an interface using `val, ok := i.(ConcreteType)` and handle the failure case safely.
- [ ] **[Type Switches]:** I can use a `switch v := i.(type)` block to handle multiple possible concrete types stored within an interface variable.
- [ ] **[The Empty Interface]:** I understand `interface{}` (or `any` in Go 1.18+) represents "zero methods," meaning _all_ types satisfy it.

### L2: The Professional (Senior)

_Goal: I design decoupled systems and avoid "Interface Pollution."_

- [ ] **[Golden Rule]:** I practice the Proverb: **"Accept Interfaces, Return Structs."** I design functions to ask for the minimum behavior they need, but return concrete types to give callers flexibility.
- [ ] **[Composition]:** I can compose larger interfaces from smaller ones (e.g., `type ReadWriter interface { Reader; Writer }`) rather than creating monolithic "God Interfaces."
- [ ] **[Standard Lib Mastery]:** I deeply understand the ubiquity of `io.Reader` and `io.Writer`. I can make my custom types compatible with the entire ecosystem simply by implementing `Read()` or `Write()`.
- [ ] **[Mocking]:** I define interfaces _at the point of use_ (in the consumer package) to make testing and mocking trivial, rather than defining them in the producer package (Java style).
- [ ] **[Nil Ambiguity]:** I can explain the "Interface Nil" trap: I know that an interface containing a `nil` concrete pointer is **not** equal to `nil`, and I can write code to prevent this bug.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I understand the runtime cost of abstraction and optimizing dispatch tables._

- [ ] **[Runtime Representation]:** I can visualize the 2-word structure of an interface value: `(tab, data)`.
  - `tab`: Pointer to the **itab** (dispatch table + type info).
  - `data`: Pointer to the concrete data.
- [ ] **[Dispatch Mechanics]:** I understand **Dynamic Dispatch**. I know that calling a method via an interface requires following the `tab` pointer to find the function address, which prevents the compiler from **Inlining** the function call.
- [ ] **[Cost Analysis]:** I can decide _against_ using an interface in a "hot loop" because the indirection and lack of inlining incurs a CPU penalty (branch prediction misses) that outweighs the abstraction benefit.
- [ ] **[Memory Allocations]:** I understand that assigning a concrete value to an interface often causes the value to **escape to the heap** (allocation) because the interface needs a pointer to the data, and that pointer must survive the function scope.
- [ ] **[Sealed Interfaces]:** I can create "Sealed Interfaces" (interfaces that cannot be implemented outside my package) by adding an unexported method (e.g., `_private()`), allowing me to restrict implementation while exposing polymorphism.
- [ ] **[Decoupling vs. Locality]:** I can argue _when_ to decouple. I resist "premature interface abstraction" (Interface Pollution). If there is only one implementation, I use the concrete type to keep code navigation simple (Code Locality).

---

## Garbage Collection (GC)

_“Go’s GC is a non-generational, concurrent, tri-color mark-and-sweep collector. It is opinionated: it sacrifices **Throughput** (CPU cycles) to minimize **Latency** (Stop-the-World pauses). An Architect does not fight the GC; they design around it.”_

This module moves from "It cleans up my mess" to minimizing the "GC Tax" that eats your CPU credits.

---

### L1: The Foundation (Junior)

_Goal: I understand that memory is managed automatically but is not infinite._

- [ ] **[Concept]:** I understand the difference between **Stack** (fast, automatic cleanup per function) and **Heap** (slower, GC-managed).
- [ ] **[Lifecycle]:** I understand that variables created in a function "live" until they are no longer referenced, at which point the GC reclaims them.
- [ ] **[Leaks]:** I understand that "Memory Leaks" in Go are usually logical errors (e.g., forgetting to remove an item from a long-lived Map or Slice) rather than unmanaged memory.
- [ ] **[Basic Tooling]:** I can run `go test -bench . -benchmem` to see the number of allocations per operation (`allocs/op`).

### L2: The Professional (Senior)

_Goal: I optimize my code to reduce GC pressure and monitor its health._

- [ ] **[GOGC Tuning]:** I can configure the `GOGC` environment variable (default 100) to trade memory usage for CPU usage (e.g., raising it to 200 reduces GC frequency but doubles heap size).
- [ ] **[Observability]:** I can use `GODEBUG=gctrace=1` to analyze GC frequency and pause times in logs.
- [ ] **[Profiling]:** I can use `go tool pprof` (specifically `alloc_space` and `alloc_objects`) to identify hot-spots generating garbage.
- [ ] **[Allocation Patterns]:** I use **Sync.Pool** to reuse complex objects and buffers, preventing them from becoming garbage in the first place.
- [ ] **[Sizing]:** I prefer strict pre-allocation (slices/maps) to avoid intermediate allocations caused by resizing.
- [ ] **[Finalizers]:** I avoid `runtime.SetFinalizer` because I know they are unpredictable, delay memory reclamation, and can resurrect objects.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I understand the internal algorithms (Pacer, Barriers) to predict system behavior under load._

- [ ] **[Tri-Color Abstraction]:** I can explain the **Tri-Color Mark & Sweep** algorithm (White/Grey/Black sets) and how the "Mark Phase" runs concurrently with my application code.
- [ ] **[Write Barriers]:** I understand that Go uses a **Hybrid Write Barrier** (Dijkstra + Yuasa) to maintain data integrity during concurrent marking. I accept that pointer writes in the heap incur a tiny CPU overhead because of this barrier.
- [ ] **[Mark Assist]:** I can identify "Mark Assist" scenarios where the GC cannot keep up with the allocation rate, forcing _my_ application goroutines to pause and do GC work, killing latency.
- [ ] **[GOMEMLIMIT]:** I can utilize `GOMEMLIMIT` (Go 1.19+) to set a soft memory cap, preventing Out-Of-Memory (OOM) kills in containerized environments (Kubernetes) without manually calculating `GOGC` ratios.
- [ ] **[Scan Cost]:** I design structs to minimize **Pointer Chasing**. I know that a `map[string]*Struct` creates significantly more work for the scanner (checking every pointer) than a `map[int]Struct` (flat data, no pointers), and I make this trade-off for high-throughput systems.
- [ ] **[Generational Hypothesis]:** I can explain that Go is **NOT** Generational (unlike Java/Python). This implies that "short-lived objects" on the Heap are expensive in Go (they must be marked/swept), whereas in Java they are cheap. Therefore, I strive to keep short-lived variables on the **Stack**.

---

## SOLID Principles in Go

_“SOLID is not a religion; it is a vocabulary for decoupling. In Go, we apply these principles through Composition and Interfaces, not Inheritance hierarchies. The goal is maintainability, not abstraction for abstraction's sake.”_

This module translates the classic OOP principles into idiomatic Go. A Junior memorizes the acronym; an Architect knows that **Go's type system flips standard SOLID advice on its head** (specifically regarding where interfaces are defined).

---

### L1: The Foundation (Junior)

_Goal: I can define the principles and identify them in code._

- [ ] **[SRP (Single Responsibility)]:** I can refactor a "God Struct" or "God Function" into smaller, focused units. I ensure a Type handles one clear domain concept (e.g., separating `UserDB` logic from `UserJSON` formatting).
- [ ] **[OCP (Open/Closed)]:** I can write code that is "Open for extension, Closed for modification" by using Interfaces. I can add a new behavior (e.g., a new `PaymentMethod`) without changing the `ProcessPayment` function.
- [ ] **[LSP (Liskov Substitution)]:** I understand that any concrete type implementing an interface must be swappable without breaking the program.
- [ ] **[ISP (Interface Segregation)]:** I prefer small interfaces (`Reader`, `Writer`) over large ones (`ReadWriteCloserSeeker`). I understand that clients should not be forced to implement methods they don't use.
- [ ] **[DIP (Dependency Inversion)]:** I can inject dependencies (like a database connection) via an interface in a constructor (`NewService(db Database)`) rather than instantiating the concrete type inside the service.

### L2: The Professional (Senior)

_Goal: I apply SOLID to create testable, idiomatic packages._

- [ ] **[SRP & Packages]:** I apply SRP at the **Package Level**. I avoid creating a generic `utils` package and instead group code by domain cohesion (e.g., `net/http` handles HTTP, not "Networking").
- [ ] **[ISP (Consumer-Defined)]:** I apply the Go-specific inversion of ISP: **"Interfaces belong to the consumer."** I define the interface in the package _using_ it, not the package _providing_ it.
- [ ] **[DIP & Mocking]:** I use Dependency Inversion specifically to enable easy unit testing. I generate mocks for dependencies to test business logic in isolation.
- [ ] **[OCP via Composition]:** I implement OCP using Struct Embedding. I can wrap an existing type to intercept its methods or add new fields without rewriting the original type.
- [ ] **[LSP & Nil]:** I ensure my interface implementations respect the contract. I avoid returning `nil` errors when the operation actually failed, and I don't implement interface methods by just `panic("not implemented")`.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I prevent "Enterprise Java" patterns from infecting Go codebases._

- [ ] **[The ISP Trade-off]:** I actively fight "Interface Pollution." I only create an interface when I have **two or more concrete implementations** (or one implementation + one mock). I reject the pattern of "One Interface per Struct" (e.g., `UserService` vs `IUserService`) because it adds cognitive load without real decoupling.
- [ ] **[DIP & Wire Complexity]:** I can weigh the cost of manual Dependency Injection vs. DI Frameworks (like Uber's `fx` or Google's `wire`). I usually prefer explicit manual wiring in `main.go` to maintain the "Go Philosophy" of transparency over magic.
- [ ] **[LSP & Behavioral Subtyping]:** I validate that implementations share **Behavioral Semantics**, not just method signatures. If one `Storage` implementation allows concurrent writes and another doesn't, they violate LSP even if they satisfy the Go interface.
- [ ] **[Config vs. Code]:** I use the **Functional Options Pattern** (an application of OCP) to design APIs that can accept new configuration parameters in the future without breaking the function signature for existing users.
- [ ] **[Refactoring Strategy]:** I do not design for SOLID upfront (YAGNI). I write concrete code first, and **refactor to SOLID** only when the code actually changes. I treat abstraction as a cost I pay only when necessary.

---

## Type Aliases

_“A Type Definition creates a new identity. A Type Alias merely creates a new name. The Architect uses the former for Type Safety and the latter for massive Refactoring.”_

This module covers the subtle but critical distinction between `type A B` (New Type) and `type A = B` (Alias). Introduced in Go 1.9, Aliases are a powerful tool for gradual code repair and architecture migrations, but they are often misunderstood as simple synonyms.

---

### L1: The Foundation (Junior)

_Goal: I can distinguish between creating a new type and renaming an existing one._

- [ ] **[Syntax]:** I know the syntactic difference between a Type Definition (`type MyInt int`) and a Type Alias (`type MyInt = int`).
- [ ] **[Identity]:** I understand that an Alias is **identical** to the underlying type. I can assign a variable of type `MyInt` (alias) to a variable of type `int` without casting.
- [ ] **[New Types]:** I understand that a Type Definition creates a distinct type. I know I _must_ explicitly cast (`T(val)`) to convert between the new type and the underlying type.
- [ ] **[Built-ins]:** I recognize that `byte` is an alias for `uint8` and `rune` is an alias for `int32` in the standard library.

### L2: The Professional (Senior)

_Goal: I use aliases to refactor packages without breaking downstream clients._

- [ ] **[The Refactoring Pattern]:** I can move a struct from `package A` to `package B`, and leave a Type Alias in `package A` pointing to the new location. This allows existing consumers of `package A` to keep compiling while the code physically moves.
- [ ] **[Method Constraints]:** I understand the **"Orphan Rule."** I know I cannot define _new_ methods on an alias if the underlying type is defined in a different package. (Methods must be defined in the same package as the type definition).
- [ ] **[Embedding Behavior]:** I understand that embedding an Alias into a struct promotes the methods of the _original_ type.
- [ ] **[Exporting]:** I can use aliases to "re-export" a type from an internal sub-package to a public API surface, keeping the implementation hidden while the interface is public.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I design migration strategies and enforce domain boundaries._

- [ ] **[Safety vs. Convenience]:** I can decide when to use a **New Type** to prevent logic errors (e.g., `type Password string` prevents accidentally logging a password as a generic string) vs. when to use an **Alias** for interoperability.
- [ ] **[Gradual Code Repair]:** I can design a multi-phase architectural migration:
  1.  **Move** the type to the new Domain Package.
  2.  **Alias** the old location to the new one.
  3.  **Deprecate** the usage of the old package alias.
  4.  **Remove** the alias once all consumers are updated.
- [ ] **[Documentation Impact]:** I understand how `go doc` treats aliases. I know that aliases often obscure the "source of truth" in generated documentation, and I weigh this cognitive load against the ease of refactoring.
- [ ] **[C-Interop]:** I use aliases to map Go types to C types (e.g., `type CInt = C.int`) inside cgo applications to abstract platform-specific width differences while maintaining strict type compatibility.
- [ ] **[Generic Constraints]:** (Go 1.18+) I can use aliases to simplify complex Generic Type constraints (e.g., `type Number = interface { int | float64 ... }`) to make function signatures readable.

---

## Defer

_“Defer is the undo button for your function’s side effects. It ensures that 'exit' means 'clean up', regardless of whether that exit was a success, an error, or a panic.”_

This module covers one of Go's most distinctive control flow features. While simple on the surface, `defer` hides complexity regarding **argument evaluation timing**, **heap allocations**, and **return value modification**.

---

### L1: The Foundation (Junior)

_Goal: I can ensure resources are cleaned up safely._

- [ ] **[Syntax]:** I can register a function call to be executed when the surrounding function returns using the `defer` keyword.
- [ ] **[LIFO Order]:** I understand that multiple deferred calls execute in **Last-In, First-Out** (stack) order. (e.g., `defer A; defer B` executes B, then A).
- [ ] **[Resource Management]:** I immediately place `defer resource.Close()` after successfully acquiring a resource (File, Network Connection, Mutex) to prevent leaks.
- [ ] **[Scope Rule]:** I understand that `defer` is bound to the **Function**, not the `{ block }`. Calling `defer` inside an `if` block or `for` loop schedules it for the _end of the function_, not the end of the block.

### L2: The Professional (Senior)

_Goal: I can manipulate return values and prevent resource exhaustion._

- [ ] **[Argument Evaluation]:** I know that arguments to a deferred function are evaluated **immediately** (at the `defer` statement), while the function body executes **later** (at return). I can write code that captures the correct values based on this rule.
- [ ] **[Named Returns]:** I can access and modify the function's **Named Return Values** inside a deferred closure. (e.g., `defer func() { if err != nil { ... } }()`).
- [ ] **[Panic Recovery]:** I can use `recover()` inside a deferred function to catch a panic, log it, and gracefully downgrade it to an error value, preventing a program crash.
- [ ] **[The Loop Trap]:** I can identify the "Defer in a Loop" memory/resource leak pattern. I know how to fix it by wrapping the loop body in an **Immediately Invoked Function Expression (IIFE)** to force early execution of the `defer`.
- [ ] **[Method Receivers]:** I understand how `defer` interacts with value-receiver vs. pointer-receiver methods (does it defer the state of the object _now_ or the state _later_?).

### L3: The Architect (Mastery & Trade-offs)

_Goal: I optimize for compiler intrinsics and zero-cost abstractions._

- [ ] **[Open-Coded Defer]:** I understand the optimization (Go 1.14+) that turns deferred calls into direct jumps at compile time ("Open-Coded"), removing the heap allocation overhead for most standard cases.
- [ ] **[Heap Allocation Costs]:** I can identify scenarios where `defer` is forced to allocate on the Heap (e.g., inside loops or when the number of defers is not known at compile time), and I can refactor hot paths to avoid this cost.
- [ ] **[Wait vs. Defer]:** In extremely latency-sensitive code (Hot Paths), I have the discipline to manually `Close()` resources to save the ~30ns overhead of the defer mechanism, provided code complexity remains manageable.
- [ ] **[Exit Bypass]:** I remember that `os.Exit()` terminates the program **immediately** without running deferred functions, and I design system shutdown hooks to handle this edge case.
- [ ] **[Traceability]:** I use `defer` to implement "Function Tracing" (logging entry and exit times) by leveraging the immediate argument evaluation for the "Start" time and the execution for the "End" time.

---

## Generics (Type Parameters)

_“Generics allow us to write algorithms once for many types. But remember: Go Generics are not C++ Templates or Java Generics. They are implemented via 'GCShape Stenciling' and Dictionaries. The Architect knows that abstraction comes with a cost: either in compile time, binary size, or runtime performance.”_

This module covers Go 1.18+ Type Parameters. While Juniors rejoice at avoiding code duplication, Architects approach Generics with extreme caution, knowing that **Interface reflection** is often cleaner for application logic, while **Generics** are best reserved for core libraries and data structures.

---

### L1: The Foundation (Junior)

_Goal: I can read and write basic generic functions to avoid copy-pasting code._

- [ ] **[Syntax]:** I can declare a generic function using square brackets: `func Map[T any](s []T, f func(T) T) []T`.
- [ ] **[Constraints]:** I understand `any` (alias for `interface{}`) and `comparable` (types that support `==`). I know I cannot use `==` on a generic type `T` unless it is constrained by `comparable`.
- [ ] **[Instantiation]:** I can call generic functions both with explicit type arguments (`Fn[int](5)`) and by relying on **Type Inference** (`Fn(5)`).
- [ ] **[Generic Structs]:** I can define data structures that hold generic values, such as `type Box[T any] struct { Val T }`.
- [ ] **[Standard Lib]:** I use the `slices` and `maps` packages (Go 1.21+) for common operations (Sorting, Binary Search) instead of writing my own loops.

### L2: The Professional (Senior)

_Goal: I create reusable libraries and understand type approximation._

- [ ] **[Custom Constraints]:** I can define interface constraints with **Type Unions** (e.g., `type Number interface { int | float64 }`) to limit valid inputs.
- [ ] **[Approximation (~)]:** I understand the tilde operator (e.g., `~int`). I use it to allow my generic function to accept not just `int`, but also `type MyID int` (custom types with `int` as the underlying type).
- [ ] **[Generic Receivers]:** I can define methods on generic types (e.g., `func (b *Box[T]) Update(v T)`), but I know I **cannot** introduce new type parameters on methods (e.g., `func (b *Box[T]) Convert[U any]()` is illegal).
- [ ] **[The "Write Concrete First" Rule]:** I follow the workflow: Write the function for `int` first. If and only if I need it for `string` _and_ the logic is identical, I refactor to Generics. I do not "Prematurely Generalize."
- [ ] **[Comparability Pitfalls]:** I understand that `comparable` is restrictive (it doesn't imply strict ordering like `<` or `>`). For sorting, I know I must explicitly require an `Ordered` constraint (often from the `cmp` package).

### L3: The Architect (Mastery & Trade-offs)

_Goal: I understand Stenciling, GCShape, and Dictionaries to manage binary bloat._

- [ ] **[Implementation Mechanics]:** I understand that Go uses a hybrid approach called **GCShape Stenciling** with Dictionaries.
  - _Stenciling:_ The compiler generates unique code for types with different memory layouts (e.g., `int` vs `string`).
  - _GCShape:_ Types with the _same_ memory layout (e.g., `*int` and `*string`, or `uint64` and `int64`) share the same machine code to reduce binary bloat, using a hidden "Dictionary" to handle runtime differences.
- [ ] **[Performance Impact]:** I can analyze the overhead. I know that "Shape-shared" instances pass a hidden dictionary pointer, which can inhibit **Inlining** and add a tiny overhead compared to a purely monomorphized C++ template.
- [ ] **[Generics vs. Interfaces]:** I can choose correctly:
  - Use **Generics** when the _implementation_ stays the same but the data type changes (e.g., `ReverseSlice`).
  - Use **Interfaces** when the _implementation_ changes based on the type (e.g., `generic.Save()` where SQL and Redis have different logic).
- [ ] **[Binary Size]:** I monitor binary size growth. I know that overuse of Generics with many distinct "Shapes" (structs of different sizes) triggers the compiler to generate many unique function bodies, bloating the executable.
- [ ] **[Zero Values]:** I can handle the "Zero Value" problem in generics. Since `return nil` doesn't work for `T` (which might be an `int`), I use `var zero T; return zero`.

---

## Error Handling vs. Panics

_“In Go, errors are not 'exceptions'—they are values, just like integers or strings. We program with them, we do not simply 'catch' them. Panics are for unrecoverable program corruption. If you panic for a missing file, you are doing it wrong.”_

This module covers the most discussed aspect of Go. The Architect understands that `if err != nil` is not boilerplate; it is the **Control Flow** of reliable software. We distinguish between "Expected Failures" (Errors) and "Bugs" (Panics).

---

### L1: The Foundation (Junior)

_Goal: I handle errors explicitly and do not crash the application._

- [ ] **[Philosophy]:** I understand that Go does not have `try/catch`. I accept that handling errors is a core part of the code logic, not an afterthought.
- [ ] **[Syntax]:** I can implement the standard check: `if err != nil { return err }`.
- [ ] **[Creation]:** I can create simple error values using `errors.New("text")` and formatted errors using `fmt.Errorf("code: %d", code)`.
- [ ] **[Panic]:** I understand that `panic()` stops the ordinary flow of control and begins **Stack Unwinding**. I know it is reserved for unrecoverable states (e.g., startup configuration missing).
- [ ] **[Recover]:** I know the syntax to `recover()` from a panic using a `defer` statement, though I rarely use it yet.

### L2: The Professional (Senior)

_Goal: I write debuggable errors using wrapping and prevent service outages._

- [ ] **[Wrapping]:** I use the `%w` verb (Go 1.13+) in `fmt.Errorf("context: %w", err)` to wrap errors. I understand this creates a linked list of errors (The Error Chain), preserving the root cause.
- [ ] **[Inspection (Is)]:** I use `errors.Is(err, target)` instead of `==` to check for specific sentinel errors (e.g., `io.EOF`), ensuring checks work even if the error is wrapped.
- [ ] **[Inspection (As)]:** I use `errors.As(err, &target)` to type-assert a wrapped error into a specific struct to extract fields (like a Status Code or Retry Delay).
- [ ] **[Custom Types]:** I can define custom error structs (e.g., `type ValidationError struct`) to carry machine-readable data (field names, violation tags) alongside the human-readable message.
- [ ] **[Boundary Protection]:** I implement "Panic Recovery Middleware" at the entry points of my application (HTTP Handlers, Queue Consumers) to ensure that a single request triggering a `nil` pointer dereference does not crash the entire server.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I design error APIs for decoupling and use Panic as an internal implementation detail._

- [ ] **[Opaque Errors]:** I design APIs that return opaque errors (interfaces) rather than concrete types. I encourage callers to assert **Behavior** (e.g., `IsTemporary() bool`) rather than **Type**, minimizing coupling between packages.
- [ ] **[Internal Panic Pattern]:** I can use `panic` and `recover` strictly as an **internal** implementation detail (e.g., inside a deep recursive parser) to simplify complex control flow, provided I catch it at the package boundary and return it as a standard `error`.
- [ ] **[Sentinel Overhead]:** I avoid creating public Sentinel Errors (`var ErrNotFound = ...`) unless necessary, as they become part of the public API contract and are hard to deprecate.
- [ ] **[Performance & Stack Traces]:** I understand that `errors.New` is cheap, but libraries that attach **Stack Traces** (like `pkg/errors`) are expensive. I decide when the debuggability of a stack trace is worth the CPU/Memory allocation cost (usually Yes for App logic, No for low-level libraries).
- [ ] **[Don't Log & Return]:** I strictly enforce the rule: **"Handle it OR Return it."** I never log an error and then return it, as this floods the logs with duplicate noise up the stack.

---

## Context

_“Context is the nervous system of a Go program. It carries the signal to stop, the deadline to die, and the identity of the request. A function without a Context is a function that cannot be controlled.”_

This module covers the `context` package. It distinguishes the Junior who treats it as "that thing I have to pass as the first argument" from the Architect who uses it to orchestrate the precise lifecycle of distributed systems.

---

### L1: The Foundation (Junior)

_Goal: I can pass context through the call stack and set basic timeouts._

- [ ] **[The First Arg]:** I follow the convention of passing `ctx context.Context` as the **first argument** to functions, never embedding it in a struct (with rare exceptions).
- [ ] **[Roots]:** I know when to use `context.Background()` (main, init, tests) vs. `context.TODO()` (when I'm unsure or refactoring).
- [ ] **[Deadlines]:** I can use `context.WithTimeout(parent, duration)` and `context.WithDeadline(parent, time)` to ensure operations don't hang forever.
- [ ] **[Cancellation]:** I can use `context.WithCancel(parent)` to create a `cancel` function, and I understand that calling it stops the context's children.
- [ ] **[Cleanup]:** I strictly adhere to the rule: **Always call the `cancel` function** (often via `defer`) returned by `WithTimeout` or `WithCancel` to release resources, even if the operation succeeds.

### L2: The Professional (Senior)

_Goal: I handle cancellation signals to prevent Goroutine Leaks and use Values safely._

- [ ] **[Listening for Stop]:** I can implement the `select { case <-ctx.Done(): ... }` pattern inside long-running loops to ensure my goroutines actually exit when the context is canceled.
- [ ] **[Error Handling]:** I check `ctx.Err()` to determine _why_ the context ended (Canceled vs. DeadlineExceeded) and return appropriate HTTP status codes (499 Client Closed Request vs 504 Gateway Timeout).
- [ ] **[Standard Lib Integration]:** I use `req.Context()` in HTTP handlers and pass context to `sql` queries (`db.QueryContext`) so database operations abort when the user disconnects.
- [ ] **[Values (Keys)]:** I understand that `context.WithValue` is for **Request-Scoped Data** (Trace IDs, Auth Tokens), _not_ for passing optional function parameters.
- [ ] **[Key Collisions]:** I prevent key collisions in `WithValue` by always using **unexported custom types** for keys, never built-in types like `string` or `int`.
- [ ] **[Go 1.20+ Causes]:** I use `context.WithCancelCause` and `context.Cause(ctx)` to attach and retrieve specific errors explaining _why_ a cancellation occurred (better debugging than just "context canceled").

### L3: The Architect (Mastery & Trade-offs)

_Goal: I optimize for propagation cost and design distributed tracing topologies._

- [ ] **[Propagation Graph]:** I visualize Context as an immutable tree. I understand that canceling a parent strictly cancels all children, but canceling a child does not affect the parent.
- [ ] **[AfterFunc Optimization]:** I utilize `context.AfterFunc` (Go 1.21+) for efficient cleanup. I know this is more efficient than spinning up a new goroutine just to wait on `<-ctx.Done()`, as it registers a callback directly in the runtime's timer heap.
- [ ] **[Value Lookup Cost]:** I understand that looking up a Value in a deep context chain is **O(N)** (linear scan up the tree). I avoid storing frequently accessed dependencies (like Loggers) deep in context chains in tight loops.
- [ ] **[Network Boundaries]:** I know how to serialize context deadlines across the wire (e.g., in gRPC metadata) so that a 500ms timeout on Service A enforces a <500ms budget on downstream Service B (Deadline Propagation).
- [ ] **[Context Detachment]:** I can safely "detach" a context using `WithoutCancel` (Go 1.21) when I need to fire-and-forget a cleanup task (like logging) that shouldn't be aborted just because the incoming HTTP request was canceled.

---

## Goroutines & The Scheduler

_“Goroutines are not threads. They are the execution units of a userspace scheduler that multiplexes tens of thousands of tasks onto a handful of OS threads. To master them is to understand the difference between Concurrency (structure) and Parallelism (execution).”_

This module covers the engine of Go. A Junior sees `go func()` as a magic "run in background" button. The Architect sees the **M:N Scheduler**, stack growth mechanics, and the cost of context switching.

---

### L1: The Foundation (Junior)

_Goal: I can spawn concurrent tasks and wait for them to finish._

- [ ] **[Syntax]:** I can start a concurrent execution using the `go` keyword (`go myFunction()`).
- [ ] **[Lifecycle]:** I understand that if the `main()` function returns, the program exits immediately, killing all other running goroutines without warning.
- [ ] **[Synchronization]:** I can use `sync.WaitGroup` (`Add`, `Done`, `Wait`) to ensure the main program waits for goroutines to complete.
- [ ] **[Closures]:** I can wrap logic in anonymous functions (`go func() { ... }()`) to execute quick tasks without defining named functions.
- [ ] **[No Return]:** I understand that a goroutine cannot "return" a value to its caller like a normal function; it must communicate results via Channels or Shared Memory.

### L2: The Professional (Senior)

_Goal: I write safe concurrent code and prevent common runtime crashes._

- [ ] **[Panic Isolation]:** I understand that a panic inside a goroutine **crashes the entire application** unless it is recovered _inside_ that specific goroutine. (A `recover` in `main` does not catch panics in spawned goroutines).
- [ ] **[Loop Variable Capture]:** I am aware of the classic "loop variable capture" bug (in Go versions < 1.22) where all goroutines printed the last index of the loop. I know how to pass arguments explicitly to the closure to fix this.
- [ ] **[The Race Detector]:** I treat `go test -race` as mandatory. I do not ship code that triggers race conditions, even if "it seems to work."
- [ ] **[Goroutine Leaks]:** I treat a "leaked goroutine" (one that never terminates) as a memory leak. I ensure every goroutine has a defined exit condition (usually via a Context cancellation or Channel close).
- [ ] **[Closure Detachment]:** I am careful when passing pointers to local variables into goroutines, understanding that this forces those variables to **escape to the heap**.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I understand the runtime scheduler (GMP) and optimize for throughput vs. latency._

- [ ] **[The GMP Model]:** I can explain the Go Scheduler components:
  - **G (Goroutine):** The code to run (starts with ~2KB stack).
  - **M (Machine):** The actual OS Thread.
  - **P (Processor):** The logical resource (Context) required to execute Go code (matches `GOMAXPROCS`).
- [ ] **[Context Switching Cost]:** I know that switching Goroutines is cheap (~200ns, registers only) compared to switching OS Threads (~1-2µs, full kernel state). I use this to justify high-concurrency designs that would fail in Java/C++.
- [ ] **[Stack Growth]:** I understand **Contiguous Stack Growth**. Goroutines start small (2KB). If they need more space, the runtime allocates a larger stack, copies data over, and updates pointers. I know this copy has a performance cost during deep recursion.
- [ ] **[Work Stealing]:** I can explain how an idle P "steals" work from the local run queue of another busy P to balance the CPU load automatically.
- [ ] **[Preemption]:** I understand that modern Go (1.14+) uses **Asynchronous Preemption**. The scheduler can interrupt a tight loop (that makes no function calls) to prevent one goroutine from starving others, ensuring fairness.
- [ ] **[Affinity]:** I design for **Data Locality**. I know that passing data between goroutines on different OS threads incurs a cache-coherency penalty (L1/L2 cache misses), so I prefer keeping related work on the same goroutine when possible.

---

## Synchronization Primitives (Sync & Atomic)

_“Channels are for passing data; Primitives are for protecting state. The Proverb says 'Share memory by communicating,' but the Architect knows that sometimes, for raw performance, you must communicate by sharing memory—safely.”_

This module covers the "Classic Concurrency" tools found in `sync` and `sync/atomic`. While Go advocates for Channels, the low-level primitives are often faster and necessary for building caches, registries, and high-performance counters.

---

### L1: The Foundation (Junior)

_Goal: I can prevent data races using basic locks and counters._

- [ ] **[Mutex Syntax]:** I can use `sync.Mutex` to protect a shared map or slice. I strictly follow the pattern: `Lock()`, perform operation, `Unlock()`.
- [ ] **[Read/Write Split]:** I understand that `sync.RWMutex` allows multiple readers (`RLock`) but only one writer (`Lock`), and I use it for read-heavy workloads.
- [ ] **[WaitGroup]:** I can synchronize multiple goroutines using `sync.WaitGroup`. I know the correct order: `Add(1)` _before_ starting the goroutine, and `defer Done()` _inside_ the goroutine.
- [ ] **[Atomic Counters]:** I can use `atomic.AddInt64` and `atomic.LoadInt64` to maintain thread-safe counters without the overhead of a full Mutex.
- [ ] **[Race Detector]:** I verify my synchronization using `go test -race` and treat any output as a critical bug.

### L2: The Professional (Senior)

_Goal: I design for correctness, preventing deadlocks and copying issues._

- [ ] **[Unlock Safety]:** I almost always use `defer mu.Unlock()` immediately after locking to ensure the lock is released even if the function panics or returns early.
- [ ] **[Copying Forbidden]:** I understand that **Mutexes must not be copied**. I pass structs containing mutexes by _pointer_, not by value, to avoid copying the lock state (which renders it useless).
- [ ] **[RWMutex Cost]:** I know that `RWMutex` has higher overhead than `Mutex`. I only use it when the "Read" hold time is significant; for very quick updates, a standard `Mutex` is often faster.
- [ ] **[Atomic Types]:** I use `atomic.Pointer[T]` (Go 1.19+) and `atomic.Value` to handle thread-safe configuration updates or swap pointers without raw `unsafe.Pointer` casting.
- [ ] **[WaitGroup Pitfall]:** I know that passing a `WaitGroup` by value to a function copies the internal counter (breaking the logic). I always pass it by **pointer** (`*sync.WaitGroup`).

### L3: The Architect (Mastery & Trade-offs)

_Goal: I optimize for CPU cache coherence and lock-free algorithms._

- [ ] **[Granularity Strategy]:** I can choose between **Coarse-Grained Locking** (one lock for the whole struct) vs. **Fine-Grained Locking** (locks per field/shard). I justify the complexity of fine-grained locking only when profiling proves contention is a bottleneck.
- [ ] **[False Sharing]:** I can identify performance degradation caused by **False Sharing** (when two atomic variables sit on the same CPU cache line but are updated by different cores). I fix this by adding padding (e.g., `_ [56]byte`) between fields.
- [ ] **[CAS Loops]:** I can implement "Optimistic Locking" using `atomic.CompareAndSwap` (CAS) loops to update values without ever blocking an OS thread, reducing latency in high-contention hot paths.
- [ ] **[Mutex Internals]:** I understand Go's Mutex **Starvation Mode**. I know that if a Goroutine waits >1ms for a lock, the Mutex switches modes to give priority to the tail of the wait queue, preventing tail-latency outliers.
- [ ] **[Memory Ordering]:** I understand that `atomic` provides **Sequential Consistency** and "Happens-Before" edges. I do not rely on "benign data races" (reading a non-atomic boolean without a lock) because I know compiler reordering can break the logic.

---

## Testing (Unit & Suites)

_“Tests are the first consumer of your API. If your code is hard to test, it is not because you are bad at testing; it is because your design is tightly coupled. We write tests to drive design, not just to find bugs.”_

This module covers the `testing` package. Unlike other languages that require heavy frameworks (JUnit, PyTest), Go pushes for a minimalist, "Table-Driven" approach. The Architect knows that a robust test suite is the only safety net that enables aggressive refactoring.

---

### L1: The Foundation (Junior)

_Goal: I can verify that my code works using the standard library._

- [ ] **[Syntax]:** I can create a test file ending in `_test.go` and define functions with the signature `func TestName(t *testing.T)`.
- [ ] **[Assertion Mechanics]:** I understand the difference between `t.Error/t.Fail` (continue execution) and `t.Fatal/t.FailNow` (stop current test immediately).
- [ ] **[Table-Driven Tests]:** I can implement the **Table-Driven Test** pattern: defining a slice of structs (inputs/outputs) and iterating over them. I accept this as the _mandatory_ Go testing idiom.
- [ ] **[Running Tests]:** I can use `go test ./...` to run all tests in a module and `go test -v` to see verbose output.
- [ ] **[Code Coverage]:** I can run `go test -cover` to see what percentage of my code is exercised by tests.

### L2: The Professional (Senior)

_Goal: I write maintainable test suites, mocks, and manage test dependencies._

- [ ] **[Subtests (Native Suites)]:** I use `t.Run("SubTestName", func(t *testing.T))` to create hierarchical tests inside a Table-Driven loop. This allows running specific cases via `go test -run TestParent/SubName`.
- [ ] **[Test Helpers]:** I use `t.Helper()` inside my utility functions to ensure that failure logs point to the _caller_ (the actual test line) rather than the helper function implementation.
- [ ] **[Setup/Teardown]:** I use `TestMain(m *testing.M)` for global package-level setup (e.g., spinning up a Docker container) and deferred closures for per-test cleanup.
- [ ] **[Mocking]:** I can generate or write Mocks for my interfaces to isolate the Unit under test. I prefer generating mocks (via `mockery` or `gomock`) over manual implementation for complex interfaces.
- [ ] **[Golden Files]:** For complex output (like JSON or HTML), I use **"Golden File"** testing: reading expected output from a `.golden` file rather than hardcoding massive strings in the Go code.
- [ ] **[Parallelism]:** I use `t.Parallel()` correctly within subtests to speed up execution, and I am aware of the "loop variable capture" risks when doing so in older Go versions.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I design for testability, fuzz inputs, and enforce performance budgets._

- [ ] **[Fuzzing]:** I utilize Go 1.18+ **Native Fuzzing** (`func FuzzName(f *testing.F)`) to automatically generate random inputs and crash-test my parsers and validators against edge cases.
- [ ] **[Benchmarking & Allocs]:** I write `func BenchmarkX(b *testing.B)` to measure performance. I strictly monitor `b.ReportAllocs()` to catch "Zero-Allocation" regressions in hot paths.
- [ ] **[Integration Tags]:** I use Build Tags (`//go:build integration`) to separate slow Integration Tests (DB/Network) from fast Unit Tests, ensuring the CI pipeline remains fast for standard commits.
- [ ] **[Race Detection]:** I strictly enforce `go test -race` in CI pipelines to catch concurrent memory access bugs that are impossible to spot during manual review.
- [ ] **[Testable Examples]:** I write `func Example()` functions that serve as both valid tests and **Live Documentation** in GoDoc, ensuring my documentation never goes out of date.
- [ ] **[White-Box vs. Black-Box]:** I make intentional decisions about testing from the same package (`package mypkg`) for access to internals vs. an external package (`package mypkg_test`) to enforce public API boundaries.

---

## Channels & Concurrency Safety

_"Do not communicate by sharing memory; instead, share memory by communicating."_

### L1: The Foundation (Junior)

_Goal: I can transmit data between Goroutines and identify obvious blockers._

- [ ] **[Syntax]:** I can declare, initialize (`make`), and distinguish between **unbuffered** and **buffered** channels.
- [ ] **[Blocking Mechanics]:** I understand that sending to a full channel or receiving from an empty unbuffered channel blocks execution, and I can explain why.
- [ ] **[Directionality]:** I can correctly use the arrow syntax (`<-ch` vs `ch<-`) to send and receive values.
- [ ] **[Range & Close]:** I can iterate over a channel using `for range` and understand that the loop terminates only when the channel is closed.
- [ ] **[Basic Deadlock]:** I can recognize the runtime panic `"fatal error: all goroutines are asleep - deadlock!"` and explain why it happens (circular dependency or no sender/receiver).
- [ ] **[Race Concept]:** I can define a "Race Condition" in my own words: _Two threads accessing the same memory, at least one is a write, and no ordering is enforced._

### L2: The Professional (Senior)

_Goal: I can build robust, leak-free pipelines and prevent runtime panics._

- [ ] **[Panic Prevention]:** I can guarantee I never write to a closed channel and never close a channel twice (adhering to the "Owner closes" principle).
- [ ] **[Leak Prevention]:** I can ensure that every started goroutine has a guaranteed exit path, preventing "orphan" goroutines blocked on channels forever.
- [ ] **[Signaling Pattern]:** I can implement the `done` channel pattern (or use `context.Done()`) to broadcast cancellation signals to multiple workers.
- [ ] **[Select Statement]:** I can use `select` to handle multiple channels non-blocking, implement timeouts using `time.After`, and handle `default` cases.
- [ ] **[Tooling]:** I strictly integrate `go run -race` into my CI/CD pipeline and can interpret the stack trace output to locate the exact lines causing a data race.
- [ ] **[Type Safety]:** I enforce directional types (e.g., `func process(in <-chan int)`) in function signatures to prevent consumers from accidentally closing or writing to input streams.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I understand the `hchan` struct, the Go Memory Model, and when NOT to use channels._

- [ ] **[Internals - hchan]:** I can describe the internal `hchan` struct, specifically how it uses a **circular ring buffer** and an **internal `sync.Mutex`** to protect it. I understand that channels are _not_ lock-free.
- [ ] **[Scheduler Interaction]:** I can explain how channels interact with the Go Scheduler (parking and unparking specific Gs on `recvq` and `sendq`) and the cost of the resulting context switches.
- [ ] **[Design Decision - Mutex vs. Channel]:** I can justify the trade-off: Using **Channels** for passing ownership of data and signaling, vs. using **Mutexes** for internal state cache coherency where performance (low latency) is critical.
- [ ] **[Memory Model]:** I can explain the "Happens-Before" guarantees provided by channel sends and receives (e.g., _A receive from an unbuffered channel happens before the send completes_) and use this to prove correctness without locks.
- [ ] **[Architecture]:** I can recognize when "Fan-In/Fan-Out" is creating a bottleneck due to lock contention on a single channel, and I can re-architect (e.g., sharding channels) to resolve it.

---

## Test-Driven Development (TDD)

_"TDD is not about testing; it is about design. If your code is hard to test, your architecture is likely wrong."_

### L1: The Foundation (Junior)

_Goal: I can write a basic test to verify my logic works._

- [ ] **[The Cycle]:** I can execute the "Red-Green-Refactor" loop: Write a failing test first, write the minimal code to pass it, then clean up.
- [ ] **[Mechanics]:** I can create a valid test file (`_test.go`) and a valid test function signature `func TestName(t *testing.T)`.
- [ ] **[Assertions]:** I understand that Go has no built-in `assert` library. I can write manual checks: `if got != want { t.Errorf(...) }`.
- [ ] **[Flow Control]:** I know the difference between `t.Error` (log and continue) and `t.Fatal` (log and stop immediately) and when to use each.
- [ ] **[Execution]:** I can run tests using `go test ./...` and `go test -v` for verbose output.

### L2: The Professional (Senior)

_Goal: I write idiomatic Table-Driven Tests that serve as documentation and regression guards._

- [ ] **[Table-Driven Tests]:** I can implement the "Table-Driven" pattern using a slice of structs to test multiple inputs/outputs in a single function, rather than copying/pasting assertions.
- [ ] **[Subtests]:** I can use `t.Run()` inside a loop to execute subtests, allowing me to isolate failures within a table-driven test.
- [ ] **[Dependency Injection]:** I can use Interfaces to inject dependencies, allowing me to swap real implementations (e.g., Database) for mocks/stubs during testing.
- [ ] **[Test Helpers]:** I can use `t.Helper()` to mark utility functions, ensuring that test failures report the line number of the _caller_ (the test), not the helper function.
- [ ] **[Coverage]:** I can run `go test -cover` to analyze code paths, but I understand that 100% coverage is a metric, not a goal in itself.
- [ ] **[Race Detection]:** I automatically include `-race` in my test execution commands to catch concurrency bugs during the test phase.

### L3: The Architect (Mastery & Trade-offs)

_Goal: I use tests to drive architectural decoupling and validate system performance._

- [ ] **[Design Pressure]:** I can recognize "Test Friction" (e.g., "I need to mock 10 things to test this function") as a signal of high coupling, and I use it to justify breaking code into smaller, isolated components.
- [ ] **[Trade-off: Mocking vs. Integration]:** I can decide when to use **Unit Tests** (fast, flaky with mocks) vs. **Integration Tests** (slow, reliable, Docker containers), avoiding the "Mocking Hell" anti-pattern where tests only verify the mock setup.
- [ ] **[Fuzzing]:** I can implement `testing.F` (Go 1.18+) to automatically generate random edge-case inputs (bit-flips, empty strings) to crash-proof my parsers and validators.
- [ ] **[Benchmarks]:** I can implement `Benchmark` functions (`func BenchmarkX(b *testing.B)`), interpret `b.ReportAllocs()`, and use the results to make data-driven optimization decisions.
- [ ] **[Golden Files]:** I can implement "Golden File" testing for complex outputs (like large JSON or HTML blobs), where the test compares current output against a stored "perfect" file, rather than brittle string assertions.
- [ ] **[Black Box Testing]:** I can use the `package foo_test` (external test package) pattern to enforce testing only the public API, ensuring I am not coupling tests to internal implementation details.

---

## Fuzz Testing (Go 1.18+ Native)

_"Unit tests check the scenarios you imagined. Fuzzing explores the scenarios you never thought possible."_

### L1: The Foundation (Junior)

_Goal: I can write a basic Fuzz Target to crash-proof a function against random input._

- [ ] **[Syntax]:** I can declare a Fuzz test using `func FuzzName(f *testing.F)` and the `f.Fuzz` method.
- [ ] **[Seed Corpus]:** I can use `f.Add()` to provide valid "seed" examples so the fuzzer knows where to start mutating.
- [ ] **[Execution]:** I can run the fuzzer from the CLI using `go test -fuzz=FuzzName` and control the duration/limit with `-fuzztime`.
- [ ] **[Argument Types]:** I understand which primitive types (string, []byte, int, bool) are currently supported by the native fuzzer as arguments.
- [ ] **[Crash Identification]:** I can recognize a crash output, locate the generated artifact in `testdata/fuzz`, and understand that this file represents the specific input that broke my code.

### L2: The Professional (Senior)

_Goal: I use specific Fuzzing Patterns (Round-trip, Differential) to verify logic, not just prevent panics._

- [ ] **[The Round-Trip Pattern]:** I can implement property-based fuzzing where `Decode(Encode(x)) == x`. If the data changes after a round trip, I fail the test.
- [ ] **[The Differential Pattern]:** I can compare my optimized implementation against a known reference implementation (e.g., `MyCustomSort(x)` vs `sort.Slice(x)`) to ensure behavior is identical.
- [ ] **[Regression Management]:** I understand that when the fuzzer finds a crash, it writes a file to `testdata`. I commit this file to Git so it becomes a permanent regression unit test.
- [ ] **[Debugging]:** I can take a specific crash file generated by the fuzzer and run _only that entry_ using standard `go test` to debug the issue without restarting the fuzzing engine.
- [ ] **[Invariant Checking]:** I write fuzz tests that don't just look for panics, but explicitly check for business rule violations (e.g., "Balance can never be negative" inside the fuzz loop).

### L3: The Architect (Mastery & Trade-offs)

_Goal: I apply fuzzing strategically to security boundaries and handle complex data structures._

- [ ] **[Internals - Coverage Guidance]:** I can explain how Go's fuzzer instruments the code to track code coverage (edges) and prefers inputs that expand coverage paths, distinguishing it from "dumb" random testing.
- [ ] **[Complex Structs]:** I can bridge the gap between the native fuzzer (which only supports primitives) and complex business structs. I can write "Transformers" that consume a stream of random bytes to deterministically hydrate a complex struct for testing.
- [ ] **[Strategic Target Selection]:** I can identify the "Attack Surface" of an application (parsers, deserializers, public APIs) and mandate fuzzing there, while explaining why fuzzing internal glue code yields diminishing returns.
- [ ] **[Continuous Fuzzing]:** I can design a pipeline using tools like **OSS-Fuzz** or **GitLab/GitHub Actions** to run fuzzers continuously/nightly, rather than just locally on a developer's machine.
- [ ] **[Cost vs. Yield]:** I can determine the "diminishing returns" point. I know that 80% of bugs are found in the first hour, and I can justify when to stop burning CPU cycles based on the criticality of the component.

---

## Benchmarking & Performance Analysis
*"In God we trust; all others must bring data. Guessing about performance is professional malpractice."*

### L1: The Foundation (Junior)
*Goal: I can write a valid benchmark to measure the execution time of a function.*

- [ ] **[Syntax]:** I can declare a valid benchmark function using `func BenchmarkName(b *testing.B)` and the correct loop structure `for i := 0; i < b.N; i++`.
- [ ] **[Execution]:** I can run benchmarks using `go test -bench=.` and interpret the basic output (`ns/op`).
- [ ] **[Unit Coexistence]:** I understand that benchmarks live alongside tests in `_test.go` files but are only executed when the `-bench` flag is provided.
- [ ] **[Variable inputs]:** I can use a simple helper function to benchmark the same logic with different input sizes (e.g., `benchmarkSort(b, 100)`, `benchmarkSort(b, 1000)`).

### L2: The Professional (Senior)
*Goal: I write accurate, noise-free benchmarks that track memory and handle setup costs.*

- [ ] **[Memory Analysis]:** I always use `b.ReportAllocs()` (or `-benchmem` CLI flag) to measure **Bytes per Operation** and **Allocations per Operation**, knowing that memory pressure often hurts performance more than raw CPU.
- [ ] **[Timer Control]:** I can use `b.ResetTimer()` and `b.StopTimer()`/`b.StartTimer()` to exclude expensive setup/teardown code (like DB connection initialization) from the measurement.
- [ ] **[Compiler Evasion]:** I prevent the Go compiler from optimizing away my function call (Dead Code Elimination) by assigning the result to a package-level variable or using a `Sink` variable pattern.
- [ ] **[Sub-Benchmarks]:** I can use `b.Run("name", ...)` to create hierarchical benchmarks, allowing me to group related performance tests (e.g., Table-Driven Benchmarks) for cleaner output.
- [ ] **[Scaling Analysis]:** I can write benchmarks that prove the Big-O complexity of an algorithm (e.g., demonstrating that a map lookup is $O(1)$ vs. a slice scan $O(n)$) by systematically increasing input size.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I use benchmarks to drive architectural refactoring, isolate compiler mechanics, and verify statistical significance.*

- [ ] **[Statistical Significance]:** I use tools like `benchstat` to compare two sets of benchmark results (Old vs. New). I ignore "improvements" that fall within the margin of error and only accept changes with a high "delta" confidence.
- [ ] **[Profile Integration]:** I can generate CPU and Memory profiles directly from my benchmarks (`-cpuprofile`, `-memprofile`) and analyze the resulting **Flame Graph** to identify bottlenecks (e.g., "70% of time is spent in GC").
- [ ] **[Inlining & Escape Analysis]:** I can explain why a function suddenly became slower or faster based on whether the compiler decided to **Inline** it or if variables **Escaped to Heap**. I use `go build -gcflags="-m"` to verify my hypotheses.
- [ ] **[False Sharing & Locality]:** I understand how CPU Cache Lines (L1/L2) impact benchmarks. I can identify "False Sharing" in concurrent benchmarks and structure data to optimize for cache locality.
- [ ] **[Cost of Abstraction]:** I can create benchmarks that specifically measure the overhead of an `interface{}` method call vs. a concrete struct call, using this data to decide when to sacrifice abstraction for raw throughput in hot paths.

---

## Advanced Mocking Strategy (Uber-Go/Mock)
*“Mocks are a powertool. Used correctly, they isolate failure. Used poorly, they calcify your implementation and make refactoring impossible.”*

*Note: This module focuses on the industry-standard `go.uber.org/mock` (formerly `golang/mock`), maintained by Uber.*

### L1: The Foundation (Junior)
*Goal: I can generate a mock and verify a simple method call.*

- [ ] **[Tooling]:** I can install `mockgen` and generate a mock file from a Go interface source file.
- [ ] **[Setup]:** I can initialize a `gomock.Controller` using `gomock.NewController(t)` and understand the importance of deferring `ctrl.Finish()` (or relying on the new Go 1.14+ cleanup behavior).
- [ ] **[Basic Expectation]:** I can write a test that expects a specific method call: `mockObj.EXPECT().Method(arg).Return(value)`.
- [ ] **[Cardinality]:** I can enforce how many times a method is called using `.Times(1)`, `.AnyTimes()`, or `.MaxTimes(n)`.
- [ ] **[Strictness]:** I understand that if a mocked method is called with arguments I didn't explicitly `EXPECT`, the test will panic/fail immediately.

### L2: The Professional (Senior)
*Goal: I can mock complex interactions, capture arguments, and integrate generation into the build pipeline.*

- [ ] **[Argument Matchers]:** I can use `gomock.Any()`, `gomock.Eq()`, or `gomock.Nil()` to match arguments flexibly rather than hardcoding exact values (essential for pointers or timestamps).
- [ ] **[Custom Matchers]:** I can implement the `gomock.Matcher` interface (Matches/String methods) to validate complex structs passed to mocks (e.g., "Match if struct field `ID` is 5, ignore the rest").
- [ ] **[Ordering]:** I can enforce the *sequence* of calls using `gomock.InOrder(...)` to verify that `Connect()` happens before `Query()`.
- [ ] **[Side Effects]:** I can use `.Do()` and `.DoAndReturn()` to execute dynamic logic or modify pointer arguments when a mock is called (simulating real behavior like filling a buffer).
- [ ] **[Go Generate]:** I do not run `mockgen` manually. I embed `//go:generate mockgen ...` directives in my code so that running `go generate ./...` automatically updates all mocks.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I know when NOT to use Mocks, preferring Fakes or Stubs to avoid "Implementation Coupling."*

- [ ] **[The "Fragile Test" Problem]:** I can explain how over-using mocks (specifically `EXPECT` with strict arguments) couples tests to the *implementation details* rather than behavior. I can recognize when a test is too brittle and refactor it.
- [ ] **[Mocks vs. Fakes]:** I can distinguish between a **Mock** (behavior verification) and a **Fake** (lightweight working implementation, e.g., an in-memory map replacing a database). I can justify choosing a Fake for Data Access Layers to allow for easier refactoring.
- [ ] **[Interface Pollution]:** I resist the urge to create Interfaces *solely* for the purpose of mocking (Mocking the Implementation). I strictly adhere to "Accept Interfaces, Return Structs" and define interfaces on the *consumer* side to keep mocking boundaries clean.
- [ ] **[Refactoring Resistance]:** I design tests such that if I change the internal sequence of private method calls but the output remains the same, the test does *not* break. I use Mocks only at the boundaries of architectural layers (e.g., between the Service and the Repository).
- [ ] **[Vendor/Fork Awareness]:** I am aware of the `golang/mock` vs `uber-go/mock` history. I understand why the fork happened (abandonment of the original) and I ensure my project uses the maintained version to support Generics.

---

## The Testing Pyramid
*"Write tests. Not too many. Mostly integration." — (Ideally adapted to: "Push tests as low as possible for speed, but high enough for confidence.")*

### L1: The Foundation (Junior)
*Goal: I understand the layers and write appropriate tests for each.*

- [ ] **[The Hierarchy]:** I can define the three main layers: **Unit** (Isolated logic, fast), **Integration** (Component interaction, database), and **E2E** (Full system, slow).
- [ ] **[The Ratio]:** I understand why it is a "Pyramid" and not a "Square": I write many cheap unit tests, fewer integration tests, and very few expensive E2E tests.
- [ ] **[Unit Tests]:** I can write unit tests in Go that run in milliseconds, testing pure functions without external dependencies (no DB, no Network).
- [ ] **[Scope Definition]:** I can distinguish when a test crosses the boundary from Unit to Integration (e.g., "If it touches the filesystem or network, it's an Integration test").
- [ ] **[Execution]:** I can run specific subsets of tests using `go test -run=UnitRegex` vs `go test -run=IntegrationRegex`.

### L2: The Professional (Senior)
*Goal: I enforce the pyramid technically using Go tooling and environment management.*

- [ ] **[Build Tags]:** I can use Go build tags (`//go:build integration`) to separate slow integration tests from fast unit tests, ensuring `go test ./...` (default) remains fast for local dev loops.
- [ ] **[Short Mode]:** I can use `if testing.Short() { t.Skip(...) }` to skip long-running tests during quick sanity checks.
- [ ] **[Ephemeral Infrastructure]:** I can use tools like **Testcontainers-go** or `dockertest` to spin up throwaway databases/caches for integration tests, ensuring isolation and reproducibility.
- [ ] **[TestMain]:** I can implement `func TestMain(m *testing.M)` to handle global setup/teardown (e.g., migration application) before running the suite.
- [ ] **[Golden Path]:** I focus E2E tests strictly on "Critical User Journeys" (The Golden Path) rather than trying to test every edge case through the UI/API.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I optimize the Return on Investment (ROI) of the test suite and detect Anti-Patterns.*

- [ ] **[Anti-Pattern Recognition]:** I can identify the **"Ice Cream Cone"** (Too many manual/E2E tests, few unit tests) and the **"Hourglass"** (No integration middle layer), and I have a strategy to refactor toward the Pyramid.
- [ ] **[The "Testing Trophy"]:** I can critique the strict Pyramid and justify the **"Testing Trophy"** model (Focus heavily on Integration) for specific architectures (e.g., CRUD apps where logic is thin but wiring is complex).
- [ ] **[Contract Testing]:** I can implement **Consumer-Driven Contract Testing** (e.g., using **Pact**) to verify microservice interactions without spinning up the entire distributed system (replacing heavy E2E tests).
- [ ] **[Cost Analysis]:** I can calculate the cost of a test suite in terms of "Feedback Time." I ruthlessly delete or demote E2E tests that are flaky or take too long, replacing them with lower-level tests.
- [ ] **[Hermetic Environments]:** I design the system architecture to allow for hermetic testing, where services can be tested in total isolation with deterministic mocked time and randomness.

***

## Containers & Docker (The Artifact)
*" 'It works on my machine' is not a valid production strategy. The container is the unit of delivery."*

### L1: The Foundation (Junior)
*Goal: I can package a Go application into a container and run it.*

- [ ] **[Dockerfile Syntax]:** I can write a basic `Dockerfile` using `FROM`, `COPY`, `RUN`, `CMD`, and `ENTRYPOINT`.
- [ ] **[Build & Run]:** I can execute `docker build` to create an image and `docker run` to start a container, mapping ports (`-p`) and environment variables (`-e`).
- [ ] **[Networking Basics]:** I understand `localhost` inside a container is not `localhost` on the host, and I can make containers talk to each other on a user-defined network.
- [ ] **[Volumes]:** I can use Bind Mounts (`-v`) to inject config files or persist data locally during development.
- [ ] **[Layer Caching]:** I order my Dockerfile instructions (e.g., copying `go.mod` before source code) to utilize build cache and speed up builds.

### L2: The Professional (Senior)
*Goal: I produce minimal, secure, production-ready images.*

- [ ] **[Multi-Stage Builds]:** I use multi-stage builds (`AS builder`) to compile the Go binary in a heavy image and copy *only* the binary to a lightweight runtime image.
- [ ] **[Distroless/Scratch]:** I can deploy Go binaries on `FROM scratch` or Google's `distroless` images to reduce attack surface (no shell, no package manager).
- [ ] **[Docker Compose]:** I can write a `docker-compose.yml` to spin up a local development environment (App + Postgres + Redis) with a single command.
- [ ] **[Signal Handling]:** I ensure my Go app correctly handles `SIGTERM` (propagated by Docker) to shut down gracefully, rather than being killed abruptly by the runtime.
- [ ] **[Tagging Strategy]:** I use semantic versioning or Git SHA tags for images (`:v1.0.1`, `:a1b2c3`), never relying on `:latest` in production.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I understand Container Runtime internals and security compliance.*

- [ ] **[Internals - Namespaces & Cgroups]:** I can explain how Linux Namespaces (isolation) and Cgroups (resource limiting) actually create the "container" illusion. I know Docker is just a wrapper around these syscalls.
- [ ] **[OCI Compliance]:** I understand the Open Container Initiative (OCI) standards and can differentiate between the Image Spec and the Runtime Spec (runc vs. crun vs. gVisor).
- [ ] **[Security Hardening]:** I enforce **Rootless Containers**. I configure the container to drop Linux Capabilities (e.g., `DROP ALL`, add back only `NET_BIND_SERVICE`) and ensure the process runs as a non-UID 0 user.
- [ ] **[Supply Chain]:** I implement Image Signing (e.g., Cosign/Sigstore) and generate SBOMs (Software Bill of Materials) to guarantee artifact integrity from build to deploy.

***

## Kubernetes (The Orchestrator)
*"Cattle, not pets. If a pod dies, we do not mourn it; we replace it."*

### L1: The Foundation (Junior)
*Goal: I can deploy a simple stateless Go service to a cluster.*

- [ ] **[Core Objects]:** I can define and apply YAML for a `Pod`, `Deployment`, and `Service`.
- [ ] **[Intervention]:** I can use `kubectl get`, `kubectl describe`, and `kubectl logs` to debug a failing application.
- [ ] **[Configuration]:** I can inject configuration using `ConfigMaps` and sensitive data using `Secrets`, mapping them as Env Vars or Files in the container.
- [ ] **[Exposure]:** I understand the difference between `ClusterIP`, `NodePort`, and `LoadBalancer` services.

### L2: The Professional (Senior)
*Goal: I configure the application for high availability and self-healing.*

- [ ] **[Probes]:** I implement correct **Liveness** (restart if dead) and **Readiness** (don't send traffic until ready) probes in my Go app, distinguishing between "I am crashing" and "I am loading."
- [ ] **[Resources]:** I strictly define `requests` (for scheduling guarantees) and `limits` (for OOM killing prevention) for CPU and Memory.
- [ ] **[Lifecycle Hooks]:** I use `preStop` hooks (or just `SIGTERM` handling) to ensure connection draining occurs before the container is killed.
- [ ] **[Helm/Kustomize]:** I use templating tools to manage environment differences (Staging vs. Prod) rather than duplicating raw YAML files.
- [ ] **[Ingress]:** I can configure an Ingress Controller (Nginx/Traefik) to route external HTTP traffic to my internal services.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I extend Kubernetes itself and architect for scale.*

- [ ] **[The Operator Pattern]:** I can use **Kubebuilder** or the **Controller-Runtime** (in Go) to write Custom Resource Definitions (CRDs) and Controllers, extending K8s logic for domain-specific tasks.
- [ ] **[Scheduling Mechanics]:** I understand how the K8s Scheduler filters and scores nodes. I use `Taints`, `Tolerations`, and `Affinity` rules to control exactly where critical workloads land (e.g., dedicated nodes for high-throughput services).
- [ ] **[Autoscaling Strategy]:** I can architect the trade-offs between **HPA** (Horizontal Pod Autoscaler) and **VPA** (Vertical), and I know why mixing them often breaks things.
- [ ] **[Service Mesh Decision]:** I can decide whether to adopt a Service Mesh (Istio/Linkerd) for mTLS and Observability, or reject it due to the "Sidecar Tax" (latency/resource overhead).
- [ ] **[GitOps]:** I implement ArgoCD or Flux to ensure the cluster state is always synchronized with a Git repository, eliminating manual `kubectl apply`.

***

## Observability (The Eyes & Ears)
*"Monitoring tells you the system is broken. Observability allows you to ask the system why."*

### L1: The Foundation (Junior)
*Goal: I can generate data that lets me debug simple issues.*

- [ ] **[The Three Pillars]:** I can define **Logs** (events), **Metrics** (aggregates), and **Traces** (request lifecycle).
- [ ] **[Structured Logging]:** I never use `fmt.Println`. I use a structured logger (Go 1.21+ `log/slog` or `zap`) to emit JSON logs so they can be parsed by machines.
- [ ] **[Basic Metrics]:** I can expose a `/metrics` endpoint (Prometheus format) and track simple counters (e.g., `http_requests_total`).
- [ ] **[Panic Handling]:** I ensure panics are logged with stack traces before the application exits.

### L2: The Professional (Senior)
*Goal: I create a system that is easy to debug in distributed environments.*

- [ ] **[Distributed Tracing]:** I can instrument my code (OpenTelemetry) to propagate `TraceID` and `SpanID` across microservice boundaries (HTTP headers/gRPC metadata).
- [ ] **[Correlation]:** I automatically inject the `TraceID` into my **Logs**, allowing me to filter all logs belonging to a single specific request across the entire fleet.
- [ ] **[The RED Method]:** I instrument every service to track **R**ate (RPS), **E**rrors (Count), and **D**uration (Latency) as the gold standard for dashboarding.
- [ ] **[Alerting]:** I write alerts on *Symptoms* (High Latency, High Error Rate), not just *Causes* (High CPU), to avoid paging the on-call engineer for non-issues.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I manage the cost and value of data at scale.*

- [ ] **[Cardinality Explosion]:** I understand the cost of high-cardinality labels in metrics (e.g., putting a UserID or IP address in a Prometheus label) and how it destroys time-series databases.
- [ ] **[Sampling Strategies]:** I can choose between **Head-Based Sampling** (random % at start) vs. **Tail-Based Sampling** (keep only errors/slow traces) to balance storage costs vs. debugging value.
- [ ] **[SLO/SLI/SLA]:** I can define **Service Level Indicators** (what we measure) and **Service Level Objectives** (the internal goal, e.g., 99.9%), and use "Error Budgets" to decide when to freeze deployments.
- [ ] **[OpenTelemetry Architecture]:** I can architect an OTel Collector pipeline to receive telemetry, process/sanitize it (removing PII), and export it to multiple backends (e.g., Datadog + S3 archive) to avoid vendor lock-in.

***

## Cloud Architecture (AWS / GCP)
*"The cloud is just someone else's computer, but with an API for everything. Your job is to code the infrastructure, not just the application."*

### L1: The Foundation (Junior)
*Goal: I can deploy a binary and use core storage/compute services via SDKs.*

- [ ] **[SDK Mastery]:** I can initialize the official Go SDKs (`aws-sdk-go-v2` or `cloud.google.com/go`) handling authentication automatically (e.g., `config.LoadDefaultConfig`).
- [ ] **[Compute - VM]:** I can provision a basic VM (EC2 / Compute Engine), SSH into it, copy my compiled Go binary, and run it as a systemd service.
- [ ] **[Compute - Serverless]:** I can write a Go function compatible with **AWS Lambda** (using the lambda handler) or **Google Cloud Run** (listening on `$PORT`), understanding the difference between "Event-driven" and "Request-driven."
- [ ] **[Object Storage]:** I can implement code to upload, download, and delete files from S3 or GCS. I handle streams (`io.Reader`) correctly to avoid loading large files entirely into RAM.
- [ ] **[Identity Basics]:** I never hardcode API keys. I rely on **IAM Roles** (AWS) or **Service Accounts** (GCP) attached to the compute resource to provide credentials implicitly.

### L2: The Professional (Senior)
*Goal: I use Infrastructure as Code (IaC) and integrate managed services for robustness.*

- [ ] **[Infrastructure as Code]:** I do not click buttons in the console. I define my infrastructure using **Terraform** or **Pulumi (Go SDK)**, ensuring my environments are reproducible.
- [ ] **[Messaging Patterns]:** I can implement reliable producer/consumer patterns using **SQS/SNS** (AWS) or **Pub/Sub** (GCP). I know how to handle visibility timeouts and dead-letter queues (DLQ) to prevent message loss.
- [ ] **[Secret Management]:** I integrate with **AWS Secrets Manager** or **GCP Secret Manager** to inject configuration at runtime, rather than storing secrets in environment variables or Git.
- [ ] **[Structured Events]:** I use **EventBridge** or **Cloud Functions** to trigger Go logic based on cloud events (e.g., "Image uploaded to bucket" -> "Trigger Resize Function").
- [ ] **[SDK Mocking]:** I can write unit tests for code that uses cloud services by mocking the interfaces provided by the SDKs (preventing my tests from actually calling AWS/GCP and incurring costs).
- [ ] **[Networking]:** I understand VPCs, Subnets (Public vs. Private), and Security Groups/Firewalls. I ensure my database and internal Go services are not accessible from the public internet.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I design for Cost, Portability, and Global Scale.*

- [ ] **[Vendor Abstraction]:** I use the **Hexagonal Architecture** (Ports & Adapters) to decouple my business logic from the cloud provider. I define a `BlobStorage` interface so I can swap S3 for GCS (or a local file mock) without changing core code.
- [ ] **[Consistency Models]:** I can choose the right database based on consistency requirements. I understand the trade-offs of **DynamoDB** (Eventual Consistency options) vs. **Cloud Spanner** (External Consistency/TrueTime) and how they impact application design.
- [ ] **[Cost Optimization]:** I can architect for **Spot Instances** / **Preemptible VMs**. I write Go applications that handle abrupt termination signals gracefully to take advantage of 90% cost savings.
- [ ] **[Identity Federation]:** I implement **Workload Identity Federation** (OIDC). I configure my CI/CD (GitHub Actions/GitLab) to assume cloud roles ephemerally, eliminating the need for long-lived static access keys.
- [ ] **[Multi-Region Strategy]:** I can design an Active-Active or Active-Passive architecture using **Global Load Balancers** and replicated databases to survive a total region failure.

---

## Cloud Architecture (AWS / GCP)
*"The cloud is just someone else's computer, but with an API for everything. Your job is to code the infrastructure, not just the application."*


### L1: The Foundation (Junior)
*Goal: I can deploy a binary and use core storage/compute services via SDKs.*

- [ ] **[SDK Mastery]:** I can initialize the official Go SDKs (`aws-sdk-go-v2` or `cloud.google.com/go`) handling authentication automatically (e.g., `config.LoadDefaultConfig`).
- [ ] **[Compute - VM]:** I can provision a basic VM (EC2 / Compute Engine), SSH into it, copy my compiled Go binary, and run it as a systemd service.
- [ ] **[Compute - Serverless]:** I can write a Go function compatible with **AWS Lambda** (using the lambda handler) or **Google Cloud Run** (listening on `$PORT`), understanding the difference between "Event-driven" and "Request-driven."
- [ ] **[Object Storage]:** I can implement code to upload, download, and delete files from S3 or GCS. I handle streams (`io.Reader`) correctly to avoid loading large files entirely into RAM.
- [ ] **[Identity Basics]:** I never hardcode API keys. I rely on **IAM Roles** (AWS) or **Service Accounts** (GCP) attached to the compute resource to provide credentials implicitly.

### L2: The Professional (Senior)
*Goal: I use Infrastructure as Code (IaC) and integrate managed services for robustness.*

- [ ] **[Infrastructure as Code]:** I do not click buttons in the console. I define my infrastructure using **Terraform** or **Pulumi (Go SDK)**, ensuring my environments are reproducible.
- [ ] **[Messaging Patterns]:** I can implement reliable producer/consumer patterns using **SQS/SNS** (AWS) or **Pub/Sub** (GCP). I know how to handle visibility timeouts and dead-letter queues (DLQ) to prevent message loss.
- [ ] **[Secret Management]:** I integrate with **AWS Secrets Manager** or **GCP Secret Manager** to inject configuration at runtime, rather than storing secrets in environment variables or Git.
- [ ] **[Structured Events]:** I use **EventBridge** or **Cloud Functions** to trigger Go logic based on cloud events (e.g., "Image uploaded to bucket" -> "Trigger Resize Function").
- [ ] **[SDK Mocking]:** I can write unit tests for code that uses cloud services by mocking the interfaces provided by the SDKs (preventing my tests from actually calling AWS/GCP and incurring costs).
- [ ] **[Networking]:** I understand VPCs, Subnets (Public vs. Private), and Security Groups/Firewalls. I ensure my database and internal Go services are not accessible from the public internet.

### L3: The Architect (Mastery & Trade-offs)
*Goal: I design for Cost, Portability, and Global Scale.*

- [ ] **[Vendor Abstraction]:** I use the **Hexagonal Architecture** (Ports & Adapters) to decouple my business logic from the cloud provider. I define a `BlobStorage` interface so I can swap S3 for GCS (or a local file mock) without changing core code.
- [ ] **[Consistency Models]:** I can choose the right database based on consistency requirements. I understand the trade-offs of **DynamoDB** (Eventual Consistency options) vs. **Cloud Spanner** (External Consistency/TrueTime) and how they impact application design.
- [ ] **[Cost Optimization]:** I can architect for **Spot Instances** / **Preemptible VMs**. I write Go applications that handle abrupt termination signals gracefully to take advantage of 90% cost savings.
- [ ] **[Identity Federation]:** I implement **Workload Identity Federation** (OIDC). I configure my CI/CD (GitHub Actions/GitLab) to assume cloud roles ephemerally, eliminating the need for long-lived static access keys.
- [ ] **[Multi-Region Strategy]:** I can design an Active-Active or Active-Passive architecture using **Global Load Balancers** and replicated databases to survive a total region failure.

***