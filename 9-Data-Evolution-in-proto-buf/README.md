# How compatible are the changes?

Protocol buffers are fully compatible, which means that it supports both backward and forward compatible.

# How are we achieving fully compatible, when there are changes which we want to do?

- Protocol buffer when add/remove fields will fall back to default value. (int to 0, string to '', etc)
- Never override a tag

# How not to override tag?

By using `reserved` keyword

```proto
message MyMessage {
  int32 id = 1;
  string first_name = 2;
}
```

Now from the above code we want to remove `first_name` and also to avoid tag overriding

```proto
message MyMessage {
  reserved 2;
  reserved "first_name";
  int32 id = 1;
}
```

This will inturn resolve conflicts in the code base

# How to use Reserve keyword

We can use range of reserve tag or field, but not mix

```proto
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

# Beware of default fields values

All though this default fields are useful for us to get code safety without any breaking changes in the existing code. There is no way to say if it was removed or default value in our code base.

### How can I be safe from default fields value

In business point of view never follow or depend on default value. That might lead to a problem. In code base use if condition to deal with it.

# Evolution of enum values

we can add, remove and reserve value as well. Here if anything is not there or not found. Default value (first value in enum) will be taken. Thus it is recommended to use default as `UNKNOWN = 0`
