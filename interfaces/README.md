### HTTP Get (Reader Interface) - explanation
```go
bs := make([]byte, 99999)
response.Body.Read(bs)
```

- response.Body is of type io.ReadCloser - a coposite interface of Reader and Closer
- This ensures whatever the implementation is, be it image, string, file, Read and Close function will be present.
- Read function takes in a byte slice and modifies the byte slice with the result.

### io.Copy (Writer Interface) - explanation
```go
io.Copy(os.Stdout, response.Body)
```
- Copy function takes in (Writer, Reader).
- Writer uses the given byte to write to the type that contains the method.

