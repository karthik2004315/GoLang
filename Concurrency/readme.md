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