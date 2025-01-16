Starting with version 1.18, Go has added support for generics, also known as type parameters.

As an example of a generic function, SlicesIndex takes a slice of any comparable type and an element of that type and returns the index of the first occurrence of v in s, or -1 if not present. The comparable constraint means that we can compare values of this type with the == and != operators. For a more thorough explanation of this type signature, see this blog post. Note that this function exists in the standard library as slices.Index.

As an example of a generic type, List is a singly-linked list with values of any type.

We can define methods on generic types just like we do on regular types, but we have to keep the type parameters in place. The type is List[T], not List.

AllElements returns all the List elements as a slice. In the next example we’ll see a more idiomatic way of iterating over all elements of custom types.

When invoking generic functions, we can often rely on type inference. Note that we don’t have to specify the types for S and E when calling SlicesIndex - the compiler infers them automatically.

