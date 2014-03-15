# Generics Generator for Go

## Issue:

Currently Go has no support for generics

## Hack:

This provides a temporary hack by providing all possible combinations of common functional list functions for Go Slices

### Currently Implemented:
1. Self (x => x)
2. FoldLeft (list z function => result)

### How To Use:
1. Change generator.go's global variable 'types' to include all the types you need functions for
2. Generated functions will be of the form FunctionType1Type2Type3... (see example code for more info)
2. Build and run, take output and use from personal package
3. Enjoy, Contribute, Be Merry!

### TODO:
 + filter/remove/partition
 + takeWhile/dropWhile/span
 + count
 + forall/exists
 + zip?
 + union?/intersect?

### Examples:

```go
import "generics"

func addStrLen(x int, y string) int {
    return x + len(y)
}

//Format:
//FoldLeftType1Type2 : f(Type2, Type1) -> Type2

func sumStringLengths(l []string) int {
	return generics.FoldLeftstringint(l, 0, addStrLen)
}

func add(x int, y int) int {
	return x + y
}

func sum(l []int) int {
	return generics.FoldLeftintint(l, 0, add)
}

func mul(x int, y int) int {
	return x * y
}

func product(l []int) int {
	return generics.FoldLeftintint(l, 1, mul)
}

func double(x int) int {
	return 2 * x
}

func doubleList(x []int) []int {
	return Mapintint(x, double)
}

func length(x string) int {
	return len(x)
}

func stringsToLengths(x []string) []int {
	return Mapstringint(x, length)
}
```
