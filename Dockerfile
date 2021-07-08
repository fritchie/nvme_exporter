FROM golang:1.16
MAINTAINER Frank R <12985912+fritchie@users.noreply.github.com>

RUN apt-get update
RUN apt-get -y install nvme-cli

WORKDIR /go/src/nvme_exporter
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9998

CMD [ "nvme_exporter" ]
