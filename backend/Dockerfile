FROM golang:1.22.5

WORKDIR /go/src

COPY go* ./
RUN go mod download

COPY . .

# FROM golang:1.22.5 as builder

# WORKDIR /app
# COPY go.* ./
# RUN go mod download

# COPY . ./

# RUN go build -v -o server

# FROM chromedp/headless-shell:latest

# COPY --from=builder /app/server /app/server

# CMD ["/app/server"]
