# Day 07

## Iteration

Learning about iteration and benchmarking in Go.

## Replace

Using an existing function to deepening my understanding of how to use Go with testing, benchmarking and documentation.

### Takeaways

1. In Go there are no while, do, until keywords, you can only use for.
2. `for` without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function.
3. You can also `continue` to the next iteration of the loop.
4. `for` with a single condition will loop repeatedly until that condition evaluates to false.
5. += is the same as `x = x + y`
6. Benchmarking is also something that is built into Go. You can run `go test -bench=.` to run benchmarks in your test files.
7. Benchmarking is a great way to see if you have made your code faster or slower.
8. It is good practice to research existing packages to see if someone has already solved your problem. You can search for packages on <https://pkg.go.dev/>
9. Researching existing packages is a great way to learn more about Go.
