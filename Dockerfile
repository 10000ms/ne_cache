FROM golang:1.14

COPY . /ne_cache
RUN make -C /ne_cache all

CMD ["/ne_cache/bin/$SERVICE_NAME", "-uuid=$UUID"]
