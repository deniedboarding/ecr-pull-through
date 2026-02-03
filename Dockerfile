FROM gcr.io/distroless/base-debian10

ENTRYPOINT [ "/mutation-webhook" ]

COPY ./bin/ /
