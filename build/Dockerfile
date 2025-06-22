# syntax=docker.io/docker/dockerfile:1.9.0

ARG GOLANG_ALP_IMAGE=golang:1.24.4-alpine3.21

FROM --platform=$BUILDPLATFORM $GOLANG_ALP_IMAGE AS builder
ARG BUILDARCH
ADD . /go/src/app
WORKDIR /go/src/app
RUN go env -w GOMODCACHE=/gomod-cache
RUN --mount=type=cache,target=/gomod-cache go mod download

FROM --platform=$BUILDPLATFORM builder AS server
ARG TARGETARCH
ARG LDFLAGS=
ARG GOAMD64=
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache GOAMD64=$GOAMD64 CGO_ENABLED=0 GOARCH=$TARGETARCH GOOS=linux go build -trimpath -tags osusergo,netgo -ldflags="-s -w -linkmode=internal -buildid= $LDFLAGS" -gcflags="all= -dwarf=false" -v -o /go/bin/app/server ./cmd/server/

FROM --platform=$BUILDPLATFORM builder AS cli
ARG TARGETARCH
ARG LDFLAGS=
ARG GOAMD64=
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache GOAMD64=$GOAMD64 CGO_ENABLED=0 GOARCH=$TARGETARCH GOOS=linux go build -trimpath -tags osusergo,netgo -ldflags="-s -w -linkmode=internal -buildid= $LDFLAGS" -gcflags="all= -dwarf=false" -v -o /go/bin/app/cli ./cmd/cli/

FROM scratch

COPY --from=server /go/bin/app/server /server
COPY --from=cli /go/bin/app/cli /cli
COPY ./assets /assets

STOPSIGNAL SIGTERM
EXPOSE $APP_HTTP_PORT
HEALTHCHECK --interval=10s --start-period=10s CMD [ "/cli", "healthcheck"]
ENTRYPOINT [ "/server" ]
