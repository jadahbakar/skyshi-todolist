CREATE TABLE activities (
  id integer AUTO_INCREMENT,
  title text,
  email text,
  created_at timestamp,
  updated_at timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE todos (
  id integer AUTO_INCREMENT,
  activity_group_id integer,
  title text,
  priority text,
  is_active boolean,
  status text,
  created_at text,
  updated_at text,
  PRIMARY KEY (id)
);