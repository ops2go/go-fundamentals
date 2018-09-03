# Go Walkthrough: encoding/json package
For better or for worse, JSON is the encoding of the Internet. Its formal definition is small enough that you could write it on the back of a napkin but yet it can encode strings, numbers, booleans, nulls, maps, and arrays. Because of this simplicity, every language has a JSON parser.

Go’s implementation is a package called encoding/json and it allows us to seamlessly add JSON encoding to our Go objects. However, with its extensive use of reflection, encoding/json is one of the least understood packages even though it’s also one of the most used. We’ll take a deep look at how this package works — not just as a user of the package — but also look at how its internals function.

This post is part of a series of walkthroughs to help you understand the standard library better. While generated documentation provides a wealth of information, it can be difficult to understand packages in a real world context. This series aims to provide context of how standard library packages are used in every day applications. If you have questions or comments you can reach me at @benbjohnson on Twitter.

## What is JSON?
JSON stands for JavaScript Object Notation and it is exactly that — the subset of the JavaScript language that defines object literals. JavaScript lacks static, declared typing so language literals needed to have implicit types. Strings are wrapped in double quotes, arrays are wrapped in brackets, and maps are wrapped in curly braces:

{"name": "mary", "friends: ["stu", "becky"], age: 30}
While this loose type information is cursed by many developers writing JavaScript, it provides an extremely easy and concise way of representing data.

## Tradeoffs of using JSON
Although JSON is easy to start using, you can run into some issues. Formats which are easy for humans to read are typically slow for computers to parse. For example, running the encoding/json benchmarks on my MacBook Pro show encoding and decoding speed at 100 MB/sec and 27 MB/sec, respectively:
```
$ go test -bench=. encoding/json
BenchmarkCodeEncoder-4     106.26 MB/s
BenchmarkCodeDecoder-4     27.76 MB/s
```
A binary decoder, however, can typically parse data at many times that speed. This problem occurs because of how data is read. A JSON number literal such as “123.45” has to be decoded in two iterative steps:

Read each byte to inspect if it is a number or a dot. If a non-number is read then the we are done scanning the number literal.
Convert the base-10 number literal into a base-2 format such as int64 or IEEE-754 floating point number representation.
This involves a lot of parsing for every incoming byte as well as a lookahead buffer on the decoder. In contrast, a binary decoder simply needs to know how many bytes to read (e.g. 2, 4, or 8) and possibly flip the endianness. These binary parsing operations also involve no branching which slows down CPU pipelining.

## When should you use JSON?
Typically JSON is used when ease-of-use is the primary goal of data interchange and performance is a low priority. Because JSON is human-readable, it is easy to debug if something breaks. Binary protocols, on the other hand, have to be decoded first before they can be analyzed.

In many applications, encoding/decoding performance is a low priority because it can easily be scaled horizontally. For example, adding additional servers to serve API endpoints is usually trivial because encoding requires no coordination with other servers. Your database, however, likely doesn’t scale as easily once you need to add servers.



## Alternative implementations
Several years ago I implemented a tool called megajson which generated type-specific encoders and decoders at compile time in order to avoid reflection entirely. This made encoding and decoding much faster. However, the tool was a proof of concept, had limited support, and was eventually abandoned.

Luckily, Paul Querna made an implementation called ffjson which does the same thing but does it much better. If you need to improve your JSON encoding and decoding performance then I highly suggest looking at his implementation.

## Summary
JSON can be a great data format when you need to get up and running quickly or you need to provide a simple API to users. Go’s implementation provides a lot of features to make it simple to use through the use of reflection.

We’ve looked at the internals of the encoding and decoding side of JSON as well as seen how we can format our JSON representation. These tools may seem simple from the outside but there’s a lot of fascinating stuff happening internally to make them as fast and efficient as possible.