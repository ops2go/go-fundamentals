## Scanning
We can also read our formatted data and parse it back into our original variables. This is called scanning. These are also broken up into similar groups as the printing functions — read from STDIN, read from an io.Reader, and read from a string.

### Disclaimer
I personally almost never use the scan functions in my applications. Most input for my applications come from CLI flags, environment variables, or API calls. These are typically already formatted as a basic primitives and I can use strconv to parse them. That being said, I’ll do my best to explain these with limited experience.

## Scanning from STDIN
The basic Scan functions operate on STDIN just as the basic Print functions operate on STDOUT. These also come in 3 types:
```
func Scan(a ...interface{}) (n int, err error)
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)
```
The Scan() function reads in space-delimited values into variables references passed into the function. It treats newlines as spaces. The Scanf() does the same but lets you specify formatting options using a format string. The Scanln() function works Scan() except that it does not treat newlines as spaces.

For example, you can use Scan() to read in successive values:
```
var name string
var age int
if _, err := fmt.Scan(&name, &age); err != nil {
        fmt.Println(err)
        os.Exit(1)
}
fmt.Printf("Your name is: %s\n", name)
fmt.Printf("Your age is: %d\n", age)
```
You can run this in your main() function and execute:
```
$ go run main.go
```
Jane 25
Your name is: Jane
Your age is: 25
Again, these functions aren’t terribly useful because you will likely pass in data via flags, environment variables, or configuration files.

Scanning from io.Reader
You can use the Fscan functions to scan from a reader besides STDIN:
```
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
```
## Scanning from a string
Finally, you can use the Sscan functions to scan from an in-memory string:

func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
Stringers
Go has a style convention where the names of interfaces are created by taking the name of the interface’s function and adding “er”. Because of this convention we get funny names like Stringer.

The Stringer interface allows objects to handle how they convert themselves into a human readable format:

type Stringer interface {
        String() string
}
This is used all over the place in the standard library such as net.IP.String() or bytes.Buffer.String(). This format is used when using the %s formatting verb or when passing a variable into Print().

## Formatting to Go
There is also a related interface called GoStringer which allows objects to encode themselves in Go syntax.

type GoStringer interface {
        GoString() string
}
This is used when the “%#v” verb is used in print functions. It’s uncommon to need to use this interface since the default implementation of that verb usually gives a good representation.

User-defined types
One of the more obscure parts of the fmt package is in user-defined formatter and scanner types. These give you full control over how an object gets formatted or scanned when using the Printf and Scanf functions.

## Formatters
Custom formatters can be added to your types by implementing fmt.Formatter:

type Formatter interface {
        Format(f State, c rune)
}
The Format() function accepts a State object which lists the options specified in the verb associated with your variable. This includes the width, precision, and other flags. The c rune specifies the character used in the verb.

## Concrete example
A concrete example will make more sense. Honestly, I couldn’t think of a good example of when to use this so we’ll make a silly one. Let’s say we have a Header type that is simply text that we want to be able to decorate using characters before and after. Here’s our example usage:

hdr := Header(“GO WALKTHROUGH”)
fmt.Printf(“%2.3s\n”, hdr)
Here we’re just printing hdr with the verb “%2.3s” meaning that we want 2 characters before the header and 3 characters after. Just for fun we’ll use “#” before the text and snowmen (☃) after the text. I know… dumb example but bear with me.

Here’s our Header type with its custom Formatter implementation:

// Header represents formattable header text.
type Header string
// Format decorates the header with pounds and snowmen.
func (hdr Header) Format(f fmt.State, c rune) {
        wid, _ := f.Width()
        prec, _ := f.Precision()
        f.Write([]byte(strings.Repeat("#", wid)))
        f.Write([]byte(hdr))
        f.Write([]byte(strings.Repeat("☃", prec)))
}
In our Format() function we’re extracting the width and precision from the verb (2 & 3, respectively) and then printing back into the State object. This will get written to our Printf() output.

You can run this example here and see the following output:
