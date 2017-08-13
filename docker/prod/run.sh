#!/usr/bin/env bash

set -e


cmd_name="zipper"


app_dir=/opt/${cmd_name}
log_dir=/var/log/${cmd_name}

echo "Start listening ..."
${app_dir}/${cmd_name}
#${app_dir}/${cmd_name} >> ${log_dir}/stdout.log 2>> ${log_dir}/stderr.log
