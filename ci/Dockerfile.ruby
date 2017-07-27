##################################################
##
## Ruby
##
##################################################
FROM analogj/libgit2-crossbuild:linux-amd64 AS base
FROM ruby:2.4 AS ruby
MAINTAINER Jason Kulatunga <jason@thesparktree.com>

RUN apt-get update && apt-get install -y --no-install-recommends \
 	apt-transport-https \
    ca-certificates \
    git \
    curl \
    gcc \
    g++ \
	&& rm -rf /var/lib/apt/lists/* \
	&& gem install rubocop \
	&& gem install rake \
	&& gem install bundler-audit


# Install GOLANG
ENV GO_VERSION 1.8.3
RUN curl -fsSL "https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz" \
	| tar -xzC /usr/local

ENV PATH="/go/bin:/usr/local/go/bin:${PATH}" \
	GOPATH="/go"

ENV PKG_CONFIG_PATH="/usr/lib/pkgconfig/:/usr/local/lib/pkgconfig/:/usr/local/lib/libgit2/lib/pkgconfig:/usr/local/lib/openssl/lib/pkgconfig:/usr/local/lib/libssh2"
RUN go get github.com/Masterminds/glide
COPY --from=base /usr/local/lib/libgit2/ /usr/local/lib/libgit2/
COPY --from=base /usr/local/lib/libssh2/ /usr/local/lib/libssh2/
COPY --from=base /usr/local/lib/openssl/ /usr/local/lib/openssl/

WORKDIR /go/src/capsulecd

COPY . .

## download glide deps & move libgit2 library into expected location.
RUN glide install \
	&& mkdir -p /go/src/capsulecd/vendor/gopkg.in/libgit2/git2go.v25/vendor/libgit2/build \
	&& cp -r /usr/local/lib/libgit2/lib/pkgconfig/. /go/src/capsulecd/vendor/gopkg.in/libgit2/git2go.v25/vendor/libgit2/build/

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

CMD ci/coverage.sh