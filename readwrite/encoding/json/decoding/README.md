## Decoding streams
Converting JSON-encoded bytes back into objects is sort of like reversing the process of encoding but with some important differences.

There are two ways to decode JSON from bytes. First is the stream-based json.Decoder which allows you decode from an io.Reader:
```
type Decoder struct {}
func NewDecoder(r io.Reader) *Decoder
func (dec *Decoder) Decode(v interface{}) error
```
Alternatively you can decode from a byte slice by using the json.Unmarshal() function:
```
func Unmarshal(data []byte, v interface{}) error
```
These decoders work in two parts: first a scanner tokenizes the input bytes and then a decodeState converts the tokens to Go objects.

## Scanning JSON
The scanner is an internal state machine used for parsing JSON. It operates in several steps. First, it checks the first byte of a value to determine the type of token to parse. If it is a “{“ then it needs to parse an object, if it’s a “[“ then it needs to parse an array. This works for simple values too. Double quotes indicate the start of a string, a “t” or an “f” indicate the start of a boolean value, and a 0–9 indicate the beginning of a number.

Once it determines the type of scanning to be done, it hands off to a type-specific function — e.g. string scan, number scan, etc. For complex objects such as maps and arrays, a stack is used to keep track of closing braces.

## Lookahead buffers
An interesting aside of scanning is the lookahead buffer. JSON is “LL(1) parseable” which means that it requires only a single byte buffer while scanning. This buffer is used to peek ahead at the next byte.

For example, the number scanning function will continue to read bytes until it finds a non-numeric character. However, since the character is already read from the stream we need to push it back on a buffer for the next scanning function to use. This is what the lookahead buffer is for.

If you’re interested in learning more about writing parsers, I wrote a Gopher Academy post called Handwriting Parsers & Lexers in Go.

### Decoding tokens
Once tokens are scanned they need to be interpreted. This is the job of the decodeState. In this phase the input value to be decoded into will be matched against each token to be processed.

For example, if you pass in a struct type then the decoder will expect to see a “{“ token. Any other token will cause the decoding to return an error. This phase of matching tokens to values involves heavy use of the reflect package, however, these decoders are not cached so reflection has to be redone on every decode.

You can also process tokens as a stream using the Decoder.Token() and Decoder.More() methods. Admittedly, I haven’t used these methods personally but it’s good to know they are available.

## Custom unmarshaling
Just like the encoding, you can specify custom implementations during decoding. The decoder will first check if a type implements json.Unmarshaler:

type Unmarshaler interface {
        UnmarshalJSON([]byte) error
}
This allows your type to receive the entire JSON value for a type and parse it itself. This can be useful if you want to write your own optimize implementation.

Next the decoder checks if the value type implements encoding.TextUnmarshaler:

type TextUnmarshaler interface {
        UnmarshalText(text []byte) error
}
This is useful if you have a string representation of your type that you want to use. One example of this is for enum types that are represented as integers internally but you want to encode/decode them as strings.

## Deferred processing
An alternative to json.Unmarshaler is the json.RawMessage type. With RawMessage, the raw JSON representation will be saved to a field that you can process after unmarshaling is complete. This can be used if you need to interpret a “type” field in your JSON object and then change JSON parsing based on the value.
```
type T struct {
        Type  string          `json:"type"`
        Value json.RawMessage `json:"value"`
}
func (t *T) Val() (interface{}, error) {
        switch t.Type {
        case "foo":
                // parse "t.Value" as Foo
        case "bar":
                // parse "t.Value" as Bar
        default:
                return nil, errors.New("invalid type")
       }
}
```
I personally find json.Unmarshaler more useful though because I don’t like to save off JSON for later interpretation.

Another way to defer processing is with JSON numbers. Because JSON doesn’t distinguish between integers and floats, the decoder will convert numbers into float64 when decoding into an interface{} field. To defer the parsing you can use the json.Number type instead.
```
type T struct {
        Value json.Number
}

if strings.Contains(t.Value, ".") {
        v, err := t.Value.Float64()
        // process as a float
} else {
        v, err := t.Value.Int64()
        // process as an integer
}
```
I don’t use json.Number too often since I usually use static types during decoding.

## Pretty printing
JSON is typically written out as one long set of bytes with no extra whitespace, however, that’s hard to read. You can set indentation in two ways. If you have an in-memory JSON-encoded byte slice then you can pass it to the json.Indent() function:

func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
The prefix argument specifies a character to write out on every line and the indent specifies the characters used for indentation. I don’t typically use the prefix but I’ll usually use a two-space or tab indent value.

There’s a helper function called json.MarshalIndent() which literally just calls json.Marshal() and then json.Indent().

If you are using the stream-based json.Encoder then you can set the indentation on it using the SetIndent() method:
```
func (enc *Encoder) SetIndent(prefix, indent string)
```
I find that many people don’t know about SetIndent() and will marshal & indent to a byte slice and then write that result to a stream.

The reverse of the indent functions is the Compact() function:

func Compact(dst *bytes.Buffer, src []byte) error
This will rewrite src to a destination buffer but remove all extraneous whitespace.
