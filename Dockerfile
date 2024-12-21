# Stage 1: Build
FROM golang:1.22 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make build-release

# Stage 2: Production
FROM scratch as production

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/<%= projectName %> /bin/<%= projectName %>

# Command to run the executable
ENTRYPOINT [ "/bin/<%= projectName %>" ]
