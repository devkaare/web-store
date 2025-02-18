FROM golang:1.23 AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate 

RUN go build -o /main cmd/api/main.go

FROM base AS dev 
RUN go install github.com/air-verse/air@latest
EXPOSE ${PORT}
CMD ["air", "-c", ".air.toml"]

FROM base AS prod
WORKDIR / 
COPY --from=base /main /main 
EXPOSE ${PORT}
CMD ["/main"]
