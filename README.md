# 简介

* 批量导出所有语雀文档

> 任何问题欢迎提Issue。 
> 
> 如果有帮助到你，可以点个Star。

# 作者

京城郭少

# 使用方式

原理：

1. 访问语雀，通过F12拿到cookie和x_csrf_token，保存到配置文件中。
2. 运行程序。

详细步骤：

1. 从这里下载适合你的程序：https://github.com/GuoFlight/yuque_export/releases
2. 从这里下载配置文件config.toml：https://github.com/GuoFlight/yuque_export
3. 访问语雀，通过F12拿到cookie和x_csrf_token，保存到配置文件中。
4. 运行程序（以arm芯片的mac为例）

```shell
vim config.toml
    cookie = "xxxxxxxx"
    x_csrf_token="xxxxxxx"
    file_dir="./output"
chmod +x yuque_export_mac_arm64
./yuque_export_mac_arm64 -c ./config.toml   # 程序会自动下载所有语雀文档
```

# 常见问题

文档默认导出到哪里？

* 文档默认会导出在output目录下，可以在配置文件中修改file_dir配置。

mac运行报错：无法打开“yuque_export_mac_arm64”，因为Apple无法检查其是否包含恶意软件。

* 原因：此程序是从互联网上下载的，若运行会被mac拦截。
* 解法：设置 -> 隐私与安全性 -> 选择"仍然允许"。



