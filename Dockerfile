FROM golang:1.16

##buat folder app
RUN mkdir /app

##test direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]