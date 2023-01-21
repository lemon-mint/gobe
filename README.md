# GOBE: Go Binary Encoding

Code Generation based Go Type Serialization Library

**Note: This library is still in active development.**

**Note: Circular reference is not supported.**

**Warning: UnmarshalGOBE() may cause a DoS vulnerability, do not use it with untrusted data.**

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
	MarshalGOBE(dst []byte) uint64
	UnmarshalGOBE(src []byte) (offset uint64, ok bool)
	SizeGOBE() uint64
}
```
