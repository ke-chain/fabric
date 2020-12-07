# 重造轮子：Fabric

## 如何安装

1. 下载docker 镜像：hyperledger/fabric-orderer:2.2
2. clone 源码
3. make orderer
4. docker run -v $PWD/build/bin:/usr/local/bin  -it --rm hyperledger/fabric-orderer:2.2 /bin/sh
5. orderer
