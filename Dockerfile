FROM golang

ARG app_env
ENV APP_ENV $app_env
ENV APP_CONFIG=/etc/myapi/config.yml

RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/user
RUN mkdir /etc/myapi

COPY ./src /go/src/github.com/user/app
WORKDIR /go/src/github.com/user/app

RUN go get github.com/gorilla/mux
RUN go get github.com/urfave/negroni
RUN go get github.com/go-sql-driver/mysql
RUN go get gopkg.in/yaml.v2

RUN go build

CMD go get github.com/pilu/fresh && fresh;

EXPOSE 8081