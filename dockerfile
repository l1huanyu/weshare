FROM alpine
COPY main /
ENTRYPOINT ["/main"]