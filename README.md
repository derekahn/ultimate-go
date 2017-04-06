# Ultimate Go 4/4/17 - 4/5/17

### Remember less is more
- Hardware is the platform; machine is the model.

- NUMA = non-uniform memory access; this is complex.

- Productivity > Performance; keep moving forward

- --Debugger; debugger == death; cannot maintain mental model

### readability = not hiding the cost; simplicity = hide complexity

- `var` always for zero value (bill standard); `var zero int` vs `zero := 0`

### If you don't understand the data, you don't understand the problem you're working on.


## Do not create values using pointer semantics

#### stack (stay on stack) vs heap (share on heap)

Go allows:
  - Mechanical sympothy - layout data and access it.
  - Data transformation.
  - Reduce pressure on GC.

- pointers are for sharing; sharing data across program boundries.
  - program boundry: fn call, moving data across 2 goroutines.

go gives 3 blocks of memory:
data segment, stack (1mb), heap

- stack frames are calculated at compile time.

- GC operates against the heap.
- values on heap are meant to be shared.
- every fn call requires new frame; every time you allocate a frame you're initializing a frame; clean before usage.

- stacks go down, memory addresses go down.

- Stacks are 2kb.

- pointer variables = * ponter type's name; always stores memory addresses;  all allocate the same amount of memory 4 || 8 bytes.
- Read and write to the value that the address points to.

- escape analysis (algorithm): Asks can we keep this value on this value on the stack frame for this fn? If shared unsafely it moves to heap; get to maintain value semantics or pointer semantics.

- sharing down the call stack doesn't create an allocation to the heap.
- sharing a value up the call stack `return &u`; it's on the heap.
- heap values once allocated do not move; unlike stack values they can grow.

```go
 count := 10
 // what's in the box, where is the box
 println( count, &count)
```
## do not create a value using pointer semantics
```go
// hides readability
u := &user {
  name: "Bill",
  email: "bill@email.env",
}
```
- You can share going down the call stack. Can't share going up.

- GC's pacing (algorithm) - monitoring live heap trying to keep it at 2mb.

- GC turning on is a 25% performance hit.

### Take command of type & Take command of semantics

- cache hides latency costs.

TLB - table lookaside buffers

- arrays = predicatble access patterns.

