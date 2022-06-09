# InnoDB和Myisam的区别

参考:

bilibili:

[【Java面试必问系列】20分钟理清MySQL锁（无敌）](https://www.bilibili.com/video/BV1AE411a7rv)

配套文档:

https://github.com/hongwen1993/all/blob/master/database/lock.md



About InnoDB And Myisam

[Mysql innodb myisam 引擎的区别](https://segmentfault.com/a/1190000021995700)

[『浅入浅出』MySQL 和 InnoDB](https://draveness.me/mysql-innodb/)

###### 如何选择存储引擎：

- Myisam：mysql 默认的存储引擎，如果主要业务以读操作和插入操作为主，且只有很少的更新和删除操作，并且对事务的完整性、并发性要求不是很高的情况下适合选择Myisam引擎。
- InnoDB：用于事务处理应用程序，支出外键，如果主要业务对事务的完整性有比较高的要求，在并发条件下要求数据的一致性，数据操作除了插入和查询以外，还包括很多的更新、删除操作，那么InnoDB存储引擎应该是比较合适的选择。



这里讲到了关于锁:

[【MySQL】MySQL 中的锁机制 ](https://www.cnblogs.com/jojop/p/13982679.html)

[『MySQL』深入理解事务的来龙去脉](https://juejin.cn/post/6844903827611582471)

[MySQL 事务](https://www.runoob.com/mysql/mysql-transaction.html)

mysql:

[数据库-MySQL 中 for update 的作用和用法](https://segmentfault.com/a/1190000023045909)

[MySQL 事务-ROLLBACK，COMMIT 用法详解](https://www.cnblogs.com/minigrasshopper/p/7803989.html)

关于粒度:

[粒度](https://baike.baidu.com/item/%E7%B2%92%E5%BA%A6/13014724)

