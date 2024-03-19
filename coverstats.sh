#!/bin/bash

# should install cloc pytest before run this scripts!!!

# 参数检查
if [ $# -ne 1 ]; then
    echo "Usage: $0 <target_language>"
    exit 1
fi

# 接受命令行参数作为目标语言
target_language="$1"

# 定义目录
coverage_dir="coverage"
statement_coverage_file="$coverage_dir/statement_coverage.out"
statement_coverage_text_file="$coverage_dir/statement_coverage.txt"
comment_coverage_file="$coverage_dir/comment_coverage.txt"

# 创建 coverage 文件夹
mkdir -p "$coverage_dir"

# 计算代码注释率
echo -e "\033[0;32mStart calculating comment coverage...\033[0m"
cloc .  --out=$comment_coverage_file
# 模糊匹配目标语言的行
language_line=$(awk -v target="$target_language" 'index(tolower($0), tolower(target)) {print $0}' "$comment_coverage_file")
# 提取comment和code的值，使用第三个和第四个数字作为关键
comment=$(echo "$language_line" | awk '{for (i=1; i<=NF; i++) if ($i ~ /^[0-9]/) {count++; if (count == 3) {print $i; break}}}')
code=$(echo "$language_line" | awk '{for (i=1; i<=NF; i++) if ($i ~ /^[0-9]/) {count++; if (count == 4) {print $i; break}}}')
if [[ -n "$comment" && -n "$code" ]]; then
    # 计算comment/code的百分比
    percentage=$(awk "BEGIN { printf \"%.2f\", ($comment / $code) * 100 }")
    echo -e "\033[0;31mComment Coverage is $percentage%\033[0m"
else
    echo -e "\033[0;31m错误：无法提取 $target_language 语言的注释和代码值。\033[0m"
fi

# 如果目标语言是 Python:
if [ "$target_language" = "python" ]; then
    echo -e "\033[0;32mStart calculating test coverage for Python...\033[0m"
    echo -e "\033[0;33m请将coverstas放在和requirements.txt同级目录或根目录下\033[0m"

    # 获取当前路径
    current_path=$(pwd)

    # 检查当前路径是否包含"kms"
    if [[ "$current_path" == *"kms"* ]]; then
        if ! coverage run -m pytest src/; then
            echo -e "\033[0;31mError: Failed to calculate Python Test Coverage\033[0m"
            exit 1
        fi
    fi

    # 检查当前路径是否包含"engine"
    if [[ "$current_path" == *"engine"* ]]; then
        # 回退到上一级目录
        cd ..
        # 执行测试覆盖率计算
        if ! coverage run -m pytest mira-mpc-engine/; then
            echo -e "\033[0;31mError: Failed to calculate Python Test Coverage for mira-mpc-engine\033[0m"
            exit 1
        fi
    fi

    # 生成测试覆盖率报告
    coverage report -m
    echo -e "\033[0;31mPython Test Coverage Calculated\033[0m"
fi

# 如果目标语言是 Go:
if [ "$target_language" = "go" ]; then
    echo -e "\033[0;32mStart calculating test coverage for Go...\033[0m"
    echo -e "\033[0;33m请将coverstas放在和go.mod同级目录下\033[0m"

    go test -coverprofile="$statement_coverage_file"  ./...
    go tool cover -func="$statement_coverage_file" -o "$statement_coverage_text_file"
    # 提取total行并提取括号内的覆盖率
    coverage_percentage=$(awk '/total:/ {print $NF}' "$statement_coverage_text_file")
    # 输出结果
    echo -e "\033[0;31mUT Statements Coverage is: ${coverage_percentage}\033[0m"
fi

# 如果目标语言是 Java:
if [ "$target_language" = "java" ]; then
    echo -e "\033[0;32mStart calculating test coverage for Java...\033[0m"
    echo -e "\033[0;33m请将coverstas放在和pom.xml同级目录或根目录下\033[0m"
    mvn clean test jacoco:report
    # 获取测试覆盖率信息
    coverage_percentage=$(grep -o 'Total.*</td>' target/jacoco-reports/index.html | grep -o '[0-9]*%' | head -n 1)
    # 输出测试覆盖率信息
    echo -e "\033[0;31mUT Statements Coverage is: ${coverage_percentage}\033[0m"
fi

