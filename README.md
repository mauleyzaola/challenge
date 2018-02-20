## Cabify Challenge
*Mauricio Leyzaola -> mauricio.leyzaola@gmail.com*

### Definition
Please read from this gist: 
https://gist.github.com/colega/ddcc8f607659e74a97c402cdb639954f

### Installation
Run within this directory the command
```$bash
go get -v -t ./...
``` 

### Execution
You need to start the API first (default port number is:8000 or `$PORT` environment variable)
```bash
cd api && go build
./api
```
Once the API has started, start the CLIENT
```bash
cd client && go build
./client
```
The first time the CLIENT program starts, it will display all the available commands 


### Tests
You can run the unit tests for the whole application (where available, not all packages have tests)
```
go test -v -cover ./...
```