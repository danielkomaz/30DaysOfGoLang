# Day 05

## Tests

Take a first step into writing tests with Go.
Tutorials followed: <https://github.com/quii/learn-go-with-tests/blob/main/hello-world.md>

## Takeaways

1. Tests are written in a separate file with the same syntax as the main program. This is because testing is built into the language and is enabled by simply importing the `testing` package.
2. Tests are run with the command `go test` in the directory of the test file.
3. To simplify writing the tests own "domain" code should be separated from "foreign" code (side-effects) like `fmt` and `os` from already tested packages. This will help leaving code DRY and make it easier to test.
4. Tests are run in parallel by default.
5. In Go, = and := are used for variable assignment, but they are used in slightly different contexts.

   - =: This is used to assign a value to an already declared variable.
   - :=: This is a shorthand for declaring and initializing a variable in one line. It infers the type of the variable from the assigned value. This operator can only be used inside a function.

   Remember, if you try to use := for a variable that's already been declared in the same scope, you'll get a compiler error. However, if the := statement is declaring at least one new variable, it's valid.

6. With `t.Fail()` we could fail a test. However, by using `t.Errorf()` we can fail a test and print a customized error message.
7. By writing the test first, we can also make sure our naming conventions are correct. For example, we can make sure that the function we are testing is named correctly (including case sensitivity).
8. TDD (Test Driven Development) Discipline:

   1. Write a test
   2. Make the compiler pass
   3. Run the test, see that it fails and check the error message is meaningful
   4. Write enough code to make the test pass
   5. Refactor

9. In Go, public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world.
10. Constants can be grouped in a block instead of declaring them each on their own line. It's a good idea to use a line between sets of related constants for readability.
