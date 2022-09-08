CRUD API
# Table of Contents
- [Installation](#installation)
- [How To Use](#howtouse)
- [Credentials](#credentials)

## Installation

### Prerequisites
- Golang(>=1.11) - Download and Install [Golang](https://golang.org/)
- Redis - Download and Install [Redis](https://redis.io/download)
- MySQL - Download and Install [MySQL](https://www.apachefriends.org/download.html)
- Docker - Download and Install [Docker](https://www.docker.com/products/docker-desktop)
- Kubectl - Download and Install [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- Virtualbox - Download and Install [Virtualbox](https://www.virtualbox.org/)
- Minikube - Download and Install [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- Postman - Download and Install [Postman](https://www.postman.com/)

# Diagram
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEh7YGefgnENMB9Oo_9Wc0RI-G4Tj4mcmhTvG7qWXBZwquzgnE1qj3CNiI94RLUs_GOnTRpLQZZBZ2Dp8BvkhRKMB-sGPe6r98omc8gZ7VdDdSHgKYmo9KZ7lFAWBcFSdeSGSkGgkX66HPudJ0_SorXKQ6cIdQUa41LetLYTufJK2DZETy4T6SM-wVJKvw)

# HowToUse

After all Prerequisites have been installed. you can follow the steps below
- Run minikube
```
    $ minikube start
```

- Change directory to this repository
```
    $ cd path/to/repository
```

- Create Secret for MySQL and Redis
```
    $ make create-secret
```
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEjJZ_zUrZhQlHeKe-t4R0l_t656StE1DzcJS4WzQqFVD9mgAe84L4jj4m2snf05Tx4qMuNMZiW3Jmeu0k0YnyY0jvt4tB47xxpiIkenk8WsS8XcgOqfOeefesYcIFk4BmbnmQGAHLwnpD0tcHuORxF2ixFo_xFy7VId2a2mu54nWP5Iuq6FPMy958yoGA)

- Create deployment
```
    $ make deploy
```
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEhDc3HX98cgXN9rNKZeTQ4dzXdZcwcpxZS31_NG2iuGkZPZ4vIz4utMVs-OvTCVP7onATRGy0HYriYSJ69IbXFlLdVO1y7AaJ3MJUojvuE1FgqfTpYYH1U__KKaI8XJ3euIED0RVH_yFemrHc0ln2Z3NhCwVFeaemHqaqRNRONhVwyQrqlfzZP3up11IA)

- Wait for all pods is running. you can check using this command and then copy pod with prefix `app-mysql-*`
```
    $ kubectl get pods
```
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEiZB7KxTfCuZKlI_0secwrxr3v8gqcm845dTYr-AVLKYKo91MiYOqGHve0w3EyMvn5CMaNlf3RbDnAIRZm8VwxLRU3mH7G__oDEiKShukZncMkOGKLp1IksMbUbUBt5Q5X3kdP8PluQEgu_7RTn1xiWseT9bS1C1rTkSlpb64e6JD3TkgPuDThL_xypFA)

- After all pods are running. enter the pod with prefix `app-mysql-*` using this command
```
    $ kubectl exec --stdin --tty pod-name-change-this -- /bin/sh 
```
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEjLVkANLIUDK2WZE2JVch8cAaG77nOyHJqkbSeIst7IJ4FXF0IpWcVqQUMy5CfC-SPM88DvJkTtZjd9wAZqTUDvcdVEEBI-sHoiePTR6QXG5feikWtL_iCzFmlTudCOSeiLdXrII82fJlsfBgo6NCGEfG3_8z0caemYeizApth7Xlcpup6q3ys1ZQZ3HA)

- Inside the pod running this command to populate table and data using this command and exit from pod
```
    $ go run migrations/migrate.go
    $ exit
```
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEg5QTvRzvaABsmIei81mGBEwQ03ZWXVrSHPjA4h6Z5p7_2aWUtsBUdcAUrPBxK5L0eJYcgJstjv1N9HhFgtk_jGZWcTs7xjT9bjdBB-AtLvNQ-YtzP5i4EPFvNWVV7K1LK5qoOPz8Th0x0QR-DCo5zKK_T-FzxHDS4zV6wwEEH3kn3uk93YaFDr0-PWDQ)

- Get base url for app using this command
```
    $ minikube service app-mysql --url
```
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEjmgcNnqLtdejNT6Xoy1kbZVNPESRAuPI-og7PBwLz_PxxkPx1PEG5NxT8JTwCL-ch2xQ6xMjHe63ka8ESOOdTXFi8n4yHPSWW7cRBbQQAowrE6S_UNM-Ksy9MlUGGMSNbLw1Wjh-iUQ6f71CPjflJxVnqeNIodGNnlwiW2y8f29fJiO2XaWhLLyXo4Hg)

- After base url is obtained. Open postman and import file `API User.postman_collection.json` and set global variable with name `localBaseURL`
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEgxjoUbtXPx8EM7uMv1EGBnCi03mUMBClkNFKhOYJqaujkm9UXFBxC_vq8iSmCCfYZ95Sc8dn623jpLvuzuFKyCRBF06zlCS3PpltS0nJkEdAN6JzzgjG6EFy34-YUuSTiW8FeQndui-dXScdzD2MtW9cVL8JMkHH9YsUyGU-4hesxckmNOEkeCjjJumQ)
![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEhZ3y7_PIAWnfrwe11MbZ2O1YW_YsqiiPhQoqQ_1fnIn-3OcfBwm3Cdagum64vnoOrPDw-H3Aqq3g3C2Xt-YyhutyldVH-MYJCR1X9qvaqEySaCi7o_69U3yNcTehOsVHdT4FjwFweQtYoH3gDScRm86tzyxt7-Kh-SD-NdUhaHVGIUhjpl4MWm0qekfg)

- Login to application using api login and save token then create global variable with name `authToken`. after that you can access another endpoint that need token.

![alt text](https://blogger.googleusercontent.com/img/a/AVvXsEjKggE2zDBH-H7RoLmDjRvgZeKCzEdD43CkHI4NahrEOirQiM4pUOgQ0jrpSKfT8WAQJTdz3Bf8VdYJiVyC51u3R313thTG6xYeKq9-VKdF2s1vAE3KgToL-kKRu_7XymhVMrT1KrFVqkSAw0wYOozhKe30tLjYTXvpz0ux_wdAvvm6HUBKS1ABJ9uJjw)
 
 # Credentials
 |     Email      |  Password |
|-------------|------:|
|admin@mail.com| adminadmin |
|user1@mail.com|useruser|
