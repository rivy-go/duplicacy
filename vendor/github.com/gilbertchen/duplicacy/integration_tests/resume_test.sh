#!/bin/bash


. ./test_functions.sh

fixture

pushd ${TEST_REPO}
${DUPLICACY} init integration-tests $TEST_STORAGE -c  4 

# Create 10 small files
add_file file1 20
add_file file2 20
rm file3; touch file3
add_file file4 20
chmod u-r file4
add_file file5 20
add_file file6 20
add_file file7 20
add_file file8 20
add_file file9 20
add_file file10 20

# Fail at the 10th chunk
env DUPLICACY_FAIL_CHUNK=10 ${DUPLICACY} backup

# Try it again to test the multiple-resume case
env DUPLICACY_FAIL_CHUNK=5 ${DUPLICACY} backup
add_file file1 20
add_file file2 20

# Fail the backup before uploading the snapshot
env DUPLICACY_FAIL_SNAPSHOT=true ${DUPLICACY} backup

# Now complete the backup
${DUPLICACY} backup
${DUPLICACY} check --files
popd
