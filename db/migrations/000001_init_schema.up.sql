CREATE TABLE activities (
  activity_id integer AUTO_INCREMENT,
  title text,
  email text,
  created_at timestamp,
  updated_at timestamp,
   PRIMARY KEY (activity_id)
);

CREATE TABLE todos (
  todo_id integer AUTO_INCREMENT,
  activity_group_id integer,
  title text,
  priority text,
  is_active boolean,
  created_at text,
  updated_at text,
  PRIMARY KEY (todo_id)
);