# Learn Go by Writing Tests
  
### Resources
- source - https://quii.gitbook.io/learn-go-with-tests
  
  
### Notes

These examples are generally for Go > 1.11 (when go modules were introduced)

Writing a test is just like writing a function, with a few rules:
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

Basic process is to write code, write test (example 1). Then:
```
go mod init hello 
```

The file go.mod will be created with similar contents:
```
module hello

go 1.19
```
Once this file is generated, it should be checked in as it will describe the dependencies your code requires. Then someone else can go mod init to retrieve the dependencies.

To test, run:
```
go test
```

