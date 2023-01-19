## Absen App

Rest API CRUD User dan User Login, dengan JWT token/refreshToken

### App Stacks
 - Backend: golang dengan framework fiber, mysql dengan orm gorm
 - Deployments: docker, kubernetes. Kubernetes jalan pada pada ubuntu desktop dengan installer dari minikube.

#### Menjalankan aplikasi dengan docker

 - dev
	```
	$ docker-compose up -d --build
	```
 - prod
	```
	$ docker build -t absen-go .
	$ docker run -d --name absen-go -e <ENV> -p 3000:3000 absen-go 
	```
	envoiremnts silahkan lihat di file docker-compose.yml

#### Menjalankan aplikasi dengan kubernetes cluster
 1. **setup mysql** 
 	```
	$ kubectl create -f k8s.mysql.secret.yaml
	$ kubectl apply -f k8s.mysql.pv.yaml
	$ kubectl apply -f k8s.mysql.pvc.yaml
	$ kubectl apply -f k8s.mysql.deployment.yml
	$ kubectl apply -f k8s.mysql.service.yml
	```

#### Mengakses API dengan Postman client
 1. download postman colections dari link ini
    [https://www.getpostman.com/collections/3797c3347deb99272049](https://www.getpostman.com/collections/3797c3347deb99272049)
 2. saat app dijalankan telah otomatis dibuat seeder data user dengan level admin. dengan credential
    - email=emixbal@gmail.com
    - password=aaaaaaaa
    dengan credential diatas gunakan request "login refresh token" untuk login
 3. ketika mendapat access access_token & refresh_token buat envoirement, dengan key
    - baseUrl, lalau isikan value link yg telah digenarate
    - jwtToken, lalu isikan dengan access_token yg didapat
    - jwtRefreshToken, lalu isikan dengan refresh_token yg didapat

#### Auth DIagram
![all pods](https://raw.githubusercontent.com/emixbal/absen-go/main/images/Picture1.jpg)
#### Refresh Token DIagram
mengutip dari [https://www.alemba.help/help/content/topics/alemba%20api/aa%20programmers%20guide.htm](https://www.alemba.help/help/content/topics/alemba%20api/aa%20programmers%20guide.htm)

![all pods](https://raw.githubusercontent.com/emixbal/absen-go/main/images/refresh%20token.jpg)