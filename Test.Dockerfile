 # @author [Gopi Karmakar]
 # @email [gopi.karmakar@monstar-lab.com]
 # @create date 2018-01-26 05:02:02
 # @modify date 2018-01-26 05:02:02
 # @desc [description]

ARG BASE_TAG
ARG IMAGE_NAME
FROM ${IMAGE_NAME}-base:${BASE_TAG}

RUN go get -u github.com/jstemmer/go-junit-report

CMD go test ./app/... -v 2>&1 | go-junit-report > report/report.xml