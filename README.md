# Service to call NBP

## How to run?


There are two ways to run this program

1. Via go run .

IF you want to run the program via go run .

### checkHost command
```
go run main.go checkHost --host="http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json" --x=5 --y=4 
```

### currencyCheck command
```
go run main.go currencyCheck --host="http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json" --x=5 --y=4 
```


The log.txt file should be under /data folder. 


2. Via docker

### checkHost command

In the main folder run to build a docker
```
docker build -t nbp_caller .
```

The next step is to make use of build docker image
```
docker run --mount source=myvol2,target=/root/data nbp_caller checkHost --host="http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json" --x=5 --y=4
```

If you want to check logs from docker volume you can make the following call
```
docker run --rm -i -v=myvol2:/root/data busybox cat /root/data/log.txt
```



### currencyCheck command

In the main folder run to build a docker
```
docker build -t nbp_caller .
```

The next step is to make use of build docker image
```
docker run --mount source=myvol2,target=/root/data nbp_caller currencyCheck --x=5 --y=4
```

If you want to check logs from docker volume you can make the following call
```
docker run --rm -i -v=myvol2:/root/data busybox cat /root/data/log.txt
```