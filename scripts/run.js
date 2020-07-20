const { execSync, execFile } = require('child_process');
const path = require('path');

execSync(`go build ${process.argv[2]}/main.go`);
execFile(path.join(process.cwd(), 'main.exe'), (err, data) =>
  err ? console.log(err) : console.log(data)
);
