FROM golang:1.16.4-alpine AS build
ENV WORKDIR_PATH /go/src/github.com/chiguhagu/amenokaguyama-mysql
WORKDIR ${WORKDIR_PATH}
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 go mod download
ADD . ${WORKDIR_PATH}
RUN CGO_ENABLED=0 go build -o /bin/amenokaguyama-mysql .

FROM alpine:3.13.5
ENV VERSION '0.10.1'
ADD https://github.com/k0kubun/sqldef/releases/download/v0.10.1/mysqldef_linux_amd64.tar.gz /tmp/
RUN cd /tmp \
    && tar zxvf mysqldef_linux_amd64.tar.gz -C /bin \
    && rm -rf mysqldef_linux_amd64.tar.gz
COPY --from=build /bin/amenokaguyama-mysql amenokaguyama-mysql
ENTRYPOINT ["./amenokaguyama-mysql"]
