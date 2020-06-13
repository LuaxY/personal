# Personal Website

### Build & Run

```shell script
docker build -t gcr.io/voidmx/personal .
docker push gcr.io/voidmx/personal
docker pull gcr.io/voidmx/personal
docker stop personal && docker rm personal
docker run --name personal -d --restart=always -p 80:80 gcr.io/voidmx/personal
```