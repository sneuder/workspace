FROM golang:1.21.1

WORKDIR /app

RUN apt update -y
RUN apt upgrade -y

RUN apt install git -y
RUN apt install make -y
RUN apt install zip -y

CMD ["sleep", "infinity"]
