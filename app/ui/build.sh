echo "Starting build..."

bash install.sh

# initialized and configures tailwind if not configured
echo "Initializing tailwind..."
if [[ ! -f "tailwind.config.js" ]]
then
    ./tailwindcss init
    sed -i '' "s|  content: \\[\\],|  content: \\['./templates/**/*.html'\\],|g" tailwind.config.js
fi

# compiles tailwind css for prod & builds project
echo "Compiling tailwindcss..."
rm -rf public static/css
./tailwindcss -i index.css -o ./static/css/index.css --minify

cp ../../playground/index.html templates/playground.html
cp -an ../../playground/static/css/ static/css
cp -an ../../playground/static/js/ static/js