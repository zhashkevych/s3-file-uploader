# File Uploading API (to S3 competible storage)  

## <a href="https://medium.com/@zhashkevych/api-%D0%BD%D0%B0-go-gin-gonic-%D0%B4%D0%BB%D1%8F-%D0%B7%D0%B0%D0%B3%D1%80%D1%83%D0%B7%D0%BA%D0%B8-%D1%84%D0%B0%D0%B9%D0%BB%D0%BE%D0%B2-99e3622a3d6a">Blog Post</a>

## Requirements
- go 1.14
- docker & docker-compose

## Run Project


Due to localstack usage, ```awscli``` tool is needed for local development.

Install it on mac via ```brew install awscli```

Before running local project, run:
```aws configure``` and set up AWS keys. Don't forget to add them to `docker-compose.yml`
Then, set up bucket using folowing commands:

`aws --endpoint-url=http://localhost:4572 s3 mb s3://file-storage`

`aws --endpoint-url=http://localhost:4572 s3api put-bucket-acl --bucket file-storage --acl public-read` 

Use ```make run``` to build and run docker containers with application itself and mongodb instance

## API:

### POST /api/upload

Used to upload image for publication

##### Input should be of type "multipart/form-data" with "file" as key to image: 

##### Example Response (Status 200 OK): 
```
{
    "status": "ok",
    "url": "http://localhost:4572/file-storage/XVlBzgbaiCMRAjWw"
}
```

##### Example Response (Status 400 Bad Request): 
```
{
    "status": "error",
    "url": "failed to open image"
}
```
