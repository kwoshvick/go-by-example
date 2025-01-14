Go supports embedding of structs and interfaces to express a more seamless composition of types. This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.

A container embeds a base. An embedding looks like a field without a name.

When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.

We can access the base’s fields directly on co, e.g. co.num.

Alternatively, we can spell out the full path using the embedded type name.

Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.

Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base.

