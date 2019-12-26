# iron/go:dev is the alpine image with the go tools added
FROM golang:1.12

ARG service_name


#COPY . /go/src/innete
ENV GO111MODULE off
WORKDIR /go/src/github.com/sakari-ai/moirai
COPY . .
RUN find ./ -name keys.generated.go -exec rm -irf {} \;
#RUN go get -d -v ./...
#RUN go install
RUN env  GOOS=linux GOARCH=amd64  go build -o ${service_name} server/build/${service_name}/main.go
