# Find Gaming Friends

Kişilerin kayıt olarak kendi oyun zevklerine göre oyun arkadaşları bulmasını kolaylaştıran bir uygulamanın api projesidir.  

## Açıklama

--

## Kullanım Bilgileri

Proje ana dizinindeyken aşağıdaki komut ile gerekli servisler ayağa kaldırılır. (App, Postgres, pgAdmin) 
```
docker compose up
``` 

Serisler ayağa kalktıktan sonra database tablolarının oluşması için aşağıdaki api metodu ile istek atılmalıdır. 

``` 
POST  -  api/v1/database
``` 

## Erişim Bilgileri

- pgAdmin:
```
localhost:5050/browser/
```

## Servis Bilgileri

- pgAdmin:
```
name : pgAdmin
password : 1234
```

- postgres:
```
host name : postgres-database
port : 5432
password : 1234
```

## Örnek Hata Response Modeli

```
{
    "code": 500,
    "message": "Internal Server Error",
    "success": false
}
```

## Neler kullandim ? 

---
