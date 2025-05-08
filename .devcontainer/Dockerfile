FROM mcr.microsoft.com/devcontainers/cpp:1-debian-11

# install dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    libconfig-dev && \
    rm -rf /var/lib/apt/lists/*