ARG appName=app

FROM golang:1.22.4-alpine3.19 AS build
ARG appName
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY source/go.mod source/go.sum source/main.go ./
COPY source/src src
COPY source/docs docs
COPY source/vendor vendor

ENV CGO_ENABLED=0
RUN go build -o ${appName}

FROM scratch
ARG appName

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy timezone data, which avaiable after installing tzdata
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta

COPY --from=build /app/${appName} /app

ENTRYPOINT [ "/app" ]