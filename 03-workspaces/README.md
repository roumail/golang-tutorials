# README

We start by creating `workspace`, `cd` into it and create a new module hello that will depend on the golang.org/x/example module.

A module is effectively a directory containing a `go.mod` file. 

In our `hello` module we can add a dependency on the `golang.org/x/example/hello/reverse` package by using `go get.` to execute `$ go get golang.org/x/example/hello/reverse`.

After creating `hello.go` to reverse the string characters, we can run the program using `go run .`.

The workspace itself can be created by running `go work init ./hello` in the workspace directory. This creates a go.work file, which has a syntax, not too different from the typical `go.mod`. 

```bash
❯ tree -C                                                                                                   
.
├── go.work
└── hello
    ├── go.mod
    ├── go.sum
    └── hello.go

```
The `use` directive tells Go that the module in the hello diretory should be main modules when doing a build. 
```bash
# go.work
go 1.23.4

use ./hello
```
This module is therefore active in any subdirectory of workspace.

Running the hello module from the `workspace` can be done using `go run ./hello`, however, we can't run the command outside the workspace doesn't work.

## Create a modified local copy of a git repo

Learning outcomes:
* Download copy of `golang.org/x/example/hello module`,
* Add it to the workspace, and then add a new function we will use from the hello program.

First we clone the project into the `workspace` directory and then use `go work use ./example/hello` to add the module to our workspace. This change is represented by the addition of the line `./example/hello` in our go.`work` file. 

Let's add a new function to reverse a number to the cloned project. This means creating an `int.go` inside the `reverse` package directory of the cloned forlder. Then we modify our hello program to use that function. 

## Release the hello module package

To make our change to the hello module we'd need to tag the commit and release it. Then we can update the go.mod file below to reference the new release

```
module golang.org/x/example/hello

go 1.19
```

Further references:
* https://go.dev/ref/mod#workspaces 
