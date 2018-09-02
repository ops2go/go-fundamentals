
## Floating-point operations
There are two floating-point types in Go — float32 & float64. They provide a way to express numbers which are not whole numbers. They provide a much larger range of available values compared to the integer types but they do so by trading off precision. They also allow you to represent NaN and ±Infinity.

Go implements the IEEE-754 specification for floating-point numbers and there are a lot of technical considerations when using it. I’m not going to get into those details here. Wikipedia has a good page on IEEE floating point if you want to read more.

## Parsing floats
Unlike integers, floating-point numbers can take on a couple different forms:

Integers such as “123”.
Numbers with a fractional part such as “123.45678”.
Numbers with an exponent such as “1.234E+56".
You can parse these using the ParseFloat() function:

func ParseFloat(s string, bitSize int) (float64, error)
This parses s and returns a value that fits within bitSize (which can be 32 or 64). If you try to parse a number that is too large or small then you’ll receive a NumError with Err set to ErrRange and the value will be +Infinity or -Infinity.

Parsing float-point in Go involves tons of optimization and bit twiddling so if you’re interested in the low level mechanics I suggest diving into the atof.go file to explore further.

## Formatting floats
Encoding floats to strings is a little more complicated than parsing them. For this task we use the FormatFloat() function:

func FormatFloat(f float64, fmt byte, prec, bitSize int) string
This function encodes f as a string. The fmt provides a couple options for how you want to display that float:

‘f’ — This encodes your float without any exponent. So 123.45 will print as “123.45”. Easy enough so far.
‘e’, ‘E’ — These encode your float by always using an exponent. In this case, 123.45 will encode as “1.2345E+02”. The case of the fmt character determines whether an “e” or “E” is used in your encoded string.
‘g’, ‘G’ — These encode your float without an exponent for small values and with an exponent for large values. What qualifies as a small value vs large value depends on your prec argument.
‘b’— This one is the most confusing. The other formats use a decimal exponent (e.g. 10ⁿ), however, this format uses a binary exponent (e.g. 2ⁿ). For example, 64.0 is formatted with a bitSize of 32 is “8388608p-17". You can convert this back to 64 by doing 8388608 × (2^-17). There is probably some fancy math stuff you use this for but I‘ve never had to use it.
Next is the prec argument to specify precision. For example, formatting 3.14159 with a precision of 2 will give you “3.14”. If you pass in a -1 then it’ll determine the precision based on the bitSize.

Finally, the bitSize specifies whether the formatting should treat f like a float32 or float64. As mentioned before, precision is affected by this.
