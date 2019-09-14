#!/bin/bash

apk update
apk add openssh -q
# mkdir ~/.ssh
# echo $SSH_KEY > ~/.ssh/id_rsa
# chmod -R 0700 ~/.ssh
# cat ~/.ssh/id_rsa
# ssh -i ~/.ssh/id_rsa -oStrictHostKeyChecking=no ci@staging.theninja.life
# exit