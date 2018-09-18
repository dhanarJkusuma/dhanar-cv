#! /bin/sh 

# Change Directory Bin
DIR="$( cd "$( dirname $0 )" >/dev/null && pwd )"
echo "[Portofolio App] Changing directory $DIR--"

echo "[Portofolio App] Checking Bin"
appBin="$DIR/bin"
if [ -f "$appBin/main" ]
then
    eval "$appBin/main"
else
    echo "[Portofolio App] Trying to Installing Portofolio Application" 
    # Change Directory App/Bin
    cdCmd="cd $appBin/../" 
    eval $cdCmd
    eval "go build -o bin/main ."
fi
