FROM ruby:2.7.8

COPY Gemfile ./
COPY Gemfile.lock ./

RUN bundle install

COPY . .

CMD ["ruby", "server.rb"]
