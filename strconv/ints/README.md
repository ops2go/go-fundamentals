## Integer operations
Computers use a binary representation of numeric values so we have to convert them if we want to work with them in decimal (or other numeral systems).

Go has two kinds of integer types, int & uint, for signed and unsigned representations. The strconv package has different functions for each kind.

## Parsing integers
If you have a string and want to convert it to one of Go’s integer types, you can use the ParseInt() or ParseUint() functions:

func ParseInt(s string, base int, bitSize int) (int64, error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
These functions read s and convert it to a numeral system based on the base argument. You can specify any base between 2 and 36 but you can also use zero to determine the base from the string. If the string contains a “0x” prefix then it’s parsed as base-16, if it contains just a “0” prefix then it’s parsed as base-8, otherwise it’s parsed as decimal.

The bitSize argument restricts the size of the integer parsed. This is important if you need to ensure that your value can fit into a smaller type such as int8, int16, or int32. You can also specify a bitSize of zero to indicate that you want it to fit into the system’s int size (i.e. 32-bit or 64-bit).

## When parsing fails
There are a couple ways that parsing can fail. The most obvious is if your string contains characters outside of the range the numeral system. For example, a “9” is not a valid number when parsing base-8. This will return a NumError with an Err field of ErrSyntax.

Parsing can also fail if the number is too large for the bitSize. For example, parsing the number “300” with a bitSize of 8 is invalid because an int8 has a maximum value of 127. This will return a NumError with an Err field of ErrRange.

## Convenient int parsing
Finally, there’s a convenience function called Atoi():

func Atoi(s string) (int, error)
Internally this is just a call to ParseInt() with a base of 10 and a bitSize of 0. Also, note that it returns an int type instead of ParseInt()’s int64 return type.

I typically use the int type for all my integers in my application unless there’s a specific need to use sized integer types (such as efficiency or to ensure 64-bit range on 32-bit systems). It reduces the clutter of having to type convert between the various sizes. Because I mainly use the int type, I primarily use the Atoi() function.

## Formatting integers
For encoding your Go integer types into a string, you can use FormatInt() and FormatUint():

func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
These functions convert i into the given base and return the string representation. The base supports anything between 2 and 36.

Formatting integers seems simple from the outside but there are a lot of optimizations internally for base-10 as well as any base which is a power of 2.

## Formatting integers (with fewer allocations)
One often overlooked part of the strconv package is its Append functions. The Format functions generally require that the returned variable is allocated each time. Allocations are the enemy of performance.

To remove these allocations for each call, we can reuse a single buffer by using the Append functions. Integer formatting provides the AppendInt() and the AppendUint() functions:
```
func AppendInt(dst []byte, i int64, base int) []byte
func AppendUint(dst []byte, i uint64, base int) []byte
```
Reusing a buffer is simple. In fact, you can create one on the stack if it’s small by using a fixed size array and then converting to a byte slice:


In this example, we have a list of int16 values called a. We can determine the buffer size by taking the number of digits of the maximum value of an int16 (which is 32,767) plus 1 extra byte for a possible negative sign. That’s 5 bytes + 1 byte which makes our buffer 6 bytes. That’s the maximum size an int16 can encode into in base-10.

Now we can loop over our list of values and append into our local buffer. We need to convert our buffer’s byte array to a byte slice so we reslice it using the [:0] notation. This just means that we want to start from the beginning of the slice but make the length zero. The capacity of our slice will be 6 so we can append into it without an allocation.

The returned b variable has a new slice header with the appropriate length set for the formatted integer but it’s data still points to the underlying buf byte array.
