#!/bin/bash

# ---------------------------
# cd <installation_dir>
# curl --silent -L https://raw.githubusercontent.com/tpkn/seqd/main/install.sh 2> /dev/null | bash
# ---------------------------

release_url="https://github.com/tpkn/seqd/releases/latest/download/seqd"

binary_name=$(cut -d '/' -f 9 <<< "$release_url")
binary_path="$PWD/$binary_name"
usr_bin_path="/usr/bin"

echo "Installing: $release_url ..."

status=$(curl --write-out %{http_code} --silent --fail -LO "$release_url")
if (( $status != 200 )); then
	echo "[x] Can't download binary: $release_url"
	exit 1
fi

echo "[✓] Downloaded: $binary_path"

if ! chmod -R 0750 "$binary_path" &> /dev/null; then
	echo "[x] Can't change permissions to '$binary_path'"
	exit 1
fi

echo "[✓] Permissions changed to 'rwxr-x---'"

while true; do
	read -p "Create alias '$binary_name' for binary? [y/n] " q
	case $q in
	[Yy]* )
			if ! command -v "$binary_name" &> /dev/null; then
				if ! ln -s "$binary_path" "$usr_bin_path/$binary_name" &> /dev/null; then
					echo "[x] Can't create alias '$binary_name' for '$binary_path'"
				fi
			else
				echo "[-] Alias '$binary_name' already exists"
			fi
		break;;
	* ) break ;;
	esac
done

echo "[✓] Done: v$($binary_path --version)"
