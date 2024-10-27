# build the binary
FROM golang:1.23.2 AS build

ENV NAME "liberty"
RUN mkdir /${NAME}
WORKDIR /opt/${NAME}

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/${NAME} ./cmd/liberty/main.go

# run the binary
FROM scratch

ENV NAME "liberty"

COPY --from=build /opt/${NAME}/bin/${NAME} /${NAME}

CMD ["./liberty"]