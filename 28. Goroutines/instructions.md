A goroutine is a lightweight thread of execution.

Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.

To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.

You can also start a goroutine for an anonymous function call.

Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).

When we run this program, we see the output of the blocking call first, then the output of the two goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.

