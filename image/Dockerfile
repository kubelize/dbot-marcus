FROM debian:trixie-slim

RUN groupadd dbot && useradd dbot -g dbot
WORKDIR /app

COPY schaebigctl /app
RUN chmod +x schaebigctl

USER dbot

ENTRYPOINT [ "./schaebigctl" ]