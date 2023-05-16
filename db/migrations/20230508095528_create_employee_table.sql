-- migrate:up
create table employee
(
    employee_id 	bigserial	not null
		constraint employee_pk 
			primary key,	
    username 		text 		not null,
    email 			text 		not null,
    password 		text 		not null,
    name 			text 		not null,
    phone_number	text 		not null,
    address 		text 		not null,
    employee_type 	text 		not null,
    manager_id 		bigint 		not null,
	
	created_at		timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at		timestamp with time zone default CURRENT_TIMESTAMP not null,
	
	constraint employee_uc
		unique (username, email)
);

-- migrate:down

