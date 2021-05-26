# pagee

usage

```go
w := pagee.Walk{
    Uri:  "https://gocn.vip/topics/excellent",
    Next: ".pagination .next a",
    Item: ".topics .topic .title a",

    Cookies: []*http.Cookie{
        &http.Cookie{Name: "_abc", Value: "123"},
    },

    LimitItems: 61,
    LimitPages: 3,
}

w.Start(func(e *pagee.Element) {
    // ....
})
```