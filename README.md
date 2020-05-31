# ASHWIN_20200529
Http web Crawler 
# Getting Started
 Application uses sqlLite db to store the crawled data .The database will be created on the fly
and stored in sqlite3 DB under .recipe.db
how to RUN
```
1.go build main.go
2.main.go



```
# How it is done 
 Used Html parser to retrieve the recipe information
 All the recipe will be crawled in parallel 
 
## Go inbuilt library used 

 *[Html] (https://golang.org/x/net/html)- To tokenize the html
 *[HTML AtoM] (https://golang.org/x/net/html/atom)-To find out each tag
 * [http](https://golang.org/pkg/net/http/)-for intiating http request 
 
 # Tests
  Added basic unit testing
