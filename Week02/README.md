我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

- sql.ErrNoRows 应该属于正常的业务查询逻辑，没必要关心底层错误，可以重新声明一个错误类型返回
- 对于其余类型异常，可以做wrap 向上返回