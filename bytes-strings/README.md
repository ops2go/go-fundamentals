Go Walkthrough: bytes + strings packages
In the previous post we covered byte streams but sometimes we need to work with bounded, in-memory byte slices instead. While working with a list of bytes seems simple enough, there are a lot of edge cases and common operations that make using the bytes package worthwhile. We’re also going to lump in the strings package in this post since its API is nearly identical although for use with strings.

This post is part of a series of walkthroughs to help you understand the Go standard library better. While generated documentation provides a wealth of information, it can be difficult to understand packages in a real world context. This series aims to provide context of how standard library packages are used in every day applications. If you have questions or comments you can reach me at @benbjohnson on Twitter.

A brief aside on bytes vs strings
Rob Pike has an excellent, thorough post on strings, bytes, runes, and characters but for the sake of this post I’d like to provide more concise definitions from an application developer standpoint.

Byte slices represent a mutable, resizable, contiguous list of bytes. That’s a mouthful so let’s understand what that means.

Given a slice of bytes:

buf := []byte{1,2,3,4}
It’s mutable so you can update elements:

buf[3] = 5  // []byte{1,2,3,5}
It’s resizable so you can shrink it or grow it:

buf = buf[:2]           // []byte{1,2}
buf = append(buf, 100)  // []byte{1,2,100}
And it’s contiguous so each byte exists one after another in memory:

1|2|3|4
Strings, on the other hand, represent an immutable, fixed-size, contiguous list of bytes. That means that you can’t update a string — you can only create new ones. This is important from a performance standpoint. In high performance code, constantly creating new strings adds a lot of load on the garbage collector.

From an application development perspective, strings tend to be easier to use when working with UTF-8 data, they can be used as map keys whereas byte slices cannot, and most APIs use strings for arguments containing character data. On the other hand, byte slices work well when you’re dealing with raw bytes such as processing byte streams. They are also good to use when you need to avoid allocations and can reuse them.

Adapting strings & slices for streams
One of the most important features of the bytes and strings packages is that it provides a way to interface in-memory byte slices and strings as io.Reader and io.Writers.

In-memory readers
Two of the most underused tools in the Go standard library are the bytes.NewReader and strings.NewReader functions:

func NewReader(b []byte) *Reader
func NewReader(s string) *Reader
These functions return an io.Reader implementation that wraps around your in-memory byte slice or string. But these aren’t just readers — they implement all the read-related interfaces in io including io.ReaderAt, io.WriterTo, io.ByteReader, io.ByteScanner, io.RuneReader, io.RuneScanner, & io.Seeker.

I frequently see code where byte slices or strings are written to a bytes.Buffer and then the buffer is used as a reader:

var buf bytes.Buffer
buf.WriteString("foo")
http.Post("http://example.com/", "text/plain", &buf)
However, this approach incurs heap allocations which will be slow and use additional memory. A better option is to use the strings.Reader:

r := strings.NewReader("foobar")
http.Post("http://example.com", "text/plain", r)
This approach also works when you have multiple strings or byte slices by using the io.MultiReader:

r := io.MultiReader(
        strings.NewReader("HEADER"),
        bytes.NewReader([]byte{0,1,2,3,4}),
        myFile,
        strings.NewReader("FOOTER"),
)
In-memory writer
The bytes package also includes an in-memory implementation of io.Writer called Buffer. It implements nearly all the io interfaces except io.Closer & io.Seeker. There’s also a helper method called WriteString() for writing a string to the end of the buffer.

I use Buffer extensively in unit tests for capturing log output from services. You can pass it as an argument to log.New() and then verify output later:

var buf bytes.Buffer
myService.Logger = log.New(&buf, "", log.LstdFlags)
myService.Run()
if !strings.Contains(buf.String(), "service failed") {
        t.Fatal("expected log message")
}
However, in production code, I rarely use Buffer. Despite its name, I don’t use it to buffer reads and writes since there’s a package called bufio specifically for that purpose.

Package organization
At first glance the bytes and strings packages appear large but they are really just a collection of simple helper functions. We can group them into a handful of categories:

