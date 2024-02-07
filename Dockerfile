ARG GOLANG_VERSION=1.22
FROM --platform=$BUILDPLATFORM golang:${GOLANG_VERSION} as builder
ARG TARGETOS TARGETARCH

WORKDIR /workspace

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -gcflags 'all=-N -l' -o access-management main.go

RUN chmod a+rx access-management

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/base:latest as serve
WORKDIR /
COPY --from=builder /workspace/access-management .

EXPOSE 8082

CMD ["/access-management"]
