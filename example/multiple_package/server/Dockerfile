FROM ruby:3.4.2-slim

# Install minimal build dependencies
RUN apt-get update && apt-get install -y \
    gcc \
    make \
    && rm -rf /var/lib/apt/lists/*

# Install gems with parallel jobs
COPY Gemfile Gemfile.lock ./
RUN bundle config set --local jobs 4 && \
    bundle install

COPY . .

CMD ["ruby", "server.rb"]
