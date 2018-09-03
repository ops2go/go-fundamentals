## Closing streams
All good things must come to an end and this is no exception when working with byte streams. The Closer interface is provided as a generic way to close streams:
```
type Closer interface {
        Close() error
}
```
There’s not much to say about Closer since it is so simple, however, I find it useful to always return an error from my Close() functions so that my types can implement Closer when it’s required. Closer isn’t always used directly but is sometimes combined with other interfaces the case of ReadCloser, WriteCloser, and ReadWriteCloser.