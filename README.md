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
2. Build and run, take output and package for personal use
3. Enjoy, Contribute, Be Merry!

### TODO:
 + map
 + filter/remove/partition
 + takeWhile/dropWhile/span
 + count
 + forall/exists
 + zip?
 + union?/intersect?

### Examples:

```
import "generics"

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
```
