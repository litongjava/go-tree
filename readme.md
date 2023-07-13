# go-tree
A simple version of the tree command implemented using the Go language

## Usage
Use example
```
go-tree -I node_modules,.vscode F:\code\frontend\project-litongjava\litongjava-vs-code-kit
```
output
```
litongjava-vs-code-kit
├── .eslintrc.json             
├── .vscodeignore              
├── .yarnrc                    
├── CHANGELOG.md               
├── README.md                  
├── extension.js               
├── jsconfig.json              
├── package.json               
├── test                       
│   ├── runTest.js             
│   ├── suite                  
│   │   ├── extension.test.js  
│   │   ├── index.js           
├── vsc-extension-quickstart.md
├── yarn.lock     
```
Command syntax
```
go-tree <argc> <argv> <dir>
```
Parameter interpretation
-I Excluding directories, use commas to separate multiple directories
dir The default is the current directory