FROM golang:1.13
LABEL maintainer="Robbie Heywood <r.a.heywood@btinternet.com>"

COPY ./ /house-app/

WORKDIR /house-app/services/house
RUN go build -o /bin/house

CMD ["house"]
EXPOSE 8080/tcp
