docker build -t gsocks5:0.2 .
docker run -d -e LOG_LEVEL=INFO -p 1080:1080 gsocks5:0.2