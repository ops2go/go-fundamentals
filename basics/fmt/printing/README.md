## Printing
The primary use of the fmt package is to format strings. These formatting functions are grouped by their output type — STDOUT, io.Writer, & string.

Each of these groups has 3 functions — default formatting, user-defined formatting, and default formatting with a new line appended.

## Printing to STDOUT
The most common use of formatting is to print to a terminal window through STDOUT. This can be done with the Print functions:

func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
The Print() function simply prints a list of variables to STDOUT with default formatting. The Printf() function allows you to specify the formatting using a format template. The Println() function works like Print() except it inserts spaces between the variables and appends a new line at the end.

I typically use Printf() when I need specific formatting options and Println() when I want default options. I almost always want a new line appended so I don’t personally use Print() much. One exception is if I am requesting interactive input from a user and I want the cursor immediately after what I print. For example, this line:

fmt.Print("What is your name? ")
Will output to the terminal with the cursor immediately after the last space:

What is your name? █

## Printing to io.Writer
If you need to print to a non-STDOUT output (such as STDERR or a buffer) then you can use the Fprint functions. The “F” in these functions comes FILE which was the argument type used in C’s fprintf() function.
```
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```
These functions are just like the Print functions except you specify the writer as the first argument. In fact, the Print functions are just small wrappers around the Fprint functions.

I typically abstract STDOUT away from my components so I use Fprint functions a lot. For example, if I have a component that logs information then I’ll add a LogOutput field:

type MyComponent struct {
        LogOutput io.Writer
}
That way I can attach STDOUT when I use it in my application:
```
var c MyComponent
c.LogOutput = os.Stdout
```
And I can attach a buffer when I use it in my tests so I can validate it:
```
var c MyComponent
var buf bytes.Buffer
c.LogOutput = &buf
c.Run()
if strings.Contains(buf.String(), "component finished") {
        t.Fatalf("unexpected log output: %s", buf.String())
}
```

## Formatting to a string
Sometimes you need to work with strings instead of writers. You could use the the Fprint functions to write to a buffer and convert it to a string but that’s a lot of work. Fortunately there are the Sprint convenience functions:

func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
The “S” here stands for “String”. These functions take the same arguments as the Print() functions except they return a string.

While these functions are convenient, they can be a bottleneck if you’re frequently generating strings. If you profile your application and find that you need to optimize it then reusing a bytes.Buffer with the Fprint() functions can be much faster.

Error formatting
One last formatting function that doesn’t quite fit into the other groups is Errorf():

func Errorf(format string, a ...interface{}) error
This is literally just a wrapper for errors.New() and Sprintf():

func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}