if [ ! -f "/c/msys64/usr/bin/bash.exe"]; then
    winget install -e --id MSYS2.MSYS2
else
    echo "You already have MSYS2, skipping installation."
fi

export PATH="/c/msys64/ucrt64/bin:$PATH"

gcc --version
sleep 5