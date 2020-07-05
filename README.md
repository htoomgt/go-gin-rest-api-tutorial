# My First Go Gin Repository

The first go lan respository for rest api with gin framework with the help of youtube tutorial named [TutorialEdge](https://www.youtube.com/watch?v=RkmvVFZJJvs).

## Expectations

- To create a sample CRUD with go using gin framework
- To test run with MVC pattern for separation of concern
- To test no-blocking feature and speed of code

## How to install Gin Web Framework

- Run this at root directory `go get -u github.com/gin-gonic/gin`

## How to create Module

- Run this `go mod init`
- That will create go.mod and can set the module name at go.mode. `module github.com/username/app-name`

## How to code from another package in golang

- At import of go file import ***"module-name/package-name"*** or ***"module-name/dir-name/package-name"***
- Example ***"github.com/htoomgt/go-gin-rest-api-tutorial/models"*** or ***"github.com/htoomgt/go-gin-rest-api-tutorial/src/models"***
- And then ***package-name.FunctionName*** or ***package-name.VariableName***
- Example `models.Person` or `configs.AppEnv`
- Used CapticalCamel case for public visibility

## To Test Run

- At root directory run this `go run src/*.go` OR
- At root directory run `go build src/main.go` and `./main`

## Reference

1. How to import another go file [link](https://stackoverflow.com/questions/26942150/importing-go-files-in-same-folder)

2. Basic markdown syntax [link](https://www.markdownguide.org/cheat-sheet/)

3. Faker for Go Lang [link](https://github.com/bxcodec/faker)

4. Go Lang gin mysql CRUD example with person [link](https://gist.github.com/rsj217/26492af115a083876570f003c64df118)

5. How to connect to mysql database [link](https://stackoverflow.com/questions/23550453/golang-how-to-open-a-remote-mysql-connection)

6. How to create mysql table [link](https://www.mysqltutorial.org/mysql-create-table/)

7. Golang cannot find package gin in sub directory [link](https://stackoverflow.com/questions/51488385/golang-can-not-find-package-gin-in-sub-directory)

8. How to get current datetime in golang by Rahul on August, 13th, 2018 [link](https://tecadmin.net/get-current-date-time-golang/)

9. Current date time in various formatting [link](https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html)

10. A tour of GO [link](https://tour.golang.org/list)
