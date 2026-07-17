# syntax=docker/dockerfile:1

FROM ruby:4.0.6-alpine AS builder

# Install minimal build dependencies for native extensions
RUN apk add --no-cache g++ make

# Pre-install gems needed to run the server (grpc-tools is a protoc
# plugin used only by `make protoc`, not at server runtime)
COPY Gemfile Gemfile.lock ./
RUN --mount=type=cache,target=/usr/local/bundle/cache \
    bundle install --jobs 4

FROM ruby:4.0.6-alpine

COPY --from=builder /usr/local/bundle /usr/local/bundle

# Copy application files
COPY . .

# Set default port
ENV APP_PORT=5001

CMD ["ruby", "server.rb"]
