#########
# Build #
#########
FROM golang:alpine AS server_builder

ARG TARGETARCH
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk add build-base

COPY . .

RUN GOOS=linux GOARCH=$TARGETARCH go build -a -o vatprc-online .

##########
# Deploy #
##########

FROM alpine

COPY --from=server_builder /build/vatprc-online /
COPY --from=server_builder /build/templates/ /templates/

ENTRYPOINT ["/vatprc-online"]

EXPOSE 9000