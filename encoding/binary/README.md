## Go Walkthrough: encoding/binary
When you need to squeeze every last bit or CPU cycle out of your protocol, one of the best tools is the encoding/binary package. It operates on the lowest level and is built for working with binary protocols. We previously looked at using encoding/json for text-based protocols but binary protocols can be a much better fit when two machines need to communicate quickly and efficiently.

Writing binary protocols seems like a daunting task but it is surprisingly simple once you understand a few concepts. We’ll look at the various ways that binary data is encoded, the different tradeoffs, and hopefully dispel any mystery around the binary package.

This post is part of a series of walkthroughs to help you understand the standard library better. While generated documentation provides a wealth of information, it can be difficult to understand packages in a real world context. This series aims to provide context of how standard library packages are used in every day applications. If you have questions or comments you can reach me at @benbjohnson on Twitter.

## What is a binary protocol?
The distinction between text protocols and binary protocols may seem silly since they both communicate using bytes. However, the important distinction is in regards to which bytes they can use and how they structure them.

Text protocols, like JSON, use only the printable set of characters in ASCII or Unicode to communicate. For example, the number “26” is represented using the “2” and “6” bytes because those are printable characters. This is great for humans to read but slow for computers to read.

With a binary protocol, the number “26” can be represented using a single byte — 0x1A in hexadecimal. That’s a 50% reduction in space and it’s already in the computer’s native binary format so it doesn’t need to be parsed. This performance difference looks insignificant for a single number but it adds up when processing millions or billions of numbers.

This weird thing called “endianness”
When you save binary data you need to decide on something called endianness. The word sounds complicated but all it means is, “the order in which you write your bytes”.

This sounds odd since we, as humans, always write our numbers the same way. When we write the number 5,273, the 5 comes first, the 2 second, 7 is next, and finally the 3. In this example, the “5” is called the most significant digit and the 3 is the least significant digit. You’d never think to write that number in reverse as 3,725. That would be ridiculous.

However, in binary encoding, you do have to make this choice with the order you write bytes. There are two kinds of endianness — big endian and little endian. Big endian is when you write your most significant byte first and little endian is when you write your least significant byte first.

For example, let’s take the decimal number 287,454,020 which is 0x11223344 in hexidecimal. The most significant byte is 0x11 and the least significant byte is 0x44.

Encoding this in big endian looks like:

11 22 33 44
and encoding in little endian looks like:

44 33 22 11
Weird, right?
So why would you ever want to use little endian? Well, little endian is used by most modern CPUs to store numbers internally. The benefit to this seemingly backwards approach is that you can change the size of your number type without moving bytes.

For example, we can easily change this 4-byte int32 number we have above to an 8-byte int64 number by simply increasing the length with four zero-padded bytes. This makes it so the int32 and int64 point to the same memory address:

44 33 22 11
44 33 22 11 00 00 00 00
This kind of operation is useful at the compiler level to extend and shrink integer variables but not terribly useful at the protocol level.

## Network byte order
Big endian encoding is the convention when operating on binary numbers over a network protocol — so much so that it’s actually referred to as network byte order.

Big endian also has an interesting property that it is lexicographically sortable. That means that you can compare two binary-encoded numbers starting from the first byte and moving to the last byte. That’s how bytes.Equal() and bytes.Compare() work. This is because the most significant bytes come first in big endian encoding.

## Fixed-length encoding
In Go we are used to integer types with a specific size. The int8, int16, int32, and int64 types have the lengths of 1, 2, 4, & 8 bytes respectively. These are called fixed-length types.

If you’ve read other walkthroughs in this series you’ll notice a common pattern that there is usually one way to process streams of bytes and a different way to process in-memory byte slices. Encoding and decoding binary data is no different.

## Encoding to byte slices
To read and write binary data from byte slices we’ll useByteOrder:

type ByteOrder interface {
        Uint16([]byte) uint16
        Uint32([]byte) uint32
        Uint64([]byte) uint64
        PutUint16([]byte, uint16)
        PutUint32([]byte, uint32)
        PutUint64([]byte, uint64)
        String() string
}
This interface has two implementations: BigEndian and LittleEndian. To write a fixed-length number to a byte slice we’ll choose an endianness and simply call the appropriate Put method:

v := uint32(500)
buf := make([]byte, 4)
binary.BigEndian.PutUint32(buf, v)
Note that the Put methods will panic if you provide a buffer that is too small to write into.

You can then read your value from the buffer by using the associated getter method:

