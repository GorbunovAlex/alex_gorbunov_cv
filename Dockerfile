# Use the latest Go version
FROM golang:latest

# Set the working directory inside the container
WORKDIR /home/alexandergorbunov/__dev__/pets/alex_gorbunov_cv

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Expose port 3000 to the outside world
EXPOSE 3000

# Define a volume to persist changes
VOLUME ["/home/alexandergorbunov/__dev__/pets/alex_gorbunov_cv"]
