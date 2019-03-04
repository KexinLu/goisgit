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