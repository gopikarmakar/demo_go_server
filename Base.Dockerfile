 # @author [Gopi Karmakar]
 # @email [gopi.karmakar@monstar-lab.com]
 # @create date 2018-01-26 05:02:02
 # @modify date 2018-01-26 05:02:02
 # @desc [description]
 
ARG GOLANG_TAG
FROM golang:${GOLANG_TAG}

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# Stuff needed for dep in alpine
RUN apk --update add git openssh curl && \
  rm -rf /var/lib/apt/lists/* && \
  rm /var/cache/apk/*

# Install dep
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# Now copy your app to the proper build path
ARG APP_NAME
RUN mkdir -p $GOPATH/src/${APP_NAME} 
WORKDIR $GOPATH/src/${APP_NAME} 
ADD app ./app
ADD Gopkg.toml Gopkg.lock ./

# Install dependencies
RUN dep ensure -vendor-only