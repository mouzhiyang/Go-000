1、我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

当系统为dao、service、api 层架构时，需要抛给上一层 service，在 service 层做下判断处理。一般来讲，在最上层是要把数据返回给 client，根据业务需要来决定是否可以返回这个空数据的错误信息，或者业务接口需要返回空数组，这时就可以在 service 将 sql.ErrNoRows 这个错误统一处理下，方便 api 调用，也方便 api 兼容不同的数据库。