Comparison functions
Inspection functions
Prefix/suffix functions
Replacement functions
Splitting & joining functions
Once we understand how the functions group together, the large API seems much more approachable.

Comparison functions
When you have two byte slices or strings you may need to ask one of two questions. First, are these two objects equal? Second, which one comes before the other when sorted?

Equality
The Equal() function answers our first question:

func Equal(a, b []byte) bool
This function only exists in the bytes package because strings can be compared with the == operator.

Although checking for equality seems easy, one common mistake is to use strings.ToUpper() to perform case-insensitive equality checks:

if strings.ToUpper(a) == strings.ToUpper(b) {
       return true
}
This is flawed because it requires 2 allocations of new strings. A better approach is to use EqualFold():

func EqualFold(s, t []byte) bool
func EqualFold(s, t string) bool
The term “Fold” refers to Unicode case-folding. It encompasses regular uppercase & lowercase rules for A-Z as well as rules for other languages such as converting φ to ϕ.

Comparison
To determine the sort order for two byte slices or strings, we’ll use Compare():

func Compare(a, b []byte) int
func Compare(a, b string) int
This function returns -1 if a is less than b, 1 if a is greater than b, and 0 if a and b are equal. This function exists in the strings package only for symmetry with the bytes package. Russ Cox even calls out in the function’s comments that “basically no one should use strings.Compare.” Instead, use the built-in < and > operators.

“Basically no one should use strings.Compare”, Russ Cox
Typically you’ll want to know if a byte slice is less than another byte slice for the purpose of sorting. The sort.Interface requires this for its Less() function. To convert from the ternary return value of Compare() to the boolean required by Less(), we simply check for equality with -1:

type ByteSlices [][]byte
func (p ByteSlices) Less(i, j int) bool {
        return bytes.Compare(p[i], p[j]) == -1
}
Inspection functions
The bytes & strings packages provide several ways to find data within your byte slices and strings.

Counting
If you are validating input from a user, it’s important to verify that certain bytes exist (or don’t exist). You can use the Contains() function to check for existence of one or more subslices or substrings:

func Contains(b, subslice []byte) bool
func Contains(s, substr string) bool
For example, you may not allow input with certain off-color words:

if strings.Contains(input, "darn") {
        return errors.New("inappropriate input")
}
If you need to obtain the exact number of times a subslice or substring was used, you can use Count():

func Count(s, sep []byte) int
func Count(s, sep string) int
Another use for Count() is to return the number of runes in a string. By passing in an empty slice or blank string as the sep argument, Count() will return the number of runes + 1. This is different from len() which will return the number of bytes. The distinction is important when dealing with multi-byte Unicode characters:

strings.Count("I ❤ ☃", "")  // 6
len("I ❤ ☃")                // 9
The first line above may seem odd because there are 5 runes but remember that Count() returns the rune count plus one.

Indexing
Asserting contents is important but sometimes you’ll need to find the exact position of a subslice or substring. You can do this using the index functions:

Index(s, sep []byte) int
IndexAny(s []byte, chars string) int
IndexByte(s []byte, c byte) int
IndexFunc(s []byte, f func(r rune) bool) int
IndexRune(s []byte, r rune) int
There are multiple index functions for different use cases. Index() finds a multi-byte subslice. IndexByte() finds a single byte within a slice. IndexRune() finds a unicode code-point within a UTF-8 interpreted byte slice. IndexAny() works like IndexRune() but searches for multiple code-points at the same time. Finally, IndexFunc() allows you to pass in a custom function to evaluate each rune in your byte slice until a match.

There’s also a matching set of functions for searching for the first instance of the end of a byte slice or string:

LastIndex(s, sep []byte) int
LastIndexAny(s []byte, chars string) int
LastIndexByte(s []byte, c byte) int
LastIndexFunc(s []byte, f func(r rune) bool) int
I don’t use the index functions much because I find that I typically need to build something more complex such as a parser.

Prefixing, suffixing, & trimming
Working with content at the beginning and end of a byte slice or string is a special case of inspection but it’s a important enough to warrant its own section.

Checking for prefixes & suffixes
Prefixes come up a lot in programming. For example, HTTP paths are typically grouped by functionality with common prefixes. Another example is special characters at the beginning of a string such as “@” for mentioning a user.

