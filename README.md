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

11. Adding database (ubuntu)
    - create database `psql -d bird_encyclopedia`
    - connect to database `\c`

12. Create interface (let say Store)
`touch store.go`
add test `touch store_test.go`

13. Create mock store to avoid test accessing database directly
`touch store_mock.go`

14. run this command to add new external library
`go get github.com/stretchr/testify/mock`

15. Modify `main.go`, then run this command to add new external library
`go get github.com/lib/pq`

Build and run the project