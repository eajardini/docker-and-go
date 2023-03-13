# docker-and-go

Exemplo de como criar container Docker com golang importando package de níveis acima do contexto.

O dockerfile deve ficar assim:

# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download


COPY ./golibs/*.go ./golibs/
COPY ./gosubapp/router/ ./gosubapp/router/
COPY ./gosubapp/*.go ./

RUN go build -o /appsubexec

EXPOSE 8080

CMD [ "/appsubexec" ]

# Para criar uma imagem docker:  sudo docker build --tag appsubexec . 
# Para rodar como demonio: sudo docker run -d -p 8090:8080 appsubexec
