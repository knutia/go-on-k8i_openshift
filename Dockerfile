FROM golang:latest 
ENV PORT 8000
EXPOSE 8000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]
