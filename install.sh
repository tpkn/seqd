#!/usr/bin/env bash
# ---------------------------
# curl --silent -L https://raw.githubusercontent.com/tpkn/seqd/main/install.sh 2> /dev/null | sudo bash
# ---------------------------
release_url="https://github.com/tpkn/seqd/releases/latest/download/seqd"

binary_name=$(cut -d '/' -f 9 <<< "$release_url")
binary_path="/usr/local/bin/$binary_name"
binary_path_alt="/usr/bin/$binary_name"

sudo -v &> /dev/null && echo "Downloading: $release_url ..." || echo "[x] You are not a sudo user"

# Check if there is a '/usr/local/bin' in $PATH
if ! grep -q '/usr/local/bin/' <<< "$PATH"; then
	binary_path=$binary_path_alt
	echo "[!] There is no '/usr/local/bin' path in $PATH, installing into '/usr/bin'"
fi

status=$(curl --fail -sLo "$binary_path" "$release_url" --write-out %{http_code})
if (( $status != 200 )); then
	echo "[x] Can't download binary: $release_url"
	exit 1
fi

echo "[✓] Installed: $binary_path"

if ! chmod 0755 "$binary_path" &> /dev/null; then
	echo "[x] Can't change permissions for '$binary_path'"
	exit 1
fi

echo "[✓] Permissions changed to 755"
echo "[✓] Done: v$($binary_path --version)"
