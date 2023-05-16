-- migrate:up
create table contract
(
    contract_id 	bigserial	not null
		constraint contract_pk 
			primary key,	
    employee_id 	bigint 		not null,
	foreign key (employee_id) references employee (employee_id),
    start_date 		timestamptz	not null,
    end_date 		timestamptz default null,
    basic_pay 		float 		not null,
    total_pto		float 		not null,
	
	created_at		timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at		timestamp with time zone default CURRENT_TIMESTAMP not null
);

-- migrate:down

