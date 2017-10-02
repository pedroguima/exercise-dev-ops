FROM debian 
ARG version
ENV version=$version

RUN apt-get update && apt-get install -y liblmdb-dev 
ADD backend_${version}_linux64 /backend

EXPOSE 9999

CMD ["/backend"]
