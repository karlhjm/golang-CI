FROM golang:1.9 
COPY . "$GOPATH/src/github.com/moandy/canyonsysu"
RUN cd "$GOPATH/src/github.com/moandy/canyonsysu" && go get -v && go install -v
WORKDIR / 
EXPOSE 7070 
VOLUME ["/data"] 
#RUN chmod +x canyonsysu
#ENTRYPOINT ["./canyonsysu"]
