#!/bin/bash

mkdir -p ~/bin
mv sync-tmp.sh ~/bin/bsync
chmod +x ~/bin/bsync

echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc