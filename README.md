# Go Asset Management & Dynamic Image Resizing
Asset Management & Dynamic Image Resizing using Go Programming language. The resized image would be stored/cached so the Go server would act like a static file router after the first resize with the same parameter.

- Tested on Go 1.24 on Windows Server
- Tested on Go 1.24 on Linux Server

## Upload Endpoint
You can modify the endpoint inside `./app/router/router.go` anytime you want. But as of now, the default is in: `http://yoursite.com/assets/upload`. With parameter:

- `path` **(string | required)** 

Your folder path in upload folder. For example, if the upload folder is in `./app/upload` then the uploaded directory path would be `./app/upload/supplied_path` or `./public/dist/supplied_path` if not an image.

- `name` **(string | optional)**

The name of the uploaded file. If not supplied, then the app will create a new random name.

- `file` **(file | required)** 

The file parameter. The default max size is 32 MB, but you can change it anytime you want in `./app/router/router.go`


## Basic Example on How To Get
For example, you wanted to dynamic resize image named **logo.jpg** with specific size such as: **width = 100px** and **height = 72px** with **non-constraint scaling** or forced resizing priority so the image is exactly on that specific size.

1. Upload your file on **Upload Endpoint**
2. The upload request would sent back the image URI if succesfully uploaded.
2. Access your dynamic images using the URI with query parameter for example: http://yoursite.com/assets/image/path/to/logo.jpg?width=200&height=72&priority=forced

As you can see, we use 100 after the letter 'width' which stand for image width and 72 after the letter 'height' which stand for image height. We also use forced after the letter 'priority' which stand for resizing priority.

## 'priority' for Priority
There are four options for this setting. The default option `width`.
- forced
The image will be forced to resize based on your width and height settings.
- width
The image will be scaled based on your width settings. For example if you wanted to resize a 1600x1200 image to 800x800 image, you will get 800x600 sized image as the image is scaling based on width.
- height
The image will be scaled based on your height settings. For example if you wanted to resize a 1600x1200 image to 800x800 image, you will get 1200x800 sized image as the image is scaling based on height.
- real
The image won't be modified and sent as it is.

## 'width' for Width
You can set your preferred width for the image on this option.

## 'height' for Height
You can set your preferred height for the image on this option.

## Configuration
.Env file with the example of env.example. To use it copy the `env.example` as `.env`

### ENV
Your default environment. It should be between `development` or `production`. Don't forget to change it to production if you wanted to build/compile the app.

### SERVER_HOST
Your server host configuration. In which host your app going to listen to the request. Typically this would be `localhost` or `127.0.0.1`

### SERVER_PORT
Your server port configuration. In which port your app going to listen to the request. Typically this would be `8080` or `Any Other Possible Port Number`

### TEMPORARY_FOLDER
Your temporary folder path from ROOTPATH. The ROOTPATH would be your root folder `./app`.
The default is `upload/temporary` which translates into `./app/upload/temporary`

### UPLOAD_FOLDER
Your upload folder path from ROOTPATH. The ROOTPATH would be your root folder `./app`.
The default is `upload` which translates into `./app/upload`

### ALLOWED_HOST
Because this app is based on a REST API, so it needed list of allowed hosts or origin to access the app. In here, you can see the multiple origins and hosts which divided by `|`.
The default is `http://localhost|http://127.0.0.1` which translates the allowed host is `http://localhost` and `http://127.0.0.1`

## ALLOWED_FILE_MIME
Golang native support on image is actually not that great. So we have to only list the available image format that can be resized. As you can see, there is also text/plain. This app can also serve non-image file and store it into `public/dist`.
The default is `image/jpeg|image/png|image/gif|text/plain` which translates the allowed file mime is `image/jpeg`, `image/png`, `image/gif`, and `text/plain`.

## ALLOWED_IMAGE_WIDTH
The list of allowed resize width which divided by `|`.
The default is `200|360|720`

## ALLOWED_IMAGE_HEIGHT
The list of allowed resize height which divided by `|`.
The default is `200|360|720`

# Dependency
- Go 1.23 or higher
- Gin Framework
- GoDotEnv
- Imaging

# Supported Image Format
JPEG, PNG, and GIF
__a bit sad, I know, but expect to support more format later!__