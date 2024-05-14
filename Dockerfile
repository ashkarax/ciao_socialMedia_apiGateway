FROM golang:1.22-alpine AS stage1
WORKDIR /project/ciao/

COPY go.* .
RUN  go mod download

COPY . .
RUN go build -o ./cmd/ciaoApiGatewayExec ./cmd/main.go

FROM alpine:latest
WORKDIR /project/ciao/


COPY --from=stage1 /project/ciao/cmd/ciaoApiGatewayExec ./cmd/
COPY --from=stage1 /project/ciao/pkg/config/envs/dev.env ./pkg/config/envs/

EXPOSE 3000
ENTRYPOINT [ "/project/ciao/cmd/ciaoApiGatewayExec" ]