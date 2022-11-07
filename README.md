# Get file URL and download it to local machine
```
go run main.go
```
This command will do a http request to ``https://php-noise.com/noise.php`` endpoint
and will construct a struct with parameter ``uri`` and then will download it and store at 
``images`` folder with random name.

## Install
```
make up
```
To create images/ folder