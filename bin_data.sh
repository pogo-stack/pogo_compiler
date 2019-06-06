if [ ! -f bindata.go ]; then
  go get -v -u github.com/jteeuwen/go-bindata/...
fi

if [ -f bindata.go ]; then
   rm bindata.go
fi

go-bindata -o bindata.go ./templates/...
