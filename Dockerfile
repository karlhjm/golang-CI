FROM golang:1.9 
COPY . "$GOPATH/src/canyonsysu" 
RUN cd "$GOPATH/src/canyonsysu" && go get -v && go install -v
WORKDIR / 
EXPOSE 8080 
VOLUME ["/data"] 
