Slices are an important data type in Go, giving a more powerful interface to 
sequences than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.

To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
We can set and get just like with arrays.

In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.

Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.

Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].

This slices up to (but excluding) s[5].
And this slices up from (and including) s[2].

We can declare and initialize a variable for slice in a single line as well.

The slices package contains a number of useful utility functions for slices.

Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.

Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.