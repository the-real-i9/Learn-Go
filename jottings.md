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

---

A *short variable declaration* may be <u>used only within a function</u>, not for package-level variables.

---

`map[type]type` explained
- `[type]` specifies the data type of the key
- `type` specifies the data type of the value

The order of iteration is not specified for a `map` range, it is random, varying from one run to another.
- The `map` type is not an ordered collection though, since it doesn't use an index.

---

A channel is a communication mechanism allows one goroutine to pass values of a specified type to another goroutines. The channel either has a <u>bounded capacity buffer</u> or a <u>zero capacity buffer</u>, in which case the channel is unbuffered
```go
// create channel (of "string" type)
ch := make(chan string) // zero capacity buffer or unbuffered channel
ch := make(chan int, 5) // bounded capacity buffer

// pass channel to functiion
func f(ch chan<- string)
f(ch)

// send data (of "type") to channel
ch <- "string" // or an expression that computes to "type"

// receive data from channel
<-ch // e.g. fmt.Print(<-ch)
```

---

The "main" function runs in a goroutine and the `go` statement creates additional goroutines that runs asynchronously (multiple `go func(...args)` runs concurrently)
- When one goroutine attempts a send or receive on a channel, it blocks until another goroutine attempts the corresponding receive or send operation, at which point the value is transferred and both goroutines proceed.

---

`fmt.Sprint(string)` variations works like `fmt.Print(string)` variations, but rather than print to os.Stdout, it formats accordingly and returns the resulting string.

`fmt.Fprint(w, string)` variations, formats accordingly writes the resulting string to `w io.Writer`

`log` works like `fmt`, but executes some system call after.

----

To make a folder into a module/package. In the folder, run this command
```bash
go mod init ${parentFolder}/${folder}
```
This creates a `go.mod` file in your folder. With the contents:
```mod
module ${parentFolder}/${folder}

go ${versionNumber}
```

- <u>**Before** your folder becomes a module,</u> all `.go` files in it are independent, in that case, multiple `.go` files can contain `func main()` and can bear any package name (including "main")
- <u>**After** your folder becomes a module,</u> all go files in it are related, in that case, only one `.go` file can contain `func main()` and all `.go` files must have the same `package ${name}`, where `${name}` can be your `customPackageName` or `main`.

There are **two kinds of module/package**: one that's a custom package and the other a special "main" package. What you get depends on the package name you declare in your `.go` file(s).\
Regardless of your module name (folder name)
- if you delare `package main`, your module becomes the main package/module, it becomes the entry point of execution, all other packages are imported into it directly or indirectly. But, 
- if you declare `package customPackageName`, it becomes a custom package, that may imported by the main package/module directly or indirectly

You can **compile** <u>only a `package main`</u> module/package/folder into one executable. 
- Custom packages that it dependes on will be compiled internally and dynamically imported.

---

Go allows a simple statement such as a local variable declaration to precede the `if` condition, which is particularly useful for error handling, or basically if condition depends on the statement.
```go
// basic if
if condition {}

// with simple statement
if statement; condition {}
// same as puting `statement` before the line
```

---

**Control flow:** `switch`
```go
switch expression {
case condition:
  evaluation
...
default:
  evaluation
}
```
Other than `if` and `for`, we have the `switch`.

`expression` is optional, if not included, it's called a *tagless* switch, which defaults to `true`.

`case`s <u>do not fallthrough from one to the next</u> as in C-like languages (where you have to use a `break` statement to go out of the control flow). *(though there is a rarely used `fallthrough` statement that overides this behaviour)*.

The `break` and `continue` statements modify the flow of control, as expected.

---

Before you embark on any new program, it's a good idea to see if packages already exist that might help you get your job done more easily.

---
## The flag package

`package flag` implements command-line flag parsing.

It uses a program's command-line arguments to set the vaues of certain variables distributed throughout the program.

It allows us define new "flag command-line arguments" and make use of them.

```go
var sep = flag.String("s", " ", "error message")

flag.Parse()

fmt.Print(strings.join(flag.Args(), *sep))
```
> First we define our command-line flags
```go
var sep = flag.String("s", " ", "error message")
/* (flagName, flagDefaultValue, flagErrorMessage) */
```
`sep` allows us to use the flag `-s` of the cmd-line in our program as a string.


> Next, we parse the command-line arguments to identify the flag and non-flag ones.
```go
flag.Parse()
```
The flag argruments are the ones we've defined and the non-flag args are the ones we didn't define 


