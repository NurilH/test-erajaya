<div id="top"></div>

  <h1 align="center">Test Erajaya Backend Engineer</h1>

</div>

<!-- DOCUMENTATION PROJECT -->
## Documentasi Project

Berikut fitur dan endpoint API yang terdapat dalam project ini :
| Method | Feature | Endpoint | Authentication
|---|---|---|---|
| POST | Create User | http://18.139.2.170/signup | No |
| POST | Login User | http://18.139.2.170/login | No |
| POST | Create product | http://18.139.2.170/jwt/products/create | yes |
| GET | Mengurutkan berdasarkan product terbaru | http://18.139.2.170/products/sort/newest | No |
| GET | Mengurutkan berdasarkan nama product (A-Z) | http://18.139.2.170/products/sort/asc/name | No |
| GET | Mengurutkan berdasarkan nama product (Z-A) | http://18.139.2.170/products/sort/desc/name | No |
| GET | Mengurutkan berdasarkan product harga termurah | http://18.139.2.170/products/sort/asc/price | No |
| GET | Mengurutkan berdasarkan product harga termahal | http://18.139.2.170/products/sort/desc/price | No |
 
 <br/>
 
 ### * Body Request
- Set Header "Conten-Type","application/json" Set Body Request <b>Create and Login User</b>
```
{
    "email": "example@example.com",
    "password": "example123"   
}
```
- Set Header "Conten-Type","application/json" Set Body Request <b>Create Product</b>
```
{
  "name": "Buku",
  "description": "Ukuran A4",
  "price": 40000,
  "quantity": 12
}
```
- Untuk create product terdapat authentication jadi diharuskan login terlebih dahulu untuk mendapatkan token
<br/>
<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

![VS Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=FFFFFF)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=FFFFFF)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=FFFFFF)&nbsp;
![JWT](https://img.shields.io/badge/JWT-05122A?style=flat&logo=JSON%20web%20tokens&logoColor=FFFFFF)&nbsp;
![Docker](https://img.shields.io/badge/Docker-05122A?style=flat&logo=docker&logoColor=FFFFFF)&nbsp;

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- Architecture -->
## Architecture
Architectur yang saya gunakan dalam mengerjakan project ini yaitu MVC (Model View Controller),
karena dalam pembuatan setiap fungsinya terstruktur dengan jelas sehingga mempermudah dalam pengembangan serta ketika ada error atau bug lebih cepat dan mudah diatasi.

<p align="right">(<a href="#top">back to top</a>)</p>


Contributor :
<br>
<p align="left">
    <a href="https://www.linkedin.com/in/nuril-huda-87b279214/" target="blank"><img align="center"
            src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/linked-in-alt.svg"
            alt="nuril huda" height="30" width="40" /></a>
    <a href="https://www.hackerrank.com/nurilhuda7337" target="blank"><img align="center"
            src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/hackerrank.svg"
            alt="nuril7337" height="30" width="40" /></a>
</p>

<p align="right">(<a href="#top">back to top</a>)</p>
<h3>
<p align="center">:copyright: 2022 | Built with :heart: from us</p>
</h3>
<!-- end -->
