## Copying bytes
Now that we can read bytes and we can write bytes, it only makes sense that we’d want to plug those two sides together and copy between readers and writers.

## Connecting readers & writers
The most basic way to copy a reader to a writer is the aptly named Copy() function:
```
func Copy(dst Writer, src Reader) (written int64, err error)
```
This function uses a 32KB buffer to read from src and then write to dst. If any error besides io.EOF occurs in the read or write then the copy is stopped and the error is returned.

One issue with Copy() is that you cannot guarantee a maximum number of bytes. For example, you may want copy a log file up to its current file size. If the log continues to grow during your copy then you’ll end up with more bytes than expected. In this case you can use the CopyN() function to specify an exact number of bytes to be written:

func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
Another issue with Copy() is that it requires an allocation for the 32KB buffer on every call. If you are performing a lot of copies then you can reuse your own buffer by using CopyBuffer() instead:

func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
I haven’t found the overhead of Copy() to be very high so I personally don’t use CopyBuffer().

## Optimizing copy
To avoid using an intermediate buffer entirely, types can implement interfaces to read and write directly. When implemented, the Copy() function will avoid the intermediate buffer and use these implementations directly.

The WriterTo interface is available for types that want to write their data out directly:
```
type WriterTo interface {
        WriteTo(w Writer) (n int64, err error)
}
```
I’ve used this in BoltDB’s Tx.WriteTo() which allows users to snapshot the database from a transaction.

On the read side, the ReaderFrom allows a type to directly read data from a reader:
```
type ReaderFrom interface {
        ReadFrom(r Reader) (n int64, err error)
}```
Adapting reader & writers
Sometimes you’ll find that you have a function that accepts a Reader but all you have is a Writer. Perhaps you need to write out data dynamically to an HTTP request but http.NewRequest() only accepts a Reader.

You can invert a writer by using io.Pipe():
```
func Pipe() (*PipeReader, *PipeWriter)
```
This provides you with a new reader and writer. Any writes to the new PipeWriter will go to the PipeReader.

I rarely use this functionality directly, however, the exec.Cmd uses this for implementing Stdin, Stdout, and Stderr pipes which can be really useful when working with command execution.
