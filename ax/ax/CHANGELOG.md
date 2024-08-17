# Changelog

## 0.6.1 (2024-08-17)


### ⚠ BREAKING CHANGES

* **xlsxsheet:** 对于使用旧字段查询方法的代码，需要更新为新方法并适应逻辑更改。Lua到JSON的解析支持现在在表格创建时作为一个参数传递。

### Features

* **ax/cmd/table:** 添加支持复杂数据类型和配置生成添加了新的数据类型解析器和代码生成器，以支持复杂的数据类型（如结构体、数组等）的配置生成。现在可以通过CLI生成带有复杂嵌套结构的配置文件。 ([fd95a9f](https://github.com/kercylan98/minotaur/commit/fd95a9f76205e39348a007b8619038be05de97c6))
* **ax/cmd:** 实现 xlsx 表到 Go 配置的转换 ([6b82d33](https://github.com/kercylan98/minotaur/commit/6b82d330700ea5d84e9c3844137a0d8c9dda3381))
* **ax:** 添加xlsx转换支持和配置生成 ([818b80c](https://github.com/kercylan98/minotaur/commit/818b80c90883f85b35ebec5515dc5cbfac72fa85))


### Code Refactoring

* **xlsxsheet:** 重构字段查询和表格创建逻辑 ([3320a8c](https://github.com/kercylan98/minotaur/commit/3320a8c8b1b9b8fec412ef86973931bd5d7e2672))
