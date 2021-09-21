# Installing the application
* Make sure you have go installed.  See golang.org for help
* This application uses postgresql as it's database, so make sure you have it installed. For help, see postgresql.org
* Clone or download the repository, cd into the repository.
* Install all packages needed with the 'go get' command.
* You can cd into the config folder inside of the project and run 'go get .' (without the quotes)
* Go to localhost:8080 in your browser to see the result

# Note

**Remember:** You need to capitalize to have a variable, type, or func exported from a package

# run the application and make a request
```
curl -i localhost:8080/books
```

```
curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
```