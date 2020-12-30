#!/bin/bash
version="12.11.18"
packagename="com.riotgames.league.teamfighttactics"
apkname="v10.21.3392173.apk"
os=android
architecture=arm64
fridagadget=frida-gadget-${version}-${os}-${architecture}.so
fridaconfig=frida-gadget-default-config.conf
uberapksignerjar="uber-apk-signer-1.1.0.jar"

# add frida-gadget.so and custom config per platform
if [ ! -d output ]; then
    mkdir output
fi

# todo: support for dex manipulation when there are no native libraries

apktool d -rs -o output/${packagename} downloads/apks/${packagename}/${apkname}

if [ ! -d output/${packagename}/lib ]; then
    mkdir output/${packagename}/lib
    mkdir output/${packagename}/lib/arm64-v8a
fi

# TODO-non-mvp: support for armeabi-v7a
if [ -d output/${packagename}/lib/armeabi-v7a ]; then
    echo "no support for armeabi-v7a"
    exit 1
fi

if [ -d output/${packagename}/lib/arm64-v8a ]; then
    echo "adding frida-gadget.so and config into apk"
    cp downloads/frida-gadget/${version}/lib/arm64/${fridagadget} output/${packagename}/lib/arm64-v8a/libfridagadget.so
    cp config/${fridaconfig} output/${packagename}/lib/arm64-v8a/libfridagadget.config.so
fi

# add gadget to a native library (lief)
echo "injecting gadget to native so"
./inject-gadget-to-native-so.py

apktool b output/${packagename} -o output/${packagename}-with-fridagadget.apk
rm -rf output/${packagename}
# sign apk

echo "re-signing apk..."
java -jar downloads/uber-apk-signer/${uberapksignerjar} --apks output/${packagename}-with-fridagadget.apk



