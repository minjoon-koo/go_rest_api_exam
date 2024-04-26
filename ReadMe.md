# SLDT Go lang Study Example code
### Fiber 프레임워크를 이용하여 간단한 restful api를 구성하고자 함

### version info
실행 전 반드시 설치
```
go : go1.22.1
fiber : 2.52.4-beta.2
```

## 설치 패키지
### fiber/v2 메인 프레임워크
```
go get github.com/gofiber/fiber/v2
```
### godotenv env 호출
```
go get github.com/joho/godotenv
```
### gorm db ORM
```
go get -u gorm.io/gorm
```
### mysql driver
```
go get gorm.io/driver/mysql
```
### jwt token module
```
go get github.com/golang-jwt/jwt
```

## local 테스트용 db
./testDB/docker-compose.yml

### local 실행
```shell
cd testDB
docker-compose up -d 
```


## local 에서 fiber rest api 바로 실행
```shell
git clone https://github.com/minjoon-koo/go_rest_api_exam.git

go run main.gp
```
