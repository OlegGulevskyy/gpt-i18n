FROM golang:1.21.1 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

RUN go build -v -o gpt-i18n
# RUN GOOS=linux GOARCH=amd64 go build -v -o gpt-i18n

# Use a minimal image to run the compiled program
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/gpt-i18n /gpt-i18n

# CMD ["/gpt-i18n"]
CMD file /gpt-i18n
