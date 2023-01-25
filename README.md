# GOBE: Go Binary Encoding

Code Generation based Go Type Serialization Library

**Note: GOBE is still under active development.**

**Note: Circular reference is not supported.**

# Usage

```go
buffer := make([]byte, obj.SizeGOBE())
obj.MarshalGOBE(buffer)
```

# Installation

```bash
go install github.com/lemon-mint/gobe@latest
```

# Custom Marshaler/Unmarshaler

```go
type GOBE_CUSTOM_TYPE interface {
	ZZMarshalGOBE(dst []byte) uint64
	ZZUnmarshalGOBE(src []byte) (offset uint64, ok bool)
	ZZSizeGOBE() uint64
}
```
