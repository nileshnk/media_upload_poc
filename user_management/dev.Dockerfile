FROM golang:1.21.3-alpine 

RUN go install github.com/cosmtrek/air@latest

WORKDIR /src

COPY . .

RUN go mod download -x

CMD ["air"]