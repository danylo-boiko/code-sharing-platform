declare @default_role varchar(50) = 'default'
declare @admin_role varchar(50) = 'admin'

declare @owned_claim varchar(30) = 'owned'
declare @foreign_claim varchar(30) = 'foreign'

declare @read_action varchar(20) = 'read'
declare @create_action varchar(20) = 'create'
declare @update_action varchar(20) = 'update'
declare @delete_action varchar(20) = 'delete'

if (select count(id) from roles) < 2
begin
    insert into roles(title, description) values (@default_role, @default_role + ' access mode'),
                                                 (@admin_role, @admin_role + ' access mode')
end

declare @default_role_id int = (select id from roles where title = @default_role)

if (select count(id) from role_claims where role_id = @default_role_id) < 4
begin
    insert into role_claims(role_id, claim_type, claim_value) values (@default_role_id, @owned_claim, @read_action),
                                                                     (@default_role_id, @owned_claim, @create_action),
                                                                     (@default_role_id, @owned_claim, @update_action),
                                                                     (@default_role_id, @owned_claim, @delete_action)
end

declare @admin_role_id int = (select id from roles where title = @admin_role)

if (select count(id) from role_claims where role_id = @admin_role_id) < 4
begin
    insert into role_claims(role_id, claim_type, claim_value) values (@admin_role_id, @foreign_claim, @read_action),
                                                                     (@admin_role_id, @foreign_claim, @create_action),
                                                                     (@admin_role_id, @foreign_claim, @update_action),
                                                                     (@admin_role_id, @foreign_claim, @delete_action)
end
