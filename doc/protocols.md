# Gogoatlak protocols

author: Hoongeun Cho(Dave Cho)

date: Nov 2, 2018

## general
### userauth

```json
{
    command: "user-auth",
    data: {
        user: "username",
        password: "password"
    }
}
```

response (success)
```json
{
    status: "success",
    data: {
       token: "abcd3wet1as3df1s2f33j",
    }
}
```

response (fail) - Authentication failed
```json
{
    status: "fail",
    data: {
        fields: {
            username: {
                status: "succuess",
                message: "Unregisted user"
            },
            password: {
                status: "fail",
                message: "Invalid username/password"
            }
        },
    }
}
```
response (error) - Alreay logined user
```json
{
    status: "error",
    message: "already logged in"
}
```

## Room
### enter room

```json
{
    command: "enter-user",
    token: "abcd3wet1as3df1s2f33j"
}
```

response (success)
```json
{
    status: "success",
    data: {
        members: [
            {
                username: "username1"
                userid: "aefw5ggs1w9h0r2xf5g"
            }, {
                username: "username2"
                userid: "zffw51gs1w9hlr23d5g"
            }
        ]
    }
}
```
broadcast (success)
```json
{
    type: "broadcast",
    command: "enter-user"
    data: {
        username: "username1"
        userid: "aefw5ggs1w9h0r2xf5g",
    }
}
```

### leave user
```json
{
    command: "leave-user",
    token: "abcd3wet1as3df1s2f33j"
}
```

response (success)
```json
{
    status: "success",
    data: {
    }
}
```

broadcast (success)
```json
{
    type: "broadcast",
    command: "leave-user"
    data: {
        userid: "aefw5ggs1w9h0r2xf5g",
    }
}
```

### fetch history
```json
{
    command: "fetch-history",
    token: "abcd3wet1as3df1s2f33j"
}
```

response (success)
```json
{
    status: "success",
    data: {
        messages: [
            {
                id: "1351241241235a1d3f2z3a",
                text: "<red>hi</red>",
                username: "username1",
                userid: "aefw5ggs1w9h0r2xf5g",
                time: 1351241241235
            },
            {
                id: "135163124123ta1zd3fdz3a",
                text: "<blue>bye</blue>",
                username: "username2",
                userid: "br3hwg6s179h0c2xfk7",
                time: 135163124123
            }
        ]
    }
}
```

response (fail)
```json
{
    status: "fail",
    data: {
        fields: {
            username: {
                data: "username",
                message: "Unregisted user"
            },
            password: {
                data: "password",
                message: "Invalid username/password"
            }
        },
    }
}
```

## Message
```json
{
    type: "message-send",
    token: "abcd3wet1as3df1s2f33j",
    data: {
        text: "<red>hi</red>"
    }
}
```

response (success)
```json
{
    status: "success",
    data: {},
}
```