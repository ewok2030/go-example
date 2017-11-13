# If building outside the image, use:
FROM scratch
COPY bin/go-example /
CMD ["/go-example"]

# If building inside the image, use the following:

# GOPATH = /go
##FROM golang

# Copy the local package files to the container's workspace.
##ENV APP github.com/ewok2030/go-example
##COPY . /go/src/$APP

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
##RUN go $APP

# Run the outyet command by default when the container starts.
## ENTRYPOINT /go/bin/go-example

# Document that the service listens on port 8080.
EXPOSE 8080