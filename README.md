# Go Binary Encoding

Code Generation-based Go Type Serialization Library

## Custom Marshaler/Unmarshaler

```go
type GOBE_CUSTOM_TYPE interface {
	MarshalGOBE(dst []byte) error
	UnmarshalGOBE(src []byte) error
	SizeGOBE() uint64
}
```
