# Dockerfile for Phantom Masternode Frontend
# Phantom Hosting - 2020-12-28

FROM photon:3.0-20201218
LABEL Name="Dockerized Phantom Nodes"
LABEL Publisher="GOSSIP Blockchain"

RUN tdnf install git -y

WORKDIR /root/phantom-hosting/

RUN git clone https://github.com/GOSSIP-Blockchain/GOSSIP-phantom-hosting.git /root/phantom-hosting/
ADD https://github.com/GOSSIP-Blockchain/GOSSIP-phantom-hosting/releases/download/v0.0.1/phantom-hosting-linux-amd64 /root/phantom-hosting/
RUN chmod a+x /root/phantom-hosting/phantom-hosting-linux-amd64

ENTRYPOINT ["/root/phantom-hosting/phantom-hosting-linux-amd64"]

EXPOSE 9999
