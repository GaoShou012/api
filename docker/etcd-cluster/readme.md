1:下载docker镜像
docker pull quay.io/coreos/etcd

2:编辑docker-compose.yaml
3:docker-compose up

4:验证集群状态
$ curl -L http://127.0.0.1:32787/v2/members
{"members":[{"id":"ade526d28b1f92f7","name":"etcd1","peerURLs":["http://etcd1:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"bd388e7810915853","name":"etcd3","peerURLs":["http://etcd3:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"d282ac2ce600c1ce","name":"etcd2","peerURLs":["http://etcd2:2380"],"clientURLs":["http://0.0.0.0:2379"]}]}

$ curl -L http://127.0.0.1:32789/v2/members
{"members":[{"id":"ade526d28b1f92f7","name":"etcd1","peerURLs":["http://etcd1:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"bd388e7810915853","name":"etcd3","peerURLs":["http://etcd3:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"d282ac2ce600c1ce","name":"etcd2","peerURLs":["http://etcd2:2380"],"clientURLs":["http://0.0.0.0:2379"]}]}

$ curl -L http://127.0.0.1:32791/v2/members
{"members":[{"id":"ade526d28b1f92f7","name":"etcd1","peerURLs":["http://etcd1:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"bd388e7810915853","name":"etcd3","peerURLs":["http://etcd3:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"d282ac2ce600c1ce","name":"etcd2","peerURLs":["http://etcd2:2380"],"clientURLs":["http://0.0.0.0:2379"]}]}

docker exec -t etcd1 etcdctl member list

curl -L http://127.0.0.1:32789/v2/keys/foo -XPUT -d value="Hello foo"
curl -L http://127.0.0.1:32789/v2/keys/foo1/foo1 -XPUT -d value="Hello foo1"
curl -L http://127.0.0.1:32789/v2/keys/foo2/foo2 -XPUT -d value="Hello foo2"
curl -L http://127.0.0.1:32789/v2/keys/foo2/foo21/foo21 -XPUT -d value="Hello foo21"

curl -L http://127.0.0.1:32787/v2/keys/foo
curl -L http://127.0.0.1:32787/v2/keys/foo2
curl -L http://127.0.0.1:32787/v2/keys/foo2?recursive=true