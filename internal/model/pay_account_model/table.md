#### kmzx_pay.pay_account 
支付账户

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | name | 名称 | varchar(64) |  | NO |  |  |
| 3 | app_key | 应用关键词 | varchar(32) | UNI | NO |  |  |
| 4 | app_secret | 应用秘钥 | varchar(32) |  | NO |  |  |
| 5 | status | 1 启用 -1 禁用 -2删除 | tinyint(1) |  | NO |  | 1 |
| 6 | create_time |  | datetime |  | YES |  |  |
| 7 | update_time |  | datetime |  | YES |  |  |
