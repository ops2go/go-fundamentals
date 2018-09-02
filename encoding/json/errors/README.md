
## Handling errors during encoding/decoding
The json package has quite a few error types within it. Here’s a list of what can go wrong while you’re encoding or decoding:

If you pass in a non-pointer value to be decoded then you are actually passing a copy of your value and the decoder can’t decode into the original value. The decoder catches this and returns an InvalidUnmarshalError.
If your data contains an invalid JSON then a SyntaxError will be returned with the byte position of the invalid character.
If an error is returned by a json.Marshaler or encoding.TextMarshaler then it will be wrapped in a MarshalerError.
If a token cannot be unmarshaled into a corresponding value then an UnmarshalTypeError is returned.
The float values of Infinity and NaN are not representable in JSON and will return an UnsupportedValueError.
Types which cannot be represented in JSON (such as functions, complex numbers, pointers, etc) will return an UnsupportedTypeError.
Before Go 1.2, invalid UTF-8 characters would cause an InvalidUTF8Error to be returned. Later versions simply convert invalid characters to U+FFFD which is the Unicode character for an “unknown character”.
While this might seem like a lot of errors, there’s not much you can do to handle them in your code other than log an error and have a human operator intervene. Also, many of them can be caught at development-time if you have unit test coverage.
