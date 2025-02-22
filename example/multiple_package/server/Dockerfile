FROM ruby:3.4.2-slim

# Install minimal build dependencies for native extensions
RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y \
    g++ \
    make \
    && rm -rf /var/lib/apt/lists/* && \
    apt-get clean

# Pre-install common gems with caching
COPY Gemfile Gemfile.lock ./
RUN --mount=type=cache,target=/usr/local/bundle \
    bundle install --jobs 4

# Copy application files
COPY . .

CMD ["ruby", "server.rb"]
