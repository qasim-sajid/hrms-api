-- migrate:up
create table transactions
(
    transaction_id 		bigserial	not null
		constraint transactions_pk 
			primary key,	
    contract_id 		bigint 		not null,	
	foreign key (contract_id) references contract (contract_id),
    transaction_date	timestamptz	not null,
    paid_amount 		float 		not null,
    availed_pto			float 		not null,
	
	created_at			timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at			timestamp with time zone default CURRENT_TIMESTAMP not null
);

-- migrate:down

