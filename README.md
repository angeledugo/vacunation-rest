# vacunation-rest
vacunation rest test

# database create
docker build . -t vacunation-rest-db
docker run -p 54321:5432 vacunation-rest-db
