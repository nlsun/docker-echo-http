FROM alpine:3.5

COPY start /start

ENTRYPOINT []
CMD ["/start", "80"]
