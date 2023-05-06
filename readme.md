# cleanarchGo
<div align="center">
  <a href="images/logo.png">
    <img src="images/logo.png" alt="Logo" width="40%">
  </a>
</div>

## 💫 About
Clean architecture starter pack for feature creation to make it faster and easier with Go.
* Clean Architecture with Echo framework, Gorm, Viper (env), Govalidator, Jwt, etc.
* Make CRUD new feature less than 1 minutes.
* Dockerfile will be generated too.

## 🚀 Import
```shell
go get -u github.com/Findryankp/cleanarchGo@latest
```

## 👨🏽‍💻 Step By Step
1. First step, add this syntax to your `main function`
```go
cleanarchGo.Init();
```
* for example :
<div align="left">
  <a href="images/01.png">
    <img src="images/01.png" alt="Logo" width="40%">
  </a>
</div>

2. Run this syntax in cmd/ terminal
```shell
go run . init
```
<div align="left">
  <a href="images/init.png">
    <img src="images/init.gif" alt="Logo" width="60%">
  </a>
</div>

* If success, the files that will be generated are :
- [x] configs database
- [x] middlewares (jwt auth, logs, cors)
- [x] environment (local.env)
- [x] etc

3. Set `local.env` with your own configuration database
<div align="left">
  <a href="images/04.png">
    <img src="images/04.png" alt="Logo" width="40%">
  </a>
</div>

## 🚀 Create new feature
* run this syntax in your cmd/terminal
```shell
go run . features featuresNames
```
* ex : `go run . features rooms`

* CRUD Features, Code all layer, route, and migratiton feature from your featureNames will be created
<div align="left">
  <a href="images/feature.png">
    <img src="images/feature.gif" alt="Logo" width="60%">
  </a>
</div>

## 🎯 Run Project
```shell
go run .
```

* try with your postman or another
<div align="left">
  <a href="images/05.png">
    <img src="images/05.png" alt="Logo" width="60%">
  </a>
</div>

## 😎 Development by
[![Findryankp](https://img.shields.io/badge/Findryankp-grey?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Findryankp)
[![findryankp](https://img.shields.io/badge/findryankp-blue?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/Findryankp/)
