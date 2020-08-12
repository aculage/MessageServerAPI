FROM golang:latest
WORKDIR /usr/home/mserverapi
RUN git clone https://github.com/arseniculage/MessageServerAPI
RUN cd MessageServerAPI/cmd/apiserver && go build
EXPOSE 9000
