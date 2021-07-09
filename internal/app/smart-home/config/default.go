package config

// nolint:lll
const defaultConfig = `
hivemq:
  port: 1883
  address: localhost
  connection: tcp
  client: server
http-server:
  address: :65432
  read-timeout: 2m
  write-timeout: 2m
  graceful-timeout: 5s
`
