FROM alpine
RUN apk add redis
COPY callisto-golang .
COPY run.sh .
RUN chmod +x run.sh
ENTRYPOINT ./run.sh
EXPOSE 8080