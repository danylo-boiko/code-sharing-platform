if (select count(id)from supported_languages) < 5
begin
    insert into supported_languages(title, file_extension) values ('Python', 'py'),
                                                                  ('C#', 'cs'),
                                                                  ('Go', 'go'),
                                                                  ('TypeScript', 'ts'),
                                                                  ('PHP', 'php')
end