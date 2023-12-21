ulimit -c unlimited
echo core > /proc/sys/kernel/core_pattern
export LD_LIBRARY_PATH=./swsds/
./hsmGoDemo