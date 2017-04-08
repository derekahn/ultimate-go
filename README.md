# 💥 Big Points 🎶

- Remember LESS is MORE.

- `readability` = not hiding the cost; `simplicity` = hide complexity

- If you don't UNDERSTAND the `DATA`, you DON'T UNDERSTAND the PROBLEM you're working on.

- Remember: `stack` (stay on stack) vs `heap` (share on heap)

- Take command of `type` & Take command of `semantics`

- Mantra: REDUCE, MINIMIZE, SIMPLIFY.

- Never share a string or slice `unless` you're marshaling or unmarshaling.

- Small is fast = because smaller data structures can stay within the cachelines.

- Memory leak in go: a reference on the heap that `isn't garbage collected` AND `not being used`.

- It's never a good reason to alter a data set after you pull it!

- Know your length when getting a slice of a slice.

- (Big NO NO) Don't mix `value semantics` WITH `pointer semantics`; vice versa.

- When in doubt USE `pointer semantics`; except with time.

- Method sets dictate interface compliance.

- We do not use embedding to reuse state! We use embedding to reuse behavior. (WE EMBED BECAUSE WE NEED BEHAVIOR)

- Group by behavior (by things that they do vs. what they are).

- There is implicit conversion in go with interface types because they are valueless and same memory model.

- Always `return` the `error` interface type

- Don't use `pointer semantics` when dealing with `error` handling, because the addresses are always different.

- You are not allowed to both: log an `error` && pass it up.

- You are not allowed to create a `go routine` unless you can determine WHEN and HOW it will be terminated

- Do not think of channels as queues, but as `signals`.

- Don't use a buffer larger than 1; Anything more and it's unsafe.
  - Using `<=1` = somewhat of a garuntee.
  - Easier to catch problems/bugs.
