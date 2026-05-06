FROM golang
LABEL  author="北稚"
WORKDIR /go/src/
COPY . .
EXPOSE 80
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy
CMD ["/bin/bash", "/go/src/script/build.sh"]    