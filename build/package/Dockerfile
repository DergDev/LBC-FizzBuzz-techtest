FROM golang:alpine

EXPOSE 8080
WORKDIR go/src/techtest
COPY ../go.mod .
COPY ../go.sum .
RUN go mod download -x
COPY .. .

CMD ["go", "run", "cmd/main.go"]