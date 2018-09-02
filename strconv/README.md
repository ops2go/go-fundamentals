# Go Walkthrough: strconv
Formatting & parsing primitive values in Go is a common task. You probably first dipped your toes into the fmt package when you started Go, however, there’s a less commonly used package for basic formatting that’s more efficient and preserves compiler type checking.

The strconv package is built for speed. It’s great for when you need to handle primitive value formatting while minimizing allocations and CPU cycles. Understanding the package also gives you a better understanding of how the fmt package itself works.

This post is part of a series of walkthroughs to help you understand the standard library better. While generated documentation provides a wealth of information, it can be difficult to understand packages in a real world context. This series aims to provide context of how standard library packages are used in every day applications. If you have questions or comments you can reach me at @benbjohnson on Twitter.

## The primitive types
There are 4 different types of primitives that strconv works with — booleans, integers, floating-point numbers, & strings. Each type has functions for formatting & parsing.

Some of these have variations for reducing allocations and some have helper functions for common options. Let’s take a look at them one by one.

## Summary
Converting Go’s boolean, numeric, and string types into human readable strings is a core component of most software. We need to see our data! While fmt is the go-to package for formatting, it can be slow and inefficient.

The strconv package gives us a way to format our primitives quickly and efficiently while providing some basic formatting options. It also preserves strong type checking for its arguments whereas fmt frequently uses interface{}.