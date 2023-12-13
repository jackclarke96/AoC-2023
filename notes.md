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

Also realise that don't need to iterate through backwards if no match is found in forwards direction.