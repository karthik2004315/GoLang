## GO Modules and Packages

* There are mainly 3 things we will discuss and those are **Repository**, **Modules**, **Packages**

##### Repositories

* Repositories is a place in version control system where the source code of project is stored.
* Go Modules are stored in Repostiory.

##### Modules
* Module is the root of a go library or go application stored in a repository. Module contains one or more packages. So a Module is well organized.
* Every Module has a globally Unique Identifier.
* A Module specifies dependencies needed for the project and set of other modules needed go.mod file

To declare a directory as a go module, we will enter this command in terminal

```go mod init <module_path>```

* when we enter the above command a **go.mod**
file is created which tracks the version and dependencies in our code.

* If we want to use a third party library in **main.go** file we will use the command to install the necessary package in go.mod file

```go mod tidy```