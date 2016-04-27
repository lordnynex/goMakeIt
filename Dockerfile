FROM golang:1.6-onbuild

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/mergermarket/app

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
WORKDIR /go/src/github.com/mergermarket/app

# Run the outyet command by default when the container starts.
CMD go test . && go build && ./app

# Document that the service listens on port 8080.
EXPOSE 8080