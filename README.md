# golang_rabbitmq
To run the project follow the below steps:
- You need to have docker desktop installed and run the command 'docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management' to have your local RabbitMQ   running.
- Open terminal with path upto publisher and run command 'go run main.go'
- Open another terminal with path upto consumer and use command 'go run main.go'

Required SQL Script
use assignment

CREATE TABLE hotel_tables(
`created_at` datetime default null,
`updated_at` datetime default null,
`deleted_at` datetime default null,
`country` varchar(255) default null,
`hotel_id` int,
`name` varchar(255) default null,
`address` varchar(255) default null,
`latitude` varchar(255) default null,
`longitude` varchar(255) default null,
`telephone` bigint default null,
`description` varchar(255) default null,
`room_count` int default null,
`currency` varchar(255) default null
)
create table amenities(
`hotel_id` varchar(255),
`amenity` varchar(255)
)
create table other_conditions(
`hotel_id` varchar(255),
`condition` varchar(255)
)
create table capacity_tables(
`hotel_id` varchar(255),
`max_adults` int,
`extra_children` int
)
create table cancellation_policy_tables(
`hotel_id` varchar(255),
`type` varchar(255),
`expires_days_before` int
)
CREATE TABLE rate_plan_tables(
`created_at` datetime default null,
`updated_at` datetime default null,
`deleted_at` datetime default null,
`hotel_id` VARCHAR(100),
`rate_plan_id` VARCHAR(100),
`name` VARCHAR(100) default null,
`meal_plan` VARCHAR(100) default null
)
CREATE table room_tables(
`created_at` datetime default null,
`updated_at` datetime default null,
`deleted_at` datetime default null,
`room_id` int,
`hotel_id` VARCHAR(100), 
`name` varchar(255) default null, 
`description` varchar(255) default null
)

Here, the project structure is as follows:
- golang_rabbitmq contains two folder namely, publisher and consumer
- publisher contains one folder named modules (which conatins package publish), json file whose data needs to be transferred and main.go
- consumer contains two folders named model (has all struct definitions realted files), modules (contains package consume) and main.go
      - package consume consists of consume.go and database.go
          - consume.go deals with getting data from RabbitMQ
          - database.go deals with connecting to SQL and inserting data into the tables
