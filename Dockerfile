FROM golang:latest as builder
WORKDIR /go/src/app
# Download dependencies
RUN go get github.com/google/uuid
RUN go get github.com/imdario/mergo
RUN go get github.com/stretchr/testify
RUN go get gorm.io/gorm
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -tags doc -a -installsuffix cgo -o /go/bin/go-docker .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir /root/src/
COPY --from=builder /go/bin/go-docker .
# EXPOSE 8080
CMD ["/root/go-docker"]
