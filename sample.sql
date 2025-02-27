create schema if not exists test;

create sequence if not exists test.seq_users;
create sequence if not exists test.seq_comments;

create table if not exists test.users
(
    id int not null default nextval('test.seq_users'::regclass),
    name varchar not null,
    email varchar not null,
    constraint "PK_users" primary key (id),
    constraint "UQ_users_email" unique (email),
    constraint "CHK_users_email" check (email like '%@%')
    );

create table if not exists test.comments
(
    id int not null default nextval('test.seq_comments'::regclass),
    id_user int not null,
    txt varchar not null,
    constraint "PK_comments" primary key (id)
    );

-- Получение юзера
create or replace function test.user_get(_id integer)
  returns json as
$BODY$
declare
_ret json;
begin
  if _id = 0 then
select array_to_json(array(
    select row_to_json(r)
      from (
        select u.id, u.name, u.email
        from test.users u
      ) r
    )) into _ret;
else
select row_to_json(r) into _ret
from (
         select u.id, u.name, u.email
         from test.users u
         where id = _id
     ) r;
end if;

return _ret;

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Добавление юзера
create or replace function test.user_ins(_params json)
  returns json as
$BODY$
declare
_newid integer;
begin
  _newid = 0;

insert into test.users (name, email)
select name, email
from json_populate_record(null::test.users, _params)
         returning id into _newid;

return json_build_object('id', _newid);

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Обновление юзера
create or replace function test.user_upd(_id integer, _params json)
  returns json as
$BODY$
begin
update test.users set
                      name = _params->>'name',
    email = _params->>'email'
where id = _id;

return json_build_object('id', _id);

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Удаление юзера
create or replace function test.user_del(_id integer)
  returns json as
$BODY$
begin
delete from test.users where id = _id;

return json_build_object('id', _id);

exception when others then

  raise notice 'Illegal operation: %', SQLERRM;

return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Просмотр комментария по id или всех, если id = 0
create or replace function test.comment_get(_id integer)
  returns json as
$BODY$
declare
_ret json;
begin
  if _id = 0 then
select array_to_json(array(
    select row_to_json(r)
      from (
        select c.id, c.id_user, c.txt
        from test.comments c
      ) r
    )) into _ret;
else
select row_to_json(r) into _ret
from (
         select c.id, c.id_user, c.txt
         from test.comments c
         where id = _id
     ) r;
end if;

return _ret;

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Редактирование текста комментария с указанным id. Авторство комментария менять нельзя.
create or replace function test.comment_upd(_id integer, _params json)
  returns json as
$BODY$
begin
update test.comments set
    txt = _params->>'txt'
where id = _id;

return json_build_object('id', _id);

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Удаление комментария с указанным id
create or replace function test.comment_del(_id integer)
  returns json as
$BODY$
begin
delete from test.comments where id = _id;

return json_build_object('id', _id);

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Добавление комментария за авторством пользователя _id_user
create or replace function test.user_comment_ins(_id_user integer, _params json)
  returns json as
$BODY$
declare
_newid integer;
begin
  _newid = 0;

insert into test.comments (id_user, txt)
select _id_user, _params->>'txt'
        returning id into _newid;

return json_build_object('id', _newid);

exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;

-- Реализовать самостоятельно:

-- Просмотр комментария по id или всех, если id = 0
create or replace function test.comment_get(_id integer)
    returns json as
$BODY$
declare
    _ret json;
begin
    if _id = 0 then
        select array_to_json(array(
                select row_to_json(r)
                from (
                         select c.id, c.id_user, c.txt
                         from test.comments c
                     ) r
                             )) into _ret;
    else
        select row_to_json(r) into _ret
        from (
                 select c.id, c.id_user, c.txt
                 from test.comments c
                 where id = _id
             ) r;
    end if;

    return _ret;

exception when others then

    return json_build_object('error', SQLERRM);
end
$BODY$
    language plpgsql volatile cost 100;

-- Редактирование текста комментария с указанным id. Авторство комментария менять нельзя.
create or replace function test.comment_upd(_id integer, _params json)
    returns json as
$BODY$
begin
    update test.comments set
        txt = _params->>'txt'
    where id = _id;

    return json_build_object('id', _id);

exception when others then

    return json_build_object('error', SQLERRM);
end
$BODY$
    language plpgsql volatile cost 100;

-- Удаление комментария с указанным id
create or replace function test.comment_del(_id integer)
    returns json as
$BODY$
begin
    delete from test.comments where id = _id;

    return json_build_object('id', _id);

exception when others then

    return json_build_object('error', SQLERRM);
end
$BODY$
    language plpgsql volatile cost 100;

-- Добавление комментария за авторством пользователя _id_user
create or replace function test.user_comment_ins(_id_user integer, _params json) returns json as
$BODY$
declare
    _newid integer;
begin
    _newid = 0;

    insert into test.comments (id_user, txt)
    select _id_user, _params->>'txt'
    returning id into _newid;

    return json_build_object('id', _newid);

exception when others then

    return json_build_object('error', SQLERRM);
end
$BODY$
    language plpgsql volatile cost 100;