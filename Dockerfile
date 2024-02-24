FROM alpine
WORKDIR /app

COPY homegpt .
CMD ["./homegpt"]