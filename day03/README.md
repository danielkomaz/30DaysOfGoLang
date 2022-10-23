# Day 03

Challenge against my son to write a program to output a string in reverse order.

## Takeaways

1. Runes

   - Rune literals are just 32-bit integer values (however they're untyped constants, so their type can change). They represent unicode codepoints. For example, the rune literal 'a' is actually the number 97.
   - Runes are used to split strings into their respective letters/symbols.
   - Since runes represent unicode codepoints, they work with all kinds of symbols (e.g. japanese characters).

2. Function len()

   - This function returs the length value of a string.

3. Variable short declaration

   - Using the short declaration := for variables will declare the variable and assign a value to it in a single step. Using only = to assign a value requires the variable to be already declared.

4. Concatinating strings

   - To concatinating strings, they can simply be "added" to each other by using the + operator. The shorthand += is also a viable option for strings.
