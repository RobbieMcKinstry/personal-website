FROM golang:alpine AS build-env
ADD . /src/github.com/RobbieMcKinstry/personal-website
RUN apk update 
RUN apk add git
ENV GOPATH=/
RUN cd /src/github.com/RobbieMcKinstry/personal-website && go get -d ./... && go build -o /bin/server

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /bin/server /app/
COPY . /app/

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

EXPOSE 8080
ENTRYPOINT ["./server"]
