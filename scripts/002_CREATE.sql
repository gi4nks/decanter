drop table vote_sent

create table vote_sent (
  id      serial primary key,
  created timestamp without time zone,
  booked    timestamp without time zone,
  sent    timestamp without time zone,
  rating  integer not null check (rating > 0 and rating < 4)
);

-- create procedure to insert
create or replace function insertonsent() returns trigger as $_$
    begin
	insert into vote_sent(id, created, rating) values (new.id, new.created, round((3 * new.rating)/5));
        return new;
    end;
$_$ language plpgsql;

DROP FUNCTION getnextvote();

create or replace function getnextvote () 
 returns table (
 voteid int,
 talkid varchar,
 userid int,
 rating int
) 
as $$
declare
    idcandidate int;
begin
 select id into idcandidate from vote_sent
 where booked is null
 and sent is null
 order by created asc
 limit 1;

 update vote_sent
 set booked = current_timestamp
 where id = idcandidate;

 return query select
  v.id,
  y.talkid,
  y.userid,
  v.rating
 from
 vote_sent v
 inner join vote y 
 on v.id = y.id
 where
 v.id = idcandidate;
end; $$ language 'plpgsql';

create or replace function resetVotes () 
 returns void
as $$
begin
 update vote_sent
 set booked = null
 where booked > NOW() - INTERVAL '10 minutes' 
 and sent is null;
end; $$ language 'plpgsql';

create or replace function archiveVote (voteid integer) 
 returns void
as $$
begin
 update vote_sent
 set sent = current_timestamp
 where id = voteid;
end; $$ language 'plpgsql';


-- create trigger on row insertion
create trigger vote_trigger after insert 
on vote
for each row 
execute procedure insertonsent();