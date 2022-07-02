FROM golang:1.18

# create directory app
RUN mkdir /e-Commerse

# set or make /app our working directory
WORKDIR /e-Commerse

# copy all files to /e-Commerse
COPY . .

RUN go build -o rest-api-e-commerse

CMD ["./rest-api-e-commerse"]
