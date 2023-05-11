-- migrate:up
ALTER TABLE employee ALTER COLUMN manager_id DROP NOT NULL;
ALTER TABLE employee
    ADD CONSTRAINT employee_manager_id_fk FOREIGN KEY (manager_id) REFERENCES employee (employee_id);

-- migrate:down

