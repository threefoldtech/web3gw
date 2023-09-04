
if ! [ -x "$(command -v cargo)" ]; then
    echo '- rust is not installed.'
    pushd /tmp
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
    popd
fi

if ! [ -x "$(command -v mdbook)" ]; then
    cargo install mdbook
    cargo install mdbook-mermaid
    #mdbook-mermaid install .
fi