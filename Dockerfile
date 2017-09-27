FROM ptpadan1246/go-glide-docker
MAINTAINER Yasushi kobayashi <ptpadan@gmail.com>

ENV BOT_ENV production
COPY ./src/app /go/src/app

WORKDIR /go/src/app
RUN glide install && \
  go build main.go

# CMD ./main &
EXPOSE 7000
