FROM ruby:3.4.4-slim

# Install minimal build dependencies for native extensions
RUN apt-get update && apt-get install -y \
    g++ \
    make \
    protobuf-compiler \
    libprotobuf-dev \
    && rm -rf /var/lib/apt/lists/* && \
    apt-get clean

# Pre-install common gems
COPY Gemfile Gemfile.lock ./
RUN bundle install --jobs 4 && \
    gem install grpc-tools

# Copy application files
COPY . .

# Set default port
ENV APP_PORT=5001

CMD ["ruby", "server.rb"]
