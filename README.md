# GOBE: Go Binary Encoding

Code Generation-based Go Type Serialization Library

## Custom Marshaler/Unmarshaler

```go
type GOBE_CUSTOM_TYPE interface {
	MarshalGOBE(dst []byte) uint64
	UnmarshalGOBE(src []byte) (offset uint64, ok bool)
	SizeGOBE() uint64
}
```
