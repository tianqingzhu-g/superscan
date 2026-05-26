# SuperScan

⚡️ **极速多语言供应链漏洞扫描器**  
一个命令，自动发现 Python 和 Node.js 依赖，查询 OSV 漏洞库，彩色输出 + JSON 报告，一键集成 CI。

## ✨ 特性

- 支持 Python (Pipfile.lock) 和 Node.js (package-lock.json)
- 10 并发查询 OSV 漏洞数据库，极速返回结果
- 彩色终端表格，一眼定位高危漏洞
- 生成 JSON 报告，方便 CI 解析
- 提供单二进制文件，无需安装额外运行时（计划中）
- GitHub Action 原生集成，PR 自动拦截高危漏洞

## 🚀 快速开始

### 下载预编译二进制（推荐）
- **Linux**:  
  ```bash
  curl -sSfL https://github.com/tiangingzhu-g/superscan/releases/latest/download/superscan-linux-amd64 -o superscan
  chmod +x superscan
  sudo mv superscan /usr/local/bin/
  ```
- **Windows**: 从 [Releases](https://github.com/tiangingzhu-g/superscan/releases) 下载 `superscan-windows-amd64.exe`，放到任意 PATH 目录。

### 从源码构建
```bash
git clone https://github.com/tiangingzhu-g/superscan.git
cd superscan
go build -o superscan ./cmd/superscan
```

### 使用
```bash
# 扫描当前目录
./superscan .

# 扫描指定路径
./superscan /path/to/your/project
```

终端会打印彩色表格，并生成 `superscan-report.json`。

## 📦 支持的包管理器

| 语言 | 锁定文件 | 状态 |
|------|----------|------|
| Python | Pipfile.lock | ✅ |
| Node.js | package-lock.json | ✅ |
| 更多... | 欢迎贡献 | 🚧 |

## 🔧 集成到 GitHub Actions

在 workflow 中添加：
```yaml
- name: Run SuperScan
  uses: tiangingzhu-g/superscan/action@main
  with:
    path: .
```
如果发现 CRITICAL 或 HIGH 漏洞，Action 会自动失败，阻止 PR 合并。

## 🤝 贡献

我们欢迎新语言解析器！如果你想添加对 `poetry.lock`, `yarn.lock`, `Gemfile.lock` 等的支持，请查看 [CONTRIBUTING.md](./CONTRIBUTING.md)（即将添加）或直接提 Issue。

## 📄 许可

MIT License
```
