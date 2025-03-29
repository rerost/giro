FROM ruby:3.4.2-slim

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

# Create empty.proto file if it doesn't exist
RUN mkdir -p /usr/include/google/protobuf && \
    if [ ! -f /usr/include/google/protobuf/empty.proto ]; then \
      echo 'syntax = "proto3";\npackage google.protobuf;\nmessage Empty {}\n' > /usr/include/google/protobuf/empty.proto; \
    fi

CMD ["ruby", "server.rb"]
