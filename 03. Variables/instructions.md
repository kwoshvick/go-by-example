
In Go, variables are explicitly declared and used by the compiler 
to e.g. check type-correctness of function calls.

var declares 1 or more variables.

Go will infer the type of initialized variables.

Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.

The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.