x := binary.BigEndian.Uint32(buf)
Again, this will panic if your buffer is too small. Also, if you’re reading from a stream it’s important to use io.ReadFull() to ensure that you’re not decoding from a partially read buffer.

## Stream processing
The binary package provides two built-in functions for reading and writing fixed-length values to streams. They are appropriately named: Read() and Write().

The Read() function works by inspecting the type of data and reading and decoding the appropriate number of bytes using the endianness specified by the order argument:

func Read(r io.Reader, order ByteOrder, data interface{}) error
Any error that occurs while reading from r will be returned. This function supports all fixed-length integer, float, and complex number types by using a fast type switch internally. For composite types, like structs and slices, it falls back to a slower reflection-based decoder. However, if you are decoding composite data types then you may want to look at more standardized, efficient protocols like Protocol Buffers.

On the encoding side, the Write() function works in the opposite way by inspecting the type of data and then encoding using the endianness specified by order and then writing that data to w:

func Write(w io.Writer, order ByteOrder, data interface{}) error
The same rules apply to this function as the Read() function.

## Variable-length encoding
The problem with fixed-length encoding is that it can take up a lot of space. If you need the range of an int64 type but most of your values consists of small numbers then you’ll have a lot of zero bytes in your data. One way to get around this limitation is by using variable-length encoding.

## How it works
The idea behind variable-length encoding is that small numbers should take up less space than big numbers. There are different schemes for doing this but the binary package uses the same implementation as Protocol Buffers.

It works by using the first bit of each byte to indicate if there are more bytes to be read and the remaining 7 bits to actually store the data. If the initial bit is 1 then continue reading the next bit, if it is 0 then stop reading. Once all the bytes are read, concatenate all the 7-bit data bits together to get your value.

For example, the number 53 (which is 110101 in binary) requires 6 bits of storage. Since we can store 7 bits per byte, only one byte is required:

00110101
The first 0 indicates that there are no remaining bytes and the remaining 0110101 is used to store the value.

However, for a number like 1,732 (which is 11011000100 in binary) this requires 11 bits of storage. Since we can store 7 bits per byte, we need two bytes. To construct this number we’ll set the first bit of the first byte to 1 which indicates that there is another byte to be read and the first bit of the second byte to 0 which indicates that there are no more bytes:

10001101 01000100
If we remove the leading bits from each byte we’ll get:

0001101 1000100
When we concatenate this together we get our original value of 11011000100 (which is 1,732 in decimal).

This can definitely seem confusing the first few times you work through it. However, you don’t need to know the internals of variable-length encoding to use it. It’s mostly just fun to understand.

## Encoding to byte slices
There are two functions for writing variable-length values to in-memory byte slices — PutVarint() & PutUvarint():

func PutVarint(buf []byte, x int64) int
func PutUvarint(buf []byte, x uint64) int
These work by encoding x into buf and returning the number of bytes written. If the buffer is too small then the function will panic, however, you can avoid this by using the MaxVarintLen constants when allocating your buffer.

There are also two complimentary functions for reading data back out — Varint() & Uvarint():

func Varint(buf []byte) (int64, int)
func Uvarint(buf []byte) (uint64, int)
These functions decode the data to the largest available signed and unsigned integer types — int64 & uint64. The number of bytes read is the second return argument.

Remember, the initial bit on each byte indicates if there’s more data so you can get into two error situations. First, if your buffer runs out before you get to a byte with a leading zero bit. In this case, the number of bytes read will return zero. The second error situation is when value exceeds 64 bits of data and cannot be returned in the 64-bit return type. In this case, the number of bytes read will be a negative number.

## Decoding from a byte stream
The binary package provides a way to read variable-length values from a byte stream but, strangely enough, doesn’t provide a complimentary way to write variable-length values to a byte stream.

There are two functions for decoding from a stream — ReadVarint() & ReadUvarint():

func ReadVarint(r io.ByteReader) (int64, error)
func ReadUvarint(r io.ByteReader) (uint64, error)
These work like the Varint() and Uvarint() functions except it pulls off bytes one at a time instead of operating on an in-memory byte slice. Errors that occur while reading from the stream will be passed through and values that cause a 64-bit overflow will return an overflow error.

## Summary
Binary protocols are a great way to provide low-level speed and efficiency to your communication. They may seem intimidating at first but they ultimately consist of only a few basic concepts. We learned how endianness determines the order in which we write our bytes and we saw how we can compress our values by using variable length encoding.

By understanding the encoding/binary package we also open up a world of existing standardized binary protocols. Everything from video formats to database files use binary encoding. I hope this post demystifies how these binary formats work.