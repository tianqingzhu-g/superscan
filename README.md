# SuperScan

⚡️ **极速多语言供应链漏洞扫描器**  
一个命令，自动发现 Python、Node.js 依赖，查询 OSV 漏洞库，彩色输出 + JSON 报告，一键集成 CI。

## 快速开始

### 安装
暂时从源码构建（计划提供预编译二进制）：
\\\ash
git clone https://github.com/tiangingzhu-g/superscan.git
cd superscan
go build -o superscan ./cmd/superscan
\\\

### 使用
\\\ash
# 扫描当前目录
./superscan .

# 扫描指定路径
./superscan /path/to/project
\\\

## 支持的包管理器
| 语言 | 锁定文件 | 状态 |
|------|----------|------|
| Python | Pipfile.lock | ✅ |
| Node.js | package-lock.json | ✅ |
| 更多... | 贡献中 | 🚧 |

## 集成到 GitHub Actions
\\\yaml
- name: Run SuperScan
  uses: tiangingzhu-g/superscan/action@main
  with:
    path: .
\\\

## 贡献
欢迎新增语言解析器！请查看 [CONTRIBUTING.md](./CONTRIBUTING.md)（即将添加）。

## 许可
MIT License
