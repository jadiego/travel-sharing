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
```
### Sample Data

### Stored Procedures
