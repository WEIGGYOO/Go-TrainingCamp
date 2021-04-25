//作业： 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

//答：dao层封装对数据库的访问,不负责处理数据库查询结果。
//	因此应该wrap sql.ErrNoRows这个错误，往上抛，由service层来处理数据,判断查询结果是否正常
if errors.Is(err, sql.ErrNoRows) {
	return errors.Wrap(err, fmt.Sprintf("sql %s find null: %v",sql, err))
}
