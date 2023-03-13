ver="1.0.9"

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o target/bookget_v${ver}_windows/bookget.exe .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/bookget_v${ver}_linux/bookget .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/bookget_v${ver}_macOS/bookget .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o target/bookget_v${ver}_macOs_M2/bookget .


cp cookie.txt target/bookget_v${ver}_linux/cookie.txt
cp cookie.txt target/bookget_v${ver}_macOS/cookie.txt
cp cookie.txt target/bookget_v${ver}_macOS_M2/cookie.txt
cp cookie.txt target/bookget_v${ver}_windows/cookie.txt


cd target/ 
7za a -t7z bookget_v${ver}_windows.7z bookget_v${ver}_windows
tar cjf bookget_v${ver}_linux.tar.bz2 bookget_v${ver}_linux
tar cjf bookget_v${ver}_macOS.tar.bz2 bookget_v${ver}_macOS
tar cjf bookget_v${ver}_macOS_M2.tar.bz2 bookget_v${ver}_macOS_M2
