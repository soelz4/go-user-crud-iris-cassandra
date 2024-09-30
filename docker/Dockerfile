FROM golang:1.22-alpine3.19 as build
RUN adduser -D nonroot
WORKDIR /app/
COPY . .
RUN apk add --update make
# Goproxy.cn (The most trusted Go module proxy in China.)
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN make build

FROM busybox:stable as run
WORKDIR /app/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /app/bin/main ./
USER nonroot
EXPOSE 9010
CMD [ "./main" ]
