# Microservice POC in GoLang

## Objectives

- Create different microservices

  - auth
  - communication
  - media
  - database
  - user_management

- Create a user gallery where a user can upload and modify its files.

User can upload files. User needs to be logged in to upload file. When user uploads a file, an email is sent to the user's mail.

### Generate protoc command

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user_management.proto
```

Note for me:
find implementation of enum value maps inside common.pb.go file. -> common.StateOptions
