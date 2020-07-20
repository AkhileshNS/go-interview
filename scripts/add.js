const fs = require('fs');
const util = require('util');
const path = require('path');

const mkdir = util.promisify(fs.mkdir);
const exists = util.promisify(fs.exists);
const writeFile = util.promisify(fs.writeFile);

(async () => {
  const name = process.argv[2];

  if (!(await exists(path.join(process.cwd(), name)))) {
    await mkdir(path.join(process.cwd(), name));
    await writeFile(
      path.join(process.cwd(), name, 'main.go'),
      `
    package main

    import "fmt"
    
    func myFunction(arg string) string {
    
        // Write the body of your function here
        
    
        return "running with " + arg
    }
    
    func main() {
    
        // Run your function through some test cases here.
        // Remember: debuggin is half the battle!
        fmt.Println(myFunction("test input"))
    }    
    `
    );
  }
})();
