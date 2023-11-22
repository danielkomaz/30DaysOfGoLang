# Day 06

## Integers

Learning how to use and return integers in Go.

### Takeaways

1. Go source files can only have one package per directory. Make sure that your files are organised into their own packages.
2. In TestAdder we are using %d as our format strings rather than %q. That's because we want it to print an integer instead of a string.
3. Go examples are executed just like tests so you can be confident examples reflect what the code actually does.
4. Examples are compiled (and optionally executed) as part of a package's test suite.
5. The `Example` function is a special function that gets called by the `go test` command. It takes one argument, a function with no arguments and no return values.
6. The `Output` field/comment of the `Example` struct is used to compare the output of the example with the expected output. If the output doesn't match, the test fails.
