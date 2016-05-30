# Travel-Sharing Application

## Overview
This travel sharing application prototype is Group 9's Final Database System for INFO340. See our application website [here](https://travel-sharing.herokuapp.com/).

Made by **John Diego**, **Mayowa Aina**, **Jake Therrien**.

## Setup
### Tables
```sql
CREATE TABLE member (
	id serial PRIMARY KEY,
	picture varchar(100) NOT NULL,
	firstname varchar(50) NOT NULL,
	lastname varchar(50) NOT NULL,
	username varchar(25) NOT NULL,
	password varchar(50) NOT NULL
);

CREATE TABLE trip (
	id serial PRIMARY KEY,
	memberID integer NOT NULL,
	name varchar(50) NOT NULL UNIQUE, 
	description text,
	FOREIGN KEY (memberID) REFERENCES member(id) ON DELETE CASCADE
);

CREATE TABLE food (
	id serial PRIMARY KEY,
	name varchar(50) NOT NULL,
	cost numeric(8,2) NOT NULL,
	rating integer,
	CHECK (cost >=0 AND rating BETWEEN 1 AND 5)
);

CREATE TABLE shelter (
	id serial PRIMARY KEY,
	type varchar(20) NOT NULL,
	name varchar(30) NOT NULL,
	cost numeric(6,2) NOT NULL,
	rating integer,
	CHECK (cost>0 AND rating BETWEEN 1 AND 5)
);

CREATE TABLE trippointlocalcontact (
	localcontactID integer REFERENCES localcontact(id) NOT NULL,
	trippointID integer REFERENCES trippoint(id) NOT NULL,
	PRIMARY KEY(localcontactID, trippointID)
);	

CREATE TABLE trippointfood (
	foodID integer REFERENCES food(id) NOT NULL,
	trippointID integer REFERENCES trippoint(id) NOT NULL,
	PRIMARY KEY(foodID, tripPointID)
);

CREATE TABLE location (
	id serial PRIMARY KEY,
	streetaddress1 varchar(50) NOT NULL,
	streetaddress2 varchar(50),
	city varchar(50) NOT NULL,
	state varchar(50),
	country varchar(50) NOT NULL,
	postalcode varchar(50),
	UNIQUE (streetaddress1, city, country)
);

CREATE TABLE transportation (
	id serial PRIMARY KEY,
	cost numeric(8,2) NOT NULL,
	type varchar(20) NOT NULL,
	Name varchar(30) NOT NULL,
	CHECK (cost >= 0)
);

CREATE TABLE trippoint (
	id serial PRIMARY KEY,
	tripID integer NOT NULL,
	shelterID integer,
	locationID integer NOT NULL,
	transportationID integer NOT NULL,
	date timestamp NOT NULL,
	description text,
	FOREIGN KEY (tripID) REFERENCES trip(id) ON DELETE CASCADE,
	FOREIGN KEY (shelterID) REFERENCES shelter(id) ON UPDATE CASCADE,
	FOREIGN KEY (locationID) REFERENCES location(id) ON UPDATE CASCADE,
	FOREIGN KEY (transportationID) REFERENCES transportation(id) ON UPDATE CASCADE
);

CREATE TABLE localcontact (
	id serial PRIMARY KEY,
	firstname varchar(50) NOT NULL,
	lastname varchar(50) NOT NULL,
	phoneNumber varchar(50),
	email varchar(50)
);
```
### Sample Data
```sql
INSERT INTO member (picture, firstname, lastname, username, password) VALUES ('https://robohash.org/errorfacilisautem.png?size=100x100&set=set1', 'George', 'Campbell', 'gcampbell0', 'P9ORqr');

INSERT INTO transportation (cost, type, name) VALUES (801.62, 'bullet train', 'Speedy');

INSERT INTO food (name, cost, rating) VALUES ('Bangers and mash', 11.49, '4');

INSERT INTO shelter (type, name, cost, rating) VALUES ('Hotel', 'Deluxe', 260.62, '5');

INSERT INTO location (streetaddress1, streetaddress2, city, state, country, postalcode) VALUES ('647 Northport Park', '7957 Melvin Road', 'Panjerrejo', null, 'Indonesia', null);

INSERT INTO localcontact (firstname, lastname, phonenumber, email) VALUES ('Betty', 'Sanders', '254-(758)994-8477', 'bsanders0@alibaba.com');

INSERT INTO trip (memberID, name, description) VALUES (1, 'Vacation to Indonesia', 'This was my first trip to Indonesia, I really enjoyed the culture and the people!');

INSERT INTO tripPoint (tripID, shelterID, locationID, transportationID, date, description) VALUES (1, 1, 1, 1, '2016-05-16 04:05:06', 'Lorem ipsum');

INSERT INTO trippointfood (foodID, trippointID) VALUES (1, 3);

INSERT INTO trippointlocalcontact (localcontactID, trippointID) VALUES (1, 1);
```
### Indices
```sql
CREATE INDEX date_index ON tripPoint(date DESC);

CREATE INDEX country_index ON location(country);
```
### Stored Procedures
```sql
CREATE OR REPLACE FUNCTION set_trip(_n varchar(50), _d text)
RETURNS integer AS $tripid$
declare
	tripid integer;
BEGIN
INSERT INTO trip(memberid, name, description) VALUES (1, _n, _d);
	SELECT id into tripid FROM trip WHERE memberid = 1 AND name = _n;
  	RETURN tripid;
END;
$tripid$ LANGUAGE plpgsql;

CREATE FUNCTION set_trippoint(_date timestamp, _trippointdescription text, _address1 varchar(50), _city varchar(50), _country varchar(50), _transportationtype varchar(20), _transportationcost numeric(8,2), _transportation varchar(30))
RETURNS void
AS
$$ 
BEGIN
	INSERT INTO transportation(cost, type, name) VALUES (_transportationcost, _transportationtype, _transportation);
	INSERT INTO location(streetaddress1, city, country) VALUES(_address1, _city, _country);
	INSERT INTO trippoint(tripid, locationid, transportationid, date, description) VALUES((SELECT max(id) FROM trip), (SELECT max(id) FROM location), (SELECT max(id) FROM transportation), _date, _trippointdescription);
END;
$$
 LANGUAGE plpgsql;
```
### Views
```sql
CREATE VIEW trip_essentials AS 
SELECT 
member.username AS member,
member.picture AS pic,
trip.name AS trip,
trip.description AS tripinfo,
trippoint.date AS trippointdate,
trippoint.description AS trippointinfo,
location.country AS country,
location.city AS city,
transportation.type AS transportation
FROM member
JOIN trip ON trip.memberID = member.id 
JOIN trippoint ON trippoint.tripID = trip.id
JOIN transportation ON transportation.id = trippoint.transportationID
JOIN location ON location.id = trippoint.locationID;
```