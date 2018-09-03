# Go Walkthrough: io package
Go is a programming language built for working with bytes. Whether you have lists of bytes, streams of bytes, or individual bytes, Go makes it easy to process. From these simple primitives we build our abstractions and services.

The io package is one of the most fundamental packages within the standard library. It provides a set of interfaces and helpers for working with streams of bytes.

This post is part of a series of walkthroughs to help you understand the standard library better. While generated documentation provides a wealth of information, it can be difficult to understand packages in a real world context. This series aims to provide context of how standard library packages are used in every day applications. If you have questions or comments you can reach me at @benbjohnson on Twitter.

## Summary
Byte streams are essential to most Go programs. They are the interface to everything from network connections to files on disk to user input from the keyboard. The io package provides the basis for all these interactions.

We’ve looked at reading bytes, writing bytes, copying bytes, and finally looked at optimizing these operations. These primitives may seem simple but they provide the building blocks for all data-intensive applications. Please take a look at the io package and consider its interfaces in your application.