FROM golang:latest

RUN useradd mark

RUN mkdir -p /home/mark/go/src/github.com/markbekhet/central-logging/http-server
RUN export GOPATH=/home/mark/go/
WORKDIR /home/mark/go/src/github.com/markbekhet/central-logging/http-server

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o http-server

EXPOSE 8080

CMD [ "./http-server" ]