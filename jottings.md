# Overview
**`for` loops** are <u>the only loops in Go</u>. They can be of **four forms**:
```go
// a regular for loop
for initilization; condition; postIncDec {
  // ...
}

// a traditional while loop
for condition {
  // ...
}

// an infinite loop
for {
  // ...
}

// iteration for loop: iterates over a range object
for value := range object { // single-value range
  // ...
}
for _, value2 := range object { // double-value range

}
```

The "initialization" portion of the `for` loop must be a simple statement that is a <u>short variable</u>, an <u>increment or assignment statement</u>, or a <u>function call</u>.


The *blank identifier* ("`_`") may be used whenever syntax requires a variable name but program logic does not, for instance to discard an unwanted loop index when we require only the element value.

A *short variable declaration* may be <u>used only within a function</u>, not for package-level variables.

`map[type]type` explained
- `[type]` specifies the data type of the key
- `type` specifies the data type of the value

The order of iteration is not specified for a `map` range, it is random, varying from one run to another.
- The `map` type is not an ordered collection though, since it doesn't use an index.

A channel is a communication mechanism allows one goroutine to pass values of a specified type to another goroutines.
```go
// create channel (of "string" type)
ch := make(chan string)

// pass channel to functiion
func f(ch chan<- string)
f(ch)

// send data (of "type") to channel
ch <- "string" // or an expression that computes to "type"

// receive data from channel
<-ch // e.g. fmt.Print(<-ch)
```
	
The "main" function runs in a goroutine and the `go` statement creates additional goroutines that runs asynchronously (multiple `go func(...args)` runs concurrently/asynchronously)

When one goroutine attempts a send or receive on a channel, it blocks until another goroutine attempts the corresponding receive or send operation, at which point the value is transferred and both goroutines proceed.

`fmt.Sprint(string)` variations works like `fmt.Print(string)` variations, but rather than print to os.Stdout, it formats accordingly and returns the resulting string.