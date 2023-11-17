# Day 2

Include an external package and call a function.

## Takeaways

1. To add a modules requirements (dependencies) and sums (authentication files), we need to run the following command after initializing the module:  
   `go mod tidy`

2. Structuring a function splits the defintion line into the following parts:  
   e.g. `func Hello(name string) string`

   - func - Declares the following block to be code of a function
   - Hello - Name of the function and how the function is called from other code parts
   - name - Name of the parameter given to the function when called
   - string - Type (Data type) of the Parameter
   - string - Type (Data type) of the return value

3. Module imports can be overwritten with version and path like in this example:

   ```Go
   require "greeting" v0.0.0
   replace "greeting" => "../greeting"
   ```
