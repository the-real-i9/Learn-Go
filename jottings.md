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