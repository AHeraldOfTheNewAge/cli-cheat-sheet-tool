# cli-cheat-sheet-tool
A linux cli cheat sheet tool built with GO(Golang).
You can create multiple cheat sheets, and view them in the terminal whenever you need them!

How to set it up:

Clone the repo: git clone https://github.com/AHeraldOfTheNewAge/cli-cheat-sheet-tool
Compile the .go file: go build -ldflags="-X main.assetsDir=/usr/local/share/cheetsheets" -o cheetsheet cheetsheet.go
Move the compiled binary: sudo mv cheetsheet /usr/local/bin
Make sure you can exectute the tool: sudo chmod +x /usr/local/bin/cheetsheet
Copy the folder cheetsheets to usr/local/share: sudo cp -r cheetsheets /usr/local/share/cheetsheets

Usage:
Edit your cheat sheets in the /usr/local/share/cheatsheets
Call the app: cheetsheet docker - for example, this way you can view your docker cheat sheet, right in your terminal!
