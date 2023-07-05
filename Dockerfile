FROM --platform=$BUILDPLATFORM golang:1.19 as builder
ENV TZ=Asia/Shanghai LANG="C.UTF-8"
ARG TARGETARCH
ARG TARGETOS

WORKDIR /workspace
COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o accesstokend cmd/main.go

FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
EXPOSE 8080
WORKDIR /
COPY --from=builder /workspace/accesstokend .
ENTRYPOINT ["/accesstokend","akt"]