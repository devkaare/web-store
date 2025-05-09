FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate
RUN curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss && \
    chmod +x tailwindcss && \
    ./tailwindcss -i views/assets/css/input.css -o views/assets/css/output.css
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

EXPOSE ${PORT}
CMD ["./main"]
