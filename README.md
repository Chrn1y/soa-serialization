# soa-serialization
Benchmark for serializations in Go

```go
Serialization and Deserialization in Go:

encoding/gob serialization
size: 48151 bytes
time: 211080 nanoseconds

encoding/gob deserialization
time: 307697 nanoseconds

json serialization
size: 63569 bytes
time: 527595 nanoseconds

json deserialization
time: 1309239 nanoseconds

xml serialization
size: 179473 bytes
time: 1469420 nanoseconds

xml deserialization
time: 7441962 nanoseconds

proto serialization
size: 54256 bytes
time: 315881 nanoseconds

proto deserialization
time: 523265 nanoseconds

yaml serialization
size: 86817 bytes
time: 5205583 nanoseconds

yaml deserialization
time: 4998681 nanoseconds

messagepack serialization
size: 50357 bytes
time: 689088 nanoseconds

messagepack deserialization
time: 737991 nanoseconds
```