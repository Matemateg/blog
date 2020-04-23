FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN go build -o /blog .

FROM heroku/heroku:16
COPY ./templates /app/templates
WORKDIR /app
COPY --from=build /blog /blog
CMD ["/blog"]