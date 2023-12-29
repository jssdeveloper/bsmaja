build docker image => sudo docker build -t bsmaja_0_1_x --no-cache .
start container => docker run -d --restart unless-stopped -p 80:80 bsmaja_0_1_x
