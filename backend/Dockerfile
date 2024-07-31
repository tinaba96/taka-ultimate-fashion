FROM golang:1.22-alpine

WORKDIR /go/src

COPY go* ./
RUN go mod download


RUN apk upgrade --update && \
    apk add --no-cache git && \
    apk add --no-cache gcc musl-dev && \
    apk add sqlite && \
    apk add wget && \
    apk add udev && \
    apk add ttf-freefont && \
    apk add chromium && \
    apk add chromium-chromedriver

RUN go get github.com/antchfx/htmlquery
RUN go get github.com/sclevine/agouti
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/aws/aws-sdk-go-v2/aws
RUN go get github.com/aws/aws-sdk-go-v2/config
RUN go get github.com/aws/aws-sdk-go-v2/service/s3
# RUN go get -u github.com/cosmtrek/air && \
#     go build -o /go/bin/air github.com/cosmtrek/air

# RUN go mod tidy
# RUN go mod download


COPY . .

CMD ["go", "run", "main.go"]

# EXPOSE 8080

