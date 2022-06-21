# readyset

readyset.go - a stupidly simple timer / benchmarker

## Usage

```go
timer := readyset.Go()

lap1 := timer.Lap() # Time elapsed since start

lap2 := timer.Lap() # Time elapsed since lap1

lap3 := timer.Lap().Milliseconds() # Time elapsed since lap2 in ms as an int64

elapsed := time.Stop() # Total time elapsed since start

elapsed2 := time.Stop().Milliseconds() # Total time elapsed since start in ms as an int64

```
