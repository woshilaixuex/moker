#!/bin/bash
#docker network create moker-etcd-network
#docker run -d --name moker-etcd \
#  --network my-etcd-network \
#  -e ETCD_NAME=node1 \
#  -e ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster \
#  -e ETCD_INITIAL_ADVERTISE_PEER_URLS=http://my-etcd:2380 \
#  -e ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380 \
#  -e ETCD_ADVERTISE_CLIENT_URLS=http://my-etcd:2379 \
#  -e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
#  bitnami/etcd:3.5.13
docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --name etcd \
  -e ETCD_NAME=my-etcd-1 \
  -e ETCD_DATA_DIR=/etcd-data \
  -e ALLOW_NONE_AUTHENTICATION=yes \
  -e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
  -e ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
  -e ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380 \
  -e ETCD_INITIAL_ADVERTISE_PEER_URLS=http://0.0.0.0:2380 \
  -e ETCD_INITIAL_CLUSTER=my-etcd-1=http://0.0.0.0:2380 \
  -e ETCD_INITIAL_CLUSTER_STATE=new \
  bitnami/etcd:3.5.13



sudo iptables -I INPUT -s 10.33.77.1 -p tcp --dport etcd_port -j ACCEPT
docker exec -it moker-etcd /bin/sh
etcdctl user add your_username --password your_password