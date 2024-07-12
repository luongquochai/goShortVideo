# goShortVideo
## Description:

Make a application perform short video like tiktok, reels.

## Backend

- [ ] **Setup Golang project**: Initialize a new `Go` module.

- [ ] **Setup web framework**: Choose and set up `Gin` or `Echo`.

- [ ] **Setup database**: Configure `GORM` and connect to `PostgreSQL/MySQL`.

- [ ] **Implement API endpoint**s: Define and implement the RESTful API endpoints.

- [ ] **Video processing**: Integrate `FFmpeg` for video processing tasks.

- [ ] **User authentication**: Implement `JWT-based authentication`.

### Project Structure:
- main.go
- /controllers
  - userController.go
  - videoController.go
- /models
  - user.go
  - video.go
- /services
  - videoService.go
  - authService.go
- /routes
  - routes.go
- /middlewares
  - authMiddleware.go
- /utils
  - ffmpeg.go
- /config
  - config.go

