# Learning to build simple web app using GO

**Step to build this project**

1. install GO

2. create directory
- `mkdir <dir-name>`
- `cd <dir-name>`

3. init the project 
- `go mod init <project-name>`

4. create main.go file
- `touch main.go`

5. run main.go project
- `go run main.go`

6. instal external libraries
- `go get github.com/gorilla/mux`

7. create unit testing
`touch main_test.go` 
- By convension files with *_test.go pattern will treat as a test

8. create static assets
- `mkdir assets`
- `touch assets/index.html #add html file`

9. run the project
- `go build`
- `./birdpedia`


10. run the test
`go test ./...`