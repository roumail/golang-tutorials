# Greetings

We run `go mod init example.com/greetings` to create our go.mod file. 

The function name is capitalized, meaning that other functions outside the package can also reference this. 

Go calls this an exported name.

`:=` declares and initializes a variable in one line, where the value on the right is used by go to infer the variables type. 

This allows use to use 

```go
message := fmt.Sprintf("Hi, %v. Welcome!", name)
```

instead of 

```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

The next part of the tutorial requires creating a directory structure like so:

```bash
├── greetings
│   ├── go.mod
│   └── greetings.go
└── hello
    ├── go.mod
    └── hello.go
```

The link to the [tutorial](https://go.dev/doc/tutorial/call-module-code) is found here.

Now, we use the greetings package inside `hello/hello.go`. However, we need to edit our `go.mod` because we don't publish our packages online. 

Therefore, we must run the following command:
`go mod edit -replace example.com/greetings=../greetings`

This action results in adding the last replace statement in our go mod file.

```
module example.com/hello

go 1.23.4

replace example.com/greetings => ../greetings
```

Once done, we now need to run `go mod tidy` to synchronize the `example.com/hello` module's dependencies adding those required by the code, but not yet tracked in the module.

Running this statement results in the following line being further added to the `go.mod` file. 
`require example.com/greetings v0.0.0-00010101000000-000000000000`

The command found the local code in the greetings directory, added the require directive above sicne we're importing greetings in `hello.go`.

Since we don't have a semantic version number for the greetings package yet, we have a pseudo-version-number that's been generated in it's place for the time being.

This takes us to the next part of the [tutorial](https://go.dev/doc/tutorial/handle-errors).

Now we're at the next stage of the [tutorial](https://go.dev/doc/tutorial/random-greeting). Here we adapt our previous scripts to learn the concept of slices and returning a random instance from the slice.

Next, we go to the next stage of the [tutorial](https://go.dev/doc/tutorial/greetings-multiple-people).

In this tutorial, we return a map so our hello function can take a multiple names and return a map where each name is associated with a greeting. Maps are made using the syntax `make(map[key-type]value-type)`.

We also use a for loop, using range on the slice of names, suppressing the first return value of the function in the for loop using `_`.

Links: 
* https://go.dev/blog/maps
* https://go.dev/doc/effective_go.html#blank

Next section: Adding a [test](https://go.dev/doc/tutorial/add-a-test)

After creating a file called greetings_test.go, we can now start writing tests. Go's convention for the `go test` command is to have file names ending with `_test.go`. This indicates that the file contains test functions.

Test functions always have the name TestName, where Name is something specific about the test. The functions take a pointer to the `testing` package's test.T type. We use this parameters methods for reporting and logging from our test. 

Link to testing package docs: https://pkg.go.dev/testing#T.Errorf

Next topic, compiling and installing the package: https://go.dev/doc/tutorial/compile-install 
