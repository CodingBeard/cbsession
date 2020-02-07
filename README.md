session
=======

[![Build Status](https://img.shields.io/shippable/5444c5ecb904a4b21567b0ff.svg)](https://travis-ci.org/codingbeard/session)
[![Go Report Card](https://goreportcard.com/badge/github.com/codingbeard/session)](https://goreportcard.com/report/github.com/codingbeard/session)
[![GoDoc](https://godoc.org/github.com/codingbeard/session?status.svg)](https://godoc.org/github.com/codingbeard/session)
[![GitHub release](https://img.shields.io/github/release/codingbeard/session.svg)](https://github.com/codingbeard/session/releases)


Provide session storage to [fasthttp](https://github.com/valyala/fasthttp).

This package follow the fasthttp philosophy, trying to avoid extra memory allocations in hot paths.

See [examples](https://github.com/codingbeard/session/tree/master/examples) to see how to use it.

## Providers

- memory
- memcache
- mysql
- postgres
- redis
- sqlite3


## Features

- Focus on the design of the code architecture and expansion.
- Provide full session storage.
- Convenient switching of session storage.
- Customizable data serialization.


## Bugs

***If you find a bug, please open new issue.***
