FROM ruby:3.4.2-slim as builder

# Install minimal build dependencies
RUN apt-get update && apt-get install -y \
    gcc \
    make \
    && rm -rf /var/lib/apt/lists/*

# Install gems
COPY Gemfile Gemfile.lock ./
RUN bundle install

FROM ruby:3.4.2-slim

# Copy only necessary files
COPY --from=builder /usr/local/bundle /usr/local/bundle
COPY . .

CMD ["ruby", "server.rb"]
