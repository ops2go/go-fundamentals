## Writer interface
The Writer interface is simply the inverse of the Reader. We provide a buffer of bytes to push out onto a stream.
```
type Writer interface {
        Write(p []byte) (n int, err error)
}
```
Generally speaking writing bytes is simpler than reading them. Readers complicate data handling because they allow partial reads, however, partial writes will always return an error.

## Duplicating writes
Sometimes you’ll want to send writes to multiple streams. Perhaps to a log file or to STDERR. This is similar to the TeeReader except that we want to duplicate writes instead of duplicating reads.

The MultiWriter comes in handy in this case:

func MultiWriter(writers ...Writer) Writer
The name is a bit confusing since it’s not the writer version of MultiReader. Whereas MultiReader concatenates several readers into one, the MultiWriter returns a writer that duplicates each write to multiple writers.

I use MultiWriter extensively in unit tests where I need to assert that a service is logging properly:
```
type MyService struct {
        LogOutput io.Writer
}
...
var buf bytes.Buffer
var s MyService
s.LogOutput = io.MultiWriter(&buf, os.Stderr)
```
Using a MultiWriter allows me to verify the contents of buf while also seeing the full log output in my terminal for debugging.

## Optimizing string writes
There are a lot of writers in the standard library that have a WriteString() method which can be used to improve write performance by not requiring an allocation when converting a string to a byte slice. You can take advantage of this optimization by using the io.WriteString() function.

The function is simple. It first checks if the writer implements a WriteString() method and uses it if available. Otherwise it falls back to copying the string to a byte slice and using the Write() method.

(Thanks to Bouke van der Bijl for pointing this one out)
