## Moving around within streams
Streams are usually a continuous flow of bytes from beginning to end but there are a few exceptions. A file, for example, can be operated on as a stream but you can also jump to a specific position within the file.

The Seeker interface is provided to jump around in a stream:
```
type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
}
```
There are 3 ways to jump around: move from on the current position, move from the beginning, and move from the end. You specify the mode of movement using the whence argument. The offset argument specifies how many bytes to move by.

Seeking can be useful if you are using fixed length blocks in a file or if your file contains an index of offsets. Sometimes this data is stored in the header so moving from the beginning makes sense but sometimes this data is specified in a trailer so you’ll need to move from the end.
