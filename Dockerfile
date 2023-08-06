FROM scratch
WORKDIR /
COPY ./backend-binary /app
ENTRYPOINT /app