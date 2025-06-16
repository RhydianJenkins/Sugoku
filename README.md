# Sugoku
Sudoku solver webapp

![image](https://github.com/RhydianJenkins/Sugoku/assets/9198690/8b926962-32ff-4880-9a43-5c3804053ecf)

## Getting started

Run and done.

```sh
docker build -t sugoku .
docker run -p 8080:8080 sugoku

# ... or just `make run` if you want to run on host
# then in a separate terminal
curl localhost:8080
```

## TODO

- [ ] More exhaustive backtracking techniques
- [ ] Prettier frontend
