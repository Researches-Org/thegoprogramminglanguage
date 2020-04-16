# Interesting features of Go programming language

* Stack size is not fixed as in the majority of programming languages the size is usually from 64 KB to 2MB. Go uses variable-size stacks that start small and grow as needed up to a limit on the order of a gigabyte. This let us use recursion safely and without worring about overflow. Fixed stack size is considered a security risk.

* Go's garbage collector recycles unused memory, but do not assume it will release unused operating system resources like open files and network connections. They should be closed explicitly.

* One thing that I particular do not like is different ways of doing the same thing, for instance, with named results, the operands of a return statement may be ommitted. This is called `bare return`, the returned values may not be clear for who is reading the code.

```
func CountWordsAndImages(url string) (words, images, int, err error) {

...
    if err != nil {
        return // ommiting operands, but it will return the error
    }
...
}
```
* Emphasis of Panic being a sure sign of a bug in  he calling code and should never happen in a well-written program.

* G's approach sets it apart from many other languages in which failures are reported using exceptions, not ordinary values. Although  Go does have an exception mechanism of sorts, it is used only for reporting truly unexpected errors that indicate a bug , not the routine errors that a robust program should be built to expect. The reason for this design is that exceptions tend to entangle the description of an error with the control flow required to handle it, often leading to an undesirable outcome: routine errors are reported to the end user in the form of an incomprehensible stack trace, full of information about the structure of the program but lacking intelligible context about what went wrong. By contrast, Go programs use ordinary control-flow mechanisms like if and return to resp ond to errors. This syyle undeniably demands that more attention be paid to error-handling logic, but that is precisely the point.

* To construct a new error message the function fmt.ErrorF formats an error message using fmt.Sprintf and returns a new error value:

```
fmt.Errorf("parsing %s as HTML: %v", url, err)
```
* Because error messages are frequently chained together, message strings should not be capitalized and new lines should be avoided. The resulting errors may be long, but they will be self-contained when found by tools like grep.

* Caveat: Capturing Iteration Variables

```
var rmdirs []func()
     for _, d := range tempDirs() {
         dir := d               // NOTE: necessary!
         os.MkdirAll(dir, 0755) // creates parent directories too
         rmdirs = append(rmdirs, func() {
             os.RemoveAll(dir)
         })
}
     // ...do some work...
     for _, rmdir := range rmdirs {
         rmdir() // clean up
}
```
* It's tempting to use a deferred call to f.Close, to close the local file, but this would be subtly wrong because os.Create opens a file for writing, creating it as needed. On many file systems, notably NFS, write errors are not reported immediately but maybe postponed until the file is closed. Failure to check the result of the close operation could cause serious data loss to go unnoticed. However, if both io.Copy and f.Close fail, we should prefer to report the error from io.Copy since it occurred first and is more likely to tell us the root cause.

* It’s go o d practice to assert that the preconditions of a function hold, but this can easily be done to excess. Unless you can provide a more informative error message or detect an error sooner, there is no point asserting a condition that the runtime will check for you.

```     
func Reset(x *Buffer) {
     if x == nil {
         panic("x is nil") // unnecessary!
     }
     x.elements = nil
}
```

* If all the methods of a named type T have a receiver type of T itself (not *T), it is safe to copy instances of that type; calling any of its methods necessarily makes a copy. For example, time.Duration values are liberally copied, including as arguments to functions. But if any method has a pointer receiver, you should avoid copying instances of T because doing so may violate internal invariants. For example, copying an instance of bytes.Buffer would cause t he original and t he copy to alias (§2.3.2) the same underlying array of bytes. Subsequent method calls would have unpredictable effects.

* In every valid method call expression, exactly one of these t hree statements is true. 

** Either the receiver argument has the same typ  as the receiver parameter, for example both have type T or both have type *T.
** Or the receiver argument is a variable of type T and the receiver parameter has type *T. The compiler implicitly takes the address of the variable.
** Or the receiver argument has type *T and the receiver parameter has type T. The compiler implicitly dereferences the receiver, in other words, loads the value 

