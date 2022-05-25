create table tb_users(
	idx serial,
	user_id 	  				varchar(64) 	  not null unique primary key,
	user_email 					varchar(64) 	  not null unique,
	user_name 					varchar(32) 	  not null,
	user_password				varchar(128) 	  not null,
	user_status					varchar(8) 	  	not null,
	create_date					varchar(64) 	  not null,
	user_phone 					varchar(32) 	  null unique,
	user_image 					varchar(128) 	  null,
	last_update					varchar(64) 	  null,
	user_login_start		varchar(64)	  	null,
	user_session				varchar(128)	  null,
	remember_me					varchar(8)			null
)