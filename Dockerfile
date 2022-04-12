FROM golang:latest as dev

WORKDIR $GOPATH/fungo

COPY go.mod .
COPY go.sum .

RUN go env -w GO111MODULE=on && \
    go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install

FROM alpine:latest as pro

RUN apk --no-cache add ca-certificates tzdata openssh

ENV TZ=Asia/Shanghai

COPY --from=dev /go/bin/fungo /bin

ENTRYPOINT ["fungo"]

CMD ["--help"]