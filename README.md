# Bookings and Reservations

This is the repository for my bookings and reservations project.

- Built-in Go version 1.16
- Uses the [chi router](htttps://github.com/go-chi/chi/v5)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs/v2) session management
- Uses [nosurf](github.com/justinas/nosurf)

## What will we need?
- an auth system
- database
- email/txt sending

## Other notes

- To get the tmpl to work, install the vs code plugin for go templates and then assign *.tpml to html in vs code settings
- On Windows, you may need to run with `go run .` from within the cmd\web directory



### Write Coverage Report
- coverage
    - go test -coverprofile="coverage.out"
    - go tool cover -html="coverage.out"


- gotests 

### Run everything
- ./run.bat
- Issues:
    - TODO: does not appear to load all pages, will investigate later

## VSCode Commands
- Go: Toggle Test Coverage in Current Package <= turn on the highlighting of the test coverage