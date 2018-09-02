# Package fmt

## Verbs
The Printing section of the fmt godoc has a lengthy explanation of all the options for these verbs but I’ll give a summary of the most useful ones:

* %v is a generic placeholder. It will automatically convert your variable into a string with some default options. This is typically useful when printing primitives such as strings or numbers and you don’t need specific options.
* %#v prints your variable in Go syntax. This means you could copy the output and paste it into your code and it’ll syntactically correct. I find this most useful when working with structs and slices because it will print out the types and field names.
* %T prints your variable’s type. This is really useful for debugging if your data is passed as an interface{} and you want to see what its concrete type.
* %d prints an integer in base-10. You can do the same with %v but this is more explicit.
* %x and %X print an integer in base-16. One nice trick though is that you can pass in a byte slice and it’ll print each byte as a two-digit hex number.
* %f prints a floating point number without an exponent. You can do the same with %v but this becomes more useful when we add width and precision flags.
* %q prints a quoted string. This is useful when your data may have invisible characters (such as zero width space) because the quoted string will print them as escape sequences.
* %p prints a pointer address of your variable. This one is really useful when you’re debugging code and you want to check if different pointer variables reference the same data.




## GO WALKTHROUGH☃☃☃
Inside the standard library the formatter is used for special math types such as big.Float & big.Int. Those seem like legitimate use cases but I really can’t think of another time to use this.

Scanners
On the scanning side there is a Scanner interface:

type Scanner interface {
        Scan(state ScanState, verb rune) error
}
This works similarly to the Formatter except that you’re reading from your state instead of writing to it.

## Review
Let’s review some of the do’s and don’ts of using the fmt package:

Do pronounce the package as “fumpt”!
Do use the %v placeholder if you don’t need formatting options.
Do use the width & precision formatting options — especially for floating-point numbers.
Don’t use Scan functions in general. There’s probably a better method of user input.
Do define String() functions on your types if the default implementation isn’t useful.
Don’t use custom formatters & scanners. There’s not a lot of good use cases unless you’re implementing mathematical types.
Don’t use fmt if you’ve found it to be a performance bottleneck. Drop down to strconv when you need to optimize. However, only do this after profiling!

## Conclusion
Printing output into a human readable format is a key part of any application and the fmt package makes it easy to do. It provides a variety of formatting options and verbs depending on the type of data you’re displaying. It also provides scanning functions and custom formatting if you ever happen to need that.