The HasPrefix() and HasSuffix() functions allow you to check for these situations:

func HasPrefix(s, prefix []byte) bool
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix []byte) bool
func HasSuffix(s, suffix string) bool
These functions may seem too simple to bother with but one common mistake I see is when developers forget to check for zero length values:

if str[0] == '@' {
        return true
}
This code looks simple enough but if str is blank then the program will panic. The HasPrefix() function includes this validation for you:

if strings.HasPrefix(str, "@") {
        return true
}
Trimming
The term “trimming” in the bytes and strings packages refers to removing bytes or runes from the beginning and/or end of a byte slice or string. The most generic function for this is Trim():

func Trim(s []byte, cutset string) []byte
func Trim(s string, cutset string) string
This will remove any runes in cutset from the beginning and end of your string. You can also trim from just the beginning or just the end of your string using TrimLeft() and TrimRight(), respectively.

But generic trimming isn’t very common. Most of the time you want to trim white space characters and you can use TrimSpace() for this:

func TrimSpace(s []byte) []byte
func TrimSpace(s string) string
You might think that trimming with a “ \n\t” cutset is enough but TrimSpace() will trim all Unicode defined white space. This includes not only the space, newline, and tab characters but also more unusual white space characters such as thin space or hair space.

TrimSpace() is actually just a thin wrapper around TrimFunc() which is a function for evaluating leading and trailing runes for trimming:

func TrimSpace(s string) string {
        return TrimFunc(s, unicode.IsSpace)
}
This makes it simple to create your own whitespace trimmer for only trailing characters:

TrimRightFunc(s, unicode.IsSpace)
Finally, if you want to trim exact prefixes or suffixes instead of character sets, there are the TrimPrefix() and TrimSuffix() functions:

func TrimPrefix(s, prefix []byte) []byte
func TrimPrefix(s, prefix string) string
func TrimSuffix(s, suffix []byte) []byte
func TrimSuffix(s, suffix string) string
These can go hand in hand with the HasPrefix() and HasSuffix() functions if you want to replace a prefix or suffix. For example, I use this to implement Bash-style home directory completion for paths my config files:

// Look up user's home directory.
u, err := user.Current()
if err != nil {
    return err
} else if u.HomeDir == "" {
    return errors.New("home directory does not exist")
}
// Replace tilde prefix with home directory.
if strings.HasPrefix(path, "~/") {
    path = filepath.Join(u.HomeDir, strings.TrimPrefix(path, "~/"))
}
Replacement functions
Basic replacement
Swapping out subslices or substrings is sometimes necessary. For the most simple cases, the Replace() function is all you need:

func Replace(s, old, new []byte, n int) []byte
func Replace(s, old, new string, n int) string
It swaps out any instance of old with new in your string. You can set n to a non-negative number to limit the number of replacements. This function is good if you have a simple placeholder in a user defined template. For example, you want to let users specify “$NOW” and have it replaced with the current time:

