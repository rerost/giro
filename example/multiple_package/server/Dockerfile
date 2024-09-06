FROM ruby:3.3.5

COPY Gemfile ./
COPY Gemfile.lock ./

RUN bundle install

COPY . .

CMD ["ruby", "server.rb"]
