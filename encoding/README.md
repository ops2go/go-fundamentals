# Go Walkthrough: encoding package
So far we’ve covered working with raw byte streams and bounded byte slices but few applications simply shuttle bytes around. Bytes alone don’t convey much meaning, however, once we encode data structures on top of those bytes then we can build truly useful applications.

This post is part of a series of walkthroughs to help you understand the Go standard library better. While generated documentation provides a wealth of information, it can be difficult to understand packages in a real world context. This series aims to provide context of how standard library packages are used in every day applications. If you have questions or comments you can reach me at @benbjohnson on Twitter.

## What is encoding exactly?
In computer science we have fancy words for simple concepts. Not only that but many times there are lots of fancy words for a single concept. Encoding is one of those words. Sometimes it’s referred to as serialization or as marshaling — it means the same thing: adding a logical structure to raw bytes.

In the Go standard library, we use the term encoding and marshaling for two separate but related ideas. An encoder in Go is an object that applies structure to a stream of bytes while marshaling refers to applying structure to bounded, in-memory bytes.

For example, the encoding/json package has a json.Encoder and json.Decoder for working with io.Writer and io.Reader streams, respectively. The package also has json.Marshaler and json.Unmarshaler for writing to and reading from byte slices.

## Two types of encoding
There is also another important distinction in encodings. Some encoding packages operate on primitives — strings, integers, etc. Strings are encoded with character encodings such as ASCII or Unicode or any number of other language specific encodings. Integers can be encoded differently based on endianness or by using variable length encoding. Even bytes themselves are often encoded using schemes like Base64 to convert them into printable characters.

Often when we think of encoding, though, we think of object encoding. This refers to converting complex structures such as structs, maps, and slices into a series of bytes. There are a lot of tradeoffs when doing this conversion and many people have developed different object encoding schemes over the years.

## Making trade-offs
Converting logical structures to bytes seems simple enough at first —these structures are already represented in-memory as bytes internally. Why not just use that format?

There’s a lot of reasons why Go’s in-memory format isn’t suitable for converting to bytes and saving to disk or sending over the network. First is compatibility. Go’s internal data structure format doesn’t match Java’s internal format so we can’t communicate between these different systems. Sometimes we need compatibility not with another programming language but with humans. CSV, JSON, and XML are all human-readable formats that can be easily viewed and edited.

Making formats human-readable introduces a trade-off though. Formats that are easy for humans to parse are slower for computers to parse. Integers are a good example — people read in base-10 format whereas computers operate in base-2. People also read variable length numbers such as 1 or 1,000 but computers operate on fixed-sized numbers such as 32-bit or 64-bit integers. The performance difference may seem trivial for a single number but it quickly becomes a big deal when parsing millions or billions of numbers.

There’s also other trade-offs we don’t think of at first. Our data structures change over time but we still need to operate on bytes that may have been encoded years ago. Some encodings, such as Protocol Buffers, allow you to write a schema for your data and version your fields — older fields can be deprecated while new fields can be added. The downside of this is that you need the schema definition in order to encode and decode objects. Go’s own gob format takes a different approach and actually includes the schema format when encoding. However, the downside of this approach is that the encoded size can be much larger.

Some formats throw caution to the wind entirely and go schema-less. JSON and MessagePack both allow you to encode structures on the fly but provide no guarantees about safely decoding structures from an older format.

We also use systems that do encoding for us but we don’t think of as encoding. Databases, for example, are a roundabout way of taking our logical data structures and eventually persisting them as bytes on disk. It may involve network calls, SQL parsing, and query planning but it’s all essentially encoding.

Finally, if you really need speed above all else, you could use Go’s internal format to save data. I even wrote a library for this called raw. It’s encoding and decoding time is literally zero seconds. Should you use it in production? Probably not.

The 4 interfaces of encoding
If you are one of the few people who has ever looked at the encoding package, you may have been underwhelmed. It is the second smallest package after the errors package and it only includes 4 interfaces.

The first two interfaces are BinaryMarshaler and BinaryUnmarshaler:

type BinaryMarshaler interface {
        MarshalBinary() (data []byte, err error)
}
type BinaryUnmarshaler interface {
        UnmarshalBinary(data []byte) error
}
These are for objects that provide a way to convert to and from a binary format. This is used in a few spots in the standard library such as time.Time.MarshalBinary(). You don’t find it more places because there’s not usually a single defined way to marshal an object to binary format. As we’ve seen, there are a multitude of serialization formats.

At the application level, however, you have probably picked a single format for marshaling. For instance, you may have chosen Protocol Buffers for all your data. There’s is typically no reason to support multiple binary formats for your application data so implementing BinaryMarshaler can make sense.

