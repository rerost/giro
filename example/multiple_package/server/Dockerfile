FROM ruby:3.3.6

COPY Gemfile ./
COPY Gemfile.lock ./

RUN bundle install

COPY . .

CMD ["ruby", "server.rb"]
