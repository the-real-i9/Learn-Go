module learn/ch1

go 1.21.3

// this is a module/package
// there are two kinds of module/package, one that's a custom package and the other a special "main" package. 
// What you get depends on the package name you declare. Regardless of your module name, if you delare "package main", your module becomes the main package/module, it becomes the entry point of execution, all other packages are imported into it directly or indirectly. But, if you declare "package moduleName", it becomes a custom package, to be imported by the main package/module directly or indirectly

// Therefore, "ch1" folder/module is the main package/module, since it has the "package main" declared. Additionally, you can compile the go files in this package/module, to one executable

// The package name you use must be the one you use in all the go files in that module.