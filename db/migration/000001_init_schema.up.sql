CREATE TABLE activities (
  activity_id  integer,
  title text,
  email text,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE todos (
  todo_id integer,
  activity_group_id integer,
  title text,
  priority text,
  is_active boolean,
  created_at text
);