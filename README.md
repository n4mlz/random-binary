# Random Binary File Generator

ランダムなバイナリファイルを生成するコマンドラインツール。

## 使い方

```bash
random-binary <サイズ> <出力ファイルパス>
```

### 引数

- `サイズ`: 生成するファイルのサイズ（単位: B, KB, MB, GB）
  - 例: `10KB`, `32MB`, `1GB`
- `出力ファイルパス`: 生成したバイナリファイルの保存先

### 使用例

```bash
# 10MBのランダムバイナリファイルを生成
random-binary 10MB output.bin

# 1GBのランダムバイナリファイルを生成
random-binary 1GB large_file.bin
```

### ビルド方法

```bash
go build
```

---

# Random Binary File Generator

A command-line tool to generate random binary files.

## Usage

```bash
random-binary <size> <output-file>
```

### Arguments

- `size`: Size of the file to generate (units: B, KB, MB, GB)
  - Example: `10KB`, `32MB`, `1GB`
- `output-file`: Path where the generated binary file will be saved

### Examples

```bash
# Generate a 10MB random binary file
random-binary 10MB output.bin

# Generate a 1GB random binary file
random-binary 1GB large_file.bin
```

### How to Build

```bash
go build
```
