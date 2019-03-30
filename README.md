# readyset

readyset.go - a stupidly simple timer / benchmarker

## Usage

```go
timer := readyset.Go()

lap1 := timer.Lap() # Time elapsed since start

lap2 := timer.Lap() # Time elapsed since lap1

elapsed := time.Stop() # Total time elapsed since start

```
