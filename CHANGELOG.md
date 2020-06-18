## Changelog

**0.3.0 - 2020 Jun 18**

- message
	- add DE125 as `*SubMessage` type

- fields
	- rename `B` field type to `B64`, because it represents a 64 length `B`
	- add `BN` as a new field, because it represents an N length `B`
	- add `Reserved`, it throws error in case of usage of reserved fields

```go
/ Usage of submessage in 0.3.0
m := &Message{
	DE125: &SubMessage{
		SE2: NewANS("Test Address"),
	},
}
result, _ := m.Encode() // err handle
```

**0.2.0 - 2020 Jun 17**

- submessage
	- add submessage which is the field representation of DE125
	- encode
	- decode

```go
// Usage of submessage in 0.2.0
sm := &SubMessage{
	SE2: NewANS("Test Address"),
}
b, _ := sm.Encode() // err handle

m := &Message{
	DE125: NewANS(string(b)),
}
result, _ := m.Encode() // err handle
```

**0.1.0 - 2020 Jun 16**

- message
	- definition of iso8583 base structure
	- encode
	- decode
	- bitmap related functions

- fields
	- definition of field types of iso8583
	- regexp to validate fields