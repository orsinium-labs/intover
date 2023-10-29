# intover

Golang package for integer arithmetic with overflow detection.

Based on [overflow](https://github.com/JohnCGriffin/overflow) package.

Features:

* Generics.
* Small and simple API.
* Can return ok or can panic.

## Installation

```bash
go get github.com/orsinium-labs/intover
```

## Usage

```go
result, ok := intover.Do(a, '+', b)
```

Or the same:

```go
result, ok := intover.Add(a, b)
```

Or if you want it to panic on overflow:

```go
result := intover.Must(intover.Add(a, b))
```
