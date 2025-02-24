FROM golang:1.22

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# Instala pacotes necessários
RUN apt-get update && \
    apt-get install -y build-essential git curl && \
    git clone --depth 1 --branch v1.9.2 https://github.com/confluentinc/librdkafka.git /tmp/librdkafka && \
    cd /tmp/librdkafka && \
    ./configure --prefix=/usr && \
    make && make install && \
    rm -rf /tmp/librdkafka


CMD ["tail", "-f", "/dev/null"]