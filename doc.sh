set -ex

rm -rf docs/_docs
rm -rf docs/_docs/v

pushd manual
bash run.sh
popd

# v fmt -w .
# v doc -m -f html . -readme -comments -no-timestamp

# mv _docs docs/v/