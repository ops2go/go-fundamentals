## String operations
Oddly enough there is also string encoding for strings. This is used for quoting strings so that control characters and non-printable characters can be displayed. It uses Go’s character escapes so it’s very specific to the Go language itself.

## Quoting strings
You can quote strings by use the Quote() function:
```
func Quote(s string) string
```
This will convert your tabs and newlines and unprintable characters using escape sequences such as \t, \n, and \uXXXX. This can be useful when displaying error messages with data since your data may include weird characters like the Backspace character (\u0008) which is invisible.

If you need to limit your string to ASCII characters only, you can use the QuoteToASCII() function:
```
func QuoteToASCII(s string) string
```
This will ensure that fancy Unicode characters will be escaped (e.g. ☃ will be displayed as “\u2603”).

There is also another function called QuoteToGraphic() for printing Unicode Graphic characters instead of escaping them. For example, one Graphic character that gets escaped by Quote() is the Ogam Space Mark (whatever that is). Honestly, I had a really hard time figuring out when you would care about the difference. The QuoteToGraphic() function isn’t even used within the standard library.

## Efficiently quoting strings
As with other strconv functions, there’s also Append functions for appending to a byte slice to reduce allocations:
```
func AppendQuote(dst []byte, s string) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte
func AppendQuoteToGraphic(dst []byte, s string) []byte
```
## Quoting individual runes
You can also quote runes by using QuoteRune(), QuoteRuneToASCII(), and QuoteRuneToGraphic():
```
func QuoteRune(r rune) string
func QuoteRuneToASCII(r rune) string
func QuoteRuneToGraphic(r rune) string
```
One difference with these is that individual runes are quoted with single quotes instead of double quotes.

## Efficiently quoting runes
Again, there’s a set of Append functions for each of these:
```
func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
```
## Unquoting strings
If you already have a quoted string value, you can parse it into a Go string by using Unquote():

func Unquote(s string) (string, error)
This will parse not only double-quoted strings but also single-quoted and backtick-quoted strings.

## Unquoting strings the hard way
If you are a masochist then you can also unquote strings one character at a time using 
```
UnquoteChar():

func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
```
This unquotes the first character of s and returns the rune value along with whether the rune is multi-byte. It also returns tail which is the remainder of the string.
