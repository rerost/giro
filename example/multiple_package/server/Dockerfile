FROM ruby:3.4.2-slim

# Install minimal build dependencies for native extensions
RUN apt-get update && apt-get install -y \
    g++ \
    make \
    && rm -rf /var/lib/apt/lists/* && \
    apt-get clean

# Pre-install common gems
COPY Gemfile Gemfile.lock ./
RUN bundle install --jobs 4

# Copy application files
COPY . .

CMD ["ruby", "server.rb"]