now := time.Now().Format(time.Kitchen)
println(strings.Replace(data, "$NOW", now, -1)
If you have multiple mappings then you’ll need to use strings.Replacer. This works by specifying old/new pairs to strings.NewReplacer():

r := strings.NewReplacer("$NOW", now, "$USER", "mary")
println(r.Replace("Hello $USER, it is $NOW"))
// Output: Hello mary, it is 3:04PM
Case replacement
You may assume that casing is simple — upper & lower case — but Go works with Unicode and Unicode is never that simple. There are 3 types of casing: upper, lower, and title case.

Uppercase and lowercase are straight foward for most languages and you can use the ToUpper() and ToLower() functions:

func ToUpper(s []byte) []byte
func ToUpper(s string) string
func ToLower(s []byte) []byte
func ToLower(s string) string
However, some languages have different rules for casing. Turkish, for example, uppercases its i as İ. For these special case languages, there are special versions of these functions:

strings.ToUpperSpecial(unicode.TurkishCase, "i")
Next we have title case and the ToTitle() function:

func ToTitle(s []byte) []byte
func ToTitle(s string) string
You may be surprised, however, when you use ToTitle() and all your characters are uppercased:

println(strings.ToTitle("the count of monte cristo"))
// Output: THE COUNT OF MONTE CRISTO
That’s because in Unicode, title case is a specific type of casing and not a way to capitalize the first character in each word. For the most part, title case and upper case are the same but there are a few code points which have differences. For example, the ǉ code point (yes, that’s one code point) is uppercased as Ǉ but title cased as ǈ.

What you’re probably looking for is the Title() function:

func Title(s []byte) []byte
func Title(s string) string
This outputs the expected result:

println(strings.Title("the count of monte cristo"))
// Output: The Count Of Monte Cristo
Mapping runes
One other function for replacing data in a bytes slice or string is Map():

func Map(mapping func(r rune) rune, s []byte) []byte
func Map(mapping func(r rune) rune, s string) string
This function lets you pass in a function to evaluate every rune and replace it. Admittedly, I didn’t even know this function existed until I started writing this post so I can’t give any personal anecdote.

Splitting & joining functions
Many times we have delimited strings that we need to break apart. For example, paths in Unix are joined with colons and the CSV file format is essentially just fields of data delimited by commas.

Substring splitting
For simple subslice or substring splitting, we have the Split() functions:

func Split(s, sep []byte) [][]byte
func SplitAfter(s, sep []byte) [][]byte
func SplitAfterN(s, sep []byte, n int) [][]byte
func SplitN(s, sep []byte, n int) [][]byte
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string
These break up byte slices or strings by a delimiter and return the subslices or substrings. The “After” functions include the delimiter at the end of the substrings. The “N” functions limit the number of splits that can occur:

strings.Split("a:b:c", ":")       // ["a", "b", "c"]
strings.SplitAfter("a:b:c", ":")  // ["a:", "b:", "c"]
strings.SplitN("a:b:c", ":", 2)   // ["a", "b:c"]
Splitting data is a very common operation, however, it’s typically done in the context of a file format such as CSV or in the context of path splitting. For these operations, I use the encoding/csv or path packages instead.

Categorical splitting
Sometimes you want to specify delimiters as a set of runes instead of a series of runes. The best example of this is breaking apart words by variable-length whitespace. Simply calling Split() using a space delimiter will give you empty substrings if you have multiple contiguous spaces. Instead you can use the Fields() function:

func Fields(s []byte) [][]byte
This will consider consecutive whitespace characters a single delimiter:

strings.Fields("hello   world")      // ["hello", "world"]
strings.Split("hello   world", " ")  // ["hello", "", "", "world"]
The Fields() function is just a simple wrapper around FieldsFunc() which lets you pass a function to evaluate each rune as a delimiter:

func FieldsFunc(s []byte, f func(rune) bool) [][]byte
Joining
Instead of breaking apart delimited data, we can join it together using the Join() function:

func Join(s [][]byte, sep []byte) []byte
func Join(a []string, sep string) string
One common mistake I‘ve seen is when developers try to implement join by hand. It looks something like:

var output string
for i, s := range a {
        output += s
        if i < len(a) - 1 {
                output += ","
        }
}
return output
The flaw in this code is that you are creating a massive number of allocations. Because strings are immutable, each iteration is generating a new string for each append. The strings.Join() function, on the other hand, uses a byte slice buffer to build upon and converts it back to a string when it returns. This minimizes heap allocations.

Miscellaneous functions
There’s two functions I couldn’t find a category for so they’re lumped in here at the bottom. First, the Repeat() function allows you generate a repeated byte slice or string. Honestly, the only time I can remember using this is to make a line to separate content in the terminal:

println(strings.Repeat("-", 80))
The other function is Runes() which returns a slice of all runes in a UTF-8 interpreted byte slice or string. I‘ve never needed to use this since the for loop over a string does the same thing but without the allocations.

Conclusion
Byte slices and strings are fundamental primitives in Go. They are the in-memory representations for series of bytes and runes. The bytes and strings packages provide a ton of useful helper functions as well as adapters to the io.Reader and io.Writer interfaces.

It’s easy to overlook many of the useful tools in these packages because of the API’s size but I hope this post has helped you to understand everything these packages have to offer.