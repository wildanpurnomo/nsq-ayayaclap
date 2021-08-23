probably won't be finished pepelaugh.\n\n

in case you're wondering, nsq is a messaging queue platform. it is similar to RabbitMQ or Apache Kefka. it can be used for communication between multiple services. for example, if main-service receives user registration, it publishes an event that can be consumed with smtp-service to send confirmation email. another example, various transaction events from main-service can be listened by log-service to be written into nosql db for business metric analysis.\n\n

needless to say, this ain't finished yet.