Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

This person struct type has name and age fields.

newPerson constructs a new person struct with the given name.

Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.

You can name the fields when initializing a struct.

Omitted fields will be zero-valued.

An & prefix yields a pointer to the struct.

It’s idiomatic to encapsulate new struct creation in constructor functions

Access struct fields with a dot.

You can also use dots with struct pointers - the pointers are automatically dereferenced.

Structs are mutable.

If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
