FROM golang:1.20

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go

RUN go build -o whatfs main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=0 /workspace/whatfs .
USER 65532:65532

ENTRYPOINT ["/whatfs"]
