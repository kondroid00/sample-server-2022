FROM golang:1.18.3-alpine3.16 as builder

ENV ROOT=/go/src/github.com/kondroid00/sample-server-2022/main
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

ADD . ${ROOT}

RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/binary main.go

WORKDIR ${ROOT}/migration
RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/migration/binary main.go


FROM scratch as runner

ENV ROOT=/go/src/github.com/kondroid00/sample-server-2022/main
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}


FROM scratch as migration

ENV ROOT=/go/src/github.com/kondroid00/sample-server-2022/main
WORKDIR ${ROOT}/migration
COPY --from=builder ${ROOT}/binary ${ROOT}