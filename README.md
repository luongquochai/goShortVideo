# goShortVideo
## Description:

### Make a simple application perform short video like tiktok, reels.

## Backend

- [x] **Setup Golang project**: Initialize a new `Go` module.

- [x] **Setup web framework**: Choose and set up `Gin` or `Echo`.

- [x] **Setup database**: Configure `GORM` and connect to `PostgreSQL/MySQL`.

- [x] **Implement API endpoint**s: Define and implement the RESTful API endpoints.

- [x] **Video processing**: Integrate `FFmpeg` for video processing tasks.

- [x] **User authentication**: Implement `JWT-based authentication`.

### Project Structure:
```bash
├── Makefile
├── README.md
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   ├── userController.go
│   │   └── videoController.go
│   ├── database
│   ├── middlewares
│   │   └── authMiddleware.go
│   ├── models
│   │   ├── user.go
│   │   └── video.go
│   ├── routes
│   │   └── routes.go
│   ├── services
│   │   ├── authService.go
│   │   └── videoService.go
│   └── utils
│       └── ffmpeg.go
├── thumbnails
│   └── 
├── tmp
│   └── build-errors.log
└── videos
    ├
    └ 
```
# Step-by-step
1. Initialize a New Go Module:
```bash
mkdir goShortVideo
cd goShortVideo
go mod init goShortVideo
```

2. Install Necessary Dependencies:
```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/aws/aws-sdk-go
```

3. Create a .env file in the root directory:
```
DATABASE_URL=your_postgres_connection_string
JWT_SECRET=your_jwt_secret_key
AWS_ACCESS_KEY_ID=your_aws_access_key_id
AWS_SECRET_ACCESS_KEY=your_aws_secret_access_key
AWS_REGION=your_aws_region
S3_BUCKET_NAME=your_s3_bucket_name

```

# Testing the Application

You can now use Postman or any other API testing tool to test the endpoints:

1. Register a user: POST /api/register

```json
{
  "username": "lqhai",
  "password": "lqhaipassword",
  "email": "lqhai@example.com"
}
```
- Result:
```json
{
    "message": "User registered successfully"
}
```

2. Login: POST /api/login to get Token
```json
{
  "email": "lqhai@example.com",
  "password": "lqhaipassword"
}

```
- Result:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjExMTQ3OTcsInVzZXJJRCI6Mn0.UGISVctHE3uHta-GtUTdxbuh2Hi-dAolxlo81AN7Eo0"
}
```
3. Upload a video: POST /api/upload

- Add an Authorization header with the JWT token obtained from the login.
- Form data: video (file), description (text).

- Result:
```json
{
    "video": {
        "ID": 1,
        "CreatedAt": "2024-07-13T14:19:06.355882953+07:00",
        "UpdatedAt": "2024-07-13T14:19:06.355882953+07:00",
        "DeletedAt": null,
        "UserID": 1,
        "FilePath": "videos/aba4fead-7ac3-477f-8856-9c146aef7eb4-videotest.mp4",
        "ThumbnailPath": "thumbnails/311255e2-8358-4334-b084-19e370cd8a8b-thumbnail.png",
        "Description": "test-video"
    }
}
```

4. Get videos: GET /api/videos
```json
{
    "videos": [
        {
            "ID": 1,
            "CreatedAt": "2024-07-13T14:19:06.355882+07:00",
            "UpdatedAt": "2024-07-13T14:19:06.355882+07:00",
            "DeletedAt": null,
            "UserID": 1,
            "FilePath": "videos/aba4fead-7ac3-477f-8856-9c146aef7eb4-videotest.mp4",
            "ThumbnailPath": "thumbnails/311255e2-8358-4334-b084-19e370cd8a8b-thumbnail.png",
            "Description": "test-video"
        }
    ]
}
```
## Frontend

- [ ] Updating...

## Perform on AWS/GCP