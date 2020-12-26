sudo rm -rf medusa
git clone https://github.com/gfes980615/medusa.git
cd medusa
pyinstaller -F main.py
cd dist
sudo rm /usr/local/medusa/dist/main
cd ..
sudo cp ./dist/main /usr/local/medusa/
sudo systemctl stop medusa
sudo systemctl daemon-reload
sudo systemctl start medusa