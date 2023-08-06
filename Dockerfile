FROM scratch
WORKDIR /
COPY ./app /app
ENTRYPOINT /app