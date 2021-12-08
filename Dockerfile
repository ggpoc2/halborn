# Simple usage with a mounted data directory:
# > docker build -t halborn .
#
# > docker run -it -p 26657:26657 -v ~/.halbornd:/root/.halbornd halborn halbornd init testing
# > docker run -it -p 26657:26657 -v ~/.halbornd:/root/.halbornd halborn halbornd start
FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/ggpoc2/halborn

# Add source files
COPY . .

# build halborn
RUN make install


# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/halbornd /usr/bin/halbornd
COPY --from=build-env /go/bin/halborncli /usr/bin/halborncli

# Run halbornd by default, omit entrypoint to ease using container with halborncli
CMD ["halbornd"]
