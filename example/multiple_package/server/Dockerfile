FROM ruby:3.4.2

COPY Gemfile ./
COPY Gemfile.lock ./

RUN bundle install

COPY . .

CMD ["ruby", "server.rb"]
