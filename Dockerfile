from golang as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/backend-spp

EXPOSE 1323
CMD ["/usr/local/bin/backend-spp"]
