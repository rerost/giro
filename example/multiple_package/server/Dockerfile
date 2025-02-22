FROM ruby:3.2-slim

# Install minimal build dependencies for native extensions
RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y \
    g++ \
    make \
    pkg-config \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/* && \
    apt-get clean

# Create and set working directory
WORKDIR /app

# Copy Gemfile first for better caching
COPY Gemfile Gemfile.lock ./

# Install gems with proper native extension support
RUN --mount=type=cache,target=/usr/local/bundle \
    bundle config set --local without 'development test' && \
    bundle install --jobs 4 --retry 3

# Copy application files
COPY . .

# Verify protobuf installation and precompile native extensions
RUN bundle exec ruby -e "require 'google/protobuf'; require 'grpc'" || exit 1

CMD ["bundle", "exec", "ruby", "server.rb"]
