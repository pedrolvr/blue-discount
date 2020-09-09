FROM golang:1.14.8-alpine3.12 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . /src/
WORKDIR /src/

RUN go build -o /bin/ ./cmd/discount/

FROM scratch
COPY --from=build /bin/discount /bin/discount
ENTRYPOINT ["/bin/discount"]
