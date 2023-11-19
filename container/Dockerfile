FROM ubuntu:latest

# Install Python
RUN apt-get update && apt-get install -y python3

#RUN apt-get install libssl-dev
RUN apt-get install curl -y

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_19.x | bash -
RUN apt-get install -y nodejs

RUN apt-get install unzip

# Install Deno
RUN curl -fsSL https://deno.land/x/install/install.sh | sh


# Add the runner binary
ADD runner /usr/local/bin

# Set the working directory
WORKDIR /app

# Run the runner binary when the container starts
ENTRYPOINT ["runner"]
