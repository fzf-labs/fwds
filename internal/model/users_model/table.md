#### fwds.users 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint unsigned | PRI | NO | auto_increment |  |
| 2 | username | 用户名称 | varchar(255) | UNI | YES |  |  |
| 3 | password | 密码 | varchar(255) |  | NO |  |  |
| 4 | nickname | 昵称 | varchar(255) |  | YES |  |  |
| 5 | email | 邮箱 | varchar(255) | UNI | YES |  |  |
| 6 | phone | 手机 | varchar(255) |  | YES |  |  |
| 7 | avatar | 头像 | varchar(255) |  | YES |  |  |
| 8 | bio |  | varchar(255) |  | YES |  |  |
| 9 | address | 地址 | varchar(255) |  | YES |  |  |
| 10 | job | 工作 | varchar(255) |  | YES |  |  |
| 11 | salt |  | varchar(32) |  | YES |  |  |
| 12 | uuid |  | varchar(255) |  | YES |  |  |
| 13 | last_ip | 最后一次登录IP | varchar(16) |  | YES |  |  |
| 14 | last_login_time | 最后一次登录时间 | timestamp |  | YES |  |  |
| 15 | type | 用户类型 | tinyint |  | YES |  | 1 |
| 16 | status |  | tinyint |  | NO |  | 1 |
| 17 | deleted_at |  | timestamp |  | YES |  |  |
| 18 | created_at |  | timestamp |  | YES |  |  |
| 19 | updated_at |  | timestamp |  | YES |  |  |
