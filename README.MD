# Hands-on go and mongoDB

This repo was created to follow the Hands on training with go and mongoDB.

The project and some notes will be stored here as I move on in the training.

## Dependencies

### MongoDB

MongoDB can be started from docker:

`docker run --name mongodb -p 27017-27017 -d mongo`


## Run the application

Before running the application, make a copy of the file `application.yml.default` and rename this copied file to `application.yml`.
Set your configuration values properly in this file.

Run the command below to start the  application.

```
go run main.go
```