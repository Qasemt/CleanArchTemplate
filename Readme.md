## build Release

If you want to build your Go application for production with a specific value for the APP_ENV environment variable, you can use the -ldflags option when running the go build command.
Here's an example of how to set the APP_ENV environment variable to production when building your Go application:

```
go build -ldflags="-X 'main.AppEnv=production'" -o myapp
```

## test msg redis

```bash
$ redis-cli

 PUBLISH channel_pub_event '{"foo": "bar"}'
```
