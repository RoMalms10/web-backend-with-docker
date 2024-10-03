# Specifies a parent image
# In this case, the latest golang image
FROM golang:latest

# Install the DB
RUN apt update
RUN apt install sqlite3

# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .

# Effectively tracks changes within your go.mod file
COPY go.mod .
 
# Installs Go dependencies
RUN go mod download

# Copies your source code into the app directory
COPY main.go .

# Build the golang executable 
# Named "app" so that it can be run with just "./app"
RUN go build -o app
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
# Without the build command, it can be run with: CMD [ "go", "run", "main.go" ]
CMD [ "./app" ]