TODO: [code::dive conference 2014 - Scott Meyers: Cpu Caches and Why You Care ](https://www.youtube.com/watch?v=WDIkqP4JbkE)

- HARDWARE `<3` arrays == mechanical sympathy.

- OOP--, because OOP == linked lists; child parent taxonomies.

### mantra: REDUCE, MINIMIZE, SIMPLIFY


## slices
- `make()` is for: arrays, slices, channels
- 3 words:
  - 1st structure is a pointer
  - length = read and write
  - capacity = total number of elements

### Never share a string or slice UNLESS you're marshaling or unmarshaling

### Small is fast = because smaller data structures can stay within the cachelines
```go
var data []string // zero value
data := []string{} // empty value
```
empty struct `struct{}`; you could create billion vals with 0 allocation.

### Memory leak in go: a reference on the heap that isn't Getting GC and not being used

- Deterining a 2nd arg in a slice[a:b] add b with how many you need.
- After a 1k elements in an array it stops doubling and only grows by a 1/4 (25%).
- Append always works on length.

- a slice of pointers is essentially a linked list.
- `append` is a mutation operation.

### Never a good reason to alter a data set after you pull it!!!

- Pointers are best passed down vs up because it would have to go in the heap.
- Value semantics mean a copy of the value is passed across program boundaries.
- Pointer semantics mean a copy of the values address is passed across program boundaries.

### Know your length when getting a slice of a slice

- data race; 2 goroutines, 1 performing a read 1 performing a write at the same time.

## Strings
- `rune` is an alias of type int32
- `byte` is an alias for uint8
- A string is immutable

> Every array in go is a slice waiting to happen.


# Design

> "Methods are valid when it is practical or reasonable for a piece of data to expose a capability." - William Kennedy

- DATA vs BEHAVIOR

- rule thumb: variable reciever should be shorter than 3 characters

### NO NO! Don't mix value and pointer semantics. example:
```go
func changeName(u *user, name string) {
	u.name = name
}
```
### WHEN IN DOUBT USE POINTER SEMANTICS, except with time

- Your factory functions dictates semantic consistency.

- Value semantics leverage the stack and take pressure off the heap.
- Pointer semantics leverage effciency and the heap for sharing.

- Applying pointer semantics to real world. Can you copy a person?  You should use a pointer.

- Concurrency is complex; Life will be simplier without using the `go` keyword

## Interfaces

- Interfaces of a contract of behavior

[sqlx > gorm](http://jmoiron.github.io/sqlx/)

# Day 2

## Interface vs concrete type

- Basically 2 type systems (interface, concrete type)
- A concrete type is you can apply a method too.
- Interfaces define behavior; interface types are `value-less`.
- A relationship of storage!
  - goes onto the heap.

### Method sets dictate interface compliance

- [Method sets wiki](https://github.com/golang/go/wiki/MethodSets)

rules:
1. If you have a *T you can call methods that have a receiver type of *T as well as methods that have a receiver type of T (the passage you quoted, Method Sets).
2. If you have a T and it is addressable you can call methods that have a receiver type of *T as well as methods that have a receiver type of T, because the method call t.Meth() will be equivalent to (&t).Meth() (Calls).
3. If you have a T and it isn't addressable, you can only call methods that have a receiver type of T, not *T.
4. If you have an interface I, and some or all of the methods in I's method set are provided by methods with a receiver of *T (with the remainder being provided by methods with a receiver of T), then *T satisfies the interface I, but T doesn't. That is because *T's method set includes T's, but not the other way around (back to the first point again).


- literal values only exist at compile time.
- Pointer semantics or value semantics matter when implementing an interface.
- If using a pointer receiver on a method, only pointer values will satisfy interface (integrity)
- Nouns as pointers?

### We do not use embedding to reuse state! We use embedding to reuse behavior. (WE EMBED BECAUSE WE NEED BEHAVIOR)

- concrete type system: algorithm (data), effciency
- interface type system: decoupling, behavior

- Start at concrete types (the data), then implement behavior with interfaces.

### Group by behavior (by things that they do vs. what they are)

- Archetecting/designing/APIs is all about that behavior.

## Rules of knowing being done
- 80% test coverage on everything. 100% on happy path.
- Decoupling from changes of today.

Design example:
[xenia transfer to pillar](https://play.golang.org/p/ES4BOnDX6O)

- Do interfaces last

Interfaces cost:
- Allocation (goes onto heap)
- Indirection (fn call)

### implicit conversion in go with interface types because valueless and same memory model

Two ways 2 shutdown program:
1. panic
2. os.Exit / log.Fatal
    - if need stack trace { log.Fatal } else os.Exit

## How to use an interface:
Declare an interface when:
* users of the API need to provide an implementation detail.
* API’s have multiple implementations they need to maintain internally.
* parts of the API that can change have been identified and require decoupling.

Don't declare an interface:
* for the sake of using an interface.
* to generalize an algorithm.
* when users can declare their own interfaces.
* if it's not clear how the ineterface makes the code better.

## Error handling

- If your user must parse the string that comes out of this implementaion, YOU FAILED THEM

- `err != nil` Is there a concrete type stored in this interface; err is not set to a zero value.

- `if` statement for negative path processing/err handling.
- `switch` > `else`.
- `if {} else{}` for assignment only.

- One `err` return per fn, this will allow for concise context.

- `err.(type) // type assertion on the concrete type` = is walking away from decoupling with interfaces;

4 things to allow for custom err types exported:
1. temporary
2. timeout
3. not-found
4. not-authorized

- custom error types are just artifacts.

### Always return the error interface type

### Don't use pointer semantics when dealing with error handling, because addresses are always different

- Go perspective logging is expensive, there is an allocation or 2.
- Logging an error anything that is not actionable is wasting resources.
- Logging is apart of error handling (not decoupled) | error is not handled until it's logged.
  - Who get's to handle the error?
- If you pass the error up the call chain ADD CONTEXT.

### You are not allowed to log an error && pass it up.

- It's best to know how to handle err logging from day 1.

[Best error package for handling context](https://github.com/pkg/errors)


## Concurrency

A thread can be in 1 of 3 states:
1. Execution state = Running on a core.
2. Runnable state = wants to execute but sitting in a queue.
3. Sleep state = Waiting for something to be moved to runnable.

- Go's runtime is a cooperative scheduler.
  -  Mimics linux scheduler.
  - Sits on top of OS sheduler.
- CPU can be in 1 of 2 modes: Kernal mode (unrestricted) OR user mode (restricted).

Scheduling decision:
1. `go` keyword
2. garbage collection
3. system call; like `fmt.Println`

- Go routine time idles so that it tricks the OS to not physically make a context switch.

- Think of concurrency like parent hood. Everytime you run a go routine you're brining a child into the world.

### You are not allowed to create a go routine unless you can determine WHEN and HOW it will be terminated

- Leverage the `sync.wg` package to manage go routines.
  - Don't use `wg.Add` within a `for` loop.

2 rules of thumbs for multi threaded software:
1. Some shared state that multiple go routines want to access at the same time
  - Data race = read and write same data same time.
2. Coordinate 2+ goroutines to

- Bill says: `atomic` > `mutex`

### Do not think of channels as queues

- channels serve one purpose `signaling`.
  - signals with data.
  - signals without data.

- Two types of channels:
  - Un buffered.
    * (Rob Pike always uses this type of channel) They all you to know 100% garuntee that the signal has been received.
      - This reduces risk
      - Signaling with the data, when the other go routine recieves it the intializer releases.
      - Cost is latency (because it's blocking)
  - Buffered.

### You are not allowed to use a buffer larger than 1, because this can give a garuntee and reduce latency

- `send` is a binary opertaion
- `receive` is a unary operation

- `ch := make(chan struct{})` unbuffered channel singaling without data

- Channel is in one of 2 states: open or closed.
- Once you `close` a channel you cannot re-open

- `select` in goroutines for cancellation

- Ask if a buffered channel is 1 > if it is a "fan out pattern".

- you'll `panic` if you try to `close` a channel twice, or if you try to send on a closed channel.
