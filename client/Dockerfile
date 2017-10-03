FROM debian

RUN apt-get update
RUN apt-get install -y curl

ADD rm.sh /rm.sh

ENTRYPOINT ["/rm.sh"]

