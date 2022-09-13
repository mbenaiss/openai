#!/bin/sh

latest=$(curl -s https://api.github.com/repos/mbenaiss/openai/releases/latest | jq -r ".tag_name")

os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m | tr '[:upper:]' '[:lower:]')
 
if [[ $os == *"darwin"* ]]; then
  if [[ $arch == *"x86_64"* ]]; then
    os_arch="darwin_amd64"
  else
    os_arch="darwin_$arch"
  fi
elif [[ $os == *"linux"* ]]; then
  os_arch="linux_$arch"
fi


echo "Downloading $latest for $os_arch"

filename="openai_${latest:1}_$os_arch.tar.gz"

# download the latest release
curl -sLO https://github.com/mbenaiss/openai/releases/download/$latest/$filename 
tar -xzf $filename -C /usr/local/bin
rm -rf $filename
