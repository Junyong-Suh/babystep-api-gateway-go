# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name
ENV SRC_DIR=/Users/junyong/Yoshi/babystep-api-gateway-go
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN go get -u github.com/gorilla/mux
RUN cd $SRC_DIR; go build main.go; go test ./tests
ENTRYPOINT ["/Users/junyong/Yoshi/babystep-api-gateway-go/main"]
