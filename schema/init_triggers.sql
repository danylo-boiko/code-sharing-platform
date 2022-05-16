create or alter trigger session_expiry_date_update
    on sessions
    after update
as
if update(expiry_date)
begin
    declare @old_expiry_date datetime2 = (select expiry_date from deleted)
    declare @new_expiry_date datetime2 = (select expiry_date from inserted)

    if(@old_expiry_date > @new_expiry_date)
    begin
        raiserror('You cannot shorten the lifetime of a session token', 12, 1)
        rollback tran
    end
end
