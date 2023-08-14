FROM golang:1.20

RUN mkdir /internal
COPY . /internal
WORKDIR /internal/sast
RUN go build -o /app .

CMD ["/app"]
