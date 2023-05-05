# cleanarchGo
Feature creation to make it faster and easier.
* Clean Architecture with Echo framework, Gorm, Viper (env), Govalidator, Jwt, etc.
* Make CRUD new feature less than 1 minutes.
* Dockerfile will be generated too.

## Import
```shell
go get -u github.com/Findryankp/cleanarchGo@latest
```

## Step By Step
1. First step, add this go syntax to your **main** function
```go
if err := cleanarchGo.Init(); err != nil {
  fmt.Println(err)
  return
}
```
for example :
<br/>
<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/01.png" alt="Logo" height="200" width="400">
  </a>
</div>

2. Run this syntax in cmd/terminal
```shell
go run . init
```

3. If success, the files that will be generated are :
* configs database
* middlewares (jwt auth, logs, cors)
* environment (local.env)
* etc
<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/02.png" alt="Logo" height="250" width="500">
  </a>
</div>  
<br/>

1. Set .env with your own configuration database

<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/env.png" alt="Logo" height="150" width="300">
  </a>
</div>

## Create new feature
run this syntax in your cmd/terminal
```shell
go run . features featuresNames
```
ex : go run . features room

* All layer, route, and migratiton feature from your featureNames will be created
<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/03.png" alt="Logo" height="180" width="400">
  </a>
</div>

### Run Project
```shell
go run .
```

try it with your postman or another
<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/04.png" alt="Logo" height="200" width="380">
  </a>
</div>

## Development by
[![Findryankp](https://img.shields.io/badge/Findryankp-grey?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Findryankp)
[![findryankp](https://img.shields.io/badge/findryankp-blue?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/Findryankp/)
