# jwt 认证

## 如果需要结合 Refresh Token 使用时

一般我们请求接口都会直接用到 Access Token，这个 Access token 是访问授权资源接口的凭证，一般会有一个过期时间，过期时间到了就需要重新获取 Access Token。

这个时候，一般我们会有 2 种解决方案：

1. 重新登录获取新的 Access Token
2. 使用 Refresh Token 获取新的 Access Token

第一种方案很好理解，就是重新登录获取新的 Access Token，但是这样会导致用户体验不好，因为用户需要重新输入用户名和密码。  
第二种方案就是使用 Refresh Token 获取新的 Access Token，这样用户就不需要重新输入用户名和密码，只需要使用 Refresh Token 就可以获取新的 Access Token。

**通常情况下， Refresh Token 的有效期会比较长，而 Access Token 的有效期比较短，当 Access Token 由于过期而失效时，可以使用 Refresh Token 就可以获取到新的 Access Token，**
**但是如果 Refresh Token 也失效了，那么用户就只能重新登录了。**

如果引入 Refresh Token 的话，请求流程就是如下：

- 客户端使用用户名和密码进行登录认证
- 服务端验证用户名和密码，如果验证通过，生成 Access Token（有效期较短，比如 10 分钟） 和 Refresh Token（有效期较长，比如 7 天），**两个 token 同时返回给客户端**
- 客户端访问需要认证的接口时，携带 Access Token
- 服务端验证 Access Token 是否有效，如果有效则返回数据
- 如果无效，则后端返回鉴权失败，比如返回 401 错误，客户端收到 401 错误后，携带 Refresh Token 去请求**刷新 Token 的接口** 申请新的 Access Token
- 服务端验证 Refresh Token 是否有效，如果有效则返回新的 Access Token，客户端再使用新的 Access Token 访问需要认证的接口，如果无效则需要用户直接重新登录

## 关于做刷新 token 的接口的思路

做法也有很多种，这里我提供了两种思路：

第一种思路：参考 `jwt_access_refresh_token.go` 文件中的代码，整体代码思路和上面的流程一致，登录成功时，会获取 2 个 token，一个是 Access Token
一个是 Refresh Token。Access Token 中包含用户的信息，Refresh Token 中不会含有用户信息，只会包含刷新 token 的过期时间这一个有用因素。
这两个 token 都需要被客户端保存到本地，访问需要认证的接口时，**只需要**携带 Access Token，如果报错 token 失效过期，则需要**同时**携带
Access Token 和 Refresh Token 去请求刷新 token 的接口，如果此时 Refresh Token 没有过期，Access Token 是因为过期而失效时，则会再次返回新的
Access Token 和 Refresh Token，这样就可以一直保持登录状态，如果 Refresh Token 过期了，则需要重新登录。

第二种思路：参考 `jwt.go` 文件中的代码，这里只会颁发一个 token，那怎么做刷新 token 的逻辑呢？

还是差不多的思路，客户端登录成功时，服务端会返回一个 token 给到客户端，客户端将这个 token 保存到本地，然后访问需要认证的接口时，携带这个 token，
但是如果报错 token 失效过期，那么则需要携带这个过期的 token 去请求刷新 token 的接口，这个刷新的接口会判断是不是因为 token 过期而失效，如果是，
那么则会解析出这个 token 中的信息，然后判断这个 token 的**首次签名时间**和当前时间对比，是不是小于**刷新 token**的时间，如果是，那么就重新颁发
一个 token，但是需要注意的是，重新颁发的这个 token 中记录的首次签名时间还是之前失效的 token 的首次签名时间，也就是首次签名时间不做变化，只更改了这个 token 的有效期。
这样就达到了通过一个 token 也可以做刷新的效果。

其实这两种方案都运用了某一个时间过期的特性，只要某一个时间能够一直固定不变，那么才能判断这个 token 的时效性。第一种思路是运用了 Refresh Token 的过期时间，
第二种思路是运用了 token 的首次签名时间不变，然后通过计算得出是否需要生成新 token。

这两种方案都是：只要 access token 是因为过期而失效，并且当前时间并没有超过刷新时间，那么都允许通过过期的 access token 去获取一个新的 access token，
当超过了刷新时间，这两种方案都可以做到过期的 access token 无法获取新的 access token。