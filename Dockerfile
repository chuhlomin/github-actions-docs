FROM golang:1.18 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=readonly -a -installsuffix cgo \
    -o /go/bin/app .


FROM gcr.io/distroless/static:966f4bd97f611354c4ad829f1ed298df9386c2ec
# latest-amd64 -> 966f4bd97f611354c4ad829f1ed298df9386c2ec
# https://github.com/GoogleContainerTools/distroless/tree/master/base

LABEL name="github-actions-docs"
LABEL repository="http://github.com/chuhlomin/github-actions-docs"
LABEL homepage="http://github.com/chuhlomin/github-actions-docs"

LABEL maintainer="Konstantin Chukhlomin <mail@chuhlomin.com>"
LABEL com.github.actions.name="GitHub Actions Docs"
LABEL com.github.actions.description="A Github action for generating docs for GitHub actions using Go template."
LABEL com.github.actions.icon="file-text"
LABEL com.github.actions.color="gray-dark"

COPY --from=build-env /go/bin/app /app

CMD ["/app"]
