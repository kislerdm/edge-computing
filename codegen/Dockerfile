FROM golang:1.17.7-alpine3.15

RUN apk update -q &&\
    apk add -q --no-cache  \
      chromium  \
      chromium-chromedriver \
      udev \
      xvfb

COPY . /src

WORKDIR /src

RUN go mod tidy && go build -o runner main.go

ENTRYPOINT ["./runner"]
