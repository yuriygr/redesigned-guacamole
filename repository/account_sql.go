package repository

const selectAccount = `
	select a.*, s.* from account as a
	left join account_to_subsite as a2s
		inner join subsite as s on s.subsite_id = a2s.subsite_id
		on a.account_id = a2s.account_id
	where s.subsite_id is not null and a2s.primary = true
`
const inserAccount = `
	insert into account (email, password, created, status)
	VALUES (:a.email, :a.password, now(), :a.status)
`

const inserAccountSubsite = `
	insert into account_to_subsite (subsite_id, account_id, primary)
	VALUES (:a.account_id, :a.account_id, '1');

	insert into subsite (subsite_id, slug, label, is_user, created, status)
	VALUES (:a.account_id, :a.slug, :a.name, '1', :a.created, :a.status);
`

const inserAccountLogin = `
	insert into account_login (account_id, ip, useragent, date)
	values (:al.account_id, :al.ip, :al.useragent, :al.date)
`

const selectLogin = `
	select
		al.*
	from account_login as al
`
