# Welcome to Lavendar77's GO Learning Series
I have tasked myself with a project for this week and I am really excited.

## A To-do List App
This is a small project to practice what I have learnt so far with my tutorials from Netninja via YouTube.

The major scope of the project is to manage a personal to-do list application containing tasks that are registered and can be marked as complete, with timestamps.

Wish me luck!

## How to run it
`go run *.go`

In this tutorial, I learnt that for running multiple files, you could just do `go run main.go secondfile.go thirdfile.go`, but I found the above shorter and precise.

## Phase 1 - Create the Todolist Object (struct in GO)
```go
type todolist struct {
    id uint64
    task string
    is_complete bool
    completed_at string
    created_at string
    updated_at string
}
```
