FROM debian

RUN apt-get update
RUN apt-get install -y curl

ADD curl.sh /curl.sh

ENTRYPOINT ["/curl.sh", "172.17.0.1"]

