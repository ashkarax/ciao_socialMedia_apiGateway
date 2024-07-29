o Ciao-SocialMedia:MicroService:

- Developed a social media application similar to Instagram using Golang and Fiber.

- Implemented API key, access token, and refresh token for authentication
- Secured communication with HTTPS for external interactions
- Features include chat via WebSocket, posts, followers, following, account creation, video/audio calls,
  likes, comments, notifications, and media sharing (photos and videos)
- Architected with microservices(5-service):
  ● API Gateway
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedia_apiGateway)
  ● Auth Service
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedia_authService)
  ● Post & Relations Service
    [GitHub Link](https://github.com/ashkarax/ciao_socilaMedia_postNrelationService)
  ● Chat & Call Service
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedia_chatNcallService)
  ● Notification Service
    [GitHub Link](https://github.com/ashkarax/ciao_socialMedai_notificationSvc)
  
- Employed PostgreSQL for core services and MongoDB for Chat & Call.
- Used gRPC and protobuf for synchronous communication between services.
- Integrated Apache Kafka for asynchronous messaging.
- Implemented Redis for caching.
- Achieved efficient concurrency for data aggregation.
- Applied rate limiting and throttling to manage API usage effectively.
- Utilized comprehensive logging with Logrus.
  
- Deployed Dockerized services in a Kubernetes cluster.
  [Kubernetes manifests](https://github.com/ashkarax/ciao_socialMedai_notificationSvc)



