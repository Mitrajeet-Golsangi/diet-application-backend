FROM golang AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy && go mod vendor

RUN go build -o ./build/output/main ./cmd/main.go

CMD ./build/output/main