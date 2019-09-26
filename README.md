# Fluent Writer
Simple Go writer for Fluentd 

Sample usage for [logrus](https://github.com/sirupsen/logrus)

```go
log.SetFormatter(&log.JSONFormatter{})

fwriter = fluentwriter.NewFluentWriter(fluentHost, fluentPort, fluentTag, 4 * time.Second, 1024)
log.SetOutput(io.MultiWriter(os.Stdout, fwriter))

```