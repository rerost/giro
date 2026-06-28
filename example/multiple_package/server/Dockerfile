# syntax=docker/dockerfile:1

FROM ruby:4.0.5-slim AS builder

# Install minimal build dependencies for native extensions
RUN apt-get update && apt-get install -y \
    g++ \
    make \
    && rm -rf /var/lib/apt/lists/* && \
    apt-get clean

# Pre-install gems needed to run the server (grpc-tools is a protoc
# plugin used only by `make protoc`, not at server runtime)
COPY Gemfile Gemfile.lock ./
RUN --mount=type=cache,target=/usr/local/bundle/cache \
    bundle install --jobs 4

FROM ruby:4.0.5-slim

COPY --from=builder /usr/local/bundle /usr/local/bundle

# Copy application files
COPY . .

# Set default port
ENV APP_PORT=5001

CMD ["ruby", "server.rb"]
