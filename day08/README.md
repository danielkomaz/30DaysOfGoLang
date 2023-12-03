# Day 08

## Arrays and Slices

Learning about arrays and slices in Go.

## Takeaways

1. Arrays are fixed length, slices are not.
2. We can initialize an array by using `var myArray [5]int` where 5 is the length.
3. We can initialize an array by using `myArray := [...]int{1, 2, 3, 4, 5}` where 1, 2, 3, 4, 5 are the values.
4. An interesting property of arrays is that the size is encoded in its type. If we try to pass an `[4]int` into a function that expects `[5]int`, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an `int`.
5. We can create a slice from an array by using `mySlice := make([]int, 5, 10)` where 5 is the length and 10 is the capacity.
6. We can create a slice from an array by using `mySlice := []int{1, 2, 3, 4, 5}` where 1, 2, 3, 4, 5 are the values.
7. Go does not let us use equality operators with slices. We could write a function to iterate over each got and want slice and check their values but for convenience sake, we can use `reflect.DeepEqual` which is useful for seeing if any two variables are the same.
8. `make` allows you to create a slice with a starting capacity of the `len` of the `numbersToSum` we need to work through.
9. We can index slices like arrays with `mySlice[N]` to get the value out or assign it a new value with `=`. We can also use `mySlice[N:M]` to get a slice of the slice.
10. We can append to a slice with `mySlice = append(mySlice, 6)` where 6 is the value you want to append.
11. Slices can be sliced! The syntax is slice[low:high]. If we omit the value on one of the sides of the : it captures everything to that side of it. In our case, we are saying "take from 1 to the end" with numbers[1:].
12. We can assign functions to variables in Go. This is useful for making our code more readable and reusable. By defining this function inside the test, it cannot be used by other functions in this package. Hiding variables and functions that don't need to be exported is an important design consideration.
13. A handy side-effect of this is this adds a little type-safety to our code. If a developer mistakenly adds a new test with checkSums(t, got, "dave") the compiler will stop them in their tracks.
