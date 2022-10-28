#### kmzx_pay.pay_account_api 
账户拥有的api

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | pay_account_id |  | bigint(20) |  | YES |  |  |
| 3 | method | 请求方式 | varchar(32) |  | YES |  |  |
| 4 | api | 请求地址 | varchar(100) |  | YES |  |  |
| 5 | status | 状态 1:启用  -1:禁用 | tinyint(4) |  | YES |  | 1 |
| 6 | create_time |  | datetime |  | YES |  |  |
| 7 | update_time |  | datetime |  | YES |  |  |
