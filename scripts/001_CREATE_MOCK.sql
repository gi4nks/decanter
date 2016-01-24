-- create the vote table
DROP TABLE VOTE;
DROP TABLE TALK;

CREATE TABLE TALK (
  talkId   varchar(10) PRIMARY KEY,
  startTime timestamp NOT NULL,
  title varchar(256)
);

CREATE TABLE VOTE (
  id      SERIAL PRIMARY KEY,
  created timestamp DEFAULT current_timestamp,
  rating  integer NOT NULL CHECK (rating > 0 AND rating < 6),
  talkid  varchar(10) NOT NULL REFERENCES TALK (talkId),
  userid  integer NOT NULL CHECK (userId > 0)
);