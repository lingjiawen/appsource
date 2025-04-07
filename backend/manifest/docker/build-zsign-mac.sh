cd zsign/build/macos
make clean && make all
sudo mkdir /usr/local/lib/
sudo cp ../../bin/libzsign.so /usr/local/lib/
sudo chmod +x /usr/local/lib/libzsign.so
sudo mkdir /usr/local/include
sudo cp -r ../../src /usr/local/include/zsign