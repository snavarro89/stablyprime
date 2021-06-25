# Stably Prime Numbers

Stably Prime Numbers is a cloud enabled app to answer which is the highest prime number lower than N.

[Demo App](https://d3n9yivspffmez.cloudfront.net/home)

API Url
https://keegq1szu9.execute-api.us-east-1.amazonaws.com/production/prime/{number}

## Tech Stack

Frontend: Angular 12 + Tailwind CSS
Backend: Golang + AWS Lambdas/Mux Framework

## Tech Setup
Stably Prime uses a number of open source projects to work 
- Golang +1.14 [Install Go](https://golang.org/doc/install)
- Angular +12 [Install Angular](https://angular.io/guide/setup-local)
- Docker +20.10 (optional for deploy and local test) [Install Docker](https://docs.docker.com/get-docker/)
- Serverless 2.2 (optional for deploy) [Install Serverless](https://www.serverless.com/framework/docs/getting-started/)

## Running the application

Download the project and install required dependencies
```sh
cd [GOSRCPATH]
mkdir github.com/snavarro89
cd github.com/snavarro89
git clone https://github.com/snavarro89/stablyprime
cd stablyprime/client
npm install
cd ../server
go build -o bin/mux
bin/mux
```
Browse for http://localhost:4200 to view the application.
The API is running on http://localhost:3001

## Testing
**Client**
Angular tests have not been created yet

**Server**
```sh
cd stablyprime/server
go test ./...
```

## Server

The server can be run on a virtual server (on prem, localhost, ec2, etc...) or as serverless functions

**Virtual Server**
The application uses [Gorilla Mux](https://github.com/gorilla/mux) package to handle http requests
To build the application run
```sh
cd stablyprime/server
go build -o bin/mux
```
To run the application
```sh
bin/mux
```
To build for production, you need to specify the target
```sh
GOOS=linux GOARCH=amd64 go build -o bin/mux
```

**Serverless**
The application also works for serverless architecture, you can build the whole app or you can build and deploy
specific functions.

Whole Application
```sh
cd stablyprime/server
scripts/build.sh
```
This script gets all the functions (individual .go files) from the aws folder and compiles them and zips them for deployment.

Single Function
```sh
cd stablyprime/server
scripts/build_one.sh [SUBFOLDER] [FUNCTION_NAME]
```
This script gets the subfolder and function name and compiles only one fuction.

AWS Server Folder structure
All lambdas are declared inside the "aws" folder, a subfolder groups functions that are relevant between them, 
for example: 
&nbsp;aws

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ user

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ get

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ create

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ update


To run your lambda locally, you can use [aws sam](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-local-start-api.html)
```sh
cd stablyprime/server
sam local start-api
```
**Deploying**
To deploy your lambdas to AWS you need to have IAM user with valid credentials, more info on [Serverless](https://www.serverless.com/framework/docs/providers/aws/guide/credentials/)
Update your ` ~/.aws/credentials` to add a profile, by default the app needs a profile called `aws-stably-test`, 
this could be change on your `serverless.yaml` file
If you have defined more than one environment file, you coud run
```sh
sls deploy --stage production
sls deploy --stage stage
sls deploy --stage dev
```

To deploy a single function
```sh
sls deploy --stage production --function prime_get
```
Where `prime_get` is the name of the function defined in the `serverless.yaml` file

## Client
The client has the following structure
&nbsp;src
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ app
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ layout
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ models
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ modules
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ services
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ shared

**Layout**
The layout module is responsible for determining the general look and feel of the application, any header
footer, navigation should be included in his module, the application has the ability to include multiple layouts. 
TODO: Add a service to automatically set layout for all the application.

**Models**
All object models, http requests and http response interfaces should be included in this folder, having a model for
everything prevents the developer to type everything with `any`

**Modules**
Inside the modules folder we should include all the different pages (components) that will be included in our application
The project includes Angular Routing so that the Layout can be lodaded first and then LazyLoad whatever module we need
based on the route. 

**Services**
The application uses rxjs to preserve state, all services for HTTP Requests and state objects should be added here. 
TODO: Implement ngrx store.

**Deploy**
To build for production and deploy, a valid [S3 static site](https://docs.aws.amazon.com/AmazonS3/latest/userguide/HostingWebsiteOnS3Setup.html) bucket should exist on your AWS Account, IAM credentials with S3 permissions should be added to ` ~/.aws/credentials`, and a valid Cloudfront Distribution should be created referencing the S3 Bucket
TODO: Create a cloudformation script to create permissions, bucket and cloudfront distribution
Once you have the required information, update the `stablyprime/client/scripts/deploy.sh` script
by default you should have an aws profile called `aws-stably-s3` and `aws-stably-cloudfront` 
Modify the script to reference your cloudfront distribution
```sh
cloudfront create-invalidation --distribution-id [YOUR CLOUDFRONT DISTRIBUTION ID] --paths
```
```sh
cd stablyprime/client
npm run build
scripts/deploy.sh
```
**TODO**
 - Add tests to Angular
 - Add Security validation on server side (Allow server to process JWT Tokens or oAuth with provider)
 - Implement ngrx on frontend
 - Decrease the build size in frontend, styles folder is too large.
 - Containirize the application
 - Create infrastructure scripts for automatic first-time deploying
 - Create CI/CD scripts for automatic test and deploy
