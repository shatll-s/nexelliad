# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/shatll-s/nexelliad

WORKDIR /go/src/github.com/Nexellia-Network/nexelliad

RUN apk add --no-cache curl git openssh binutils gcc musl-dev
#RUN apk add --no-cache git
#RUN mkdir gord
RUN git clone https://github.com/shatll-s/nexelliad.git .

COPY go.mod .
COPY go.sum .


# Cache nexellia dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o nexellia .
RUN chmod 777 nexellia

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/shatll-s/nexelliad/nexelliad /app/
COPY --from=build /go/src/github.com/shatll-s/nexelliad/infrastructure/config/sample-nexelliad.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
