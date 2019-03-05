Check if directory is git

```
 isGit, err := IsGitDir(path string)
 if err {
    handle err ...
 }
 if isGit {
    do ...
 }
```

You can change the path for git binary by setting
GitCmdStr
eg. `goisgit.GitCmdStr = "D:/ProgramFiles/git/git.exe"`
