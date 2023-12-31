# VSCode issue with multiple projects

'gopls requires a module at the root of your workspace.
You can work with multiple modules by opening each one as a workspace folder.
Improvements to this workflow will be coming soon (https://github.com/golang/go/issues/32394),
and you can learn more here: https://github.com/golang/go/issues/36899.'

To solve this:

1. Create module using `go mod init module_name` 
2. Create a `go.work` file using `go work init`. Add the module to go.work in the parent directory

Maybe learn about '...':

```go
// findAllString with -1 says get all in a slice
// append takes two slices and merges them
// ... spreads them before append puts them back into the slice.
// In this context, the ellipsis is used to expand the second slice argument in the append function. 
// Since append takes a variable number of arguments, the ellipsis is necessary when appending a slice to another slice.
// equivalent of append([1,2,3], 4, 5, 6)
matches := append(reWords.FindAllString(myReverseSubstr, -1), reDigits.FindAllString(myReverseSubstr, -1)...)

```

## Problem day 1

run using

`go run main.go part2.go 1` or `go run main.go part2.go 2`

depending on which problem you would like to run the code for.

## Concurrency

* Remember the closure loop vars that cause inconsistent concurrent behaviour

# Process


## Days 

Dys 1-3. Just get it done. Don't care about efficiency, just learn the language a bit. Use concurrency

Days 4-7: Start to think more abou data structure choice. Enhance readability and lookup times etc. Good example for readability day 7. Also think about efficiency using algorithms

# Things learnt

Algorithms!

Sorting(day 7), taking minima and maxim (day 5), binary search (day 5), 