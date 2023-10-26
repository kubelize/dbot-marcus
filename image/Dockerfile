FROM debian:trixie-slim

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*
RUN groupadd dbot && useradd dbot -g dbot
WORKDIR /app

COPY schaebigctl /app
RUN chmod +x schaebigctl

USER dbot

ENTRYPOINT [ "./schaebigctl" ]