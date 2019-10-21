FROM golang:latest
RUN apt-get update && apt-get upgrade && \
    apt-get install bash git -y && \
    apt-get install openssh-server -y

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN go build -v ./cmd/apiserver

EXPOSE 8080

CMD ["sleep", "infinity"]