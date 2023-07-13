# go-tree
使用Go语言实现的一个简易版本的tree命令

## Usage
使用示例
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
命令语法
go-tree <argc> <argv> <dir>
参数解释
-I 自定排除目录,多个目录之间使用逗号分割
dir 默认是当前目录