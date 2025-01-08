Go supports recursive functions. 

This fact function calls itself until it reaches the base case of fact(0).

Anonymous functions can also be recursive, but this requires explicitly declaring a variable with
var to store the function before it’s defined.

Since fib was previously declared in main, Go knows which function to call with fib here.