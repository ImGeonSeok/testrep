FROM goLang:alpine as builder

WORKDIR /usr/src/app
COPY . .

ENV CGO_ENABL=0 \
    GOOS=linux \
    GOARCH=amd64
    
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

RUN go build -o main .

CMD ["./main"]
