build docker image => sudo docker build -t bsmaja{version} .
start container => docker run -d --restart unless-stopped -p 80:80 bsmaja_0_0_1
