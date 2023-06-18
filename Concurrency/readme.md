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

    ```goroutine
    go <function name>
    ```
* here we used a special keyword called go when we used it is considered as go routine