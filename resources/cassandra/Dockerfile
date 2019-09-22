FROM cassandra:3.11.4

EXPOSE 9042

RUN apt update -y && apt upgrade -y
RUN apt install -y net-tools

COPY dump/ /opt/data
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]