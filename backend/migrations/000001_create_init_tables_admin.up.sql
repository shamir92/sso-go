
CREATE TABLE region
(
  id INT NOT NULL,
  name VARCHAR(65535) NOT NULL,
  description VARCHAR(65535) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  uuid uuid NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (uuid)
);

CREATE TABLE languages
(
  id INT NOT NULL,
  name VARCHAR(65535) NOT NULL,
  display_name VARCHAR(65535) NOT NULL,
  uuid uuid NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (uuid)
);

CREATE TABLE role
(
  id INT NOT NULL,
  name VARCHAR(65535) NOT NULL,
  description TEXT NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE member
(
  id INT NOT NULL,
  name VARCHAR(65535) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  user_metada DATE NOT NULL,
  app_metadata DATE NOT NULL,
  created_at INT NOT NULL,
  updated_at INT NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE member_m2m_fa
(
  id INT NOT NULL,
  name VARCHAR(65535) NOT NULL,
  confirmed INT NOT NULL,
  number INT NOT NULL,
  key INT NOT NULL,
  method INT NOT NULL,
  member_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (member_id) REFERENCES member(id)
);

CREATE TABLE tenant
(
  id INT NOT NULL,
  uuid uuid NOT NULL,
  name INT NOT NULL,
  environment_stage INT NOT NULL,
  friendly_name VARCHAR(255) NOT NULL,
  logo_url VARCHAR(65535) NOT NULL,
  support_email VARCHAR(65535) NOT NULL,
  support_url VARCHAR(65535) NOT NULL,
  error_pages_url VARCHAR(65535) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  region_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (region_id) REFERENCES region(id),
  UNIQUE (uuid)
);

CREATE TABLE tenant_language
(
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  tenant_id INT NOT NULL,
  language_id INT NOT NULL,
  FOREIGN KEY (tenant_id) REFERENCES tenant(id),
  FOREIGN KEY (language_id) REFERENCES languages(id)
);

CREATE TABLE tenant_member
(
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  member_id INT NOT NULL,
  role_id INT NOT NULL,
  tenant_id INT NOT NULL,
  FOREIGN KEY (member_id) REFERENCES member(id),
  FOREIGN KEY (role_id) REFERENCES role(id),
  FOREIGN KEY (tenant_id) REFERENCES tenant(id)
);