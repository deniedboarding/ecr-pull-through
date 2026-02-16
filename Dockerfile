FROM gcr.io/distroless/static

ENTRYPOINT [ "/mutation-webhook" ]

COPY ./bin/ /
