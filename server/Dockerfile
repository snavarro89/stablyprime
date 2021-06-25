FROM golang:1.14.3-alpine AS build-go
RUN apk --no-cache add git
WORKDIR /server
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/stablyapi

FROM scratch
COPY --from=build-go /server/bin/stablyapi /server/
EXPOSE 3001
ENTRYPOINT [ "./server/stablyapi" ]

