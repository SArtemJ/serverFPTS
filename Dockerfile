FROM golang:alpine
RUN apk add --no-cache --update \
    git

RUN mkdir -p ${GOPATH}/src/${PROJECT_PATH}
WORKDIR ${GOPATH}/src/${PROJECT_PATH}
COPY . .

ENTRYPOINT [ "./serverFPTS" ]

EXPOSE 8099