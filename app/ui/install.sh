echo "Starting installation..."

if [[ ! -f "htmx.min.js" ]]
then
    echo "--Installing htmx..."
    curl -sLO https://unpkg.com/htmx.org/dist/htmx.min.js
    mv htmx.min.js static/js/htmx.min.js
fi

if [[ ! -f "_hyperscript.min.js" ]]
then
    echo "--Installing hyperscript..."
    curl -sLO https://unpkg.com/hyperscript.org@0.9.8/dist/_hyperscript.min.js
    mv _hyperscript.min.js static/js/_hyperscript.min.js
fi

# TODO: Check other paths for binary executable ?
if [[ ! -f "tailwindcss" ]]
then
    # checks os and architecture for correct release
    # https://stackoverflow.com/a/8597411 
    echo "--Installing tailwind..."
    ASSET="tailwindcss"

    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        ASSET="$ASSET-linux"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        ASSET="$ASSET-macos"
    fi
    if [[ "$(uname -m)" == "x86_64"* ]]; then
        ASSET="$ASSET-x64"
    elif [[ "$(uname -m)" == "arm64"* ]]; then
        ASSET="$ASSET-arm64"
    fi

    curl -sLO "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/${ASSET}"
    chmod +x $ASSET
    mv $ASSET tailwindcss        
fi