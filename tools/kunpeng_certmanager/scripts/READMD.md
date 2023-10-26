# 鲲鹏TEE certmanager

- [x] 1. 需先向华为申请certmanager TA开发者证书
- [x] 2. 脚本执行按照01-05顺序执行
- [x] 3. 05为生成TA开发者证书，不需要向华为申请

# Run Scripts
- 首先需要将root.key/config/manifest.txt配置到scripts/certmanager目录下（该目录如果没有需要手工创建）
- 运行脚本[01_download_itrustee_depends.sh](01_download_itrustee_depends.sh)下载itrustee依赖库
- 运行脚本[02_configure_root_crt.sh](02_configure_root_crt.sh) 生成root证书，并提取root公钥，配置到certmanager TA程序中（硬编码）
- 运行脚本[03_compile_certmanager.sh](03_compile_certmanager.sh)生成certmanager CA/TA程序
- 运行脚本[04_configure_second_crt.sh](04_configure_second_crt.sh)生成certmanager 二级证书并导入（需要uncomment脚本中的部分代码）
- 运行脚本[05_issue_ta_crt.sh](05_issue_ta_crt.sh)生成TA开发者证书