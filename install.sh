#!/bin/sh

latest=$(curl -s https://api.github.com/repos/mbenaiss/openai/releases/latest | jq -r ".tag_name")

os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m | tr '[:upper:]' '[:lower:]')
 
if [[ $os == *"darwin"* ]]; then
    os_arch="darwin_$arch"
elif [[ $os == *"linux"* ]]; then
  os_arch="linux_$arch"
fi


echo "Downloading $latest for $os_arch"

filename="openai_${latest:1}_$os_arch.tar.gz"

# download the latest release
curl -sL https://github.com/mbenaiss/openai/releases/download/$latest/$filename -o $filename 
tar -xzf $filename -C /usr/local/bin
rm -rf $filename
