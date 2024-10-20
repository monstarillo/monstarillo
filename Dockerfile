# syntax=docker/dockerfile:1

FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download



# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
RUN ls /app
COPY . .
RUN ls -la /app


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /monstarillo



#RUN mkdir -m777 /usr/local/monstarillo
#RUN mkdir -m777 /usr/local/monstarillo/templates
#RUN mkdir -m777 /usr/local/monstarillo/output
#RUN mkdir -m777 /usr/local/monstarillo/output3

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
#EXPOSE 8080

# Run
RUN addgroup --gid 1000 user
RUN adduser --disabled-password --gecos '' --uid 1000 --gid 1000 user
USER user

ENTRYPOINT ["/monstarillo"]