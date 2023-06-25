# Concurrency
----------------------------------------------------------------------------------------------------------------------------

### Sequential Processing vs Concurrent Programming

 * sequential processing means excecuting tasks in CPU one by one. where as the drawback in this is CPU is super fast so we will waste the remaining memory when we use sequential processing

 * this is where concurrency comes into picture when we use concurrency comes into place, we execute multiple tasks at a time so that CPU memory is efficiently utilized. 

 * concurrency is where multiple processes are execute at one time.
 * **Multi processing system** : let us say there are two processes and two cpus now these two cpus are working in a way like both cpu execute both processes by exchanging in between them. this is called Multi processing system
 * **Parallel processing** : by taking same example as above where as here two cpus execute one process each parallelly
 ![concVSparallel](https://th.bing.com/th/id/OIP.YodILe1zI4wCSo_vVDrdCQHaEo?w=287&h=180&c=7&r=0&o=5&dpr=1.3&pid=1.7)
----------------------------------------------------------------------------------------------------------------------------

### Go routines

* it is considered as light weight thread in which every go routine executes independently. other go routines will execute concurrently. these go routines are managed by Go runtime scheduler

syntax: 

```
    go <function name>
```
* here we used a special keyword called go when we used it is considered as go routine
----------------------------------------------------------------------------------------------------------------------------

To understand the differnce between the sequential progamming and concurrency I have written a sequential program and a concurrency program

* Here when we execute sequential program it takes more time where as when we execute concurrent program each printing is done in seperate go routine so it takes less time than sequential program

#### Main go routine

Main go routine is start of all other go routines and there are no children for all go routines. and the output is non-deterministic so we cant say which go routine executes first

#### Anonymous go routine

Anonymous go routines are functions which doesnt have any name it just executes like any other go routine
 
Example:

```
    func main(){
        go func{
            fmt.Println("this is anonymous function")
        }()
        time.Sleep(1 * time.Second)
    }
```
 
----------------------------------------------------------------------------------------------------------------------------

## Go Runtime

We know that CPU has kernel to manage threads but where as in Go lang we have our own Go runtime for Go routines. and function of this Go runtime is it schedules our Go routines into CPU threads(also known as m to n multiplexing)

Heres the over view

* OS contains multiple CPU's and each CPU has some threads called OS threads now what happens is a LRQ or Local run Queue is assigned to each OS thread.
* Now each LRQ contains Go routines to be executed. The Go runtime scheduler maintains a GRQ or Global run queue which has Go routines which are yet to be executed
* So the task of Go runtime Scheduler is, it assigns the Go routines in GRQ to LRQ

* Go runtime is cooperative scheduler which means it is non preemptive scheduler it will not interrupt the ongoing process for some other process like context switching

Some examples of context switching: Function calls, on using go keyword etc..

### Go routine VS Thread

* Go routine takes less space than thread so its cheaper
* Go routines are multiplexed to fewer OS threads( like 1 OS thread can contain more than 1000 GO routines)
* Go routines are faster
* Go routines communicate using channels to prevent race condition from happening

### Wait Groups

* The main problem with Go routines is the main program is being terminated before the Go routines are completed.
* This package acts like a counter and it blocks the execution in structured way until its internal counter becomes 0

syntax:

    ```
    import "sync"

    var <name> sync.WaitGroup
    ```
* we have three methods in wait group
    * **wg.Add(int)** - this indicates the number of available go routines to wait for. integer param is a counter
    * **wg.Wait()**   - this blocks the execution of code until the internal counter reduced to 0
    * **wg.Done()**   - this decreases the internal counter by 1
* lets see how that works lets say we have 2 go routines so we will call add method with param 2 **wg.Add(2)** now we will 
 call wait method to block the execution of main method **wg.Wait()** now we can execute go routines.
* when we called add method we have internal counter of count = 2 now when 1 go routine finished it calls **wg.Done()** method then internal counter count = 1 then second go routine finishes and internal count reduces to 0 so now the main method is executed
* this is all about wait groups

### Channels

* Channels are means through which go routines communicate
* it is bidirectional by default means we can send or recieve values from channels
syntax:

```
var c chan int
```
* what we did is we created a channel called "c" and of type integer
* channels as many operations such as 
    * recieve a value
    * send a value
    * closing a channel
    * querying buffer over a channel
    * querying length of a channel
* #### sending a value over a channel:

    syntax:
    ```
        ch <- v
    ```
    * "<-" this is send operator
* #### recieving a value over a channel

    syntax:
    ```
    val := <- v
    ```
    * this is same as sending a value but we are recieving a value so we have to assign it to a variable
* #### closing a channel

    syntax:
    ```
    close(ch)
    ```
* #### querying a buffer of a channel

    syntax:
    ```
    cap(ch)
    ```
* #### querying length of the buffer of a channel

    syntax:
    ```
    len(ch)
    ```
* channels are two types buffered and unbuffered
    * **UnBuffered channels**: the channel that needs a reciever as soon as a message is emitted to the channel, no capacity to hold, do not stores data

    * ** Buffered channels** : it has a capacity to hold the data
        * sending to a channel blocks the go routine, only if the buffer is full
        * recieving from a channel blocks only when a channel is empty 

syntax of a buffered channel:
    ```
        ch := make(chan <data type>, capacity)
    ```
Example:
    ```
        ch := make(chan int, 10)
    ```
* length of a unbufferd channel is always 0
* closing a channel means no more data can be sent to that channel and we can do that by using inbuilt function close()
* we can check if a channel is closed or not by assigning second variable when recieving data
    Example:
        ```
        val, ok := <- ch
        ```
### Panic situations
* panic situation is like an exception it occurs at runtime
* some panic situations are like
    * when closing a closed channel
    * when sending a data to a closed channel

* we can accept values of a channel by using for - range method also

### Select statement
* select statement is like a switch statement
* it lets a go routine wait on multiple communication operations
* in select each of the case statement waits for sending or recieving data from a channel
syntax:
```
select{
    case channel send  or recieve:
        //statement
    case channel send or recieve:
        // statement
    default: 
        // statement
}
```
### Go routines leaks
* whenever we declared a go routine we must make sure that it is terminated or it occupies memory and its memory is reserved. this kind of memory leak is called go routine leak

### concurrency practices
* whenever we use closures we have to make sure that value passed to the go routine or function should be as a parameter not like scope variable because go routines follow parallelism and they may start execution after certain point of time and we may get another value as output

* the main problem is when to use buffered channel and unbuffered channel. so buffered channels are useful when you know how many go routines you have launched, and want to limit the go routines you have launched, or want to limit the amount of work queued up
* **Time out code**: most interactive programs should return the response within a certain time so to achieve that we use time out code that is we only block the code for certain period of time and it can be achieved by using "After" function the time package 