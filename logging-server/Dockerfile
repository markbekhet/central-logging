FROM golang:latest

RUN useradd mark

RUN mkdir -p /home/mark/go/src/github.com/markbekhet/central-logging/logging-server
RUN export GOPATH=/home/mark/go/
WORKDIR /home/mark/go/src/github.com/markbekhet/central-logging/logging-server

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o logging-server

EXPOSE 5051

CMD [ "./logging-server" ]