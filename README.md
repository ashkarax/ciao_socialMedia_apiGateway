<h2> Go+Fiber Ciao-SocialMedia:MicroService: </h2>
<h5>A social media application's backend ,similar to the functionalities of Instagram, built using Golang and Fiber.</h5> 
<br><br>
 <h4>- Architected with 5-services: </h4>
 <br>
  ● API Gateway 
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedia_apiGateway)  <br>
  ● Auth Service
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedia_authService) <br>
  ● Post & Relations Service
    [GitHub Link](https://github.com/ashkarax/ciao_socilaMedia_postNrelationService) <br>
  ● Chat & Call Service
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedia_chatNcallService) <br>
  ● Notification Service
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedai_notificationSvc) <br>
    <hr>
  ● Kubernetes Manifests:
    [Kubernetes manifests](https://github.com/ashkarax/ciao_socialMedai_notificationSvc) <br><br>


- Implemented API key, access token, and refresh token for authentication
- Secured communication with HTTPS for external interactions
- Features include chat via WebSocket, posts, followers, following, account creation, video/audio calls,
  likes, comments, notifications, and media sharing (photos and videos)
- Employed PostgreSQL for core services and MongoDB for Chat & Call.
- Used gRPC and protobuf for synchronous communication between services.
- Integrated Apache Kafka for asynchronous messaging.
- Implemented Redis for caching.
- Achieved efficient concurrency for data aggregation.
- Applied rate limiting and throttling to manage API usage effectively.
- Utilized comprehensive logging with Logrus.
- Deployed Dockerized services in a Kubernetes cluster.



