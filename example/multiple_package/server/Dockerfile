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

# Create symbolic link for google/protobuf directory
RUN mkdir -p /usr/include/google/protobuf && \
    ln -sf /usr/include/google/protobuf/empty.proto /usr/include/google/protobuf/empty.proto

CMD ["ruby", "server.rb"]
