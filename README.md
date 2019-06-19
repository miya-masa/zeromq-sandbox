# Sandbox of ZeroMQ

* https://github.com/zeromq/goczmq



## install 

```
# install zeromq libzmq
brew install zmq

# install libsodium-dev
brew install libsodium

# install czmq
sudo apt-get update
sudo apt-get install -y \
    git build-essential libtool \
    pkg-config autotools-dev autoconf automake cmake \
    uuid-dev libpcre3-dev libsodium-dev valgrind

# only execute this next line if interested in updating the man pages as well (adds to build time):
sudo apt-get install -y asciidoc

# in the git@github.com:zeromq/libzmq.git
git clone git://github.com/zeromq/libzmq.git
cd libzmq
./autogen.sh
# do not specify "--with-libsodium" if you prefer to use internal tweetnacl security implementation (recommended for development)
./configure --with-libsodium
make check
sudo make install
sudo ldconfig
cd ..

git clone git://github.com/zeromq/czmq.git
cd czmq
./autogen.sh && ./configure && make check
sudo make install
sudo ldconfig
cd ..
```
