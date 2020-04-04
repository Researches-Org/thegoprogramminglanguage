# Interesting features of Go programming language

## Stack size is not fixed as in the majority of programming languages the size is usually from 64 KB to 2MB. Go uses variable-size stacks that start small and grow as needed up to a limit on the order of a gigabyte. This let us use recursion safely and without worring about overflow. Fixed stack size is considered a security risk.

## Go's garbage collector recycles unused memory, but do not assume it will release unused operating system resources like open files and network connections. They should be closed explicitly.

## One thing that I particular do not like is different ways of doing the same thing, for instance, with named results, the operands of a return statement may be ommitted.

```
func CountWordsAndImages(url string) (words, images, int, err error) {

...
    if err != nil {
        return // ommiting operands, but it will return the error
    }
...
}
```

