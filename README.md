# Go-TOTP
RFC 6232 (TOTP)的Golang实现

# QuickStart
获取go-totp：
```go
go get w12lab.com/go-totp
```

初始化：
```go
// 适用于 Google Authenticator 的快速实现
client := totp.GetGoogle2FAAuth()

// 适用于 RFC 6232 的实现
// Default: T0:0  TI: 30 Digital: 6
client := totp.NewTotp(totp.DefaultT0,totp.DefaultTI,totp.DefaultDigital)
```

获得秘钥：
```go
key := client.GenerateSecret() // "TK7EF4QNPE3IX2PPGBO7UE3BCOHL4XLN"
```

获取验证码：
```go
code := client.GenerateCode(key) // "511055"
```

验证验证码：
```go
res := client.VerifyCode(secret, code) //true
```

# 咋用？
把秘钥导入到
> Google Authenticator

或

> FreeOtp

>Notice: Google Authenticator仅支持默认值的T0,TI,Digital！

# License：
本程序基于WPFPL开源协议开源

>             DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
>                      Version 2, December 2004
>  
>   Copyright (C) 2020 W12Lab <muyang.li@binqsoft.com>
>  
>   Everyone is permitted to copy and distribute verbatim or modified
>   copies of this license document, and changing it is allowed as long
>   as the name is changed.
>  
>              DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
>     TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
>  
>    0. You just DO WHAT THE FUCK YOU WANT TO.

***
Attention：

This program is free software. It comes without any warranty, to the extent permitted by applicable law. You can redistribute it
 and/or modify it under the terms of the Do What The Fuck You Want
 To Public License, Version 2, as published by Sam Hocevar. See
 http://www.wtfpl.net/ for more details.