The next two interfaces are TextMarshaler and TextUnmarshaler:

type TextMarshaler interface {
        MarshalText() (text []byte, err error)
}
type TextUnmarshaler interface {
        UnmarshalText(text []byte) error
}
These two interfaces are similar to the binary marshaling interfaces except that they output in a UTF-8 format.

Some formats have their own marshaling interfaces, such as json.Marshaler, which follow the same naming style.

Overview of encoding packages
There are a lot of useful encoding packages baked into the standard library. We’ll cover these in more detail in future posts but I’d like to give an overview first. Some of these are subpackages of encoding while others are scattered in different locations.

## Primitive encodings
The first package you probably used when you started with Go is the fmt package (pronounced “fumpt”). It uses C-style printf() conventions to encode and decode numbers, strings, bytes, and even includes limited support for object encoding. The fmt package is a great, simple way to build human-readable strings from templates but the template parsing can add overhead.

If you need better performance then you can avoid templating by using the string conversion package — strconv. This low-level package provides basic formatting and scanning for strings, integers, floats, and booleans and is generally pretty fast.

These packages, along with Go itself, assume that you’re encoding strings using UTF-8. The near total lack of non-Unicode character encoding support in the standard library could be because the Internet has quickly converged on a standard of UTF-8 over the last several years or it could be because Rob Pike is a coauthor of Go & UTF-8. Who knows? I’ve been lucky enough to not have to deal with any non UTF-8 encodings in Go so far, however, there is some encoding support in unicode/utf16, encoding/ascii85, and the golang.org/x/text package tree. The “x” package tree contains a wealth of awesome packages that are part of the Go project but are not covered under the Go 1 compatibility requirements.

For integer encoding, the encoding/binary package provides big endian and little endian encodings as well as variable length encodings. Endianness refers to the order that bytes are written to disk. For example, the uint16 representation of 1,000 (which is 0x03E8 in hex) is composed of 2 bytes: 03 & E8. With big endian encoding, the bytes are written in that order “03 E8”. In little endian, the order is reversed: “E8 03”. Many common CPUs architectures use little endian. However, big endian is typically used when sending bytes over the network. Big endian is even called network byte order.

Finally, for byte encoding there are a couple packages available. Byte encoding is typically used to convert bytes into a printable format. The encoding/hex package, for example, can be used if you need to view binary data in hexidecimal format. I’ve personally only used it for debugging purposes. On the other hand, sometimes you need a printable format because you need to transport data over protocols with historically limited binary support (such as email). The encoding/base32 and encoding/base64 packages are an example of this. Another example is the encoding/pem package which is used for encoding TLS certificates.

## Object encodings
We find fewer packages within the standard library for object encodings. However, in practice, these packages are many times all we need.

In case you’ve been living under a rock for the past decade, you’ve probably noticed that JSON has become the default object encoding of the Internet. As mentioned above, JSON has its flaws but it’s easy to use and it has library support in every language so adoption has skyrocketed. The encoding/json package provides great support for this protocol and there are also third party implementations for generating faster parsers such as ffjson.

While JSON has dominated as a protocol between machines, the CSV format is a more common protocol for exporting data to humans. The encoding/csv package provides a good interface for exporting tabular data in this format.

If you’re interacting with a system built circa 2000 then you probably need to use XML. The encoding/xml package provides a SAX-style interface with an additional tag-driven marshaler/unmarshaler that’s similar to the json package. If you’re looking for more complex features like DOM, XPath, XSD, or XSLT then you should probably use libxml2 via cgo.

Go also has its own stream encoding called gob. This package is used by the net/rpc package for implementing a remote procedure call interface between two Go services. Gob is easy to use, however, it does not have any cross language support. gRPC seems to be a popular alternative if you need to communicate between different languages.

Finally, there’s a package called encoding/asn1. There’s limited information in the documentation and the only link in the package points to a layman’s guide to ASN.1 which is a 25 page wall of text. ASN.1 is a complex object encoding scheme that is most notably used by X.509 certificates in SSL/TLS.

## Summary
Encoding provides the fundamental basis for layering information on top of our bytes. Without it we wouldn’t have strings or data structures or databases or any useful applications. What seems like a relatively simple concept has a rich history of implementations and a wide variety of tradeoffs.

In this post we looked at an overview of the different encoding implementations within the standard library and some of their tradeoffs. We saw how these primitive and object encoding packages built on our knowledge of byte streams and slices. In the next several posts we’ll take a deeper dive into these packages to see how to use them in a real world context.