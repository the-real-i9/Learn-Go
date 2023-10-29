module learn/helpers

go 1.21.3

// this is a custom package/module, since the go files in it have "package moduleName" declared.

// custom package/module functions to be exported must begin with uppercase, others must begin with lowercase and must be used.

// in other packages/modules (including "main"), import this package/module with `import "learn/helpers"`, you can access it with "packageName.ExportedFunc()"