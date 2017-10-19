FROM debian 
ARG version
ENV version=$version

ADD server_${version}_linux64 /server
ADD testdata/ /testdata/

EXPOSE 8888 

CMD ["/server"]
