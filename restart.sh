lsof -i:8000 | awk 'NR==2 {print $2}' | uniq | xargs kill -9
rm -f main
go build main.go
chmod 777 main
nohup ./main > log