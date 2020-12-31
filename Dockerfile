FROM golang:1.14

COPY . /ne_cache
RUN make -C /ne_cache all

RUN chmod -R a+x /ne_cache/bin
