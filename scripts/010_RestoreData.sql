truncate table vote;
truncate table vote_sent;
truncate table talk;

INSERT INTO talk (talkid, starttime, title) 
VALUES ('TALK_1', current_timestamp, 'TITLE TALK 1');

INSERT INTO vote(
            id, created, rating, talkid, userid)
    VALUES (1, current_timestamp, 3, 'TALK_1', 1);

INSERT INTO vote(
            id, created, rating, talkid, userid)
    VALUES (2, current_timestamp, 4, 'TALK_1', 2);


INSERT INTO vote(
            id, created, rating, talkid, userid)
    VALUES (3, current_timestamp, 5, 'TALK_1', 3);


update vote_sent
 set booked = null
 where booked > NOW() - INTERVAL '10 minutes' 
 and sent is null;



select * from vote_sent

select resetVotes () 

update vote_sent
set status = 'booked'