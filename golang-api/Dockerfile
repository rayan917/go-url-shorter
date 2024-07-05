
FROM golang:latest


WORKDIR /app


COPY . .


RUN go get -t url-short


RUN go test -v

RUN go build -o bin .

RUN chmod a+x /bin


EXPOSE 8080


ENTRYPOINT [ "/app/bin" ] 