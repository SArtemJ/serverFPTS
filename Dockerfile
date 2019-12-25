FROM golang:alpine
RUN apk add --no-cache --update \
    git

RUN mkdir -p ${GOPATH}/src/${PROJECT_PATH}
WORKDIR ${GOPATH}/src/${PROJECT_PATH}
COPY . .
RUN go build -o fpts .

ENTRYPOINT [ "./fpts" ]

EXPOSE 8099