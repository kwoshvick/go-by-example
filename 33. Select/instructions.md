
Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

For our example we’ll select across two channels.

Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.

We’ll use select to await both of these values simultaneously, printing each one as it arrives.

We receive the values "one" and then "two" as expected.

Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

