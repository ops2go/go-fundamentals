## Encoding streams
The json package has two ways to encode values to JSON. First is the stream-based json.Encoder which will encode the value to an io.Writer:
```
type Encoder struct {}
func NewEncoder(w io.Writer) *Encoder
func (enc *Encoder) Encode(v interface{}) error
```
The second option is json.Marshal() which will return an in-memory byte slice of your encoded value:
```
func Marshal(v interface{}) ([]byte, error)
```
When you pass a value into these encoders, the underlying JSON library goes through a complex process of inspecting type definitions, compiling encoders, and recursively processing values in your data. Let’s look at each of these in detail.

## Type inspection
When you pass a value into the encoder, the first step is to look up the value’s type encoder. Types are inspected using Go’s reflect package and the json package holds an internal mapping of these reflect.Type values. For built-in types such as int, string, map, struct, and slice, there are hardcoded implementations in the json package. These are fairly simple — the stringEncoder wraps string values in double quotes and escapes characters as needed, the intEncoder converts integers to a string format, etc.

Note: the use of the reflect library in Go is a touchy subject. On one hand it makes generic runtime encoders such as encoding/json possible and on the other hand it can be abused by developers who use it instead of using static type checking constructs. I find the use of reflect at the application level to generally be a poor choice.

## Encoder compilation
For types that are not built-in, an encoder is built on the fly and then cached for reuse. First, the encoder will check if the type implements json.Marshaler:
```
type Marshaler interface {
        MarshalJSON() ([]byte, error)
}
```
If it does then the marshaling is deferred to the type. This is really useful if one of your types has a special JSON representation and shouldn’t be handled by the json package’s reflection-based encoder.

Next, the encoder checks if a type implements encoding.TextMarshaler:
```
type TextMarshaler interface {
        MarshalText() (text []byte, err error)
}
```
If it does then it will generate a value from that function and then encode the result as a JSON string. You see this all the time when using time.Time. Because time.Time has a MarshalText() method, the JSON encoder will automatically encode a time.Time value as a RFC 3339 formatted string.

Finally, if neither interface is implemented then it will recursively build an encoder based on the primitive encoders. For example, a type which consists of a struct which contains an int field and a string field will generate a structEncoder which has an intEncoder and stringEncoder inside. Again, building this encoder only happens once and the resulting encoder will be cached for future use.

## Per-field options
One important note about the struct encoder is that it reads field tags to determine per-field options for encoding. Tags are the backticked strings you sometimes see at the end of structs.

For example:
```
type User struct {
        Name    string `json:"name"`
        Age     int    `json:"age,omitempty"`
        Zipcode int    `json:"zipcode,string"`
}
```
These options include:

Renaming the field’s key. A lot of JSON keys are camel cased so it can be important to change the name to match.
The omitempty flag can be set which will remove any non-struct fields which have an empty value.
The string flag can be used to force a field to encode as a string. For example, forcing an int field to be encoded as a quoted string.
Recursive processing
Finally, when encoding is performed it is written to an internal buffer called encodeState. This object is handed off to each encoder your value needs so that the encoder can append its bytes. For the json.Marshal() call, a reference to this buffer’s bytes is returned.

When using a json.Encoder, a sync.Pool is used internally to reuse these encodeState buffers. That helps minimize the number of heap allocations required by the encoder so for stream processing always use a json.Encoder.
