FROM golang:1.26

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download && \
go install github.com/go-task/task/v3/cmd/task@latest && \
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest && \
go install github.com/air-verse/air@latest && \
go install github.com/uudashr/gocognit/cmd/gocognit@latest && \
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

COPY . .

EXPOSE 8080

CMD ["air"]
