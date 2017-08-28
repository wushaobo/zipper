#!/usr/bin/env bash

set -e

cd "$(dirname $0)"

registry=wushaobo/zipper
tag_latest=${registry}:latest

function push_tag_version () {
    git_tag=$1
    tag=${registry}:${git_tag}
    docker tag ${tag_latest} ${tag}
    docker push ${tag}
    docker rmi ${tag} ${tag_latest}
}

function push_latest_version () {
    docker push ${tag_latest}
    docker rmi ${tag_latest}
}

function push_to_registry () {
    if [ -z "${TRAVIS_TAG}" ]; then
        push_latest_version
    else
        push_tag_version ${TRAVIS_TAG}
    fi
}

function build_bin () {
    image=golang:1.8
    docker run --rm -v ${PWD}:/tmp/zipper -w /tmp/zipper/ ${image} docker/dev/build_bin.sh
}

function build_image () {
    build_bin

    docker_file=docker/prod/Dockerfile
    context_path=.

    docker build --force-rm \
        -f ${docker_file} \
        -t ${tag_latest} \
        ${context_path}
}

# add new cmd entry here
cmds=( \
build-image \
push-to-registry \
)

function do_command () {
    case $1 in
        build-image)
            build_image
            ;;
        push-to-registry)
            push_to_registry
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
