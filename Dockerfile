FROM golang:1.21.1 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -v -o gpt-i18n

# Use a minimal image to run the compiled program
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/gpt-i18n /gpt-i18n

RUN uname -m

CMD ["/gpt-i18n"]
