FROM golang:1.17-alpine AS build-env
WORKDIR /app
RUN go mod init app-go
COPY . ./
RUN go build -o /app/app-go
FROM  golang:1.17-alpine AS runtime
WORKDIR /app
COPY --from=build-env /app/templates ./templates
COPY --from=build-env /app/app-go ./
CMD ["/app/app-go"]

# FROM  golang:1.17-alpine AS runtime
# WORKDIR /app
# RUN go mod init app-go
# COPY . ./
# RUN go build -o /app/app-go
# CMD ["/app/app-go"]