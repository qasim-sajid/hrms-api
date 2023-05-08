-- migrate:up
create table requests
(
    request_id 	bigserial	not null
		constraint request_pk 
			primary key,	
    employee_id 	bigint 		not null,	
	foreign key (employee_id) references employee (employee_id),
    start_date 		timestamptz	not null,
    end_date 		timestamptz not null,
    action_at 		timestamptz not null,
    is_approved 	boolean 	not null,
	
	created_at		timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at		timestamp with time zone default CURRENT_TIMESTAMP not null
);

-- migrate:down

