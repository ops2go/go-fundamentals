## Boolean operations
Parsing booleans
To parse boolean values, we can use the ParseBool() function:

func ParseBool(str string) (bool, error)
This function has a set list of true and false values for str. True values consist of “1”, “t”, “T”, “true”, “True”, & “TRUE”. False values consist of “0”, “f”, “F”, “false”, “False”, & “FALSE”. Anything else will return an error.

I don’t typically use ParseBool() simply because I like my inputs to be specific (e.g. either “true” or “false”).

Formatting booleans
We can perform the reverse operation and format a boolean using FormatBool():

func FormatBool(b bool) string
This returns “true” or “false” depending on the value of b. Easy peasy so far.

There’s also another option for formatting booleans when you’re using byte slices called AppendBool():

func AppendBool(dst []byte, b bool) []byte
This will append “true” or “false” to the end of dst and return the new byte slice.

Despite this function only being 3 lines, I do find myself using it. No reason to rewrite 3 lines of code that’s already in the standard library.
