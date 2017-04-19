# mastering-go-concurrency
Go Concurrency 101


- Concurrency is the independent executions of goroutines.
- Functions are created as goroutines with the keyword go.
- Goroutines are executed within the scope of a logical processor that owns a single operating system thread and run queue.
- A race conditions is when two or more goroutines attempt to access the same resource.
- Atomic functions and mutexes provide a safe way to protect against race conditions.
- Channels provide an intrinsic way to safely share data between two goroutines.
- Unbuffered channels provide a guarantee between an exchange of data. Buffered channels do not.