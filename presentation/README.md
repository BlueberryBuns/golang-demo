# Presentation

**Table of content:**
1. [Why is it worth to choose Go lang?](#wgl)
    1. [Advantages](#wgl-a)
    2. [Disadvantages](#wgl-d)
2. [Go vs Python - comparison](#govspython)
3. [Docker Example and Running code](#dockerexample)
    1. [Build for your platform and other platforms](#dockerdifferenplatform)
4. [Writing HTTP server](#httpserver)
5. [Testing in Go](#testing)
6. [Deploying package](#deployingpackage)

<a id="wgl"></a>
## Why is it worth to choose Go lang?

One may consider why is it worth to even bother choosing other languages than what they are familiar with (as in our case: `Python`, `Java`, `Node.js`, `PHP`). So here some advantages of why go might be good choice for you.

1. **Performance**: Go is a compiled language, which means it generally performs better than interpreted languages like JavaScript (Node.js) and Python. It's also often faster than Java because it doesn't have the overhead of a JVM. Here are some comparison benchmarks to prove the point, huge shout out to @antonputra for providing benchmarks with real life use cases:
    - [Go (Gin) vs Python (Flask)](https://www.youtube.com/watch?v=vJsqDqq1R0Y) **Result: Go** :crown:
    - [Go (Fiber) vs Node.js](https://www.youtube.com/watch?v=ntMKNlESCpM) **Result: Go** :crown:
    - [Go (Fiber) vs Java (Spring Boot)](https://www.youtube.com/watch?v=8CiErLxdaA8)  **Result: TIE** (slight Go advantage on low resource)
        - :japanese_ogre:/:coffee: Go **IN MOST CASES (not all)** uses less resources
        - :japanese_ogre: Go has faster response times (Java dropped requests - lack of resources)
        - :coffee: Go has ran out of resources in second test


2. **Concurrency**: Go has built-in support for concurrent programming with goroutines and channels. This makes it easier to write efficient, concurrent programs compared to languages like Python and Java.

3. **Simplicity**: Go has a smaller and simpler syntax compared to Java and JavaScript, which can make it easier to learn and use. It also enforces a single, idiomatic way to do things, which can lead to more readable and maintainable code.

4. **Static Typing:** Go is statically typed, which can help catch errors at compile time rather than at runtime. This can lead to more robust code compared to dynamically typed languages like JavaScript and Python.

5. **Standard Library:** Go has a rich standard library that covers a lot of areas, reducing the need for external libraries. This can make dependency management simpler compared to languages like JavaScript (Node.js) that rely heavily on external packages.

6. **Cross-Compilation:** Go makes it easy to cross-compile your code for different platforms. This can be a big advantage if you need to distribute your software for different operating systems.

<a id="wgl-a"></a>
### Advantages
1. It's a simple language, maybe not as simple as Python, but on the other hand it doesn't hide as many surprises as python and it's really fast to get up to speed with go.
```go
func convertData(){
    return "XD"
}
```
2. Built in concurrency comes in handy
3. A lot has changed after generics were introduced
4. Write once and have nothing to worry until performance becomes an issue (but then congrats :tada::clap: you have succeeded in your IT journey)
5. Go can be easily functional by choice (especially with upcoming iterators)
6. Actually fun to code. Feels arcade-ish :cake:
7. Single binary as output
8. Compile time
9. Language made to work with server stuff
<a id="wgl-d"></a>
### Disadvantages
1. Error handling is fu**ing terrible. You can see how errors are handled in an example shown below.
```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

It doesn't seem terrible, right? In fact this might seem readable. So how about that?

```go
//**Side note**: This function is written in terrible way
func foo() error {
    res, err := acquireSomething()
    if err != nil {
        return fmt.Errorf("error acquiring sth: %w", err)
    }
    err = useSomething(res)
    if err != nil {
        err = fmt.Errorf("error using sth: %w", err)
    }
    releaseErr := releaseSomething(res)
    if releaseErr != nil {
        err = fmt.Errorf("%w; failed to release sth: %w", err, errW)
    }
    return err
  }
```
<p align="center">
  <img width="460" height="300" src="assets/mr_incredible_err.png">
</p>

2. Syntax might seem strange **(although it's not my opinion)**
3. Changes made recently will definitely change how Go code is handled, making old tutorials obsolete.
    - Go standard HTTP library
    - Gorilla Mux package - [source of drama](https://www.reddit.com/r/golang/comments/zh0w0p/gorilla_web_toolkit_is_now_in_archive_only_mode/)
4. Easy to mess up project structure, making totally spaghetti code **(skill issue)**
5. Pointers might seem scarry (they are not)

<a id="govspython"></a>
## Go vs Python - comparison

Examples are attached in the section [go-python-comparison](presentation/go-python-comparison/)

<a id="dockerexample"></a>
## Running your code in docker

<a id="dockerdifferenplatform"></a>
### Build for your platform vs to other platforms

<a id="httpserver"></a>
## Writing HTTP server

<a id="testing"></a>
## Testing in Go

<a id="deployingpackage"></a>
## Deploying package