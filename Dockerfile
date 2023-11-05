# Builder
FROM golang:1.20-alpine as builder

WORKDIR /src/

# Dependencies fetching will be done firstly before build. With this approach, multilayer caching feature in Docker could be implemented.
COPY go.mod go.sum ./
RUN  go mod download

COPY . ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /out/app -ldflags="-w -s" -a -installsuffix cgo

# Distribution
FROM gcr.io/distroless/static as final

ENV TZ=Asia/Jakarta

WORKDIR /

COPY --from=builder /out/app .

ENTRYPOINT ["./app"]