> Finally, we then use the flag:
```go
fmt.Print(strings.join(flag.Args(), *sep))
```

`flag.Args()` is the range[] of all "non-flag" args
- `sep` identifies the flag in our program, while `-s` identifies the flag in the command-line
  - the value of the flag is what (string) immediately follows `-s`, we use `*sep` to get it in our program, as `sep` holds the address of the variable containg the value.
  - if the flag isn't specified, its default value is used with the above line of code

When our control encounters this line of code, it replaces `*sep` with the value specified in the cmd-line as the separator
  - if we specify `-s \`, then "`\`" is the separator
  - if the flag is not specified the default value of the flag is used.

Both `sep` and `-s` are same, the former helps use it in our program, while the latter helps use it in the command-line

If an undefined flag is used, the error message specified with our flags will throw.

> 😎✨This is a major tool in building command-line applications. For example,\
> To display help message for your program, you can create a boolean flag that checks if `-h` or `-help` is specified in the command-line arguments, follow by `fmt.Println("Help message")`.

---

# Program Structure

## Pointers

A *pointer* value is the *address* of a variable. A pointer is thus the location at which a value is stored.

With a pointer, <u>we can read or update the value of a variable indirectly</u>.

Each component of a variable of aggregage type (struct field or array element) is also a variable and thus has an address too.

Expressions that denote variables are the only expressions t which the *address-of* operator `&` may be applied.

> Since a pointer to a variable, is itself a variable, that holds the address of another variable. That means <u>a pointer itself a variable and is addressable</u>.
```go
var x = 10

var px = &x // pointer to x

var ppx = &px // pointer to the pointer variable of x

&px // address of pointer to x

*&px == &x // true
*ppx == &x // true
```

The zero value for a pointer is `nil`

Pointers are comparable; <u>two pointers are equal if and only if they point to the same variable or both are `nil`</u>

It is perfectly safe for a function to return the address of a local variable. The local variable will remain in existence even after the function has returned, and a pointer can refer to it.
- Comparing the function call with `==` will not equal each other.

Because a pointer contains the address of a variable, passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed. *You're passing the actual memory box*.

Pointers are key to the `flag` package, check [The flag package](#the-flag-package) above.

---

## The `new` Function
Use the `new(T)` to create a *pointer to an unnamed variable* of type `T`. This is just a syntactic convenience, such that, you don't have to first declare a dummy unnamed variable before you can then create a pointer to it.
```go
// instead of
var dummy int;
var p = &dummy
*p = 2

// you can do
var p = new(int)
*p = 2
```

## Lifetime of Variables
The lifetime of a package-level variable is the entire execution of the program.
- You can access a package-level variable before it is declared.

The lifetime of a local variable starts each time its function is invoked, and ends when it becomes *unreachable* (in normal cases, when the function returns).
- A local variable is not available for access until it is declared.

A compiler may choose to allocate variables on the heap or on the stack.
  - If you assign a local variable to a package-level pointer variable, the local variable will be allocated on the heap rather than on the stack. This means you *escaped the local variable*, and it's still reachable.
  - This is good to keep in mind during performance optimization, since each variable that escapes requires an extra memory allocation.

Keeping unnecessary pointers to short-lived objects within long-lived objects, especially global variables, will prevent the garbage collector from reclaiming the short-lived objects.

## Assignments
Expressions with multiple results must have multiple vaiables declared on the LHS, as many as the results.
- We can assign unwanted values to the blank identifier ( `-` ).

```go
f, err = os.Open("foo.txt")

_, err = io.Copy(os.Stdout, req.Body) // discard byte count
```

## Assignability
An assignment, explicit or implicit, is always legal if <u>the LHS (the variable) and the RHS (the value) have the same type</u>.

Whether two values may be compared with `==` and `!=` is related to assignability: <u>in any comparison, the first operand must be assignable to the type of the second operand, or vice versa</u>.

## Type Declarations
There are variables that share the same representation but signify very different concepts.

A `type` declaration defines a new *named type* that has the same *underlying type* as an existing type.
```go
type Celcius float64
type Farenheit float64
```

They  don't change the value or representaion in any way, but they make the change of meaning explicit.

A named type may provide notational convinience if it helps avoid writing out complex types over and over again. e.g. *struct*s.

```go

```

Named types also make it possible to define new behaviours for values of the type. These behaviors are expressed as a set of functions associated with the type, called the type's *methods*.
```go
func (c Celcius) toString() string {
  return fmt.Stprintf("%g", c)
}
```

