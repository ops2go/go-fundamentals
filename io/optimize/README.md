## Optimizing for Data Types
Reading and writing in chunks can be tedious if all you need is a single byte or rune. Go provides some interfaces for making this easier.

Working with individual bytes
The ByteReader and ByteWriter interfaces provide a simple interface for reading and writing single bytes:
```
type ByteReader interface {
        ReadByte() (c byte, err error)
}
type ByteWriter interface {
        WriteByte(c byte) error
}
```
You’ll notice that there’s no length arguments since the length will always be either 0 or 1. If a byte is not read or written then an error is returned.

The ByteScanner interface is also provided for working with buffered byte readers:
```
type ByteScanner interface {
        ByteReader
        UnreadByte() error
}
```
This allows you to push the previously read byte back onto the reader so it can be read the next time. This is particularly useful when writing LL(1) parsers since it allows you to peek at the next available byte.

## Working with individual runes
If you are parsing Unicode data then you’ll need to work with runes instead of individual bytes. In that case, the RuneReader and RuneScanner are used instead:
```
type RuneReader interface {
        ReadRune() (r rune, size int, err error)
}
type RuneScanner interface {
        RuneReader
        UnreadRune() error
}
```