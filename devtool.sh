#!/usr/bin/env bash

set -e

cd "$(dirname $0)"

function build_and_run () {
    ./build.sh build-image
    docker-compose -f docker/dev/docker-compose.yml up --force-recreate zipper
}

function start_deps () {
    docker-compose -f docker/dev/docker-compose.yml up -d redis
}

# add new cmd entry here
cmds=( \
build-and-run \
start-deps \
)

function do_command () {
    case $1 in
        start-deps)
            start_deps
            ;;
        build-and-run)
            build_and_run
            ;;
        *)
            echo "Please select what you want to do:"
            ;;
    esac
}

# functional codes for this shell, you can ignore
function select_cmd () {
    echo "Please select what you want to do:"
    select CMD in ${cmds[*]}; do
        if [[  $(in_array $CMD ${cmds[*]}) = 0 ]]; then
            do_command $CMD
            break
        fi
    done
}

function select_arg () {
    cmd=$1
    shift 1
    options=("$@")

    echo "Please select which arg you want:"
    select ARG in ${options[*]}; do
        if [[  $(in_array ${ARG} ${options[*]}) = 0 ]]; then
            ${cmd} ${ARG}
            break
        fi
    done
}

function in_array () {
    TARGET=$1
    shift 1
    for ELEMENT in $*; do
        if [[ "$TARGET" == "$ELEMENT" ]]; then
            echo 0
            return 0
        fi
    done
    echo 1
    return 1
}

function main () {
    if [[ $1 != "" && $(in_array $1 ${cmds[*]}) = 0 ]]; then
        do_command $*
    else
        select_cmd
    fi
}

main $*
