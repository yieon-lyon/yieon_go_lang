echo '--- ex5.1.1.go'
go run ex5.1.1.go
echo '--- ex5.1.3.go'
go run ex5.1.3.go
echo '--- ex5.1.4.go'
go run ex5.1.4.go
echo '--- ex5.1.5.go'
go run ex5.1.5.go
echo '\n--- ex5.2.2.go'
go run ex5.2.2.go << EOF
3 4
EOF
go run ex5.2.2.go << EOF
Hello 4
EOF
go run ex5.2.2.go << EOF
4 Hello
EOF
echo '--- ex5.2.3.go'
go run ex5.2.3.go << EOF
3 4
EOF
go run ex5.2.3.go << EOF
Hello 4
EOF
go run ex5.2.3.go << EOF
4 Hello
EOF
echo '--- ex5.2.4.go'
go run ex5.2.4.go << EOF
3 4
EOF
go run ex5.2.4.go << EOF
Hello 4
EOF
go run ex5.2.4.go << EOF
4 Hello
EOF
echo '--- ex5.3.go'
go run ex5.3.go << EOF
3 4
EOF