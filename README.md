# Cartola Data Scrapper

The `cartola-data-scrapper` service polls data from the CartolaFC's REST API and sends the data to our own service, the [cartola-rest-api](https://github.com/dijckstra/cartola-rest-api/). It uses [gocron](https://github.com/jasonlvhit/gocron) to periodically request the latest player information.

To run it, you need to install [Golang](https://golang.org/doc/install). Follow the link for installation instructions.

After that, run the following commands:

``` bash
dep ensure
go run main.go
```

This service executes a request every Monday; you can edit the code to make it run right away. Make sure you run the REST API (and point this service to it) before doing so.