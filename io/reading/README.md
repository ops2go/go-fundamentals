## Reading bytes
There are two fundamental operations when working with bytes: reading & writing. Let’s take a look at reading bytes first.

## Reader interface
The basic construct for reading bytes from a stream is the Reader interface:
```
type Reader interface {
        Read(p []byte) (n int, err error)
}
```
This interface is implemented throughout the standard library by everything from network connections to files to wrappers for in-memory slices.

The Reader works by passing a buffer, p, to the Read() method so that we can reuse the same bytes. If Read() returned a byte slice instead of accepting one as an argument then the reader would have to allocate a new byte slice on every Read() call. That would wreak havoc on the garbage collector.

One problem with the Reader interface is that it comes with some subtle rules. First, it returns an io.EOF error as a normal part of usage when the stream is done. This can be confusing for beginners. Second, your buffer isn’t guaranteed to be filled. If you pass an 8-byte slice you could receive anywhere between 0 and 8 bytes back. Handling partial reads can be messy and error prone. Fortunately there are helpers functions for this problem.

## Improving reader guarantees
Let’s say you have a protocol you’re parsing and you know you need to read an 8-byte uint64 value from a reader. In this case it’s preferable to use io.ReadFull() since you have a fixed size read:

func ReadFull(r Reader, buf []byte) (n int, err error)
This function ensures that your buffer is completely filled with data before returning. If your buffer is partially read then you’ll receive an io.ErrUnexpectedEOF back. If no bytes are read then an io.EOF is returned. This simple guarantee simplifies your code tremendously. To read 8 bytes you only need to do this:
```
buf := make([]byte, 8)
if _, err := io.ReadFull(r, buf); err == io.EOF {
        return io.ErrUnexpectedEOF
} else if err != nil {
        return err
}
```
There are also many higher level parsers such as binary.Read() which handle parsing specific types. We’ll cover those in future walkthroughs within different packages.

Another lesser used helper function is ReadAtLeast():

func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
This function will read additional data into your buffer if it is available but will always return a minimum number of bytes. I haven’t found a need for this function personally but I can see it being useful if you need to minimize Read() calls and you’re willing to buffer additional data.

## Concatenating streams
Many times you’ll encounter instances where you need to combine multiple readers together. You can combine these into a single reader by using the MultiReader:
```
func MultiReader(readers ...Reader) Reader
```
For example, you may be sending a HTTP request body that combines an in-memory header with data that’s on-disk. Many people will try to copy the header and file into an in-memory buffer but that’s slow and can use a lot of memory.

Here’s a simpler approach:
```
r := io.MultiReader(
        bytes.NewReader([]byte("...my header...")),
        myFile,
)
http.Post("http://example.com", "application/octet-stream", r)
```
The MultiReader let’s the http.Post() consider the two readers as one single concatenated reader.

## Duplicating streams
One issue you may run across when using readers is that once a reader is read, the data cannot be reread. For example, your application may fail to parse an HTTP request body and you’re unable to debug the issue because the parser has already consumed the data.

The TeeReader is a great option for capturing the reader’s data while not interfering with the consumer of the reader.
```
func TeeReader(r Reader, w Writer) Reader
```
This function constructs a new reader that wraps your reader, r. Any reads from the new reader will also get written to w. This writer can be anything from an in-memory buffer to a log file to STDERR.

For example, you can capture bad requests like this:
```
var buf bytes.Buffer
body := io.TeeReader(req.Body, &buf)
// ... process body ...
if err != nil {
        // inspect buf
        return err
}
```
However, it’s important that you restrict the request body that you’re capturing so that you don’t run out of memory.

## Restricting stream length
Because streams are unbounded they can cause memory or disk issues in some scenarios. The most common example is a file upload endpoint. Endpoints typically have size restrictions to prevent the disk from filling, however, it can be tedious to implement this by hand.

The LimitReader provides this functionality by producing a wrapping reader that restricts the total number of bytes read:

func LimitReader(r Reader, n int64) Reader
One issue with LimitReader is that it won’t tell you if your underlying reader exceeds n. It will simply return io.EOF once n bytes are read from r. One trick you can use is to set the limit to n+1 and then check if you’ve read more than n bytes at the end.

Writing bytes
Now that we’ve covered reading bytes from streams let’s look at how to write them to streams.
