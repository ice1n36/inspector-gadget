#!/bin/bash
version="12.11.18"
os=android
architecture=arm64
fridagadget=frida-gadget-${version}-${os}-${architecture}.so
# download latest frida-gadget
if [ ! -f ${fridagadget}.xz ]; then
    echo "downloading..."
	wget -q "https://github.com/frida/frida/releases/download/${version}/${fridagadget}.xz"
fi
if [ -f ${fridagadget}.xz ]; then
	unxz --keep ${fridagadget}.xz
fi

echo "moving over to downloads/frida-gadget/${version}/lib/${architecture}/${fridagadget}"
mv ${fridagadget} downloads/frida-gadget/${version}/lib/${architecture}/${fridagadget}

echo "cleaning up..."
rm ${fridagadget}.xz

