FROM ruby:3.4.2-slim

# Install minimal build dependencies for native extensions
RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y \
    g++ \
    make \
    pkg-config \
    && rm -rf /var/lib/apt/lists/* && \
    apt-get clean

# Create and set working directory
WORKDIR /app

# Copy Gemfile first for better caching
COPY Gemfile Gemfile.lock ./

# Install gems with proper native extension support
RUN --mount=type=cache,target=/usr/local/bundle \
    bundle config set --local build.google-protobuf --with-cflags="-fPIC" && \
    bundle install --jobs 4

# Copy application files
COPY . .

# Verify protobuf installation
RUN ruby -e "require 'google/protobuf'" || exit 1

CMD ["ruby", "server.rb"]
