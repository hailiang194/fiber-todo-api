# Fiber Todo API
Stupid a\*ss h\*le Todo API using Golang\'s Fiber

## Execution
```
go run main.go
```

## API GUIDE
|Method|Path|Description|
|------|----|-----------|
|```GET```|```/api/todos/```|Get all todos|
|```POST```|```/api/todos/```|Create new todo and append to list|
|```GET```|```/api/todos/:id/```|Get the todo whose ID is ```id```|
|```POST```|```/api/todos/:id/```|Update the todo whose ID is ```id```|
|```DELETE```|```/api/todos/:id```|Delete the todo whose ID is ```id```|