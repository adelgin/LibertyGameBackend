# build the binary
FROM golang:1.23.2 AS build

RUN useradd -u 10001 gopher

ENV NAME "liberty"
RUN mkdir /${NAME}
WORKDIR /opt/${NAME}

ENV GOPROXY="https://proxy.golang.org,direct"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/${NAME} ./cmd/tap/main.go

# run the binary
FROM scratch

ENV NAME "liberty"

COPY --from=build /etc/passwd /etc/passwd

USER gopher

COPY --from=build /opt/${NAME}/bin/${NAME} /${NAME}

CMD ["./liberty"]