FROM ubuntu:21.10

# Set ENVs for our GO Project to Read
ENV GOLANG_VERION 1.16.6
ENV GOLANG_CHECKSUM be333ef18b3016e9d7cb7b1ff1fdb0cac800ca0be4cf2290fe613b3d069dfe0d

# Get Utilities (cURL)
RUN apt-get update && apt-get -y install curl

# Get our source file from the Google Repo
RUN \
      curl "https://golang.org/dl/go$GOLANG_VERION.linux-amd64.tar.gz" -L > go.tar.gz \
      && echo "$GOLANG_CHECKSUM go.tar.gz" | sha256sum -c - \
      && mkdir -p /usr/local/go \
      && tar xzf go.tar.gz -C /usr/local/go --strip-components=1

# Add Go to our Path
ENV PATH $PATH:/usr/local/go/bin

