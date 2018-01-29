 # @author [Gopi Karmakar]
 # @email [gopi.karmakar@monstar-lab.com]
 # @create date 2018-01-26 05:02:02
 # @modify date 2018-01-26 05:02:02
 # @desc [description]

ARG BASE_TAG
ARG IMAGE_NAME
FROM ${IMAGE_NAME}-base:${BASE_TAG}

# Compile binary
ARG APP_NAME
RUN go install ${APP_NAME}/app

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD curl -f http://localhost/healthz || exit 1

# Run program
ENTRYPOINT ["/go/bin/app"]

# Document that the service listens on port 80.
EXPOSE 80